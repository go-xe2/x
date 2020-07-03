package xtimer

import (
	_type "github.com/go-xe2/x/sync/type"
	"github.com/go-xe2/x/sync/xsafeStack"
	"time"
)

// 单层时间轮
type twheel struct {
	timer      *TTimer                    // 所属定时器
	level      int                        // 所属分层索引号
	slots      []*xsafeStack.TSafeStackQe // 所有的循环任务项, 按照Slot Number进行分组
	number     int64                      // Slot Number=len(slots)
	ticks      *_type.TInt64              // 当前时间轮已转动的刻度数量
	totalMs    int64                      // 整个时间轮的时间长度(毫秒)=number*interval
	createMs   int64                      // 创建时间(毫秒)
	intervalMs int64                      // 时间间隔(slot时间长度, 毫秒)
}

// 开始循环
func (w *twheel) start() {
	go func() {
		ticker := time.NewTicker(time.Duration(w.intervalMs) * time.Millisecond)
		for {
			select {
			case <-ticker.C:
				switch w.timer.status.Val() {
				case STATUS_RUNNING:
					w.proceed()

				case STATUS_STOPPED:
				case STATUS_CLOSED:
					ticker.Stop()
					return
				}

			}
		}
	}()
}

// 执行时间轮刻度逻辑
func (w *twheel) proceed() {
	n := w.ticks.Add(1)
	l := w.slots[int(n%w.number)]
	length := l.Size()
	if length > 0 {
		go func(l *xsafeStack.TSafeStackQe, nowTicks int64) {
			entry := (*TEntry)(nil)
			nowMs := time.Now().UnixNano() / 1e6
			for i := length; i > 0; i-- {
				if v := l.PopFront(); v == nil {
					break
				} else {
					entry = v.(*TEntry)
				}
				// 是否满足运行条件
				runnable, addable := entry.check(nowTicks, nowMs)
				if runnable {
					// 异步执行运行
					go func(entry *TEntry) {
						defer func() {
							if err := recover(); err != nil {
								if err != mPANIC_EXIT {
									panic(err)
								} else {
									entry.Close()
								}
							}
							if entry.Status() == STATUS_RUNNING {
								entry.SetStatus(STATUS_READY)
							}
						}()
						entry.job()
					}(entry)
				}
				// 是否继续添运行, 滚动任务
				if addable {
					entry.wheel.timer.doAddEntryByParent(entry.rawIntervalMs, entry)
				}
			}
		}(l, n)
	}
}
