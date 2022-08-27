package router

import (
	"fmt"

	"github.com/chrissgon/goddd/interfaces"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func Start(port string) {
	m := martini.Classic()

	m.Use(render.Renderer())

	m.Group("/users", func(r martini.Router) {
		r.Post("", interfaces.User{}.Add)
		r.Get("/:ID", interfaces.User{}.Get)
		r.Get("", interfaces.User{}.GetAll)
	})

	m.RunOnAddr(fmt.Sprintf(":%s", port))
}
