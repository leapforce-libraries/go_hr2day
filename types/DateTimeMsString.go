package hr2day

import (
	"encoding/json"
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
)

const (
	dateTimeMSFormat string = "2006-01-02T15:04:05.999+0000"
)

type DateTimeMsString time.Time

func (d *DateTimeMsString) UnmarshalJSON(b []byte) error {
	var returnError = func() error {
		errortools.CaptureError(fmt.Sprintf("Cannot parse '%s' to DateTimeMsString", string(b)))
		return nil
	}

	var s string

	err := json.Unmarshal(b, &s)
	if err != nil {
		fmt.Println("DateTimeMsString", string(b))
		return returnError()
	}

	if s == "" {
		d = nil
		return nil
	}

	_t, err := time.Parse(dateTimeMSFormat, s)
	if err != nil {
		return returnError()
	}

	*d = DateTimeMsString(_t)
	return nil
}

func (d *DateTimeMsString) ValuePtr() *time.Time {
	if d == nil {
		return nil
	}

	_d := time.Time(*d)
	return &_d
}

func (d DateTimeMsString) Value() time.Time {
	return time.Time(d)
}
