package routes

import (
	dbconnection "event-system-backend/internal/db"
	authcontroller "event-system-backend/pkg/controller/authentication"
	eventcontroller "event-system-backend/pkg/controller/event"
	healthcheckcontroller "event-system-backend/pkg/controller/health_check"
	"event-system-backend/pkg/handler"
	eventrepository "event-system-backend/pkg/repository/event"
	userrepository "event-system-backend/pkg/repository/user"
	authenticationservice "event-system-backend/pkg/service/authentication"
	eventservice "event-system-backend/pkg/service/event"
	userservice "event-system-backend/pkg/service/user"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.Use(handler.ErrorHandler())

	db := dbconnection.GetDB()

	userRepository := userrepository.NewUserRepository(db)
	eventRepository := eventrepository.NewEventRepository(db)

	userService := userservice.NewUserService(userRepository)
	eventService := eventservice.NewEventService(userService, eventRepository)
	authenticationService := authenticationservice.NewAuthenticationService(userService)

	eventController := eventcontroller.NewEventController(eventService)
	authController := authcontroller.NewAuthenticationController(authenticationService)

	api := r.Group("/api")
	healthcheckcontroller.SetupHealthCheckRoutes(api)
	eventcontroller.SetupEventRoutes(api, eventController)
	authcontroller.SetupAuthenticationRoutes(api, authController)

	port := os.Getenv("PORT")
	log.Printf("Application running on port : %s", port)
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to run backend")
	}
}
