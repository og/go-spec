package dict

import "testing"

func ExampleModelNewsCreate() {
	ServiceNewsCreate(QueryNewsCreate{
		Range: QueryNewsCreate{}.Dict().Range.Wechat,
		Title: "a",
		Mobile: "13888888888",
	})
}

func TestModelNewsCreate(t *testing.T) {
	ExampleModelNewsCreate()
}