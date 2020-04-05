package elastic_scheduler

import (
	"fmt"
	"github.com/pingcap/advanced-statefulset/pkg/apis/apps/v1/helper"
	"github.com/pingcap/tidb-operator/pkg/apis/pingcap/v1alpha1"
	"github.com/pingcap/tidb-operator/pkg/label"
	operatorUtils "github.com/pingcap/tidb-operator/pkg/util"
	appsv1 "k8s.io/api/apps/v1"
	"time"
)

func (esm *elasticSchedulerManager) syncTiKV(tc *v1alpha1.TidbCluster, tac *v1alpha1.TidbClusterAutoScaler) error {
	if tac.Spec.TiKV == nil {
		return nil
	}
	sts, err := esm.stsLister.StatefulSets(tc.Namespace).Get(operatorUtils.GetStatefulSetName(tc, v1alpha1.TiKVMemberType))
	if err != nil {
		return err
	}
	if !checkAutoScalingPrerequisites(tc, sts, v1alpha1.TiKVMemberType) {
		return nil
	}
	instances := filterTiKVInstances(tc)
	currentReplicas := int32(len(instances))
	targetReplicas := esm.targetReplicas
	targetReplicas = limitTargetReplicas(targetReplicas, tac, v1alpha1.TiKVMemberType)
	if targetReplicas == tc.Spec.TiKV.Replicas {
		return nil
	}
	return syncTiKVAfterCalculated(tc, tac, currentReplicas, targetReplicas, sts)
}

// syncTiKVAfterCalculated would check the Consecutive count to avoid jitter, and it would also check the interval
// duration between each auto-scaling. If either of them is not meet, the auto-scaling would be rejected.
// If the auto-scaling is permitted, the timestamp would be recorded and the Consecutive count would be zeroed.
// The currentReplicas of TiKV calculated in auto-scaling is the count of the StateUp TiKV instance, so we need to
// add the number of other state tikv instance replicas when we update the TidbCluster.Spec.TiKV.Replicas
func syncTiKVAfterCalculated(tc *v1alpha1.TidbCluster, tac *v1alpha1.TidbClusterAutoScaler, currentReplicas, recommendedReplicas int32, sts *appsv1.StatefulSet) error {

	intervalSeconds := tac.Spec.TiKV.ScaleInIntervalSeconds
	if recommendedReplicas > tc.Spec.TiKV.Replicas {
		intervalSeconds = tac.Spec.TiKV.ScaleOutIntervalSeconds
	}
	ableToScale, err := checkStsAutoScalingInterval(tac, *intervalSeconds, v1alpha1.TiKVMemberType)
	if err != nil {
		return err
	}
	if !ableToScale {
		return nil
	}
	return updateTcTiKVIfScale(tc, tac, currentReplicas, recommendedReplicas, sts)
}

//TODO: fetch tikv instances info from pdapi in future
func filterTiKVInstances(tc *v1alpha1.TidbCluster) []string {
	var instances []string
	for _, store := range tc.Status.TiKV.Stores {
		if store.State == v1alpha1.TiKVStateUp {
			instances = append(instances, store.PodName)
		}
	}
	return instances
}

// we record the auto-scaling out slot for tikv, in order to add special hot labels when they are created
func updateTcTiKVIfScale(tc *v1alpha1.TidbCluster, tac *v1alpha1.TidbClusterAutoScaler, currentReplicas, recommendedReplicas int32, sts *appsv1.StatefulSet) error {
	tac.Annotations[label.AnnTiKVLastAutoScalingTimestamp] = fmt.Sprintf("%d", time.Now().Unix())
	if recommendedReplicas > currentReplicas {
		newlyScaleOutOrdinalSets := helper.GetPodOrdinals(recommendedReplicas, sts).Difference(helper.GetPodOrdinals(currentReplicas, sts))
		if newlyScaleOutOrdinalSets.Len() > 0 {
			if tc.Annotations == nil {
				tc.Annotations = map[string]string{}
			}
			existed := operatorUtils.GetAutoScalingOutSlots(tc, v1alpha1.TiKVMemberType)
			v, err := operatorUtils.Encode(newlyScaleOutOrdinalSets.Union(existed).List())
			if err != nil {
				return err
			}
			tc.Annotations[label.AnnTiKVAutoScalingOutOrdinals] = v
		}
	}
	tc.Spec.TiKV.Replicas = recommendedReplicas
	return nil
}
