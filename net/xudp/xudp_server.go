package xudp

import (
	"errors"
	"github.com/go-xe2/x/os/xlog"
	"github.com/go-xe2/x/sync/xsafeMap"
	"github.com/go-xe2/x/type/t"
	"net"
)

const (
	mDEFAULT_SERVER = "default"
)

// tcp server结构体
type Server struct {
	conn    *TConn // UDP server connection object.
	address string // Listening address.
	handler func(*TConn)
}

// Server表，用以存储和检索名称与Server对象之间的关联关系
var serverMapping = xsafeMap.NewStrAnyMap()

// 获取/创建一个空配置的UDP Server
// 单例模式，请保证name的唯一性
func GetServer(name ...interface{}) *Server {
	serverName := mDEFAULT_SERVER
	if len(name) > 0 {
		serverName = t.String(name[0])
	}
	if s := serverMapping.Get(serverName); s != nil {
		return s.(*Server)
	}
	s := NewServer("", nil)
	serverMapping.Set(serverName, s)
	return s
}

// 创建一个tcp server对象，并且可以选择指定一个单例名字
func NewServer(address string, handler func(*TConn), names ...string) *Server {
	s := &Server{
		address: address,
		handler: handler,
	}
	if len(names) > 0 {
		serverMapping.Set(names[0], s)
	}
	return s
}

// 设置参数 - address
func (s *Server) SetAddress(address string) {
	s.address = address
}

// 设置参数 - handler
func (s *Server) SetHandler(handler func(*TConn)) {
	s.handler = handler
}

// Close closes the connection.
// It will make server shutdowns immediately.
func (s *Server) Close() error {
	return s.conn.Close()
}

// 执行监听
func (s *Server) Run() error {
	if s.handler == nil {
		err := errors.New("start running failed: socket handler not defined")
		xlog.Error(err)
		return err
	}
	addr, err := net.ResolveUDPAddr("udp", s.address)
	if err != nil {
		xlog.Error(err)
		return err
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		xlog.Error(err)
		return err
	}
	s.conn = NewConnByNetConn(conn)
	s.handler(s.conn)
	return nil
}
