package financial

import (
	"context"
	"math"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/vflopes/nihil/pkg/analysis"
)

type FinancialReturn struct {
	valueGetter analysis.DataPointValueGetter
	log         zerolog.Logger
}

func NewFinancialReturn() *FinancialReturn {
	return &FinancialReturn{
		valueGetter: analysis.GetAbsoluteValue,
		log:         log.With().Str("analysis_financial", "financial_return").Logger(),
	}
}

func (f *FinancialReturn) WithValueGetter(valueGetter analysis.DataPointValueGetter) *FinancialReturn {
	f.valueGetter = valueGetter
	return f
}

func (f *FinancialReturn) Do(ctx context.Context, parameters *analysis.FinancialReturnParameters, src *analysis.ThreadSafeSeries) *analysis.ThreadSafeSeries {

	dst := analysis.NewThreadSafeSeries()
	periods := int(parameters.Periods)

	for i := range src.DataPoints {

		if i == 0 {
			srcDataPoint := src.GetDataPoint(i)
			dstDataPoint := &analysis.DataPoint{
				Timestamp:     srcDataPoint.Timestamp,
				AbsoluteValue: 0,
			}
			dst.AppendDataPoint(dstDataPoint)
			continue
		}

		if i%periods == 0 {
			srcDataPoint := src.GetDataPoint(i)
			currentValue := f.valueGetter(srcDataPoint)
			previousValue := f.valueGetter(src.DataPoints[i-periods])

			dstDataPoint := &analysis.DataPoint{
				Timestamp:     srcDataPoint.Timestamp,
				AbsoluteValue: 0,
			}

			if previousValue != 0 {
				switch parameters.Function {
				case analysis.FinancialReturnParameters_RATIO:
					dstDataPoint.AbsoluteValue = (currentValue / previousValue) - 1
				case analysis.FinancialReturnParameters_NATURAL_LOGARITHM:
					dstDataPoint.AbsoluteValue = math.Log(currentValue / previousValue)
				default:
					f.log.Fatal().
						Int32("function_enum", int32(parameters.Function)).
						Msg("invalid function to calculate financial return")
				}
			}

			dst.AppendDataPoint(dstDataPoint)

		}

	}

	dst.Unit = analysis.Series_RATIO
	dst.Name = "financial_return"

	return dst

}
