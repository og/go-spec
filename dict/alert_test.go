package dict

import "testing"

func ExampleAlert(){
	Alert(AlertData{
		Type: AlertData{}.Dict().Type.Danger,
		Msg: "user id is empty",
	})
}

func TestAlert(t *testing.T) {
	ExampleAlert()
}