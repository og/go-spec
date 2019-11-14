package dict


type QueryNewCreate struct {
	Range string
	Title string
	Mobile string
}
func (self QueryNewCreate) Dict() (dict struct{
	Range struct {
		Wechat string
		All string
	}
}) {
	// 一般情况下可以直接指向到 News{}.Dict()
	dict.Range = News{}.Dict().Range
	return
}
func ServiceNewsCreate(query QueryNewCreate) {
	ModelNewsCreate(News{
		Range: query.Range,
		Title: query.Title,
	})
	NewsPublishNotice(query.Mobile)
}
func NewsPublishNotice (mobile string) {
	// send sms
}
type News struct {
	Range string `db:"range"`
	Title string `db:"title"`
}
func (news News) Dict() (dict struct{
	Range struct {
		Wechat string
		All string
	}
}) {
	dict.Range.Wechat = "wechat"
	dict.Range.All = "all"
	return
}
func ModelNewsCreate (news News) {

}