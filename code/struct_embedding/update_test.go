package emailS

import "testing"

func TestUpdate(t *testing.T) {
	Update(DataUpdate{
		ID: "",
		DataForm: DataForm{
			Title: "",
			Content: "",
		},
	})
}
