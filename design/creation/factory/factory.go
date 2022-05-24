package factory

import (
	"fmt"
	"reflect"
)

//简单工厂
type FoodCreateMachine interface {
	CreateFood(rawMaterial string)
}

type breadCreateMachine struct {
}

func (bread *breadCreateMachine) CreateFood(rawMaterial string) {
	fmt.Printf("breadCreateMachine CreateFood rawMaterial is %s, create bread\n", rawMaterial)
}

func (bread *breadCreateMachine) CreateHotFood(rawMaterial string) {
	fmt.Printf("breadCreateMachine CreateHotFood rawMaterial is %s, create hot bread\n", rawMaterial)
}

type milkCreateMachine struct {
}

func (milk *milkCreateMachine) CreateFood(rawMaterial string) {
	fmt.Printf("milkCreateMachine CreateFood rawMaterial is %s, create milk\n", rawMaterial)
}

func (milk *milkCreateMachine) CreateHotFood(rawMaterial string) {
	fmt.Printf("milkCreateMachine CreateHotFood rawMaterial is %s, create hot milk\n", rawMaterial)
}

func GetFoodCreateMachine(rawMaterial string) FoodCreateMachine {
	switch rawMaterial {
	case "flour":
		return &breadCreateMachine{}
	case "milk":
		return &milkCreateMachine{}
	}
	return nil
}

//工厂方法
type FoodCreateMachineFactory interface {
	CreateFoodCreateMachine() FoodCreateMachine
	CreateHotFoodCreateMachine() HotFoodCreateMachine
}

type breadCreateMachineFactory struct {
}

func (bread *breadCreateMachineFactory) CreateFoodCreateMachine() FoodCreateMachine {
	fmt.Printf("get breadCreateMachineFactory Create breadCreateMachine\n")
	return &breadCreateMachine{}
}

func (bread *breadCreateMachineFactory) CreateHotFoodCreateMachine() HotFoodCreateMachine {
	fmt.Printf("get breadCreateMachineFactory Create breadCreateMachine\n")
	return &breadCreateMachine{}
}

type milkCreateMachineFactory struct {
}

func (milk *milkCreateMachineFactory) CreateFoodCreateMachine() FoodCreateMachine {
	fmt.Printf("get milkCreateMachineFactory Create milkCreateMachine\n")
	return &milkCreateMachine{}
}

func (milk *milkCreateMachineFactory) CreateHotFoodCreateMachine() HotFoodCreateMachine {
	fmt.Printf("get milkCreateMachineFactory Create milkCreateMachine\n")
	return &milkCreateMachine{}
}

func GetFoodCreateMachineFactory(rawMaterial string) FoodCreateMachineFactory {
	switch rawMaterial {
	case "flour":
		return &breadCreateMachineFactory{}
	case "milk":
		return &milkCreateMachineFactory{}
	}
	return nil
}

//抽象工厂
type HotFoodCreateMachine interface {
	CreateHotFood(rawMaterial string)
}

// type hotDogFoodCreateMachine struct {
// }

// func (hotDog *hotDogFoodCreateMachine) CreateHotFood(rawMaterial string) {
// 	fmt.Printf("hotDogFoodCreateMachine CreateHotFood hot dog, rawMaterial is %s ", rawMaterial)
// }

// DI，依赖注入框架
// 使用反射实现的: https://github.com/uber-go/dig
// 使用 generate 实现的: https://github.com/google/wire
// Container DI 容器

type Container struct {
	// 假设一种类型只能有一个 provider 提供
	providers map[reflect.Type]provider

	// 缓存以生成的对象
	results map[reflect.Type]reflect.Value
}

type provider struct {
	value reflect.Value

	params []reflect.Type
}

// New 创建一个容器
func New() *Container {
	return &Container{
		providers: map[reflect.Type]provider{},
		results:   map[reflect.Type]reflect.Value{},
	}
}

// isError 判断是否是 error 类型
func isError(t reflect.Type) bool {
	if t.Kind() != reflect.Interface {
		return false
	}
	return t.Implements(reflect.TypeOf(reflect.TypeOf((*error)(nil)).Elem()))
}

// Provide 对象提供者，需要传入一个对象的工厂方法，后续会用于对象的创建
func (c *Container) Provide(constructor interface{}) error {
	v := reflect.ValueOf(constructor)

	// 仅支持函数 provider
	if v.Kind() != reflect.Func {
		return fmt.Errorf("constructor must be a func")
	}

	vt := v.Type()

	// 获取参数
	params := make([]reflect.Type, vt.NumIn())
	for i := 0; i < vt.NumIn(); i++ {
		params[i] = vt.In(i)
	}

	// 获取返回值
	results := make([]reflect.Type, vt.NumOut())
	for i := 0; i < vt.NumOut(); i++ {
		results[i] = vt.Out(i)
	}

	provider := provider{
		value:  v,
		params: params,
	}

	// 保存不同类型的 provider
	for _, result := range results {
		// 判断返回值是不是 error
		if isError(result) {
			continue
		}

		if _, ok := c.providers[result]; ok {
			return fmt.Errorf("%s had a provider", result)
		}

		c.providers[result] = provider
	}

	return nil
}

// Invoke 函数执行入口
func (c *Container) Invoke(function interface{}) error {
	v := reflect.ValueOf(function)

	// 仅支持函数 provider
	if v.Kind() != reflect.Func {
		return fmt.Errorf("constructor must be a func")
	}

	vt := v.Type()

	// 获取参数
	var err error
	params := make([]reflect.Value, vt.NumIn())
	for i := 0; i < vt.NumIn(); i++ {
		params[i], err = c.buildParam(vt.In(i))
		if err != nil {
			return err
		}
	}

	v.Call(params)

	// 获取 providers
	return nil
}

// buildParam 构建参数
// 1. 从容器中获取 provider
// 2. 递归获取 provider 的参数值
// 3. 获取到参数之后执行函数
// 4. 将结果缓存并且返回结果
func (c *Container) buildParam(param reflect.Type) (val reflect.Value, err error) {
	if result, ok := c.results[param]; ok {
		return result, nil
	}

	provider, ok := c.providers[param]
	if !ok {
		return reflect.Value{}, fmt.Errorf("can not found provider: %s", param)
	}

	params := make([]reflect.Value, len(provider.params))
	for i, p := range provider.params {
		params[i], err = c.buildParam(p)
	}

	results := provider.value.Call(params)
	for _, result := range results {
		// 判断是否报错
		if isError(result.Type()) && !result.IsNil() {
			return reflect.Value{}, fmt.Errorf("%s call err: %+v", provider, result)
		}

		if !isError(result.Type()) && !result.IsNil() {
			c.results[result.Type()] = result
		}
	}
	return c.results[param], nil
}
