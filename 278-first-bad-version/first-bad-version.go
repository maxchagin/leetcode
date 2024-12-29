package firstbadversion

// Global variable for mocking isBadVersion
var badVersion int

// Mock implementation of isBadVersion API
func isBadVersion(version int) bool {
	return version >= badVersion
}

func firstBadVersion(n int) int {

	return 0
}
