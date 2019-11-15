package emailCtrl_test

import (
	emailS "github.com/og/golang-spec/code/struct_embedding"
	emailCtrl "github.com/og/golang-spec/code/struct_embedding/ctrl"
	"testing"
)

func TestCreate(t *testing.T) {
	emailCtrl.Create(emailCtrl.ReqCreate{
		DataCreate: emailS.DataCreate{
			DataForm: emailS.DataForm{
				Title: "abc",
				Content: "this is a text",
			},
		},
	})
}
