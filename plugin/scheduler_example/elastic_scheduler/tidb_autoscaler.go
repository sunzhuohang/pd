package elastic_scheduler

import (
	"fmt"
	"github.com/pingcap/tidb-operator/pkg/apis/pingcap/v1alpha1"
	"github.com/pingcap/tidb-operator/pkg/label"
	operatorUtils "github.com/pingcap/tidb-operator/pkg/util"
	appsv1 "k8s.io/api/apps/v1"
	"time"
)

func (esm *elasticSchedulerManager) syncTiDB(tc *v1alpha1.TidbCluster, tac *v1alpha1.TidbClusterAutoScaler) error {
	if tac.Spec.TiDB == nil {
		return nil
	}
	sts, err := esm.stsLister.StatefulSets(tc.Namespace).Get(operatorUtils.GetStatefulSetName(tc, v1alpha1.TiDBMemberType))
	if err != nil {
		return err
	}
	if !checkAutoScalingPrerequisites(tc, sts, v1alpha1.TiDBMemberType) {
		return nil
	}
	currentReplicas := tc.Spec.TiDB.Replicas
	targetReplicas := esm.targetReplicas
	targetReplicas = limitTargetReplicas(targetReplicas, tac, v1alpha1.TiDBMemberType)
	if targetReplicas == tc.Spec.TiDB.Replicas {
		return nil
	}
	return syncTiDBAfterCalculated(tc, tac, currentReplicas, targetReplicas, sts)
}

// syncTiDBAfterCalculated would check the Consecutive count to avoid jitter, and it would also check the interval
// duration between each auto-scaling. If either of them is not meet, the auto-scaling would be rejected.
// If the auto-scaling is permitted, the timestamp would be recorded and the Consecutive count would be zeroed.
func syncTiDBAfterCalculated(tc *v1alpha1.TidbCluster, tac *v1alpha1.TidbClusterAutoScaler, currentReplicas, recommendedReplicas int32, sts *appsv1.StatefulSet) error {
	intervalSeconds := tac.Spec.TiDB.ScaleInIntervalSeconds
	if recommendedReplicas > currentReplicas {
		intervalSeconds = tac.Spec.TiDB.ScaleOutIntervalSeconds
	}
	ableToScale, err := checkStsAutoScalingInterval(tac, *intervalSeconds, v1alpha1.TiDBMemberType)
	if err != nil {
		return err
	}
	if !ableToScale {
		return nil
	}
	return updateTcTiDBIfScale(tc, tac, recommendedReplicas)
}

// Currently we didnt' record the auto-scaling out slot for tidb, because it is pointless for now.
func updateTcTiDBIfScale(tc *v1alpha1.TidbCluster, tac *v1alpha1.TidbClusterAutoScaler, recommendedReplicas int32) error {
	tac.Annotations[label.AnnTiDBLastAutoScalingTimestamp] = fmt.Sprintf("%d", time.Now().Unix())
	tc.Spec.TiDB.Replicas = recommendedReplicas
	return nil
}




