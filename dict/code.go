package dict

import (
	"encoding/json"
	"log"
)

type Code struct {
	value string
}

func CodeDict() (dict struct{
	NotLogin Code
	PasswordError Code
}) {
	dict.NotLogin = Code{"notLogin"}
	dict.PasswordError = Code{"passwordError"}
	return
}
func Fail(code Code) {
	jsons , err := json.Marshal(map[string]interface{}{
		"code": code.value,
	})
	if err != nil { panic(err) }
	log.Print(jsons)
}
