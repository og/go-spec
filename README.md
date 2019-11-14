# golang spec


## dict

当结构体的字段是选项时需定义对应的 dict

> 增加 dict 能让调用函数时候明确可选参数，不需要写文档说明有哪些参数。也不需要记忆参数。能避免一些写错字符串的低级错误。
> 可以将 dict 理解成选项，函数实现放需要定义参数结构体，参数结构体在有选项字段时候应该增加  Dict() 方法，以此提供一种选项，而不是让调用函数方查看文档复制固定的内容传参。

```go
type AlertData struct {
	Type string
	Msg string
}
func (self AlertData) Dict () (dict struct {
	Type struct {
		Danger string
		Info string
	}
}) {
	dict.Type.Danger = "danger"
	dict.Type.Info = "info"
	return
}
func Alert(data AlertData) {
	log.Printf("[%s] %s", data.Type, data.Msg)
	if data.Type == data.Dict().Type.Danger {
		log.Print("!!!!!!!!!!!!!!!!!!")
	}
}
```

```go
Alert(AlertData{
  Type: AlertData{}.Dict().Type.Danger,
  Msg: "user id is empty",
})
```

虽然
```go
Alert(AlertData{
  Type: "danger",
  Msg: "user id is empty",
})
```
也能运行，但是需要查看文档才知道能使用 `"danger"` ,并且如果手误写成 `"dange"` 只会在运行时候才能检测到错误。
如果写成了 `AlertData{}.Dict().Type.Dange` 在编译期就会报错。
当 Alert 函数不在支持 type=danger 时，在 `AlertData{}.Dict()` 中去掉 Danger ，就能在编译时发现一些错误。

```diff
func (self AlertData) Dict () (dict struct {
	Type struct {
		Danger string
		Info string
	}
}) {
-	dict.Type.Danger = "danger"
	dict.Type.Info = "info"
	return
}
```

```go
Alert(AlertData{
  Type: AlertData{}.Dict().Type.Danger, // 此行会报错，因为 Danger 已经被删除
  Msg: "user id is empty",
})
```  

---

当两个结构体共同继承同一个结构体自然就公用同一个字典


```go
type QueryFrom struct {
	Status string
	Type string
	Title string
	Content string
}
func (self QueryFrom) Dict() (dict struct {
	Status struct {
		Normal string
		CheckPending string
	}
	Type struct {
		Exigency string
		Log string
	}
}) {
	dict.Status.Normal = "normal"
	dict.Status.CheckPending = "checkPending"

	dict.Type.Exigency = "exigency"
	dict.Type.Log = "log"
	return
}
type QueryCreate struct {
	QueryFrom
	ID string `db:"id"`
}
```
```go
Create(QueryCreate{
	QueryFrom: QueryFrom{
		Status: QueryFrom{}.Dict().Status.Normal,
		Type: QueryFrom{}.Dict().Type.Exigency,
		Title: "a",
		Content:"some",
	},
	ID: "1",
})
Update(QueryUpdate{
	QueryFrom: QueryFrom{
		Status: QueryFrom{}.Dict().Status.Normal,
		Type: QueryFrom{}.Dict().Type.Exigency,
		Title: "a",
		Content:"some",
	},
})
```

在常见的 model + service 分层中即使字段是直接传递赋值的，也依然需要字典

```go
package dict


type QueryNewsCreate struct {
	Range string
	Title string
	Mobile string
}
func (self QueryNewsCreate) Dict() (dict struct{
	Range struct {
		Wechat string
		All string
	}
}) {
	// 一般情况下可以直接指向到 News{}.Dict()
	dict.Range = News{}.Dict().Range
	return
}
func ServiceNewsCreate(query QueryNewsCreate) {
	ModelNewsCreate(News{
		Range: query.Range,
		Title: query.Title,
	})
	NewsPublishNotice(query.Mobile)
}
func NewsPublishNotice (mobile string) {
	// send sms
}
type News struct {
	Range string `db:"range"`
	Title string `db:"title"`
}
func (news News) Dict() (dict struct{
	Range struct {
		Wechat string
		All string
	}
}) {
	dict.Range.Wechat = "wechat"
	dict.Range.All = "all"
	return
}
func ModelNewsCreate (news News) {

}
```

```go
ServiceNewsCreate(QueryNewsCreate{
	Range: QueryNewsCreate{}.Dict().Range.Wechat,
	Title: "a",
	Mobile: "13888888888",
})
```

只所有没有直接用 `News{}.Dict().Range.Wechat` 
是因为 `ServiceNewsCreate()` 可能会支持 `QueryNewsCreate{}.Dict().Range.All`

当只支持 `All` 时候只需将修改字典 

```go
func (self QueryNewsCreate) Dict() (dict struct{
	Range struct {
		All string
	}
}) {
	dict.Range.All = News{}.Dict().Range.All
	return
}
```

修改后所有调用 `QueryNewCreate{}.Dict().Range.Wechat` 在编译期就会报错，这样可以更好的重构和迭代函数。

在 `QueryNewCreate` 的字段年终中使用 `News{}.Dict()` 是错误的用法

```go
QueryNewsCreate{
	Range: News{}.Dict().Range.Wechat
}
```

虽然这样暂时也能运行，但是让 `QueryNewCreate` 与 `News` 在外部意外耦合了。（QueryNewCreate不期望调用方知道 News 的存在）
为了达到高内聚，低耦合。需要使用 `QueryNewsCreate{}.Dict().Range.Wechat`。

如果 `New` 可以直接当做 `QueryNewCreate` 的字段，（QueryNewCreate要求调用方知道 News 的存在）则可直接使用 `New{}.Dict()`

例如

```go
// 正确的用法
type QueryNewCreate struct {
	News
	Mobile string
}
QueryNewCreate{
	News News{
		Range: News{}.Dict().Range.Wechat
		title: "a"
	},
	Mobile: "13641822109",
}
```

---

还有一种情况是函数的调用方的参数都是通过代码生成，而不是外部请求。此时可使用更明确的字典。


```go

type Code struct {
	value string
}

func CodeDict() (dict struct{
	NotLogin Code
	PasswordError Code
}) {
	dict.NotLogin = Code{"notLogin"}
	dict.PasswordError = Code{"passwordError"}
	return
}
func Fail(code Code) {
	jsons , err := json.Marshal(map[string]interface{}{
		"code": code.value,
	})
	if err != nil { panic(err) }
	log.Print(jsons)
}
```

```go
Fail(CodeDict().PasswordError)
// Fail("passwordError") // 此行代码会报错，因为必须通过 CodeDict() 返回的 CodeItem 结构体作为 Fail() 的第一个参数
```

`Fail()` 函数不使用 `enum` 的原因是需求要求返回 `{"code":"notLogin"}` 这样的 json ，而不只是作为函数内部的判断条件变量。

dict 是利用类型系统增加代码可维护性方法，虽然看起来麻烦了，但实际上写完代码过一段时间再来看会发现使用 dict 的代码比没有使用 dict 的代码更易于维护。
