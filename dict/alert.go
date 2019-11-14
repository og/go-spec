package dict

import "log"

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
