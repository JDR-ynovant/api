package repository

type UserRepository struct {
	tableName string
}

func NewUserRepository() UserRepository {
	return UserRepository{
		tableName: "users",
	}
}

func (UserRepository) FindById(id string) interface{} {
	panic("implement me")
}
