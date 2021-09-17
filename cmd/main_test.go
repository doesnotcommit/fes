package main

import (
	"bytes"
	"compress/gzip"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var samplePath = filepath.Join("..", "sample", "paths.csv.gz")

func Test_E2E(t *testing.T) {
	sf, err := os.Open(samplePath)
	if err != nil {
		t.Fatal(err)
	}
	defer sf.Close()
	gzsf, err := gzip.NewReader(sf)
	if err != nil {
		t.Fatal(err)
	}
	defer gzsf.Close()
	buf := bytes.Buffer{}
	const maxWorkers = 1
	if err := run(maxWorkers, &buf, gzsf); err != nil {
		t.Fatal(err)
	}
	wantStr := "1,11.34\n2,13.10\n3,33.84\n4,3.47\n5,22.78\n6,9.41\n7,30.01\n8,9.21\n9,6.35\n"
	assert.Equal(t, wantStr, buf.String())
}
