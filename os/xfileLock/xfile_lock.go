package xfileLock

import (
	"github.com/go-xe2/third/github.com/theckman/go-flock"
	"github.com/go-xe2/x/os/xfile"
)

type TFileLocker struct {
	flock *flock.Flock // Underlying file locker.
}

func New(file string) *TFileLocker {
	dir := xfile.TempDir() + xfile.Separator + "xflock"
	if !xfile.Exists(dir) {
		_ = xfile.Mkdir(dir)
	}
	path := dir + xfile.Separator + file
	lock := flock.NewFlock(path)
	return &TFileLocker{
		flock: lock,
	}
}

func (l *TFileLocker) Path() string {
	return l.flock.Path()
}

func (l *TFileLocker) IsLocked() bool {
	return l.flock.Locked()
}

func (l *TFileLocker) IsRLocked() bool {
	return l.flock.RLocked()
}

func (l *TFileLocker) TryLock() bool {
	ok, _ := l.flock.TryLock()
	return ok
}

func (l *TFileLocker) TryRLock() bool {
	ok, _ := l.flock.TryRLock()
	return ok
}

func (l *TFileLocker) Lock() (err error) {
	return l.flock.Lock()
}

func (l *TFileLocker) Unlock() (err error) {
	return l.flock.Unlock()
}

func (l *TFileLocker) RLock() (err error) {
	return l.flock.RLock()
}

func (l *TFileLocker) RUnlock() (err error) {
	return l.flock.Unlock()
}
