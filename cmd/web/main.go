package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"social-media/internal/config"
	http_delivery "social-media/internal/delivery/http"
	"social-media/internal/delivery/http/route"
	"social-media/internal/repository"
	"social-media/internal/use_case"
)

func main() {
	fmt.Println("Web started.")

	errEnvLoad := godotenv.Load()
	if errEnvLoad != nil {
		panic(fmt.Errorf("error loading .env file: %w", errEnvLoad))
	}

	envConfig := config.NewEnvConfig()
	databaseConfig := config.NewDatabaseConfig(envConfig)

<<<<<<< HEAD
	searchRepository := repository.NewSearchRepository(databaseConfig)
	userRepository := repository.NewUserRepository(databaseConfig)

	userUseCase := use_case.NewUserUseCase(userRepository)
	searchUseCase := use_case.NewSearchUseCase(searchRepository)

	userController := http_delivery.NewUserController(userUseCase, searchUseCase)
=======
	userRepository := repository.NewUserRepository(databaseConfig)
	repositoryConfig := config.NewRepositoryConfig(
		userRepository,
	)

	userUseCase := use_case.NewUserUseCase(repositoryConfig)
	useCaseConfig := config.NewUseCaseConfig(
		userUseCase,
	)

	userController := http_delivery.NewUserController(useCaseConfig)
	controllerConfig := config.NewControllerConfig(
		userController,
	)
>>>>>>> 558e4ba51ee27b0b0fb2cb0334760a582a7e6f35

	router := mux.NewRouter()
	userRoute := route.NewUserRoute(router, controllerConfig)
	rootRoute := route.NewRootRoute(
		router,
		userRoute,
	)

	rootRoute.Register()

	address := fmt.Sprintf(
		"%s:%s",
		envConfig.App.Host,
		envConfig.App.Port,
	)
	listenAndServeErr := http.ListenAndServe(address, rootRoute.Router)
	if listenAndServeErr != nil {
		panic(listenAndServeErr)
	}

	fmt.Println("Web finished.")
}
