package xentity

import "github.com/go-xe2/x/xf/ef/xqi"

// 数据分页数据结构
type tQueryPageInfo struct {
	page  int
	size  int
	count int
	total int
}

func NewQueryPageInfo(pageIndex, pageSize, pageCount, total int) xqi.QueryPageInfo {
	return &tQueryPageInfo{
		page:  pageIndex,
		size:  pageSize,
		count: pageCount,
		total: total,
	}
}

func (qpi *tQueryPageInfo) PageIndex() int {
	return qpi.page
}

func (qpi *tQueryPageInfo) PageSize() int {
	return qpi.size
}

func (qpi *tQueryPageInfo) PageCount() int {
	return qpi.count
}

func (qpi *tQueryPageInfo) Total() int {
	return qpi.total
}
