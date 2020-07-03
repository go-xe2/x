package xproc

import (
	"github.com/go-xe2/x/sync/xsafeMap"
	"os"
)

// 进程管理器
type TProcessManager struct {
	processes *xsafeMap.TIntAnyMap // 所管理的子进程map
}

// 创建一个进程管理器
func NewProcessManager() *TProcessManager {
	return &TProcessManager{
		processes: xsafeMap.NewIntAnyMap(),
	}
}

// 创建一个进程(不执行)
func (m *TProcessManager) NewProcess(path string, args []string, environment []string) *TProcess {
	p := NewProcess(path, args, environment)
	p.Manager = m
	return p
}

// 获取当前进程管理器中的一个进程
func (m *TProcessManager) GetProcess(pid int) *TProcess {
	if v := m.processes.Get(pid); v != nil {
		return v.(*TProcess)
	}
	return nil
}

// 添加一个已存在进程到进程管理器中
func (m *TProcessManager) AddProcess(pid int) {
	if m.processes.Get(pid) == nil {
		if process, err := os.FindProcess(pid); err == nil {
			p := m.NewProcess("", nil, nil)
			p.Process = process
			m.processes.Set(pid, p)
		}
	}
}

// 移除进程管理器中的指定进程
func (m *TProcessManager) RemoveProcess(pid int) {
	m.processes.Remove(pid)
}

// 获取所有的进程对象，构成列表返回
func (m *TProcessManager) Processes() []*TProcess {
	processes := make([]*TProcess, 0)
	m.processes.RLockFunc(func(m map[int]interface{}) {
		for _, v := range m {
			processes = append(processes, v.(*TProcess))
		}
	})
	return processes
}

// 获取所有的进程pid，构成列表返回
func (m *TProcessManager) Pids() []int {
	return m.processes.Keys()
}

// 等待所有子进程结束
func (m *TProcessManager) WaitAll() {
	processes := m.Processes()
	if len(processes) > 0 {
		for _, p := range processes {
			p.Wait()
		}
	}
}

// 关闭所有的进程
func (m *TProcessManager) KillAll() error {
	for _, p := range m.Processes() {
		if err := p.Kill(); err != nil {
			return err
		}
	}
	return nil
}

// 向所有进程发送信号量
func (m *TProcessManager) SignalAll(sig os.Signal) error {
	for _, p := range m.Processes() {
		if err := p.Signal(sig); err != nil {
			return err
		}
	}
	return nil
}

// 向所有进程发送消息
func (m *TProcessManager) Send(data []byte) {
	for _, p := range m.Processes() {
		p.Send(data)
	}
}

// 向指定进程发送消息
func (m *TProcessManager) SendTo(pid int, data []byte) error {
	return Send(pid, data)
}

// 清空管理器
func (m *TProcessManager) Clear() {
	m.processes.Clear()
}

// 当前进程总数
func (m *TProcessManager) Size() int {
	return m.processes.Size()
}
