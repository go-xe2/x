package xfilePool

import (
	"fmt"
	"github.com/go-xe2/x/container/xpool"
	"github.com/go-xe2/x/os/xfileNotify"
	_type "github.com/go-xe2/x/sync/type"
	"github.com/go-xe2/x/sync/xsafeMap"
	"os"
	"sync"
)

type TFilePool struct {
	id     *_type.TInt  // 指针池ID，用以识别指针池是否需要重建
	pool   *xpool.TPool // 底层对象池
	initEd *_type.TBool // 是否初始化(在执行第一次执行File方法后初始化，主要用于文件监听的添加，但是只能添加一次)
	expire int          // 过期时间
}

// 文件指针池指针
type File struct {
	*os.File              // 底层文件指针
	mu       sync.RWMutex // 互斥锁
	pool     *TFilePool   // 所属池
	poolId   int          // 所属池ID，如果池ID不同表示池已经重建，那么该文件指针也应当销毁，不能重新丢到原有的池中
	flag     int          // 打开标志
	perm     os.FileMode  // 打开权限
	path     string       // 绝对路径
}

var (
	// 全局文件指针池Map, 不过期
	pools = xsafeMap.NewStrAnyMap()
)

// 获得文件对象，并自动创建指针池(过期时间单位：毫秒)
func Open(path string, flag int, perm os.FileMode, expire ...int) (file *File, err error) {
	fpExpire := 0
	if len(expire) > 0 {
		fpExpire = expire[0]
	}
	pool := pools.GetOrSetFuncLock(fmt.Sprintf("%s&%d&%d&%d", path, flag, expire, perm), func() interface{} {
		return New(path, flag, perm, fpExpire)
	}).(*TFilePool)

	return pool.File()
}

// Deprecated.
// See Open.
func OpenFile(path string, flag int, perm os.FileMode, expire ...int) (file *File, err error) {
	return Open(path, flag, perm, expire...)
}

// 创建一个文件指针池，expire = 0表示不过期，expire < 0表示使用完立即回收，expire > 0表示超时回收，默认值为0表示不过期。
// 注意过期时间单位为：毫秒。
func New(path string, flag int, perm os.FileMode, expire ...int) *TFilePool {
	fpExpire := 0
	if len(expire) > 0 {
		fpExpire = expire[0]
	}
	p := &TFilePool{
		id:     _type.NewInt(),
		expire: fpExpire,
		initEd: _type.NewBool(),
	}
	p.pool = newFilePool(p, path, flag, perm, fpExpire)
	return p
}

// 创建文件指针池
func newFilePool(p *TFilePool, path string, flag int, perm os.FileMode, expire int) *xpool.TPool {
	pool := xpool.New(expire, func() (interface{}, error) {
		file, err := os.OpenFile(path, flag, perm)
		if err != nil {
			return nil, err
		}
		return &File{
			File:   file,
			pool:   p,
			poolId: p.id.Val(),
			flag:   flag,
			perm:   perm,
			path:   path,
		}, nil
	}, func(i interface{}) {
		_ = i.(*File).File.Close()
	})
	return pool
}

// 获得一个文件打开指针
func (p *TFilePool) File() (*File, error) {
	if v, err := p.pool.Get(); err != nil {
		return nil, err
	} else {
		f := v.(*File)
		stat, err := os.Stat(f.path)
		if f.flag&os.O_CREATE > 0 {
			if os.IsNotExist(err) {
				if file, err := os.OpenFile(f.path, f.flag, f.perm); err != nil {
					return nil, err
				} else {
					f.File = file
					if stat, err = f.Stat(); err != nil {
						return nil, err
					}
				}
			}
		}
		if f.flag&os.O_TRUNC > 0 {
			if stat.Size() > 0 {
				if err := f.Truncate(0); err != nil {
					return nil, err
				}
			}
		}
		if f.flag&os.O_APPEND > 0 {
			if _, err := f.Seek(0, 2); err != nil {
				return nil, err
			}
		} else {
			if _, err := f.Seek(0, 0); err != nil {
				return nil, err
			}
		}
		// 优先使用 !p.inited.Val() 原子读取操作判断，保证判断操作的效率；
		// p.inited.Set(true) == false 使用原子写入操作，保证该操作的原子性；
		if !p.initEd.Val() && p.initEd.Set(true) == false {
			_, _ = xfileNotify.Add(f.path, func(event *xfileNotify.TEvent) {
				// 如果文件被删除或者重命名，立即重建指针池
				if event.IsRemove() || event.IsRename() {
					// 原有的指针都不要了
					p.id.Add(1)
					// Clear相当于重建指针池
					p.pool.Clear()
					// 为保证原子操作，但又不想加锁，
					// 这里再执行一次原子Add，将在两次Add中间可能分配出去的文件指针丢弃掉
					p.id.Add(1)
				}
			}, false)
		}
		return f, nil
	}
}

// 关闭指针池
func (p *TFilePool) Close() {
	p.pool.Close()
}

// 获得底层文件指针(返回error是标准库io.ReadWriteCloser接口实现)
func (f *File) Close() error {
	if f.poolId == f.pool.id.Val() {
		f.pool.pool.Put(f)
	}
	return nil
}
