package foundation

// SliceContainsStr return target slice contains s or not
func SliceContainsStr(s string, target []string) bool {
	for _, item := range target {
		if s == item {
			return true
		}
	}
	return false
}
