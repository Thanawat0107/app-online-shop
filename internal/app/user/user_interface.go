package user

type UserRepository interface {
	Create(user *UserEntity) (*UserEntity, error)
	FindById(userId string) (*UserEntity, error)
}
