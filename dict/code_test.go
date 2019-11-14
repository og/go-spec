package dict

import "testing"

func ExampleCodeDict() {
	Fail(CodeDict().PasswordError)
	// Fail("passwordError") // 此行代码会报错，因为必须通过 CodeDict() 返回的 CodeItem 结构体作为 Fail() 的第一个参数
}
func TestCodeDict(t *testing.T) {
	ExampleCodeDict()
}
