package dict


type QueryFrom struct {
	Status string
	Type string
	Title string
	Content string
}
func (self QueryFrom) Dict() (dict struct {
	Status struct {
		Normal string
		CheckPending string
	}
	Type struct {
		Exigency string
		Log string
	}
}) {
	dict.Status.Normal = "normal"
	dict.Status.CheckPending = "checkPending"

	dict.Type.Exigency = "exigency"
	dict.Type.Log = "log"
	return
}
type QueryCreate struct {
	QueryFrom
	ID string `db:"id"`
}

func Create(query QueryCreate) {

}
type QueryUpdate struct {
	QueryFrom
}
func Update(query QueryUpdate) {

}
