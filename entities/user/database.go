package user

import "github.com/chrissgon/goddd/infra"

type PostgresUserRepo struct {
	DB *infra.FakePostgresRepo
}

func (repo *PostgresUserRepo) Add(e *Entity) {
	repo.DB.Add(e.ID, e.Name)
}

func (repo *PostgresUserRepo) Get(ID string) Entity {
	return Entity{
		ID:   ID,
		Name: repo.DB.Get(ID),
	}
}

func (repo *PostgresUserRepo) GetAll() Entities {
	var Users Entities

	for name, ID := range repo.DB.GetAll() {
		Users = append(Users, Entity{ID: ID, Name: name})
	}

	return Users
}

func NewPostgresUserRepo() Repo {
	return &PostgresUserRepo{DB: infra.NewFakePostgresRepo()}
}
