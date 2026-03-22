package auth

import "github.com/Thanawat0107/app-online-shop/internal/app/user"

type authGoogleUsecaseImpl struct {
	userRepo user.UserRepository
}

func NewAuthGoogleUsecase(userRepo user.UserRepository) AuthGoogleUsecase {
	return &authGoogleUsecaseImpl{
		userRepo: userRepo,
	}
}

func (u *authGoogleUsecaseImpl) UserLogin(userReq *UserLoginRequest) error {
	user := &user.UserEntity{
		UserId:   userReq.ID,
		FullName: userReq.FullName,
		Email:    userReq.Email,
		Picture:  userReq.Picture,
	}

	_, err := u.userRepo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *authGoogleUsecaseImpl) UserExists(userId string) bool {
	user, err := u.userRepo.FindById(userId)
	if err != nil {
		return false
	}
	return user != nil
}
