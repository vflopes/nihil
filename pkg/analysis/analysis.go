package analysis

import (
	sync "sync"
)

type DataPointValueGetter func(*DataPoint) float64

func GetAbsoluteValue(dp *DataPoint) float64 {
	return dp.AbsoluteValue
}

func GetCandlestickHigh(dp *DataPoint) float64 {
	return dp.Candlestick.High
}

func GetCandlestickOpen(dp *DataPoint) float64 {
	return dp.Candlestick.Open
}

func GetCandlestickClose(dp *DataPoint) float64 {
	return dp.Candlestick.Close
}

func GetCandlestickLow(dp *DataPoint) float64 {
	return dp.Candlestick.Low
}

type ThreadSafeSeries struct {
	*Series
	lock *sync.RWMutex
}

func NewThreadSafeSeries() *ThreadSafeSeries {
	return &ThreadSafeSeries{
		Series: &Series{},
		lock:   new(sync.RWMutex),
	}
}

func (x *ThreadSafeSeries) GetDataPoint(i int) *DataPoint {
	x.lock.RLock()
	if len(x.DataPoints) <= i {
		x.lock.RUnlock()
		return nil
	}
	dp := x.DataPoints[i]
	x.lock.RUnlock()
	return dp
}

func (x *ThreadSafeSeries) AppendDataPoint(dp *DataPoint) {
	x.lock.Lock()
	if x.DataPoints == nil {
		x.DataPoints = []*DataPoint{}
	}
	x.DataPoints = append(x.DataPoints, dp)
	x.lock.Unlock()
}

func (x *Series) ToThreadSafe() *ThreadSafeSeries {
	t := NewThreadSafeSeries()
	t.Series = x
	return t
}

func (x *Axis) FindSeriesByName(name string) *Series {
	for _, series := range x.Series {
		if series.Name == name {
			return series
		}
	}
	return nil
}

func (x *Axis) PutSeries(series *Series) {
	for i, s := range x.Series {
		if series.Name == s.Name {
			x.Series[i] = series
			return
		}
	}
	x.Series = append(x.Series, series)
	return
}

func (x AnalysisParameters_ValueGetter) GetValueGetter() DataPointValueGetter {
	switch x {
	case AnalysisParameters_CANDLESTICK_HIGH:
		return GetCandlestickHigh
	case AnalysisParameters_CANDLESTICK_LOW:
		return GetCandlestickLow
	case AnalysisParameters_CANDLESTICK_OPEN:
		return GetCandlestickOpen
	case AnalysisParameters_CANDLESTICK_CLOSE:
		return GetCandlestickClose
	}
	return GetAbsoluteValue
}
