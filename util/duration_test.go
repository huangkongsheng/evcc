package util

import (
	"math"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseDuration(t *testing.T) {
	maxSeconds := math.MaxInt64 / int64(time.Second)

	for _, tc := range []struct {
		name    string
		in      string
		want    time.Duration
		wantErr string
	}{
		{name: "zero", in: "0", want: 0},
		{name: "seconds", in: "30", want: 30 * time.Second},
		{name: "hour", in: "3600", want: time.Hour},
		{name: "max", in: strconv.FormatInt(maxSeconds, 10), want: time.Duration(maxSeconds) * time.Second},
		{name: "negative", in: "-1", wantErr: "invalid duration: -1"},
		{name: "overflow", in: strconv.FormatInt(maxSeconds + 1, 10), wantErr: "duration overflow: " + strconv.FormatInt(maxSeconds+1, 10)},
		{name: "empty", in: "", wantErr: "strconv.Atoi: parsing \"\": invalid syntax"},
		{name: "not a number", in: "1s", wantErr: "strconv.Atoi: parsing \"1s\": invalid syntax"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseDuration(tc.in)
			if tc.wantErr != "" {
				assert.EqualError(t, err, tc.wantErr)
				assert.Equal(t, time.Duration(0), got)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
