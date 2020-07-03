package xtcp

import (
	"crypto/tls"
	"errors"
	"github.com/go-xe2/x/os/xlog"
	"github.com/go-xe2/x/sync/xsafeMap"
	"github.com/go-xe2/x/type/t"
	"net"
)

const (
	mDEFAULT_SERVER = "default"
)

// TCP Server.
type Server struct {
	listen    net.Listener
	address   string
	handler   func(*TConn)
	tlsConfig *tls.Config
}

// Map for name to server, for singleton purpose.
var serverMapping = xsafeMap.NewStrAnyMap()

// GetServer returns the TCP server with specified <name>,
// or it returns a new normal TCP server named <name> if it does not exist.
// The parameter <name> is used to specify the TCP server
func GetServer(name ...interface{}) *Server {
	serverName := mDEFAULT_SERVER
	if len(name) > 0 {
		serverName = t.String(name[0])
	}
	return serverMapping.GetOrSetFuncLock(serverName, func() interface{} {
		return NewServer("", nil)
	}).(*Server)
}

// NewServer creates and returns a new normal TCP server.
// The parameter <name> is optional, which is used to specify the instance name of the server.
func NewServer(address string, handler func(*TConn), name ...string) *Server {
	s := &Server{
		address: address,
		handler: handler,
	}
	if len(name) > 0 {
		serverMapping.Set(name[0], s)
	}
	return s
}

// NewServerTLS creates and returns a new TCP server with TLS support.
// The parameter <name> is optional, which is used to specify the instance name of the server.
func NewServerTLS(address string, tlsConfig *tls.Config, handler func(*TConn), name ...string) *Server {
	s := NewServer(address, handler, name...)
	s.SetTLSConfig(tlsConfig)
	return s
}

// NewServerKeyCrt creates and returns a new TCP server with TLS support.
// The parameter <name> is optional, which is used to specify the instance name of the server.
func NewServerKeyCrt(address, crtFile, keyFile string, handler func(*TConn), name ...string) *Server {
	s := NewServer(address, handler, name...)
	if err := s.SetTLSKeyCrt(crtFile, keyFile); err != nil {
		xlog.Error(err)
	}
	return s
}

// SetAddress sets the listening address for server.
func (s *Server) SetAddress(address string) {
	s.address = address
}

// SetHandler sets the connection handler for server.
func (s *Server) SetHandler(handler func(*TConn)) {
	s.handler = handler
}

// SetTlsKeyCrt sets the certificate and key file for TLS configuration of server.
func (s *Server) SetTLSKeyCrt(crtFile, keyFile string) error {
	tlsConfig, err := LoadKeyCrt(crtFile, keyFile)
	if err != nil {
		return err
	}
	s.tlsConfig = tlsConfig
	return nil
}

// SetTlsConfig sets the TLS configuration of server.
func (s *Server) SetTLSConfig(tlsConfig *tls.Config) {
	s.tlsConfig = tlsConfig
}

// Close closes the listener and shutdowns the server.
func (s *Server) Close() error {
	return s.listen.Close()
}

// Run starts running the TCP Server.
func (s *Server) Run() (err error) {
	if s.handler == nil {
		err = errors.New("start running failed: socket handler not defined")
		xlog.Error(err)
		return
	}
	if s.tlsConfig != nil {
		// TLS Server
		s.listen, err = tls.Listen("tcp", s.address, s.tlsConfig)
		if err != nil {
			xlog.Error(err)
			return
		}
	} else {
		// Normal Server
		addr, err := net.ResolveTCPAddr("tcp", s.address)
		if err != nil {
			xlog.Error(err)
			return err
		}
		s.listen, err = net.ListenTCP("tcp", addr)
		if err != nil {
			xlog.Error(err)
			return err
		}
	}
	for {
		if conn, err := s.listen.Accept(); err != nil {
			xlog.Error(err)
			return err
		} else if conn != nil {
			go s.handler(NewConnByNetConn(conn))
		}
	}
}
