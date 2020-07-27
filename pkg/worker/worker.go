package worker

import (
	context "context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/vflopes/nihil/pkg/analysis"
)

type Worker struct {
	log zerolog.Logger
}

func NewWorker() *Worker {
	return &Worker{
		log: log.With().Str("component", "worker").Logger(),
	}
}

func (w *Worker) Execute(ctx context.Context, pipeline *analysis.Pipeline) *analysis.Axis {
	var axis *analysis.Axis
	var sequence []*analysis.AnalysisParameters

	for i, step := range pipeline.Steps {

		if i == 0 {
			axis = step.AxisPipeline.Source
		}
		sequence = step.AxisPipeline.ParametersSequence

		for t, parameters := range sequence {
			locallog := w.log.With().Int("step_index", i).Int("sequence_index", t).Logger()
			locallog.Debug().Msg("executing algorithm")

			outAxis := Do(ctx, parameters, axis)

			for _, series := range outAxis.Series {

				if step.Rename != nil {
					if newName, ok := step.Rename[series.Name]; ok {
						locallog.Debug().Str("old_name", series.Name).Str("new_name", newName).Msg("renaming series")
						series.Name = newName
					}
				}

				axis.PutSeries(series)

			}

		}

	}

	return axis
}
