package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	errors "github.com/tavomartinez88/marvel/internal/error"
	"github.com/tavomartinez88/marvel/internal/utils"
	"github.com/tavomartinez88/marvel/pkg/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	if os.Getenv(utils.ApiKey) == "" || os.Getenv(utils.PrivateApiKey) == ""  {
		log.Fatalf("Must set the environment variables %v and %v", utils.ApiKey, utils.PrivateApiKey)
	}

	handlerCharacter := handlers.NewCharacterHandler()
	handlerCollaborator := handlers.NewCollaboratorHandler()

	router := gin.Default()
	marvelRoute := router.Group("/marvel")
	{
		marvelRoute.GET("/collaborators/:name", func(context *gin.Context) {
			name := context.Param("name")
			response, err := handlerCollaborator.GetCollaborators(name)

			if err != nil {
				errClient := errors.ClientError{}
				_ = json.Unmarshal([]byte(err.Error()), &errClient)
				context.JSON(errClient.HttpStatus, err)
				return
			}

			context.JSON(http.StatusOK, response)
		})

		marvelRoute.GET("/characters/:name", func(context *gin.Context) {
			name := context.Param("name")
			response, err := handlerCharacter.GetCharacters(name)

			if err != nil {
				errClient := errors.ClientError{}
				_ = json.Unmarshal([]byte(err.Error()), &errClient)
				context.JSON(errClient.HttpStatus, err)
				return
			}

			context.JSON(http.StatusOK, response)
		})
	}

	_ = router.Run()
}
