package util

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func ParseDuration(s string) (time.Duration, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	// Compare in int64 so this compiles on 32-bit platforms, where the
	// max-seconds constant does not fit in an int.
	secs := int64(v)
	if secs < 0 {
		return 0, fmt.Errorf("invalid duration: %d", v)
	}

	if secs > math.MaxInt64/int64(time.Second) {
		return 0, fmt.Errorf("duration overflow: %d", v)
	}

	return time.Duration(secs) * time.Second, nil
}
