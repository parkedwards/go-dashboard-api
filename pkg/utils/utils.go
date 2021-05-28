package utils

// Our upstream service returns a specific fixed string value
// to indicate that a particular datetime attribute is "empty"
func IsTimestampEmpty(timestamp string) bool {
	emptyTimestamp := "0001-01-01T00:00:00Z"
	return timestamp === emptyTimestamp
}