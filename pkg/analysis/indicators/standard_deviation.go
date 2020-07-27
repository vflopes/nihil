package indicators

import (
	"context"
	"math"

	"github.com/vflopes/nihil/pkg/analysis"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type StandardDeviation struct {
	valueGetter analysis.DataPointValueGetter
	log         zerolog.Logger
}

func NewStandardDeviation() *StandardDeviation {
	return &StandardDeviation{
		valueGetter: analysis.GetAbsoluteValue,
		log:         log.With().Str("analysis_indicator", "moving_average").Logger(),
	}
}

func (m *StandardDeviation) WithValueGetter(valueGetter analysis.DataPointValueGetter) *StandardDeviation {
	m.valueGetter = valueGetter
	return m
}

func (m *StandardDeviation) Do(ctx context.Context, parameters *analysis.StandardDeviationParameters, src *analysis.ThreadSafeSeries, mean *analysis.ThreadSafeSeries) *analysis.ThreadSafeSeries {

	var (
		i            int
		srcDataPoint *analysis.DataPoint
		currentValue float64
	)

	dst := analysis.NewThreadSafeSeries()

	currentMeanIndex := 0
	currentMean := mean.GetDataPoint(currentMeanIndex)
	meanValue := m.valueGetter(currentMean)
	periods := int(parameters.Periods)
	currentPeriods := 0
	previousValues := make([]float64, periods)

	for i = range src.DataPoints {
		if currentPeriods < periods {
			currentPeriods++
		}

		srcDataPoint = src.GetDataPoint(i)
		currentValue = m.valueGetter(srcDataPoint)
		if i >= periods {
			analysis.ShiftValueSlice(previousValues, currentValue)
		} else {
			previousValues[i] = currentValue
		}

		if i == 0 {
			dstDataPoint := &analysis.DataPoint{
				Timestamp:     srcDataPoint.Timestamp,
				AbsoluteValue: 0,
			}
			dst.AppendDataPoint(dstDataPoint)
			continue
		}

		nextMean := mean.GetDataPoint(currentMeanIndex + 1)

		if nextMean != nil && nextMean.Timestamp >= srcDataPoint.Timestamp {

			currentMeanIndex++
			currentMean = nextMean
			meanValue = m.valueGetter(currentMean)

		}

		dstDataPoint := &analysis.DataPoint{
			Timestamp:     srcDataPoint.Timestamp,
			AbsoluteValue: math.Sqrt(analysis.GetSliceVariance(previousValues, currentPeriods-1, meanValue)),
		}
		dst.AppendDataPoint(dstDataPoint)

	}

	dst.Unit = src.Unit
	dst.Name = "standard_deviation"

	return dst

}
