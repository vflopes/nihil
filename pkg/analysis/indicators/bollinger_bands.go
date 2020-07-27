package indicators

import (
	"context"

	"github.com/vflopes/nihil/pkg/analysis"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type BollingerBands struct {
	valueGetter analysis.DataPointValueGetter
	log         zerolog.Logger
}

func NewBollingerBands() *BollingerBands {
	return &BollingerBands{
		valueGetter: analysis.GetAbsoluteValue,
		log:         log.With().Str("analysis_indicator", "bollinger_bands").Logger(),
	}
}

func (m *BollingerBands) WithValueGetter(valueGetter analysis.DataPointValueGetter) *BollingerBands {
	m.valueGetter = valueGetter
	return m
}

func (m *BollingerBands) Do(ctx context.Context, parameters *analysis.BollingerBandsParameters, mean *analysis.ThreadSafeSeries, stddev *analysis.ThreadSafeSeries) *analysis.ThreadSafeSeries {

	dst := analysis.NewThreadSafeSeries()
	factor := parameters.Factor

	for i := range mean.DataPoints {
		meanDataPoint := mean.GetDataPoint(i)
		stddevDataPoint := stddev.GetDataPoint(i)

		meanValue := m.valueGetter(meanDataPoint)
		stddevValue := m.valueGetter(stddevDataPoint) * factor

		dstDataPoint := &analysis.DataPoint{
			Timestamp: meanDataPoint.Timestamp,
			Candlestick: &analysis.Candlestick{
				High: meanValue + stddevValue,
				Low:  meanValue - stddevValue,
			},
		}
		dst.AppendDataPoint(dstDataPoint)

	}

	dst.Unit = mean.Unit
	dst.Name = "bollinger_bands"

	return dst

}
