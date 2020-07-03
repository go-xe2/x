package xquery

import . "github.com/go-xe2/x/xf/ef/xqi"

func (qe *tQueryExp) Limit(rows int, offset ...int) Query {
	qe.checkMainTableSet()
	n := 0
	if len(offset) > 0 {
		n = offset[0]
	}
	qe.com.Limit(rows, n)
	return qe
}

func (qe *tQueryExp) Page(size int, index ...int) Query {
	qe.checkMainTableSet()
	n := 0
	if len(index) > 0 {
		n = index[0]
	}
	qe.com.Page(size, n)
	return qe
}
