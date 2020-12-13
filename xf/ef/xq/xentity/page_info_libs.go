/*****************************************************************
* Copyright©,2020-2022, email: 279197148@qq.com
* Version: 1.0.0
* @Author: yangtxiang
* @Date: 2020-09-14 10:33
* Description:
*****************************************************************/

package xentity

var minPageIndex int = 0
var minPageSize int = 1
var maxPageSize int = 10

func SetMinPageIndex(pi int) {
	minPageIndex = pi
}

func SetMinPageSize(size int) {
	minPageSize = size
}

func SetMaxPageSize(size int) {
	maxPageSize = size
}

func GetMinPageIndex() int {
	return minPageIndex
}

func GetMinPageSize() int {
	return minPageSize
}

func GetMaxPageSize() int {
	return maxPageSize
}

func CheckPageIndex(pi *int) {
	if pi == nil {
		pi = &minPageIndex
	}
	if *pi < minPageIndex {
		*pi = minPageIndex
	}
}

func CheckPageSize(size *int) {
	if size == nil {
		size = &minPageSize
	}
	if *size < minPageSize {
		*size = minPageSize
	}
	if *size > maxPageSize {
		*size = maxPageSize
	}
}
