package calculator

import (
	"fmt"
	"strconv"
)

func Add(x, y string) (int64, error) {
	xInt, err := strconv.ParseInt(x, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("converting x failed: %v", err)
	}

	yInt, err := strconv.ParseInt(y, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("converting y failed: %v", err)
	}

	return xInt + yInt, nil
}
