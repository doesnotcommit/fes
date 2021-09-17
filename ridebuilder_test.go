package fes

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BuildRideHappy(t *testing.T) {
	rr := RawRide{
		ID: "1",
		Points: []RawPoint{
			{Lat: "37.966660", Lng: "23.728308", Timestamp: "1405594957"},
			{Lat: "37.966627", Lng: "23.728263", Timestamp: "1405594966"},
			{Lat: "37.966660", Lng: "23.728308", Timestamp: "1405594957"},
			{Lat: "37.966627", Lng: "23.728263", Timestamp: "1405594966"},
			{Lat: "37.966625", Lng: "23.728263", Timestamp: "1405594974"},
			{Lat: "37.966613", Lng: "23.728375", Timestamp: "1405594984"},
			{Lat: "37.966203", Lng: "23.728597", Timestamp: "1405594992"},
			{Lat: "37.966195", Lng: "23.728613", Timestamp: "1405595001"},
			{Lat: "37.966195", Lng: "23.728613", Timestamp: "1405595009"},
			{Lat: "37.966195", Lng: "23.728613", Timestamp: "1405595017"},
		},
	}
	ride, err := rr.BuildRide()
	if err != nil {
		t.Error(err)
	}
	wantRide := Ride{
		ID: 1,
		Segments: []Segment{
			{
				Distance: 0.005387608950152276,
				Speed:    2.15504358006091,
				Points: [2]Point{
					{
						Lat:       37.96666,
						Lng:       23.728308,
						Timestamp: 1405594957,
					},
					{
						Lat:       37.966627,
						Lng:       23.728263,
						Timestamp: 1405594966,
					},
				},
			},
			{
				Distance: 0.005387608950152276,
				Speed:    2.15504358006091,
				Points: [2]Point{
					{
						Lat:       37.966627,
						Lng:       23.728263,
						Timestamp: 1405594966,
					},
					{
						Lat:       37.96666,
						Lng:       23.728308,
						Timestamp: 1405594957,
					},
				},
			},
			{
				Distance: 0.005387608950152276,
				Speed:    2.15504358006091,
				Points: [2]Point{
					{
						Lat:       37.96666,
						Lng:       23.728308,
						Timestamp: 1405594957,
					},
					{
						Lat:       37.966627,
						Lng:       23.728263,
						Timestamp: 1405594966,
					},
				},
			},
			{
				Distance: 0.0002223898541684477,
				Speed:    0.10007543437580146,
				Points: [2]Point{
					{
						Lat:       37.966627,
						Lng:       23.728263,
						Timestamp: 1405594966,
					}, {
						Lat:       37.966625,
						Lng:       23.728263,
						Timestamp: 1405594974,
					},
				},
			},
			{
				Distance: 0.009908475188669013,
				Speed:    3.5670510679208447,
				Points: [2]Point{
					{
						Lat:       37.966625,
						Lng:       23.728263,
						Timestamp: 1405594974,
					}, {
						Lat:       37.966613,
						Lng:       23.728375,
						Timestamp: 1405594984,
					},
				},
			},
			{
				Distance: 0.049569929021084466,
				Speed:    22.306468059488008,
				Points: [2]Point{
					{
						Lat:       37.966613,
						Lng:       23.728375,
						Timestamp: 1405594984,
					}, {
						Lat:       37.966203,
						Lng:       23.728597,
						Timestamp: 1405594992,
					},
				},
			},
			{
				Distance: 0.0016609132234105446,
				Speed:    0.6643652893642179,
				Points: [2]Point{
					{
						Lat:       37.966203,
						Lng:       23.728597,
						Timestamp: 1405594992,
					}, {
						Lat:       37.966195,
						Lng:       23.728613,
						Timestamp: 1405595001,
					},
				},
			},
			{
				Distance: 0,
				Speed:    0,
				Points: [2]Point{
					{
						Lat:       37.966195,
						Lng:       23.728613,
						Timestamp: 1405595001,
					}, {
						Lat:       37.966195,
						Lng:       23.728613,
						Timestamp: 1405595009,
					},
				},
			},
			{
				Distance: 0,
				Speed:    0,
				Points: [2]Point{
					{
						Lat:       37.966195,
						Lng:       23.728613,
						Timestamp: 1405595009,
					}, {
						Lat:       37.966195,
						Lng:       23.728613,
						Timestamp: 1405595017,
					},
				},
			},
		},
	}
	assert.Equal(t, wantRide, ride, "unexpected ride")
}

func Test_BuildRideWithInvalidSegment(t *testing.T) {
	rr := RawRide{
		ID: "1",
		Points: []RawPoint{
			{Lat: "37.966660", Lng: "23.728308", Timestamp: "1405594957"},
			{Lat: "37.966627", Lng: "23.728263", Timestamp: "1405594966"},
			{Lat: "37.966660", Lng: "23.728308", Timestamp: "1405594957"},
			{Lat: "37.966627", Lng: "23.728263", Timestamp: "1405594966"},
			{Lat: "37.966625", Lng: "23.728263", Timestamp: "1405594974"},
			{Lat: "37.966613", Lng: "23.728375", Timestamp: "1405594984"},
			{Lat: "37.966203", Lng: "23.728597", Timestamp: "1405594992"},
			{Lat: "37.966195", Lng: "23.728613", Timestamp: "1405595001"},
			{Lat: "37.966195", Lng: "23.728613", Timestamp: "1405595009"},
			{Lat: "37.976195", Lng: "23.738613", Timestamp: "1405595017"}, // bad point here
		},
	}
	ride, err := rr.BuildRide()
	if err != nil {
		t.Error(err)
	}
	wantRide := Ride{
		ID: 1,
		Segments: []Segment{
			{
				Distance: 0.005387608950152276,
				Speed:    2.15504358006091,
				Points: [2]Point{
					{
						Lat:       37.96666,
						Lng:       23.728308,
						Timestamp: 1405594957,
					},
					{
						Lat:       37.966627,
						Lng:       23.728263,
						Timestamp: 1405594966,
					},
				},
			},
			{
				Distance: 0.005387608950152276,
				Speed:    2.15504358006091,
				Points: [2]Point{
					{
						Lat:       37.966627,
						Lng:       23.728263,
						Timestamp: 1405594966,
					},
					{
						Lat:       37.96666,
						Lng:       23.728308,
						Timestamp: 1405594957,
					},
				},
			},
			{
				Distance: 0.005387608950152276,
				Speed:    2.15504358006091,
				Points: [2]Point{
					{
						Lat:       37.96666,
						Lng:       23.728308,
						Timestamp: 1405594957,
					},
					{
						Lat:       37.966627,
						Lng:       23.728263,
						Timestamp: 1405594966,
					},
				},
			},
			{
				Distance: 0.0002223898541684477,
				Speed:    0.10007543437580146,
				Points: [2]Point{
					{
						Lat:       37.966627,
						Lng:       23.728263,
						Timestamp: 1405594966,
					}, {
						Lat:       37.966625,
						Lng:       23.728263,
						Timestamp: 1405594974,
					},
				},
			},
			{
				Distance: 0.009908475188669013,
				Speed:    3.5670510679208447,
				Points: [2]Point{
					{
						Lat:       37.966625,
						Lng:       23.728263,
						Timestamp: 1405594974,
					}, {
						Lat:       37.966613,
						Lng:       23.728375,
						Timestamp: 1405594984,
					},
				},
			},
			{
				Distance: 0.049569929021084466,
				Speed:    22.306468059488008,
				Points: [2]Point{
					{
						Lat:       37.966613,
						Lng:       23.728375,
						Timestamp: 1405594984,
					}, {
						Lat:       37.966203,
						Lng:       23.728597,
						Timestamp: 1405594992,
					},
				},
			},
			{
				Distance: 0.0016609132234105446,
				Speed:    0.6643652893642179,
				Points: [2]Point{
					{
						Lat:       37.966203,
						Lng:       23.728597,
						Timestamp: 1405594992,
					}, {
						Lat:       37.966195,
						Lng:       23.728613,
						Timestamp: 1405595001,
					},
				},
			},
			{
				Distance: 0,
				Speed:    0,
				Points: [2]Point{
					{
						Lat:       37.966195,
						Lng:       23.728613,
						Timestamp: 1405595001,
					}, {
						Lat:       37.966195,
						Lng:       23.728613,
						Timestamp: 1405595009,
					},
				},
			},
		},
	}
	assert.Equal(t, wantRide, ride, "unexpected ride")
}

func Test_BuildRideWithMissingPoints(t *testing.T) {
	rr1 := RawRide{
		ID: "1",
		Points: []RawPoint{
			{Lat: "37.966660", Lng: "23.728308", Timestamp: "1405594957"},
		},
	}
	if _, err := rr1.BuildRide(); !errors.Is(err, ErrInvalidRide) {
		t.Errorf("want err: %v, got err: %v", ErrInvalidRide, err)
	}
	rr2 := RawRide{
		ID:     "1",
		Points: []RawPoint{},
	}
	if _, err := rr2.BuildRide(); !errors.Is(err, ErrInvalidRide) {
		t.Errorf("want err: %v, got err: %v", ErrInvalidRide, err)
	}
}
