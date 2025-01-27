// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics

import (
	"net/http"

	// nolint:gosec
	_ "net/http/pprof"

	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/util/typeutil"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

const (
	milvusNamespace = "milvus"

	AbandonLabel = "abandon"

	SearchLabel = "search"
	QueryLabel  = "query"

	CacheHitLabel  = "hit"
	CacheMissLabel = "miss"
)

var (
	//DataCoordDataNodeList records the num of regsitered data nodes
	DataCoordDataNodeList = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: milvusNamespace,
			Subsystem: typeutil.DataCoordRole,
			Name:      "list_of_data_node",
			Help:      "List of data nodes registered within etcd",
		}, []string{"status"},
	)
)

//RegisterDataCoord registers DataCoord metrics
func RegisterDataCoord() {
	prometheus.MustRegister(DataCoordDataNodeList)
}

var (
	// DataNodeFlushSegmentsCounter counts the num of calls of FlushSegments
	DataNodeFlushSegmentsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: milvusNamespace,
			Subsystem: typeutil.DataNodeRole,
			Name:      "flush_segments_total",
			Help:      "Counter of flush segments",
		}, []string{"type"})

	// DataNodeWatchDmChannelsCounter counts the num of calls of WatchDmChannels
	DataNodeWatchDmChannelsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: milvusNamespace,
			Subsystem: typeutil.DataNodeRole,
			Name:      "watch_dm_channels_total",
			Help:      "Counter of watch dm channel",
		}, []string{"type"})
)

//ServeHTTP serves prometheus http service
func ServeHTTP() {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		if err := http.ListenAndServe(":9091", nil); err != nil {
			log.Error("handle metrics failed", zap.Error(err))
		}
	}()
}
