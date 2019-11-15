package emailCtrl

import emailS "github.com/og/golang-spec/code/struct_embedding"

type ReqUpdate struct {
	emailS.DataUpdate
}
func Update(req ReqUpdate) {

}