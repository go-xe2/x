package xproc

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// 子进程
type TProcess struct {
	exec.Cmd
	Manager *TProcessManager // 所属进程管理器
	PPid    int              // 自定义关联的父进程ID
}

// 创建一个进程(不执行)
func NewProcess(path string, args []string, environment ...[]string) *TProcess {
	var env []string
	if len(environment) > 0 {
		env = make([]string, 0)
		for _, v := range environment[0] {
			env = append(env, v)
		}
	} else {
		env = os.Environ()
	}
	env = append(env, fmt.Sprintf("%s=%s", mPROC_TEMP_DIR_ENV_KEY, os.TempDir()))
	p := &TProcess{
		Manager: nil,
		PPid:    os.Getpid(),
		Cmd: exec.Cmd{
			Args:       []string{path},
			Path:       path,
			Stdin:      os.Stdin,
			Stdout:     os.Stdout,
			Stderr:     os.Stderr,
			Env:        env,
			ExtraFiles: make([]*os.File, 0),
		},
	}
	// 当前工作目录
	if d, err := os.Getwd(); err == nil {
		p.Dir = d
	}
	if len(args) > 0 {
		start := 0
		if strings.EqualFold(path, args[0]) {
			start = 1
		}
		p.Args = append(p.Args, args[start:]...)
	}
	return p
}

// 开始执行(非阻塞)
func (p *TProcess) Start() (int, error) {
	if p.Process != nil {
		return p.Pid(), nil
	}
	p.Env = append(p.Env, fmt.Sprintf("%s=%d", mPROC_ENV_KEY_PPID_KEY, p.PPid))
	if err := p.Cmd.Start(); err == nil {
		if p.Manager != nil {
			p.Manager.processes.Set(p.Process.Pid, p)
		}
		return p.Process.Pid, nil
	} else {
		return 0, err
	}
}

// 运行进程(阻塞等待执行完毕)
func (p *TProcess) Run() error {
	if _, err := p.Start(); err == nil {
		return p.Wait()
	} else {
		return err
	}
}

// PID
func (p *TProcess) Pid() int {
	if p.Process != nil {
		return p.Process.Pid
	}
	return 0
}

// 向进程发送消息
func (p *TProcess) Send(data []byte) error {
	if p.Process != nil {
		return Send(p.Process.Pid, data)
	}
	return errors.New("invalid process")
}

func (p *TProcess) Release() error {
	return p.Process.Release()
}

func (p *TProcess) Kill() error {
	if err := p.Process.Kill(); err == nil {
		if p.Manager != nil {
			p.Manager.processes.Remove(p.Pid())
		}
		return nil
	} else {
		return err
	}
}

func (p *TProcess) Signal(sig os.Signal) error {
	return p.Process.Signal(sig)
}
