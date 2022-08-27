package infra

type FakePostgresRepo struct{}

var database  = map[string]string{}

func (FakePostgresRepo) Add(key, value string) {
	database[key] = value
}

func (FakePostgresRepo) Get(key string) string {
	return database[key]
}

func (FakePostgresRepo) GetAll() map[string]string {
	return database
}

func NewFakePostgresRepo() *FakePostgresRepo {
	return &FakePostgresRepo{}
}
