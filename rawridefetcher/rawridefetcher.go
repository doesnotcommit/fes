package rawridefetcher

import (
	"errors"
	"fes"
	"fmt"
	"io"
)

type Fetcher struct {
	readRec fes.RecReader
	buf     chan fes.RawRide
}

func New(readRec fes.RecReader) (Fetcher, error) {
	handleErr := func(err error) (Fetcher, error) {
		return Fetcher{}, fmt.Errorf("new ride fetcher: %w", err)
	}
	rec, err := readRec()
	if err != nil {
		return handleErr(err)
	}
	if len(rec) < 4 {
		return handleErr(fes.ErrInvalidInput)
	}
	id, point := getRawPoint(rec)
	buf := make(chan fes.RawRide, 1)
	buf <- fes.RawRide{
		ID:     id,
		Points: []fes.RawPoint{point},
	}
	return Fetcher{
		readRec: readRec,
		buf:     buf,
	}, nil
}

func (f Fetcher) FetchRawRide() (fes.RawRide, bool, error) {
	r := <-f.buf
	for {
		rec, err := f.readRec()
		if errors.Is(err, io.EOF) {
			return r, false, nil
		} else if err != nil {
			return r, false, err
		}
		id, point := getRawPoint(rec)
		if id == r.ID {
			r.Points = append(r.Points, point)
			continue
		}
		f.buf <- fes.RawRide{
			ID:     id,
			Points: []fes.RawPoint{point},
		}
		return r, true, nil
	}
}

func getRawPoint(rec []string) (string, fes.RawPoint) {
	return rec[csvCols["id_ride"]], fes.RawPoint{
		Lat:       rec[csvCols["lat"]],
		Lng:       rec[csvCols["lng"]],
		Timestamp: rec[csvCols["timestamp"]],
	}
}

var csvCols = map[string]int{
	"id_ride":   0,
	"lat":       1,
	"lng":       2,
	"timestamp": 3,
}
