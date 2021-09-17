package users

import "AwsServerLessCleanCodeArchitecture/entity"

type UseCase interface {
	Signup(input entity.User) error
	SignIn(input entity.User) (token string, err error)
	SayHello(input entity.User) (message string, err error)
}
