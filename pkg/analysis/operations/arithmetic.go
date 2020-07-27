package operations

import (
	"context"
	"math"

	"github.com/vflopes/nihil/pkg/analysis"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Arithmetic struct {
	valueGetter analysis.DataPointValueGetter
	log         zerolog.Logger
}

func NewArithmetic() *Arithmetic {
	return &Arithmetic{
		valueGetter: analysis.GetAbsoluteValue,
		log:         log.With().Str("analysis_operations", "arithmetic").Logger(),
	}
}

func (m *Arithmetic) WithValueGetter(valueGetter analysis.DataPointValueGetter) *Arithmetic {
	m.valueGetter = valueGetter
	return m
}

func (m *Arithmetic) Do(ctx context.Context, parameters *analysis.ArithmeticOperationsParameters, src *analysis.ThreadSafeSeries) *analysis.ThreadSafeSeries {

	dst := analysis.NewThreadSafeSeries()
	value := parameters.Value
	switch parameters.Operation {
	case analysis.ArithmeticOperationsParameters_TIMES_SQRT:
		value = math.Sqrt(value)
	}

	for i := range src.DataPoints {

		srcDataPoint := src.GetDataPoint(i)
		currentValue := m.valueGetter(srcDataPoint)

		dstDataPoint := &analysis.DataPoint{
			Timestamp: srcDataPoint.Timestamp,
		}

		switch parameters.Operation {
		case analysis.ArithmeticOperationsParameters_TIMES_SQRT:
			dstDataPoint.AbsoluteValue = currentValue * value
		default:
			m.log.Fatal().
				Int32("operation_enum", int32(parameters.Operation)).
				Msg("invalid arithmetic operation")
		}

		dst.AppendDataPoint(dstDataPoint)

	}

	dst.Unit = src.Unit
	dst.Name = "arithmetic_operation"

	return dst

}
