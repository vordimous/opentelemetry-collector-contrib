// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper/scrapertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testDataSet int

const (
	testDataSetDefault testDataSet = iota
	testDataSetAll
	testDataSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name        string
		metricsSet  testDataSet
		resAttrsSet testDataSet
		expectEmpty bool
	}{
		{
			name: "default",
		},
		{
			name:        "all_set",
			metricsSet:  testDataSetAll,
			resAttrsSet: testDataSetAll,
		},
		{
			name:        "none_set",
			metricsSet:  testDataSetNone,
			resAttrsSet: testDataSetNone,
			expectEmpty: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := scrapertest.NewNopSettings(scrapertest.NopType)
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, tt.name), settings, WithStartTime(start))

			expectedWarnings := 0

			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSystemNetworkConnectionsDataPoint(ts, 1, AttributeProtocolTcp, "state-val")

			allMetricsCount++
			mb.RecordSystemNetworkConntrackCountDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordSystemNetworkConntrackMaxDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSystemNetworkDroppedDataPoint(ts, 1, "device-val", AttributeDirectionReceive)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSystemNetworkErrorsDataPoint(ts, 1, "device-val", AttributeDirectionReceive)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSystemNetworkIoDataPoint(ts, 1, "device-val", AttributeDirectionReceive)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSystemNetworkPacketsDataPoint(ts, 1, "device-val", AttributeDirectionReceive)

			res := pcommon.NewResource()
			metrics := mb.Emit(WithResource(res))

			if tt.expectEmpty {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			assert.Equal(t, res, rm.Resource())
			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if tt.metricsSet == testDataSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if tt.metricsSet == testDataSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "system.network.connections":
					assert.False(t, validatedMetrics["system.network.connections"], "Found a duplicate in the metrics slice: system.network.connections")
					validatedMetrics["system.network.connections"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of connections.", ms.At(i).Description())
					assert.Equal(t, "{connections}", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("protocol")
					assert.True(t, ok)
					assert.Equal(t, "tcp", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("state")
					assert.True(t, ok)
					assert.Equal(t, "state-val", attrVal.Str())
				case "system.network.conntrack.count":
					assert.False(t, validatedMetrics["system.network.conntrack.count"], "Found a duplicate in the metrics slice: system.network.conntrack.count")
					validatedMetrics["system.network.conntrack.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The count of entries in conntrack table.", ms.At(i).Description())
					assert.Equal(t, "{entries}", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "system.network.conntrack.max":
					assert.False(t, validatedMetrics["system.network.conntrack.max"], "Found a duplicate in the metrics slice: system.network.conntrack.max")
					validatedMetrics["system.network.conntrack.max"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The limit for entries in the conntrack table.", ms.At(i).Description())
					assert.Equal(t, "{entries}", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "system.network.dropped":
					assert.False(t, validatedMetrics["system.network.dropped"], "Found a duplicate in the metrics slice: system.network.dropped")
					validatedMetrics["system.network.dropped"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of packets dropped.", ms.At(i).Description())
					assert.Equal(t, "{packets}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("device")
					assert.True(t, ok)
					assert.Equal(t, "device-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.Equal(t, "receive", attrVal.Str())
				case "system.network.errors":
					assert.False(t, validatedMetrics["system.network.errors"], "Found a duplicate in the metrics slice: system.network.errors")
					validatedMetrics["system.network.errors"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of errors encountered.", ms.At(i).Description())
					assert.Equal(t, "{errors}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("device")
					assert.True(t, ok)
					assert.Equal(t, "device-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.Equal(t, "receive", attrVal.Str())
				case "system.network.io":
					assert.False(t, validatedMetrics["system.network.io"], "Found a duplicate in the metrics slice: system.network.io")
					validatedMetrics["system.network.io"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of bytes transmitted and received.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("device")
					assert.True(t, ok)
					assert.Equal(t, "device-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.Equal(t, "receive", attrVal.Str())
				case "system.network.packets":
					assert.False(t, validatedMetrics["system.network.packets"], "Found a duplicate in the metrics slice: system.network.packets")
					validatedMetrics["system.network.packets"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of packets transferred.", ms.At(i).Description())
					assert.Equal(t, "{packets}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("device")
					assert.True(t, ok)
					assert.Equal(t, "device-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.Equal(t, "receive", attrVal.Str())
				}
			}
		})
	}
}
