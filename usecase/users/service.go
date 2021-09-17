package users

// The UseCase layer depend on the entity so ( if entity updated will lead to update the UseCase)

// But it will not affect any higher circle
import "github.com/rs/zerolog"

type Service struct {
	Repository Repository
	Logger     *zerolog.Logger
}

func LoadService(repository Repository, logger *zerolog.Logger) *Service {
	return &Service{
		Repository: repository,
		Logger:     logger,
	}
}
