package fareestimator

import (
	"context"
	"fes"
	"fmt"

	"golang.org/x/sync/errgroup"
)

type Estimator struct {
	rf         fes.RawRideFetcher
	few        fes.FareEstimateWriter
	maxWorkers int
}

func New(rf fes.RawRideFetcher, few fes.FareEstimateWriter, maxWorkers int) Estimator {
	return Estimator{
		rf:         rf,
		few:        few,
		maxWorkers: maxWorkers,
	}
}

func (e Estimator) GetFareEstimates(ctx context.Context) error {
	handleErr := func(err error) error {
		return fmt.Errorf("get fare estimates: %w", err)
	}
	sem := make(chan struct{}, e.maxWorkers)
	g, ctx := errgroup.WithContext(ctx)
	for {
		rawRide, more, err := e.rf.FetchRawRide()
		if err != nil {
			return handleErr(err)
		}
		sem <- struct{}{}
		g.Go(func() error {
			defer func() {
				<-sem
			}()
			ride, err := rawRide.BuildRide()
			if err != nil {
				return err
			}
			if err := e.few.WriteFareEstimate(ride.ID, ride.CalcFare()); err != nil {
				return err
			}
			return nil
		})
		if !more {
			break
		}
	}
	return g.Wait()
}
