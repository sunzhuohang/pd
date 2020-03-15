package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pingcap/pd/v4/server/schedule"
	"github.com/pingcap/pd/v4/server/schedule/opt"
)

type userBaseScheduler struct {
	opController *schedule.OperatorController
}

const (
	// MaxScheduleInterval is
	MaxScheduleInterval = time.Second * 5
	// MinScheduleInterval is
	MinScheduleInterval = time.Millisecond * 10
	// ScheduleIntervalFactor is
	ScheduleIntervalFactor = 1.3
)

// 函数-newUserBaseScheduler 创建一个 userBaseScheduler 并传入 opController
func newUserBaseScheduler(opController *schedule.OperatorController) *userBaseScheduler {
	return &userBaseScheduler{opController: opController}
}

func (s *userBaseScheduler) Prepare(cluster opt.Cluster) error { return nil }

func (s *userBaseScheduler) Cleanup(cluster opt.Cluster) {}

func (s *userBaseScheduler) GetMinInterval() time.Duration {
	return MinScheduleInterval
}

func (s *userBaseScheduler) GetNextInterval(interval time.Duration) time.Duration {
	return minDuration(time.Duration(float64(interval)*ScheduleIntervalFactor), MaxScheduleInterval)
}

// TODO 函数-EncodeConfig 未实现
func (s *userBaseScheduler) EncodeConfig() ([]byte, error) {
	return schedule.EncodeConfig(nil)
}

// TODO 函数-ServeHTTP 未实现
func (s *userBaseScheduler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not implements")
}

func minDuration(a, b time.Duration) time.Duration {
	if a < b {
		return a
	}
	return b
}
