package users

import "AwsServerLessCleanCodeArchitecture/entity"

type Repository interface {
	GetUser(username string) (entity.User, error)
	AddUser(user entity.User) error
}
