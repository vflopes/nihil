package indicators

import (
	"context"

	"github.com/vflopes/nihil/pkg/analysis"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type MovingAverage struct {
	valueGetter analysis.DataPointValueGetter
	log         zerolog.Logger
}

func NewMovingAverage() *MovingAverage {
	return &MovingAverage{
		valueGetter: analysis.GetAbsoluteValue,
		log:         log.With().Str("analysis_indicator", "moving_average").Logger(),
	}
}

func (m *MovingAverage) WithValueGetter(valueGetter analysis.DataPointValueGetter) *MovingAverage {
	m.valueGetter = valueGetter
	return m
}

func (m *MovingAverage) Do(ctx context.Context, parameters *analysis.MovingAverageParameters, src *analysis.ThreadSafeSeries) *analysis.ThreadSafeSeries {

	dst := analysis.NewThreadSafeSeries()
	periods := int(parameters.Periods)
	granularity := int(parameters.Granularity)

	sum := float64(0)

	for i := range src.DataPoints {
		srcDataPoint := src.GetDataPoint(i)
		currentValue := m.valueGetter(srcDataPoint)
		totalItems := periods
		if i < periods {
			totalItems = i + 1
		}
		if i >= periods {
			previousDataPoint := src.GetDataPoint(i - periods)
			sum = sum - m.valueGetter(previousDataPoint)
		}
		sum += currentValue

		if i == 0 {
			srcDataPoint := src.GetDataPoint(i)
			dstDataPoint := &analysis.DataPoint{
				Timestamp:     srcDataPoint.Timestamp,
				AbsoluteValue: currentValue,
			}
			dst.AppendDataPoint(dstDataPoint)
			continue
		}

		if i%granularity == 0 {
			dstDataPoint := &analysis.DataPoint{
				Timestamp: srcDataPoint.Timestamp,
			}

			switch parameters.Mode {
			case analysis.MovingAverageParameters_SIMPLE:
				dstDataPoint.AbsoluteValue = sum / float64(totalItems)
			default:
				m.log.Fatal().
					Int32("mode_enum", int32(parameters.Mode)).
					Msg("invalid mode to calculate moving average")
			}
			dst.AppendDataPoint(dstDataPoint)
		}

	}

	dst.Unit = src.Unit
	dst.Name = "moving_average"

	return dst

}
