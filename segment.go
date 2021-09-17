package fes

import "time"

const maxValidSpeedKPH = 100

func NewSegment(prevPoint, point Point) Segment {
	dist, speed := calcKMAndKPH(prevPoint, point)
	return Segment{
		Distance: dist,
		Speed:    speed,
		Points: [2]Point{
			prevPoint,
			point,
		},
	}
}

type Segment struct {
	Distance float64 // km
	Speed    float64 // kph
	Points   [2]Point
}

func (s Segment) isValid() bool {
	return s.Speed <= maxValidSpeedKPH
}

func (s Segment) calcFare() float64 {
	if s.Speed <= maxIdleSpeedKPH {
		return float64(s.Points[1].Timestamp-s.Points[0].Timestamp) / secondsInHour * baseIdleFarePerHour
	}
	h := time.Unix(int64(s.Points[0].Timestamp), 0).Hour()
	if h >= 0 && h <= 5 {
		return s.Distance * baseMovingFarePerKMNight
	}
	return s.Distance * baseMovingFarePerKMDay
}
