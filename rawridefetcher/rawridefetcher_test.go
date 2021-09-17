package rawridefetcher

import (
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"fes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
)

var samplePath = filepath.Join("..", "sample", "paths.csv.gz")

func Test_Fetcher(t *testing.T) {
	wantRawReader, err := os.Open(samplePath)
	if err != nil {
		t.Fatal(err)
	}
	defer wantRawReader.Close()
	wantBytesReader, err := gzip.NewReader(wantRawReader)
	if err != nil {
		t.Fatal(err)
	}
	wantBytes, err := io.ReadAll(wantBytesReader)
	if err != nil {
		t.Fatal(err)
	}
	sampleReadRec, cleanup, err := newCSVRecReader(samplePath)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()
	f, err := New(sampleReadRec)
	var rr []fes.RawRide
	for {
		r, more, err := f.FetchRawRide()
		if err != nil {
			t.Fatal(err)
		}
		rr = append(rr, r)
		if !more {
			break
		}
	}
	gotBuf := bytes.Buffer{}
	for _, r := range rr {
		for _, p := range r.Points {
			row := fmt.Sprintf("%s,%s,%s,%s\n", r.ID, p.Lat, p.Lng, p.Timestamp)
			gotBuf.WriteString(row)
		}
	}
	if bytes.Compare(wantBytes, gotBuf.Bytes()) != 0 {
		t.Fatal("raw ride fetcher is broken")
	}
}

func newCSVRecReader(path string) (fes.RecReader, func(), error) {
	handleErr := func(err error) (fes.RecReader, func(), error) {
		return nil, nil, fmt.Errorf("create csv rec reader: %w", err)
	}
	f, err := os.Open(path)
	if err != nil {
		return handleErr(err)
	}
	gzf, err := gzip.NewReader(f)
	if err != nil {
		_ = f.Close()
		return handleErr(err)
	}
	cleanup := func() {
		_ = gzf.Close()
		_ = f.Close()
	}
	r := csv.NewReader(gzf)
	if err != nil {
		cleanup()
		return handleErr(err)
	}
	return r.Read, cleanup, nil
}
