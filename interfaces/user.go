package interfaces

import (
	"net/http"
	"time"

	"github.com/chrissgon/goddd/app"
	"github.com/chrissgon/goddd/entities/user"
	"github.com/go-martini/martini"
	"github.com/google/uuid"
	"github.com/goombaio/namegenerator"
	"github.com/martini-contrib/render"
)

type User struct{}

func (User) Add(render render.Render, r *http.Request) {
	e := user.Entity{
		ID:   uuid.New().String(),
		Name: namegenerator.NewNameGenerator(time.Now().UTC().UnixNano()).Generate(),
	}

	app.NewUserApp(user.NewPostgresUserRepo()).Add(&e)

	render.JSON(201, e)
}

func (User) Get(render render.Render, r *http.Request, params martini.Params) {
	e := app.NewUserApp(user.NewPostgresUserRepo()).Get(params["ID"])
	render.JSON(200, e)
}

func (User) GetAll(render render.Render, r *http.Request) {
	es := app.NewUserApp(user.NewPostgresUserRepo()).GetAll()
	render.JSON(200, es)
}
