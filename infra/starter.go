package infra

import (
	"github.com/sirupsen/logrus"
	"reflect"
)

type StarterContext map[string]interface{}

// 资源启动器，每个应用少不了依赖其他资源，比如数据库，缓存，消息中间件等等服务
// 启动器实现类，不需要实现所有方法，只需要实现对应的阶段方法即可，可以嵌入@BaseStarter
// 通过实现资源启动器接口和资源启动注册器，友好的管理这些资源的初始化、安装、启动和停止。
// Starter对象注册器，所有需要在系统启动时需要实例化和运行的逻辑，都可以实现此接口
// 注意只有Start方法才能被阻塞，如果是阻塞Start()，同时StartBlocking()要返回true
type Starter interface {
	// 资源初始化，通常把一些准备资源放在这里运行
	Init(StarterContext)
	// 资源的安装，所有启动需要的具备条件，使得资源达到可以启动的就备状态
	Setup(StarterContext)
	// 启动资源，达到可以使用的状态
	Start(StarterContext)
	// 用于释放资源和终止资源的使用，
	// 通常要优雅的释放，等待正在进行的任务继续，但不再接受新的任务
	Stop(StarterContext)
	// 说明该资源启动器开始启动服务时，是否会阻塞
	// 如果存在多个阻塞启动器时，只有最后一个阻塞，之前的会通过goroutine来异步启动
	// 所以，需要规划好启动器注册顺序
	StartBlocking() bool
}

type BaseStarter struct {
}

func (b BaseStarter) Init(context StarterContext) {}

func (b BaseStarter) Setup(context StarterContext) {}

func (b BaseStarter) Start(context StarterContext) {}

func (b BaseStarter) Stop(context StarterContext) {}

func (b BaseStarter) StartBlocking() bool { return false }

// 服务启动注册器
// 不用需外部构造，全局只有一个
type starterRegister struct {
	nonBlockingStarters []Starter
	blockingStarters    []Starter
}

func (s *starterRegister) AllStarters() []Starter {
	starters := make([]Starter, 0)
	starters = append(starters, s.nonBlockingStarters...)
	starters = append(starters, s.blockingStarters...)
	return starters
}

// 注册starter
func (s *starterRegister) Register(starter Starter) {
	if starter.StartBlocking() {
		s.blockingStarters = append(s.blockingStarters, starter)
	} else {
		s.nonBlockingStarters = append(s.nonBlockingStarters, starter)
	}

	typ := reflect.TypeOf(starter)
	logrus.Infof("Register starter: %s", typ.String())
}

var StarterRegister *starterRegister = &starterRegister{}

// 注册starter
func Register(starter Starter) {
	StarterRegister.Register(starter)
}

// 获取所有注册的starter
func GetStarters() []Starter {
	return StarterRegister.AllStarters()
}
