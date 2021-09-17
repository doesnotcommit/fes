package fes

import (
	"fmt"
	"strconv"
)

type RawRide struct {
	ID     string
	Points []RawPoint
}

type segmentValidator func(s Segment) bool

func (r RawRide) BuildRide() (Ride, error) {
	handleErr := func(err error) (Ride, error) {
		return Ride{}, fmt.Errorf("build ride: %w", err)
	}
	if len(r.Points) < 2 {
		return handleErr(ErrInvalidRide)
	}
	var segments []Segment
	prevPoint, err := r.Points[0].Parse()
	if err != nil {
		return handleErr(ErrInvalidRide)
	}
	for i := 1; i < len(r.Points); i++ {
		point, err := r.Points[i].Parse()
		if err != nil {
			return handleErr(ErrInvalidRide)
		}
		seg := NewSegment(prevPoint, point)
		if !seg.isValid() {
			continue
		}
		segments = append(segments, seg)
		prevPoint = point
	}
	if len(segments) == 0 {
		return Ride{}, ErrInvalidRide
	}
	id, err := strconv.Atoi(r.ID)
	if err != nil {
		return handleErr(err)
	}
	return Ride{
		ID:       id,
		Segments: segments,
	}, nil
}
