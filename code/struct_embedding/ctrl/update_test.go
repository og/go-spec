package emailCtrl_test

import (
	emailS "github.com/og/golang-spec/code/struct_embedding"
	emailCtrl "github.com/og/golang-spec/code/struct_embedding/ctrl"
	"testing"
)

func TestUpdate(t *testing.T){
	emailCtrl.Update(emailCtrl.ReqUpdate{
		DataUpdate: emailS.DataUpdate{
			ID: "1",
			DataForm: emailS.DataForm{
				Title: "abc",
				Content: "This is a text",
			},
		},
	})
}