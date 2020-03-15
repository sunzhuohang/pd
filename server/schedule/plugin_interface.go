// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package schedule

import (
	"github.com/pingcap/pd/v4/server/schedule/opt"
	"path/filepath"
	"plugin"
	"sync"
	"time"

	"github.com/pingcap/log"
	"go.uber.org/zap"
)

// ScheduleOperationList 是存放所有调度的信息的列表，即调度计划
var OpList []MoveRegionOp

// ScheduleListLock 是操作 ScheduleOperationList 时要上的锁
var OpListLock = sync.RWMutex{}

// Schedule_MoveRegion 是一次调度Region的操作
type MoveRegionOp struct {
	RegionIDs []uint64
	StoreIDs  []uint64
	StartTime time.Time
	EndTime   time.Time
}

// IsPredictedHotRegion 判断该region是否已在调度计划-OpList
// 这个函数是留给其他调度器（如balance_leader.go）执行的
func IsPredictedHotRegion(cluster opt.Cluster, regionID uint64) bool {
	for _, oneScheduleOperation := range OpList {
		currentTime := time.Now()
		if currentTime.After(oneScheduleOperation.EndTime) || currentTime.Before(oneScheduleOperation.StartTime) {
			continue
		}
		for _, id := range oneScheduleOperation.RegionIDs {
			if id == regionID {
				log.Info("region is predicted hot", zap.Uint64("region-id", regionID))
				return true
			}
		}
	}
	return false
}

// PluginInterface is used to manage all plugin.
type PluginInterface struct {
	pluginMap     map[string]*plugin.Plugin
	pluginMapLock sync.RWMutex
}

// NewPluginInterface create a plugin interface
func NewPluginInterface() *PluginInterface {
	return &PluginInterface{
		pluginMap:     make(map[string]*plugin.Plugin),
		pluginMapLock: sync.RWMutex{},
	}
}

// GetFunction gets func by funcName from plugin(.so)
func (p *PluginInterface) GetFunction(path string, funcName string) (plugin.Symbol, error) {
	p.pluginMapLock.Lock()
	defer p.pluginMapLock.Unlock()
	if _, ok := p.pluginMap[path]; !ok {
		//open plugin
		filePath, err := filepath.Abs(path)
		if err != nil {
			return nil, err
		}
		log.Info("open plugin file", zap.String("file-path", filePath))
		plugin, err := plugin.Open(filePath)
		if err != nil {
			return nil, err
		}
		p.pluginMap[path] = plugin
	}
	//get func from plugin
	f, err := p.pluginMap[path].Lookup(funcName)
	if err != nil {
		log.Error("Lookup func error!")
		return nil, err
	}
	return f, nil
}
