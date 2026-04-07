package utils

import "strconv"

func LimitAndOffsetConverter(limitStr, offsetStr string) (int32, int32) {
	// set defaults
	limit := 10
	offset := 0

	if limitStr != "" {
		if v, err := strconv.Atoi(limitStr); err == nil {
			limit = v
		}
	}

	if offsetStr != "" {
		if v, err := strconv.Atoi(offsetStr); err == nil {
			offset = v
		}
	}

	return int32(limit), int32(offset)
}