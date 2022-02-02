package designpattren

import "sync"

// 单例模式，写一个非线程安全的单例和一个线程安全的单例模式

type singleton struct {
}

// 非线程安全的懒汉模式
var instance1 *singleton

func GetInstance1() *singleton {
	if instance1 == nil {
		return new(singleton)
	}
	return instance1
}

// 线程安全的单例模式
var instance2 *singleton
var once sync.Once

func GetInstance2() *singleton {
	once.Do(func() {
		instance2 = new(singleton)
	})
	return instance2
}
