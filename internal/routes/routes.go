package routes

import (
	dbconnection "event-system-backend/internal/db"
	authcontroller "event-system-backend/pkg/controller/authentication"
	eventcontroller "event-system-backend/pkg/controller/event"
	healthcheckcontroller "event-system-backend/pkg/controller/health_check"
	usercontroller "event-system-backend/pkg/controller/user"
	"event-system-backend/pkg/handler"
	eventrepository "event-system-backend/pkg/repository/event"
	userrepository "event-system-backend/pkg/repository/user"
	authenticationservice "event-system-backend/pkg/service/authentication"
	eventservice "event-system-backend/pkg/service/event"
	userservice "event-system-backend/pkg/service/user"
	"fmt"
	"log"
	"os"
	"strings"

	_ "event-system-backend/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run() {
	r := gin.Default()
	r.Use(handler.ErrorHandler())

	// swagger API
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	allowOrigins := os.Getenv("CORS_ALLOW_ORIGINS")
	allowMethods := os.Getenv("CORS_ALLOW_METHOD")

	// cors configuration
	config := cors.Config{
		AllowOrigins:     strings.Split(allowOrigins, ","),
		AllowMethods:     strings.Split(allowMethods, ","),
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}
	r.Use(cors.New(config))

	db := dbconnection.GetDB()
	userRepository := userrepository.NewUserRepository(db)
	eventRepository := eventrepository.NewEventRepository(db)

	userService := userservice.NewUserService(userRepository)
	eventService := eventservice.NewEventService(userService, eventRepository)
	authenticationService := authenticationservice.NewAuthenticationService(userService)

	eventController := eventcontroller.NewEventController(eventService)
	authController := authcontroller.NewAuthenticationController(authenticationService)
	userController := usercontroller.NewUserController(userService)

	api := r.Group("/api")
	healthcheckcontroller.SetupHealthCheckRoutes(api)
	eventcontroller.SetupEventRoutes(api, eventController)
	authcontroller.SetupAuthenticationRoutes(api, authController)
	usercontroller.SetupUserRoutes(api, userController)

	port := os.Getenv("PORT")
	log.Printf("Application running on port : %s", port)
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to run backend")
	}
}
