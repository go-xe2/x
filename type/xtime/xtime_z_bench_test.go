// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package xtime_test

import (
	"github.com/go-xe2/x/type/xtime"
	"testing"
)

func Benchmark_Second(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xtime.Second()
	}
}

func Benchmark_Millisecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xtime.Millisecond()
	}
}

func Benchmark_Microsecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xtime.Microsecond()
	}
}

func Benchmark_Nanosecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xtime.Nanosecond()
	}
}

func Benchmark_StrToTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xtime.StrToTime("2018-02-09T20:46:17.897Z")
	}
}

func Benchmark_ParseTimeFromContent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xtime.ParseTimeFromContent("2018-02-09T20:46:17.897Z")
	}
}

func Benchmark_NewFromTimeStamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xtime.NewFromTimeStamp(1542674930)
	}
}

func Benchmark_Date(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xtime.Date()
	}
}

func Benchmark_Datetime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xtime.Datetime()
	}
}

func Benchmark_SetTimeZone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xtime.SetTimeZone("Asia/Shanghai")
	}
}

func TestTime_UnmarshalJSON(t *testing.T) {
	s := "\"2020-06-12 12:24:30\""
	t1 := xtime.New()
	err := t1.UnmarshalJSON([]byte(s))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("t1:", t1)
}
