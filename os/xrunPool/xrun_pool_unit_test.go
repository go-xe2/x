package xrunPool

import (
	"github.com/go-xe2/x/container/xarray"
	"github.com/go-xe2/x/xtest"
	"sync"
	"testing"
	"time"
)

func Test_Basic(t *testing.T) {
	xtest.Case(t, func() {
		wg := sync.WaitGroup{}
		array := xarray.NewArray()
		size := 100
		wg.Add(size)
		for i := 0; i < size; i++ {
			Add(func() {
				array.Append(1)
				wg.Done()
			})
		}
		wg.Wait()
		time.Sleep(100 * time.Millisecond)
		xtest.Assert(array.Len(), size)
		xtest.Assert(Jobs(), 0)
		xtest.Assert(Size(), 0)
	})
}

func Test_Limit1(t *testing.T) {
	xtest.Case(t, func() {
		wg := sync.WaitGroup{}
		array := xarray.NewArray()
		size := 100
		pool := New(10)
		wg.Add(size)
		for i := 0; i < size; i++ {
			pool.Add(func() {
				array.Append(1)
				wg.Done()
			})
		}
		wg.Wait()
		xtest.Assert(array.Len(), size)
	})
}

func Test_Limit2(t *testing.T) {
	xtest.Case(t, func() {
		wg := sync.WaitGroup{}
		array := xarray.NewArray()
		size := 100
		pool := New(1)
		wg.Add(size)
		for i := 0; i < size; i++ {
			pool.Add(func() {
				array.Append(1)
				wg.Done()
			})
		}
		wg.Wait()
		xtest.Assert(array.Len(), size)
	})
}

func Test_Limit3(t *testing.T) {
	xtest.Case(t, func() {
		array := xarray.NewArray()
		size := 1000
		pool := New(100)
		xtest.Assert(pool.Cap(), 100)
		for i := 0; i < size; i++ {
			pool.Add(func() {
				array.Append(1)
				time.Sleep(2 * time.Second)
			})
		}
		time.Sleep(time.Second)
		xtest.Assert(pool.Size(), 100)
		xtest.Assert(pool.Jobs(), 900)
		xtest.Assert(array.Len(), 100)
		pool.Close()
		time.Sleep(2 * time.Second)
		xtest.Assert(pool.Size(), 0)
		xtest.Assert(pool.Jobs(), 900)
		xtest.Assert(array.Len(), 100)
		xtest.Assert(pool.IsClosed(), true)
		xtest.AssertNE(pool.Add(func() {}), nil)

	})
}
