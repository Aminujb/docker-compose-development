package main

import (
	"daily-standup/config"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"

	_entity "daily-standup/app/entity"
	_handler "daily-standup/app/handler"
	_repo "daily-standup/app/repository"

	"github.com/spf13/viper"

)

func viperConfigVariable(key string) string {

	// name of config file (without extension)
	viper.SetConfigName("config")
	// look for config in the working directory
	viper.AddConfigPath("./")
  
	// Find and read the config file
	err := viper.ReadInConfig()
  
	if err != nil {
	  log.Fatalf("Error while reading config file %s", err)
	}
  
	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	
	value, ok := viper.Get(key).(string)
  
	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
	  log.Fatalf("Invalid type assertion")
	}
	return value
}

func sanityCheck() {
	envProps := []string{
		"STANDUP_COLLECTION",
		"DATABASE_NAME",
		"DATABASE_ADDRESS",
		"PORT",
	}

	for _, k := range envProps {
		if viperConfigVariable(k) == "" {
			log.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}

	fmt.Println("Environment variables successfully loaded. Starting application...")
}

func main() {
	sanityCheck()
	// gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.TestMode)

	app := gin.Default()



	DATABASE_NAME, _ := viper.Get("DATABASE_NAME").(string)
	DATABASE_ADDRESS, _ := viper.Get("DATABASE_ADDRESS").(string)
	STANDUP_COLLECTION, _ := viper.Get("STANDUP_COLLECTION").(string)

	standupCollection := config.GetEntityDbCollection(DATABASE_NAME, DATABASE_ADDRESS, STANDUP_COLLECTION)

	repo := _repo.NewReportRepository(standupCollection)
	entityPost := _entity.NewReportEntity(repo)

	api := app.Group("/api/v1")

	_handler.NewReportHandler(api, entityPost)

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:4001"}

	app.Use(cors.New(config))

	//Port
	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))
	vPORT, _ := viper.Get("PORT").(string)

    if PORT == ":" {
        PORT = vPORT
    } 

	app.Run(PORT)
}
