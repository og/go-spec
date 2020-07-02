# golang spec

## 参数嵌套



## dict

> 当结构体的字段是选项时需定义对应的 dict

增加 dict 能让调用函数时候明确可选参数，不需要写文档说明有哪些参数。
也不需要记忆参数。能避免一些写错字符串的低级错误。
可以将 dict 理解成选项，函数实现放需要定义参数结构体，参数结构体在有选项字段时候应该增加  Dict() 方法，以此提供一种选项，而不是让调用函数方查看文档复制固定的内容传参。

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

----

参数嵌套时的字典 

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

**严格字典**

不限定的情况下


```go
type OneRange struct {
	Type string
	Start time.Time
	End time.Time
}
func (self OneRange) Dict() (dict struct{
	Type struct{
		Year string
		Month string
		Day string
	}
}) {
	dict.Type.Year = "year"
	dict.Type.Month = "month"
	dict.Type.Day = "day"
	return
}
```

限定的情况
```go
type rangeType struct {
	value string
}
type TwoRange struct {
	Type rangeType
	Start time.Time
	End time.Time
}
func (self TwoRange) Dict() (dict struct{
	Type struct{
		Year rangeType
		Month rangeType
		Day rangeType
	}
}) {
	dict.Type.Year = rangeType{"year"}
	dict.Type.Month = rangeType{"month"}
	dict.Type.Day = rangeType{"day"}
	return
}
```

如果 struct 需要跟请求所绑定（类似 json.Unmarshal 的操作），则 Range.Type应该是个 string 。如果不需要绑定则可以使用严格字典。这样可以去报使用者必须通过 dict 赋值 type

```go
TwoRange{
	Type: Range{}.Dict().Type.Day,
	Start: gtime.Parse(Day, "2018-01-01")),
	End:   gtime.Parse(Day, "2018-01-05")),
}
```

因为如果是 string, 有时候我们会为了偷懒不通过dict赋值type 比如：

```go
query := struct {
	Type string `json:"type"`
}
gjson.Parse(reuqest.Body, &query)
OneRange{
	Type: query.Type,
	Start: gtime.Parse(Day, "2018-01-01")),
	End:   gtime.Parse(Day, "2018-01-05")),
}
```
这种情况下如果自己约定了 query.type 必须是 "day" "month" "year"，并且一切按约定运行则没问题。但如果 Range 改了字典值，会导致请 query 与 range不一致出现无法预料的结果。所以如果你的结构体不需要 json xml 等做 decode 操作则通过让字段是私有结构体，并且只能通过字典获取私有结构体来避免因为”偷懒“导致的埋雷。

如果字典有多个字段，应该设置不同的类型

```go
type dateTypeDict struct {v string}
type dateStatusDict struct {v string}
type Date struct {
	Type dateTypeDict
	Status dateStatusDict
}
func (self Date) Dict() (dict struct {
	Type struct {
		Month dateTypeDict
		Day dateTypeDict
	}
	Status struct{
		Pass dateStatusDict
		Fail dateStatusDict
	}
}) {
	dict.Type.Month = dateTypeDict{"month"}
	dict.Type.Day = dateTypeDict{"day"}
	dict.Status.Pass = dateStatusDict{"pass"}
	dict.Status.Fail = dateStatusDict{"fail"}
	return
}
```

## 值继承的字典
```go
type User struct {
	ID string
	Name string
	Kind string
}
func (User) Dict() (dict struct {
	Kind struct{
		Normal string
		VIP string
	}
}) {
	dict.Kind.Normal = "normal"
	dict.Kind.VIP = "vip"
	return
}
type Group struct {
	ID string
	Name string
	UserKind string
}
func (Group) Dict()(dict struct{
	Kind struct{
		Normal string
		VIP string
	}
}){
	dict.Kind = User{}.Dict().Kind
	return 
}
```

> 匿名结构体可以赋值具名结构体

## part model

当 sql 只查出部分字段时，应当建立新的结构体作为 model 返回，而不是公用字段完整的 model 结构体。并且注意函数命名。

有时我们只需要 sql 查询部分字段。请看下面的代码

```go
type User struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Age      int    `db:"age"`
	Integral int	`db:"integral"`
}
// 不好的示例
func ListPartUser() (userList []User) {
	db.Select(&userList, `SELECT id, name from user`)
	return userList
}
```

`ListPartUser()` 返回的 `[]User` 中只有 `id` `name` 有数据。

调用方可能使用 `ListPartUser()[0].Age`，此时 `age` 是[zero value](https://studygolang.com/articles/15145?fr=sidebar) 使用它是不"安全"的。

并且在 `ListPartUser()` 内部增加或删除字段 

```diff
- SELECT id,name FROM user
+ SELECT id FROM user
``` 

 `ListPartUser()` 的调用方并不知道没有了 `name` 必须修改 `ListPartUser()` 的人肉检查所有调用方。

我们想利用类型系统去检查就需要定义出一个新的类型




```go

type User struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Age      int    `db:"age"`
	Integral int	`db:"integral"`
}
type PartUser_ID_Name struct {
	ID string
	Name string
}
func ListPartUser_ID_Name() (partUserList []PartUser_ID_Name) {
	userList := []User{}
	db.Select(&userList, `SELECT id, name FROM user`)
	for _, user := range userList {
		partUserList = append(partUserList, PartUser_ID_Name{
			ID: user.ID,
			Name: user.Name,
		})
	}
	return partUserList
}
// 单个字段直接使用 ListUser{字段} 并返回 []字段类型
func ListUserID() (userIDList []string) {
	userList := []User{}
	db.Select(&userIDList, `SELECT id FROM user`)
	for _, user := range userList {
		userIDList = append(userIDList, user.ID)
	}
	return userIDList
}
func main() {
	partUserList := ListPartUser_ID_Name()
	log.Print(partUserList[0].Name)
	// 此时使用 partUserList[0].Age 会报错，因为 PartUser_ID_Name 中没有 Age
}
```

此时使用变量 `partUserList` 会有"安全的"类型提示，不用担心字段是 zero value 。

注意如果在 `ListPartUser_ID_Name` 中新增了一个字段必须改变函数名，例如 新增 `age` 改名 `ListPartUser_ID_Name_Age`，接着又删除了 Name 改名 `ListPartUser_ID_Age`。

还有一种情况：最初有 `A()` `B()` `C()` 三个函数都调用了 `ListPartUser_ID_Name()` ，随着业务的变化 `C()` 需要调用一个新的函数 `ListPartUser_ID_Name_Age()` ,此时不应该将 `ListPartUser_ID_Name()` 名和内部实现修改为 `ListPartUser_ID_Name_Age`，而是要新增一个 `ListPartUser_ID_Name_Age()`。因为 `A()` `B()` 只需要 `id` 和 `name` 。（不要因为"偷懒"降低代码可维护性和性能）

## 封装switch

一段不好的代码

```go
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
```

虽然 `Bad()` 中的 `switch` 使用了 `dict`防止写错单词，但是如果 `Res` 的 `Type` 新增了一个 `Wait` 字段,需要检查所有 `switch res.Type` 的地方.
一旦项目多处 `switch res.Type` 很容易遗漏。

```go
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
```

SwitchType 函数中的 _Pass bool 和 _Fail string 之所有是第一个是 bool 第二个是 string 是
为了防止调用 SwitchType 时写错回调函数顺序

当 Type 非常多时可以通过如下顺序配置回调函数

```
int
bool
string
[]int
[]bool
[]string
map[int]int
map[int]bool
map[int]string
```

顺序按照字数少的类型到字数大的， 只使用 int bool string 作为类型， 逐渐增加 slice map

```go
func (res Res) SwitchType(
    Pass func (_Pass int),
    Fail func (_Fail bool),
    Wait func(_Wait string),
    Danger func(_Danger []int),
) {}
```