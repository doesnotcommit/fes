package fes

import (
	"fmt"
	"strconv"
)

type Point struct {
	Lat       float64
	Lng       float64
	Timestamp int
}

type RawPoint struct {
	Lat       string
	Lng       string
	Timestamp string
}

func (rp RawPoint) Parse() (Point, error) {
	lat, err := strconv.ParseFloat(rp.Lat, 64)
	if err != nil {
		return Point{}, fmt.Errorf("parse lat: %q: %w", rp.Lat, err)
	}
	lng, err := strconv.ParseFloat(rp.Lng, 64)
	if err != nil {
		return Point{}, fmt.Errorf("parse lng: %q: %w", rp.Lng, err)
	}
	ts, err := strconv.Atoi(rp.Timestamp)
	if err != nil {
		return Point{}, fmt.Errorf("parse timestamp: %q: %w", rp.Timestamp, err)
	}
	return Point{
		Lat:       lat,
		Lng:       lng,
		Timestamp: ts,
	}, nil
}
