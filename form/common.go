// common
package form

type RequestForm struct {
	IndexId int64  `form:"index_id"`
	Name    string `form:"name"`
	Day     string `form:"day"`
	Message string `form:"message"`
	Skip    int    `form:"skip"`
	Limit   int    `form:"limit"`
}
