package app

import "github.com/chrissgon/goddd/entities/user"

type UserApp struct {
	repo user.Repo
}

func (app *UserApp) Add(e *user.Entity) {
	app.repo.Add(e)
}

func (app *UserApp) Get(ID string) user.Entity {
	return app.repo.Get(ID)
}

func (app *UserApp) GetAll() user.Entities {
	return app.repo.GetAll()
}

func NewUserApp(repo user.Repo) *UserApp {
	return &UserApp{repo: repo}
}
