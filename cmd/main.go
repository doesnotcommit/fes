package main

import (
	"context"
	"encoding/csv"
	"fes/fareestimatewriter"
	"fes/fareestimator"
	"fes/rawridefetcher"
	"io"
	"os"
)

const maxWorkers = 500

func main() {
	defer os.Stdin.Close()
	defer os.Stdout.Close()
	if err := run(maxWorkers, os.Stdout, os.Stdin); err != nil {
		panic(err)
	}
}

func run(maxWorkers int, w io.Writer, r io.Reader) error {
	csvReader := csv.NewReader(r)
	rrf, err := rawridefetcher.New(csvReader.Read)
	if err != nil {
		return err
	}
	csvWriter := csv.NewWriter(w)
	defer csvWriter.Flush()
	few := fareestimatewriter.New(csvWriter.Write)
	fe := fareestimator.New(rrf, few, maxWorkers)
	ctx := context.Background()
	if err := fe.GetFareEstimates(ctx); err != nil {
		return err
	}
	return nil
}
