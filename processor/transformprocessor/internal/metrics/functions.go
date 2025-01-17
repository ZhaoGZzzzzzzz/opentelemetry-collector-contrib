// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/metrics"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/common"
)

func DataPointFunctions() map[string]ottl.Factory[ottldatapoint.TransformContext] {
	functions := common.Functions[ottldatapoint.TransformContext]()

	datapointFunctions := ottl.CreateFactoryMap[ottldatapoint.TransformContext](
		newConvertSumToGaugeFactory(),
		newConvertGaugeToSumFactory(),
		newConvertSummarySumValToSumFactory(),
		newConvertSummaryCountValToSumFactory(),
	)

	for k, v := range datapointFunctions {
		functions[k] = v
	}

	return functions
}

func MetricFunctions() map[string]ottl.Factory[ottlmetric.TransformContext] {
	return common.Functions[ottlmetric.TransformContext]()
}
