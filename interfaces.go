package fes

type RecReader func() ([]string, error)
type RecWriter func([]string) error

//go:generate go run github.com/matryer/moq -fmt goimports -pkg mocks -out mocks/rawridefetcher_mock.go . RawRideFetcher:RawRideFetcherMock
type RawRideFetcher interface {
	FetchRawRide() (RawRide, bool, error)
}

//go:generate go run github.com/matryer/moq -fmt goimports -pkg mocks -out mocks/fareestimatewriter_mock.go . FareEstimateWriter:FareEstimateWriterMock
type FareEstimateWriter interface {
	WriteFareEstimate(id int, estimate float64) error
}
