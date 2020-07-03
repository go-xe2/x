package xqi

type QueryPageInfo interface {
	PageIndex() int
	PageSize() int
	PageCount() int
	Total() int
}
