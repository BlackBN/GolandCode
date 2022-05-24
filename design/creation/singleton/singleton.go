package singleton

import "sync"

type Singleton struct {
}

var singleton *Singleton

// 单例模式采用了 饿汉式 和 懒汉式 两种实现。
// 个人其实更倾向于饿汉式的实现，简单，并且可以将问题及早暴露。
// 懒汉式虽然支持延迟加载，但是这只是把冷启动时间放到了第一次使用的时候，并没有本质上解决问题，并且为了实现懒汉式还不可避免的需要加锁。

// 饿汉式
func init() {
	singleton = &Singleton{}
}

func GetSingleton() *Singleton {
	return singleton
}

// 懒汉式
var (
	lazySingleton *Singleton
	//使用 sync.Once 保证只执行一次初始化函数，变量初始化过程中，所有读都被阻塞，直到初始化完成
	once = &sync.Once{}
)

func GetLazySingleton() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
