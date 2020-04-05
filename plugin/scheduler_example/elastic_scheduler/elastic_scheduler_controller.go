package elastic_scheduler

import (
	"fmt"
	perrors "github.com/pingcap/errors"
	"github.com/pingcap/tidb-operator/pkg/apis/pingcap/v1alpha1"
	"github.com/pingcap/tidb-operator/pkg/client/clientset/versioned"
	informers "github.com/pingcap/tidb-operator/pkg/client/informers/externalversions"
	listers "github.com/pingcap/tidb-operator/pkg/client/listers/pingcap/v1alpha1"
	"github.com/pingcap/tidb-operator/pkg/controller"
	"github.com/pingcap/tidb-operator/pkg/controller/autoscaler"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	eventv1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	"time"
)

type Controller struct {
	control  autoscaler.ControlInterface
	taLister listers.TidbClusterAutoScalerLister
	queue    workqueue.RateLimitingInterface
}

func NewController(
	kubeCli kubernetes.Interface,
	cli versioned.Interface,
	informerFactory informers.SharedInformerFactory,
	kubeInformerFactory kubeinformers.SharedInformerFactory,
	targetReplicas int32,
	) *Controller {
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(klog.Infof)
	eventBroadcaster.StartRecordingToSink(&eventv1.EventSinkImpl{
		Interface: eventv1.New(kubeCli.CoreV1().RESTClient()).Events("")})
	recorder := eventBroadcaster.NewRecorder(v1alpha1.Scheme, corev1.EventSource{Component: "tidbclusterautoscaler"})
	autoScalerInformer := informerFactory.Pingcap().V1alpha1().TidbClusterAutoScalers()
	esm := newElasticSchedulerManager(cli, informerFactory, kubeInformerFactory, recorder, targetReplicas)

	tac := &Controller{
		control:  newDefaultElasticSchedulerControl(recorder, esm),
		taLister: autoScalerInformer.Lister(),
		queue: workqueue.NewNamedRateLimitingQueue(
			workqueue.DefaultControllerRateLimiter(),
			"tidbclusterautoscaler"),
	}
	controller.WatchForObject(autoScalerInformer.Informer(), tac.queue)
	return tac
}

func (tac *Controller) Run()  {
	key, quit := tac.queue.Get()
	if quit {
		return
	}
	defer tac.queue.Done(key)
	if err := tac.sync(key.(string)); err != nil {
		if perrors.Find(err, controller.IsRequeueError) != nil {
			klog.Infof("TidbClusterAutoScaler: %v, still need sync: %v, requeuing", key.(string), err)
		} else {
			utilruntime.HandleError(fmt.Errorf("TidbClusterAutoScaler: %v, sync failed, err: %v", key.(string), err))
		}
		tac.queue.AddRateLimited(key)
	} else {
		tac.queue.Forget(key)
	}
	return
}

func (tac *Controller) sync(key string) error {
	startTime := time.Now()
	defer func() {
		klog.V(4).Infof("Finished syncing TidbClusterAutoScaler %q (%v)", key, time.Since(startTime))
	}()

	ns, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	ta, err := tac.taLister.TidbClusterAutoScalers(ns).Get(name)
	if errors.IsNotFound(err) {
		klog.Infof("TidbClusterAutoScaler has been deleted %v", key)
		return nil
	}
	if err != nil {
		return err
	}

	return tac.control.ResconcileAutoScaler(ta)
}