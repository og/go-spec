package dict

import "testing"

func ExampleModelNewsCreate() {
	ServiceNewsCreate(QueryNewCreate{
		Range: QueryNewCreate{}.Dict().Range.Wechat,
		Title: "a",
		Mobile: "13888888888",
	})
}

func TestModelNewsCreate(t *testing.T) {
	ExampleModelNewsCreate()
}