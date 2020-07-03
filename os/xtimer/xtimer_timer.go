package xtimer

import (
	_type "github.com/go-xe2/x/sync/type"
	"github.com/go-xe2/x/sync/xsafeStack"
	"time"
)

// 定时器/分层时间轮
type TTimer struct {
	status     *_type.TInt // 定时器状态
	wheels     []*twheel   // 分层时间轮对象
	length     int         // 分层层数
	number     int         // 每一层Slot Number
	intervalMs int64       // 最小时间刻度(毫秒)
}

// 创建分层时间轮
func New(slot int, interval time.Duration, level ...int) *TTimer {
	length := mDEFAULT_WHEEL_LEVEL
	if len(level) > 0 {
		length = level[0]
	}
	t := &TTimer{
		status:     _type.NewInt(STATUS_RUNNING),
		wheels:     make([]*twheel, length),
		length:     length,
		number:     slot,
		intervalMs: interval.Nanoseconds() / 1e6,
	}
	for i := 0; i < length; i++ {
		if i > 0 {
			n := time.Duration(t.wheels[i-1].totalMs) * time.Millisecond
			w := t.newWheel(i, slot, n)
			t.wheels[i] = w
			t.wheels[i-1].addEntry(n, w.proceed, false, mDEFAULT_TIMES, STATUS_READY)
		} else {
			t.wheels[i] = t.newWheel(i, slot, interval)
		}
	}
	t.wheels[0].start()
	return t
}

// 创建自定义的循环任务管理对象
func (t *TTimer) newWheel(level int, slot int, interval time.Duration) *twheel {
	w := &twheel{
		timer:      t,
		level:      level,
		slots:      make([]*xsafeStack.TSafeStackQe, slot),
		number:     int64(slot),
		ticks:      _type.NewInt64(),
		totalMs:    int64(slot) * interval.Nanoseconds() / 1e6,
		createMs:   time.Now().UnixNano() / 1e6,
		intervalMs: interval.Nanoseconds() / 1e6,
	}
	for i := int64(0); i < w.number; i++ {
		w.slots[i] = xsafeStack.New()
	}
	return w
}

// 添加循环任务
func (t *TTimer) Add(interval time.Duration, job JobFunc) *TEntry {
	return t.doAddEntry(interval, job, false, mDEFAULT_TIMES, STATUS_READY)
}

// 添加定时任务
func (t *TTimer) AddEntry(interval time.Duration, job JobFunc, singleton bool, times int, status int) *TEntry {
	return t.doAddEntry(interval, job, singleton, times, status)
}

// 添加单例运行循环任务
func (t *TTimer) AddSingleton(interval time.Duration, job JobFunc) *TEntry {
	return t.doAddEntry(interval, job, true, mDEFAULT_TIMES, STATUS_READY)
}

// 添加只运行一次的循环任务
func (t *TTimer) AddOnce(interval time.Duration, job JobFunc) *TEntry {
	return t.doAddEntry(interval, job, true, 1, STATUS_READY)
}

// 添加运行指定次数的循环任务。
func (t *TTimer) AddTimes(interval time.Duration, times int, job JobFunc) *TEntry {
	return t.doAddEntry(interval, job, true, times, STATUS_READY)
}

// 延迟添加循环任务。
func (t *TTimer) DelayAdd(delay time.Duration, interval time.Duration, job JobFunc) {
	t.AddOnce(delay, func() {
		t.Add(interval, job)
	})
}

// 延迟添加循环任务, 支持完整的参数。
func (t *TTimer) DelayAddEntry(delay time.Duration, interval time.Duration, job JobFunc, singleton bool, times int, status int) {
	t.AddOnce(delay, func() {
		t.AddEntry(interval, job, singleton, times, status)
	})
}

// 延迟添加单例循环任务
func (t *TTimer) DelayAddSingleton(delay time.Duration, interval time.Duration, job JobFunc) {
	t.AddOnce(delay, func() {
		t.AddSingleton(interval, job)
	})
}

// 延迟添加只运行一次的循环任务
func (t *TTimer) DelayAddOnce(delay time.Duration, interval time.Duration, job JobFunc) {
	t.AddOnce(delay, func() {
		t.AddOnce(interval, job)
	})
}

// 延迟添加只运行一次的循环任务
func (t *TTimer) DelayAddTimes(delay time.Duration, interval time.Duration, times int, job JobFunc) {
	t.AddOnce(delay, func() {
		t.AddTimes(interval, times, job)
	})
}

// 启动定时器
func (t *TTimer) Start() {
	t.status.Set(STATUS_RUNNING)
}

// 定制定时器
func (t *TTimer) Stop() {
	t.status.Set(STATUS_STOPPED)
}

// 关闭定时器
func (t *TTimer) Close() {
	t.status.Set(STATUS_CLOSED)
}

// 添加定时任务
func (t *TTimer) doAddEntry(interval time.Duration, job JobFunc, singleton bool, times int, status int) *TEntry {
	return t.wheels[t.getLevelByIntervalMs(interval.Nanoseconds()/1e6)].addEntry(interval, job, singleton, times, status)
}

// 添加定时任务，给定父级Entry, 间隔参数参数为毫秒数.
func (t *TTimer) doAddEntryByParent(interval int64, parent *TEntry) *TEntry {
	return t.wheels[t.getLevelByIntervalMs(interval)].addEntryByParent(interval, parent)
}

// 根据intervalMs计算添加的分层索引
func (t *TTimer) getLevelByIntervalMs(intervalMs int64) int {
	pos, cmp := t.binSearchIndex(intervalMs)
	switch cmp {
	// intervalMs与最后匹配值相等, 不添加到匹配得层，而是向下一层添加
	case 0:
		fallthrough
	// intervalMs比最后匹配值小
	case -1:
		i := pos
		for ; i > 0; i-- {
			if intervalMs > t.wheels[i].intervalMs && intervalMs <= t.wheels[i].totalMs {
				return i
			}
		}
		return i

	// intervalMs比最后匹配值大
	case 1:
		i := pos
		for ; i < t.length-1; i++ {
			if intervalMs > t.wheels[i].intervalMs && intervalMs <= t.wheels[i].totalMs {
				return i
			}
		}
		return i
	}
	return 0
}

// 二分查找当前任务可以添加的时间轮对象索引.
func (t *TTimer) binSearchIndex(n int64) (index int, result int) {
	min := 0
	max := t.length - 1
	mid := 0
	cmp := -2
	for min <= max {
		mid = int((min + max) / 2)
		switch {
		case t.wheels[mid].intervalMs == n:
			cmp = 0
		case t.wheels[mid].intervalMs > n:
			cmp = -1
		case t.wheels[mid].intervalMs < n:
			cmp = 1
		}
		switch cmp {
		case -1:
			max = mid - 1
		case 1:
			min = mid + 1
		case 0:
			return mid, cmp
		}
	}
	return mid, cmp
}
