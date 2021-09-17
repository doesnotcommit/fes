package fareestimatewriter

import (
	"fes"
	"strconv"
)

type Writer struct {
	writeRec fes.RecWriter
}

func New(writeRec fes.RecWriter) Writer {
	return Writer{
		writeRec: writeRec,
	}
}

func (w Writer) WriteFareEstimate(id int, est float64) error {
	return w.writeRec([]string{
		strconv.Itoa(id),
		strconv.FormatFloat(est, 'f', 2, 64),
	})
}
