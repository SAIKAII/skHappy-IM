package infra

import (
	"github.com/sirupsen/logrus"
	"reflect"
)

type BootApplication struct {
	starterCtx StarterContext
}

// 构造启动器
func New(ctx StarterContext) *BootApplication {
	a := &BootApplication{starterCtx: ctx}
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
		logrus.Infof("Initializing type: %s", typ.String())
		v.Init(b.starterCtx)
	}
}

// setup 程序安装
func (b *BootApplication) setup() {
	logrus.Info("Setup starters...")
	for _, v := range GetStarters() {
		typ := reflect.TypeOf(v)
		logrus.Info("Setup: ", typ.String())
		v.Setup(b.starterCtx)
	}
}

// start 程序运行，开始接受调用
func (b *BootApplication) start() {
	logrus.Info("Starting starters...")
	for i, v := range GetStarters() {
		typ := reflect.TypeOf(v)
		logrus.Info("Starting: ", typ.String())
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
		logrus.Info("Stopping: ", typ.String())
		v.Stop(b.starterCtx)
	}
}
