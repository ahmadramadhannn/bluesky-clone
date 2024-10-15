package user

import "context"

type userService struct {
	userRepository UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{
		userRepository: repo,
	}
}

func (us *userService) GetUserByID(ctx context.Context, id int) (username string, err error) {
	return us.userRepository.GetUserByID(ctx, id)
}
