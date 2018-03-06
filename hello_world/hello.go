package main

import "fmt"
//import (
//	"os"
//	"time"

// package 别名
// io "fmt"

// 省略调用 包前面加"."


//函数名,常量,变量等 首字母大写为public 小写为private

//)


//Go基本类型

//布尔型:bool 1字节
//整型: int/uint 根据平台可能为32或者64
//8位整型: int8/uint8 8字节整形
//int32 别名 rune
//浮点型: float32/float64
//字节型: byte (uint8别名)

//定义常量
const PI = 3.14

//全局变量声明与赋值
var name = "张三"

//一般类型声明
type newType int

//结构声明
type gopher struct {

}

//接口声明
type golang interface {

}

func main() {
	//最简结构声明
	b :=1
	fmt.Print("hello你好" + name + " -- ")
	fmt.Print(b)

}
