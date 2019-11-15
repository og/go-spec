package emailS

import "testing"

func TestCreate(t *testing.T) {
	Create(DataCreate{
		DataForm: DataForm{
			Title: "",
			Content: "",
		},
	})
}