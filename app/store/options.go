package store

type ListOption interface {
	Apply(*listSettings)
}

// MaxResults
func WithMaxResults(maxResults int) ListOption {
	return withMaxResults(maxResults)
}

type withMaxResults int

func (w withMaxResults) Apply(ls *listSettings) {
	ls.maxResults = int(w)
}

// Offset
func WithOffset(offset int) ListOption {
	return withOffset(offset)
}

type withOffset int

func (w withOffset) Apply(ls *listSettings) {
	ls.offset = int(w)
}

// Filters
func WithFilters(filters []string) ListOption {
	return withFilters(filters)
}

type withFilters []string

func (w withFilters) Apply(ls *listSettings) {
	ls.filters = make([]string, len(w))
	copy(ls.filters, w)
}

// Live
func WithLive(live bool) ListOption {
	return withLive(live)
}

type withLive bool

func (w withLive) Apply(ls *listSettings) {
	ls.live = bool(w)
}

// Gaming
func WithGaming(gaming bool) ListOption {
	return withGaming(gaming)
}

type withGaming bool

func (w withGaming) Apply(ls *listSettings) {
	ls.gaming = bool(w)
}
