package common

func IsOutTime(start, end int64, min int64) bool {
	if end-start >= min*60 {
		return true
	}
	return false
}
