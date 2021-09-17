package users

import (
	"AwsServerLessCleanCodeArchitecture/entity"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const encodingString = "t0k3n"

func (service *Service) Signup(input entity.User) error {

	//validate user not exist
	details, err := service.Repository.GetUser(input.Username)

	if err != nil {
		service.Logger.Error().Msgf("failed to get user with error: %s", err)
		return err
	}

	if details.Username == input.Username {
		return entity.ErrUserAlreadyExists
	}

	// hash password
	hashed, err := HashAndSalt(input.Password)
	if err != nil {
		service.Logger.Error().Msgf("failed to encrypt password with error %s", err)
		return err
	}

	input.Password = hashed

	err = service.Repository.AddUser(input)
	if err != nil {
		service.Logger.Error().
			Msgf("failed to add user with error: %s", err)
		return err
	}
	return nil
}

func (service *Service) SignIn(input entity.User) (token string, err error) {
	details, err := service.Repository.GetUser(input.Username)

	if err != nil {
		service.Logger.Error().Msgf("failed to get user with error: %s", err)
		return token, err
	}
	if details.Username == "" {
		return token, entity.ErrUserDoesNotExist
	}
	if !ComparePasswords(details.Password, input.Password) {
		return token, entity.ErrUserDoesNotExist
	}
	token, err = CreateToken(details.Username)
	if err != nil {
		service.Logger.Error().
			Msgf("failed to create token with error: %s", err)
		return token, err
	}
	return token, nil
}

func (service *Service) SayHello(input entity.User) (message string, err error) {
	details, err := service.Repository.GetUser(input.Username)
	if err != nil {
		service.Logger.Error().
			Msgf("failed to retrieve user data: %s", err)
		return message, err
	}

	message = fmt.Sprintf("Hello %s", details.FirstName)

	return message, nil
}
func CreateToken(userName string) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	now := time.Now().Local()
	token.Claims = jwt.MapClaims{
		"username": userName,
		"iat":      now.Unix(),
		"exp":      now.Add(time.Hour * time.Duration(1)).Unix(),
	}
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(encodingString))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ComparePasswords(hashed, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	if err != nil {
		return false
	}
	return true
}
func HashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
