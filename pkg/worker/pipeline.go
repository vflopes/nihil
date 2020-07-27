package worker

import (
	"context"

	"github.com/vflopes/nihil/pkg/analysis"
	"github.com/vflopes/nihil/pkg/analysis/financial"
	"github.com/vflopes/nihil/pkg/analysis/indicators"
	"github.com/vflopes/nihil/pkg/analysis/operations"
)

func Do(ctx context.Context, parameters *analysis.AnalysisParameters, axis *analysis.Axis) *analysis.Axis {

	valueGetter := parameters.ValueFrom.GetValueGetter()

	switch params := parameters.Parameters.(type) {
	case *analysis.AnalysisParameters_BollingerBandsParameters:
		bolbands := indicators.NewBollingerBands().WithValueGetter(valueGetter)
		mean := axis.FindSeriesByName(params.BollingerBandsParameters.MeanName)
		stddev := axis.FindSeriesByName(params.BollingerBandsParameters.StandardDeviationName)
		dst := bolbands.Do(ctx, params.BollingerBandsParameters, mean.ToThreadSafe(), stddev.ToThreadSafe())
		return &analysis.Axis{
			Series: []*analysis.Series{dst.Series},
		}
	case *analysis.AnalysisParameters_ArithmeticOperationsParameters:
		arithm := operations.NewArithmetic().WithValueGetter(valueGetter)
		src := axis.FindSeriesByName(params.ArithmeticOperationsParameters.SourceName)
		dst := arithm.Do(ctx, params.ArithmeticOperationsParameters, src.ToThreadSafe())
		return &analysis.Axis{
			Series: []*analysis.Series{dst.Series},
		}
	case *analysis.AnalysisParameters_FinancialReturnParameters:
		finret := financial.NewFinancialReturn().WithValueGetter(valueGetter)
		src := axis.FindSeriesByName(params.FinancialReturnParameters.SourceName)
		dst := finret.Do(ctx, params.FinancialReturnParameters, src.ToThreadSafe())
		return &analysis.Axis{
			Series: []*analysis.Series{dst.Series},
		}
	case *analysis.AnalysisParameters_MovingAverageParameters:
		movavg := indicators.NewMovingAverage().WithValueGetter(valueGetter)
		src := axis.FindSeriesByName(params.MovingAverageParameters.SourceName)
		dst := movavg.Do(ctx, params.MovingAverageParameters, src.ToThreadSafe())
		return &analysis.Axis{
			Series: []*analysis.Series{dst.Series},
		}
	case *analysis.AnalysisParameters_StandardDeviationParameters:
		stddev := indicators.NewStandardDeviation().WithValueGetter(valueGetter)
		src := axis.FindSeriesByName(params.StandardDeviationParameters.SourceName)
		mean := axis.FindSeriesByName(params.StandardDeviationParameters.MeanName)
		dst := stddev.Do(ctx, params.StandardDeviationParameters, src.ToThreadSafe(), mean.ToThreadSafe())
		return &analysis.Axis{
			Series: []*analysis.Series{dst.Series},
		}
	}
	return axis
}
