package utils

import (
	"errors"
	"fmt"
	"github.com/metamatex/metamate/gen/v0/mql"
	"math"
	"time"
)

func FromUnixTimeStamp(ts mql.UnixTimestamp) (t time.Time, err error) {
	if ts.Value == nil {
		err = errors.New("UnixTimestamp.value is nil")

		return
	}

	d, err := FromDurationScalar(*ts.Value)
	if err != nil {
	    return
	}

	t = time.Unix(int64(d/time.Second), 0)

	return
}

func FromDurationScalar(s mql.DurationScalar) (d time.Duration, err error) {
	if s.Value == nil {
		err = errors.New("DurationScalar.value is nil")

		return
	}

	if s.Unit == nil {
		err = errors.New("DurationScalar.unit is nil")

		return
	}

	switch *s.Unit {
	case mql.DurationUnit.Ns:
		d = time.Duration(int64(math.Round(*s.Value))) * time.Nanosecond
	case mql.DurationUnit.Ms:
		d = time.Duration(int64(math.Round(*s.Value * 10e6))) * time.Nanosecond
	case mql.DurationUnit.S:
		d = time.Duration(int64(math.Round(*s.Value * 10e9))) * time.Nanosecond
	case mql.DurationUnit.M:
		d = time.Duration(int64(math.Round(*s.Value * 60 * 10e9))) * time.Nanosecond
	case mql.DurationUnit.H:
		d = time.Duration(int64(math.Round(*s.Value * 60 * 60 * 10e9))) * time.Nanosecond
	case mql.DurationUnit.D:
		d = time.Duration(int64(math.Round(*s.Value * 24 * 60 * 60 * 10e9))) * time.Nanosecond
	case mql.DurationUnit.W:
		d = time.Duration(int64(math.Round(*s.Value * 7 * 24 * 60 * 60 * 10e9))) * time.Nanosecond
	case mql.DurationUnit.Y:
		d = time.Duration(int64(math.Round(*s.Value * 365 * 24 * 60 * 60 * 10e9))) * time.Nanosecond
	default:
		err = errors.New(fmt.Sprintf("can't handle duration unit %v", *s.Unit))

		return
	}

	return
}
