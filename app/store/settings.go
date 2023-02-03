package store

type listSettings struct {
	maxResults int
	offset     int
	filters    []string
	live       bool
	gaming     bool
}

func (ls *listSettings) GetMaxResults() int {
	if ls.maxResults > 0 {
		return ls.maxResults
	}
	return 20
}

func (ls *listSettings) GetOffset() int {
	if ls.offset > 0 {
		return ls.offset
	}
	return 0
}

func (ls *listSettings) GetFilters() []string {
	return ls.filters
}

func (ls *listSettings) GetLive() bool {
	return ls.live
}

func (ls *listSettings) GetGaming() bool {
	return ls.gaming
}
