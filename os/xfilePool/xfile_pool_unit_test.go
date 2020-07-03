package xfilePool

import (
	"fmt"
	"github.com/go-xe2/x/os/xfile"
	"github.com/go-xe2/x/xtest"
	"os"
	"testing"
	"time"
)

func TestOpen(t *testing.T) {
	testFile := start("filePoolTest.txt")
	xtest.Case(t, func() {
		f, err := Open(testFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		xtest.AssertEq(err, nil)
		f.Close()

		f2, err1 := Open(testFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		xtest.AssertEq(err1, nil)
		xtest.AssertEq(f, f2)
		f2.Close()

		// Deprecated test
		f3, err2 := OpenFile(testFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		xtest.AssertEq(err2, nil)
		xtest.AssertEq(f, f3)
		f3.Close()

	})

	stop(testFile)
}

func TestOpenErr(t *testing.T) {
	xtest.Case(t, func() {
		testErrFile := "errorPath"
		_, err := Open(testErrFile, os.O_RDWR, 0666)
		xtest.AssertNE(err, nil)

		// delete file error
		testFile := start("TestOpenDeleteErr.txt")
		pool := New(testFile, os.O_RDWR, 0666)
		_, err1 := pool.File()
		xtest.AssertEq(err1, nil)
		stop(testFile)
		_, err1 = pool.File()
		xtest.AssertNE(err1, nil)

		// append mode delete file error and create again
		testFile = start("TestOpenCreateErr.txt")
		pool = New(testFile, os.O_CREATE, 0666)
		_, err1 = pool.File()
		xtest.AssertEq(err1, nil)
		stop(testFile)
		_, err1 = pool.File()
		xtest.AssertEq(err1, nil)

		// append mode delete file error
		testFile = start("TestOpenAppendErr.txt")
		pool = New(testFile, os.O_APPEND, 0666)
		_, err1 = pool.File()
		xtest.AssertEq(err1, nil)
		stop(testFile)
		_, err1 = pool.File()
		xtest.AssertNE(err1, nil)

		// trunc mode delete file error
		testFile = start("TestOpenTruncErr.txt")
		pool = New(testFile, os.O_TRUNC, 0666)
		_, err1 = pool.File()
		xtest.AssertEq(err1, nil)
		stop(testFile)
		_, err1 = pool.File()
		xtest.AssertNE(err1, nil)
	})
}

func TestOpenExpire(t *testing.T) {
	testFile := start("TestOpenExpire.txt")

	xtest.Case(t, func() {
		f, err := Open(testFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666, 100)
		xtest.AssertEq(err, nil)
		f.Close()

		time.Sleep(150 * time.Millisecond)
		f2, err1 := Open(testFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666, 100)
		xtest.AssertEq(err1, nil)
		//xtest.AssertNE(f, f2)
		f2.Close()

		// Deprecated test
		f3, err2 := Open(testFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666, 100)
		xtest.AssertEq(err2, nil)
		xtest.AssertEq(f2, f3)
		f3.Close()
	})

	stop(testFile)
}

func TestNewPool(t *testing.T) {
	testFile := start("TestNewPool.txt")

	xtest.Case(t, func() {
		f, err := Open(testFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		xtest.AssertEq(err, nil)
		f.Close()

		pool := New(testFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		f2, err1 := pool.File()
		// pool not equal
		xtest.AssertEq(err1, nil)
		//xtest.AssertNE(f, f2)
		f2.Close()

		pool.Close()
	})

	stop(testFile)
}

// test before
func start(name string) string {
	testFile := os.TempDir() + string(os.PathSeparator) + name
	if xfile.Exists(testFile) {
		xfile.Remove(testFile)
	}
	content := "123"
	xfile.PutContents(testFile, content)
	return testFile
}

// test after
func stop(testFile string) {
	if xfile.Exists(testFile) {
		err := xfile.Remove(testFile)
		if err != nil {
			fmt.Println(err)
		}
	}
}
