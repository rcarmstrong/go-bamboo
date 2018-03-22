package bamboo

func emptyStrings(strings ...string) bool {
	for _, s := range strings {
		if s == "" {
			return true
		}
	}
	return false
}

// Pagination used to specify the start and limit indexes of a paginated API resource
type Pagination struct {
	Start int
	Limit int
}
