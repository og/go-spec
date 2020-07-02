package main

import (
	"errors"
	"log"
)

type Res struct {
	Type string
}
func (Res) Dict() (dict struct {
	Type struct {
		Pass string
		Fail string
	}
}) {
	dict.Type.Pass = "pass"
	dict.Type.Fail = "fail"
	return
}
func Bad(res Res) {
	dict := res.Dict().Type
	switch res.Type {
	default:
		panic(errors.New("type error"))
	case dict.Pass:
		log.Print("pass")
	case dict.Fail:
		log.Printf("fail")
	}
}

_ = `
虽然 Bad() 中的 switch 使用了 dict，但是如果 Res 的 Type 新增了一个 Wait 字段,需要检查所有 switch res.Type 的地方.
一旦项目多处 switch res.Type 很容易遗漏。
`

func (v Res) SwitchType(
	Pass func (_Pass int),
	Fail func (_Fail bool),
	) {
	dict := v.Dict().Type
	switch v.Type {
	default:
		panic(errors.New("Res Type can not match(" + v.Type +")"))
	case dict.Pass:
		Pass(1)
	case dict.Fail:
		Fail(true)
	}
}
/*
通过封装 SwitchType 并使用回调函数代替 case ,
SwtichType 内部用switch 实现，当 Res 新增 Type 时候只需要在
SwitchType 修改即可，并且参数新增函数就能在编译器自动检查所有 switch res.Type 的地方
利用 IDE 的自动补全能够快速写出 SwitchType 调用代码
 */
func Good(res Res) {
	res.SwitchType(
	   func(_Pass int) {
		log.Print("pass")
	}, func(_Fail bool) {
		log.Print("fail")
	})
}

_ = `
SwitchType 函数中的 _Pass bool 和 _Fail string 之所有是第一个是 bool 第二个是 string 是
为了防止调用 SwitchType 时写错回调函数顺序

当 Type 非常多时可以通过如下顺序配置回调函数

int
bool
string
[]int
[]bool
[]string
map[int]int
map[int]bool
map[int]string

顺序按照字数少的类型到字数大的， 只使用 int bool string 作为类型， 逐渐增加 slice map

Pass func (_Pass int),
Fail func (_Fail bool),
Wait func(_Wait string),
Danger func(_Danger []int),


`