// router
package core

import (
	"birthday/form"
	"birthday/handler"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func initRouter(app *martini.ClassicMartini) {
	//404处理
	app.NotFound(func(r render.Render) {
		r.JSON(404, "Not Found")
		return
	})
	//公用模块
	app.Group("/comment", func(rt martini.Router) {
		rt.Post("/set", binding.Bind(form.RequestForm{}), handler.Set)
		rt.Get("/list", binding.Bind(form.RequestForm{}), handler.List)
		rt.Get("/get", binding.Bind(form.RequestForm{}), handler.Get)
		rt.Get("/node", binding.Bind(form.RequestForm{}), handler.Node)
		rt.Get("/deploy", binding.Bind(form.RequestForm{}), handler.Deploy)
	})

}
