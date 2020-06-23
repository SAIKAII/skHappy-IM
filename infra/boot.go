package infra

import (
	"github.com/sirupsen/logrus"
	"reflect"
)

type BootApplication struct {
	starterCtx StarterContext
}

// 构造启动器
func New() *BootApplication {
	a := &BootApplication{starterCtx: StarterContext{}}
	return a
}

// 启动程序
func (b *BootApplication) Start() {
	// 初始化
	b.init()
	// 安装
	b.setup()
	// 启动
	b.start()
}

// init 程序初始化
func (b *BootApplication) init() {
	logrus.Info("Initializing starters...")
	for _, v := range GetStarters() {
		typ := reflect.TypeOf(v)
		logrus.Debug("Initializing type: %s", typ.String())
		v.Init(b.starterCtx)
	}
}

// setup 程序安装
func (b *BootApplication) setup() {
	logrus.Info("Setup starters...")
	for _, v := range GetStarters() {
		typ := reflect.TypeOf(v)
		logrus.Debug("Setup: ", typ.String())
		v.Setup(b.starterCtx)
	}
}

// start 程序运行，开始接受调用
func (b *BootApplication) start() {
	logrus.Info("Starting starters...")
	for i, v := range GetStarters() {
		typ := reflect.TypeOf(v)
		logrus.Debug("Starting: ", typ.String())
		if v.StartBlocking() {
			if i+1 == len(GetStarters()) {
				v.Start(b.starterCtx)
			} else {
				go v.Start(b.starterCtx)
			}
		} else {
			v.Start(b.starterCtx)
		}
	}
}

// Stop 关闭程序
func (b *BootApplication) Stop() {
	logrus.Info("Stopping starters...")
	for _, v := range GetStarters() {
		typ := reflect.TypeOf(v)
		logrus.Debug("Stopping: ", typ.String())
		v.Stop(b.starterCtx)
	}
}
