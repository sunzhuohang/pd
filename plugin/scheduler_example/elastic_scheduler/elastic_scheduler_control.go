package elastic_scheduler

import (
	"github.com/pingcap/tidb-operator/pkg/apis/pingcap/v1alpha1"
	"k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/client-go/tools/record"
)

type ControlInterface interface {
	ResconcileAutoScaler(ta *v1alpha1.TidbClusterAutoScaler) error
}

type defaultElasticSchedulerControl struct {
	recoder           record.EventRecorder
	esm ElasticSchedulerManager
}

func newDefaultElasticSchedulerControl(recorder record.EventRecorder, esm ElasticSchedulerManager) ControlInterface{
	return &defaultElasticSchedulerControl{
		recoder:    recorder,
		esm: 		esm,
	}
}

func (tac *defaultElasticSchedulerControl) ResconcileAutoScaler(ta *v1alpha1.TidbClusterAutoScaler) error {
	var errs []error
	if err := tac.reconcileAutoScaler(ta); err != nil {
		errs = append(errs, err)
	}
	return errors.NewAggregate(errs)
}

func (tac *defaultElasticSchedulerControl) reconcileAutoScaler(ta *v1alpha1.TidbClusterAutoScaler) error {
	return tac.esm.Sync(ta)
}