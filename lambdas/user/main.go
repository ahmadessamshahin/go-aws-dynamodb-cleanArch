package main

import (
	"AwsServerLessCleanCodeArchitecture/api/handler"
	ddb "AwsServerLessCleanCodeArchitecture/repository/dynamodb"
	"AwsServerLessCleanCodeArchitecture/usecase/users"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var (
	ginLambda *ginadapter.GinLambda
	srv       *users.Service
	secret    string
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := log.With().Logger()

	secret = os.Getenv("JWT_SECRET")
	if secret == "" {
		logger.Panic().
			Msg("jwt secret environment varialbe is missing")
	}

	sess, err := session.NewSession()
	if err != nil {
		logger.Panic().
			Err(err).
			Msg("unable to create aws session")
	}

	repo := ddb.NewDynamoDB(dynamodb.New(sess))
	srv = users.LoadService(repo, &logger)

	gin.SetMode(gin.ReleaseMode)
	r := handler.NewGinHandler(srv, secret)
	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
