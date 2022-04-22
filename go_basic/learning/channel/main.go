package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
	"unsafe"
)

func writeData(data chan<- int) {
	for i := 0; i < 50; i++ {
		data <- i
		//time.Sleep(time.Second)
	}
	close(data)
}

func readData(data <-chan int, exitChannel chan<- bool) {
	for c := range data {
		fmt.Printf("read data : %d\n", c)
	}
	exitChannel <- true
	close(exitChannel)
}

func main() {
	fmt.Println("ab")

	data := make(chan int, 10)
	exitChannel := make(chan bool, 1)
	// var aa chan<- int
	// aa = make(chan int, 23)
	// aa <- 2
	//基础类型 整形 浮点型 布尔型 string
	go writeData(data)
	go readData(data, exitChannel)
	for {
		_, ok := <-exitChannel
		if ok {
			break
		}
	}
	// var a chan<- int = make(chan int, 10)
	// a <- 20
	// var c <-chan int = make(<-chan int)
	// num := <-c
	var a int = 3 //根据操作系统位数，一般都是64位
	fmt.Printf("a type : %T\raa\n", a)

	fmt.Printf("a unicode byte : %d\n", unsafe.Sizeof(a))
	var c float32 = 23.222 //浮点数都是有符号位的  符号位 + 指数位 + 尾数位
	fmt.Printf("c type: %T, size : %d\n", c, unsafe.Sizeof(c))
	var d float64 = 23.2323
	fmt.Printf("c type: %T, size : %d\n", d, unsafe.Sizeof(d))
	var f float32 = -123.0000901 //精度会损失
	fmt.Println("f : ", f)
	var e float64 = -123.0000902342341341324134 //精度会损失
	fmt.Println("e : ", e)
	var k = .1244 // 默认float64 , 不受os系统位数限制
	fmt.Printf("k : %f , %T\n", k, k)
	var l = 4.555e-4
	fmt.Println("l : ", l)
	var m = -3.444e-3
	fmt.Println("m : ", m)

	var qq byte = 'a'
	var ww byte = 'b'
	fmt.Println("qq : ", qq)
	fmt.Println("ww : ", ww)
	fmt.Printf("qq: %c,ww:%c, qq:%T,size : %d\n", qq, ww, qq, unsafe.Sizeof(qq))
	var ee int16 = '毓'
	fmt.Println("ee:", ee)
	fmt.Printf("ee %T, size: %d\n", ee, unsafe.Sizeof(ee))
	var rr int = 'c'
	fmt.Println("rr : ", rr)

	//字符使用的UTF-8编码 英文字符1个字节 ， 汉子3个字节

	var tt bool = false
	fmt.Printf("tt: %t, %T,%d\n", tt, tt, unsafe.Sizeof(tt))

	var uu = "abcdefghijklmnopqrstuvwxyz毓" //不可变
	fmt.Printf("%T,%d\n", uu, unsafe.Sizeof(uu))
	uu = "毓"
	fmt.Printf("%c\n", uu[0])
	fmt.Println(uu[0])
	str1 := `
	sdfad
	adf
	adf
	asdf
	ads
	f
	ad
	f
	adfa`
	fmt.Println(str1)

	str2 := "adfaf" +
		"adfa"
	str2 += "xujiayu"
	fmt.Println(str2)

	//基本类型转换 T(v) T : 类型 , v : 变量名

	var ii int32 = 900000
	var oo float64 = float64(ii)
	fmt.Println("oo: ", oo)
	var iii byte = byte(ii)
	fmt.Println("iii:", iii)
	fmt.Printf("%T\n", iii)
	str4 := fmt.Sprintf("%d", iii)
	fmt.Printf("%T,%s%q\n", str4, str4, "''")
	cccc := strconv.Itoa(int(ii))
	fmt.Printf("%s,%T,%d\n", cccc, cccc, unsafe.Sizeof(cccc))

	str3 := "1230404"
	pppp, err := strconv.Atoi(str3)
	if err != nil {
		fmt.Println("dddd")
	}
	fmt.Println(pppp)
	ppp, err := strconv.ParseInt(str3, 10, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%T,%v,%d\n", ppp, ppp, unsafe.Sizeof(ppp))

	//指针

	var kint int = 123
	var pointOne *int = &kint
	fmt.Printf("%T,%d\n", pointOne, unsafe.Sizeof(pointOne))
	//var pointTwo **int = &pointOne
	changeNum(pointOne)
	fmt.Printf("%v\n", kint)

	//算数运算符 + - * / %
	// ++ -- 只能作为独立语言使用 不允许 b:=a++ ，只能 a++
	//赋值 = += -= *= /= %=   <<= >>= &= ^= |=

	//比较|关系  == != > < >= <=

	//逻辑  && || !
	//位 << >> & | ^
	// >> 右移 符号位不变 并用符号位(0或1)补溢出的高位
	// << 左移 符号位不变 低位补0

	//其它运算符 & *
	//go没有三元运算符
	nn := 1
	nn++
	if nn > 1 {
		fmt.Println("ok")
	}
	aaaa := 20
	bbbb := 30
	aaaa = aaaa + bbbb
	bbbb = aaaa - bbbb
	aaaa = aaaa - bbbb
	fmt.Println(aaaa, bbbb)

	//获取终端输入
	var name string
	var age byte
	var sal float64
	var isPass bool
	// fmt.Scanln(&name)
	// fmt.Scanln(&age)
	// fmt.Scanln(&sal)
	// fmt.Scanln(&isPass)
	// fmt.Println(name, age, sal, isPass)
	//fmt.Scanf("%s,%d,%f,%t", &name, &age, &sal, &isPass)
	fmt.Println(name, age, sal, isPass)
	var er int64 = 1223434133333
	fmt.Printf("%b\n", er)
	var ba int = 0x2374
	fmt.Printf("%d\n", ba)
	// 二进制 八进制 十进制 十六进制
	// 二进制、八进制、十六进制转十进制   每一位数字乘以进制数的（位数减一）次方，然后相加
	// 十进制转二进制、八进制、十六进制   不断的除以进制数得到的余数反向连在一块
	// 二进制转八进制  每三位算
	// 二进制转十六进制 每四位算
	// 八进制转二进制 每一位的数字用三位二进制表示
	// 十六进制转二进制 每一位的数字用四位二进制表示
	//位运算
	//有符号位
	//源码 反码 补码
	// 正数 源码 反码 补码都一样
	// 负数 反码 = 符号位不变，其余位数取反  补码 = 反码+1
	// 计算机储存整数时，是用该整数（正负都是）的补码进行储存的。
	//|---------------------------------|-------------------|
	//|9 - 17 = -8                      |计算机实际存储的二进制 |
	//|---------------------------------|-------------------|
	//|9   源码 00001001  反码 00001001  | 补码  00001001     |
	//|-17 源码 10010001  反码 11101110  | 补码  11101111     |
	//|---------------------------------|-------------------|
	//|    源码 10011010  反码 11110111  | 补码  11111000     |
	//|---------------------------------|-------------------|
	//|-8  源码 10001000  反码 11110111  | 补码  11111000     |
	//|---------------------------------|-------------------|
	// 补码的补码就是源码 11111000 源码 11111000  反码 10000111 补码 10001000 = -8的源码
	// 负数源码加补码总等于 1 0000 0000 字节数溢出一位

	var leftMove int8 = 23 << 2
	fmt.Printf("%d\n", leftMove)
	var rightMove int8 = 23 >> 2
	fmt.Printf("%d\n", rightMove)

	var leftFMove int8 = -23 << 2
	fmt.Printf("%d\n", leftFMove)
	var rightFMove int8 = -23 >> 2
	fmt.Printf("%d\n", rightFMove)
	if leftFMove > 0 {
		fmt.Println("aa")
	} else if leftFMove < 1 {
		fmt.Println("bb")
	} else {
		fmt.Println("cc")
	}

	//switch
	switch leftFMove {
	case 18:
		fmt.Println("ee")
	case 10:
		fmt.Println("ff")
	default:
		fmt.Println("gg")
	}
	var agedd int = 10

	//做if else用
	switch {
	case agedd == 10:
		fmt.Println("age == 10")
		//默认只能穿透一层 不要多用
		fallthrough
	case agedd == 20:
		fmt.Println("age == 20")
		fallthrough
	default:
		fmt.Println("default")
	}
	fmt.Println(agedd)

	//判断类型
	var xxd interface{}
	var yddd = 10.0
	xxd = yddd
	switch i := xxd.(type) {
	case float64:
		fmt.Printf("%T\n", i)
	}
	// for 循环
	for i := 1; i < 11; i++ {
		fmt.Println("test for")
	}

	mnb := "abcda虚假毓"

	//这种方式是一个一个字节读的，一个中文字符3个字节，所以有乱码
	for i := 0; i < len(mnb); i++ {
		fmt.Printf("%c\n", mnb[i])
	}
	//转成切片
	mnb2 := []rune(mnb)
	fmt.Println(len(mnb2))
	for i := 0; i < len(mnb2); i++ {
		fmt.Printf("%c\n", mnb2[i])
	}
	//或用range遍历
	for i, v := range mnb {
		fmt.Printf("%d,%c\n", i, v)
	}

	var xin byte = 'x'
	var empty byte = ' '
	var leng int = 12
	for i := 1; i <= (leng+1)/2; i++ {
		innerLng := (leng+1)/2 - i
		for j := 0; j < innerLng; j++ {
			fmt.Printf("%c", empty)
		}
		for j := 0; j < 2*i-1; j++ {
			fmt.Printf("%c", xin)
		}
		fmt.Println()
	}
	rand.Seed(time.Now().Unix())
	ax := rand.Intn(10) + 1
	fmt.Println(ax)
	//ddd:
	for i := 1; i < ax+1; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%d*%d=%d  ", i, j, i*j)
			//break ddd
			//continue ddd
		}
		fmt.Println()
	}
	// goto
	fmt.Println(hypot(12.3, 12.3))

}

func changeNum(point *int) {
	fmt.Printf("%v\n", *point)
	*point = 10
	fmt.Printf("%v\n", *point)

}

func TwoToTen() {}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}
