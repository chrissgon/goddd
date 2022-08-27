package user

type Repo interface {
	Add(*Entity)
	Get(string) Entity
	GetAll() Entities
}
