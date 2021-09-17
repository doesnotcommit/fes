package fes

const maxIdleSpeedKPH = 10

const (
	flagAmount               = 1.30
	minFare                  = 3.47
	baseIdleFarePerHour      = 11.90
	baseMovingFarePerKMDay   = 0.74
	baseMovingFarePerKMNight = 1.30
)

type Ride struct {
	ID       int
	Segments []Segment
}

func (r Ride) CalcFare() float64 {
	fare := flagAmount
	for _, seg := range r.Segments {
		fare += seg.calcFare()
	}
	if fare < minFare {
		return minFare
	}
	return fare
}
