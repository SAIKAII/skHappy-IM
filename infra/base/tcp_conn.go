package base

import (
	coma "github.com/SAIKAII/go-conn-manager"
	"github.com/SAIKAII/skHappy-IM/infra"
	"sync"
)

var connManager *ConnManager

func ConnectionManager() *ConnManager {
	return connManager
}

type ConnStarter struct {
	infra.BaseStarter
}

func (s *ConnStarter) Setup(ctx infra.StarterContext) {
	connManager = &ConnManager{
		Conns: make(map[string]*coma.Conn),
		lock:  sync.RWMutex{},
	}
}

type ConnManager struct {
	Conns map[string]*coma.Conn
	lock  sync.RWMutex
}

// StoreConn 存储名字为name的连接到管理器中，若该名字已关联连接，则覆盖
func (t *ConnManager) StoreConn(name string, conn *coma.Conn) {
	t.lock.Lock()

	delete(t.Conns, name)
	t.Conns[name] = conn

	t.lock.Unlock()
}

// DeleteConn 删除名字为name的连接
func (t *ConnManager) DeleteConn(name string) {
	t.lock.Lock()

	delete(t.Conns, name)

	t.lock.Unlock()
}

// GetConn 获取名字为name的连接
func (t *ConnManager) GetConn(name string) *coma.Conn {
	t.lock.RLock()
	defer t.lock.RUnlock()

	return t.Conns[name]
}
