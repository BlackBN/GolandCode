package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//单元测试
//以 TestXxx 开头
//参数为 t *testing.T
//在当前目录下执行 go test 命令，会执行以 _test.go 为结尾的文件里的单元测试函数
func TestGetSingleton(t *testing.T) {
	assert.Equal(t, GetSingleton(), GetSingleton())
}

//压力测试
//以 BenchmarkXxx 开头
//参数为 b *testing.B
//如需使用 go test 执行到压力测试函数，则需要带上参数 -test.bench，如 go test -test.bench=".*" 表示测试全部的压力测试函数
func BenchmarkGetSingletonParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetSingleton() != GetSingleton() {
				b.Errorf("test fail")
			}
		}
	})
}

func TestGetLazySingleton(t *testing.T) {
	assert.Equal(t, GetLazySingleton(), GetLazySingleton())
}

func BenchmarkGetLazySingletonParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazySingleton() != GetLazySingleton() {
				b.Errorf("test fail")
			}
		}
	})
}
