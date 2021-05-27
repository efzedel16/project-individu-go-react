package routes

import (
	"github.com/gin-gonic/gin"
	"silih_a3/auth"
	"silih_a3/config"
	"silih_a3/handlers"
	"silih_a3/middleware"
	"silih_a3/repositories"
	"silih_a3/services"
)

var (
	userDb         = config.ConnectDB()
	userRepository = repositories.NewUserRepository(userDb)
	userService    = services.NewUserService(userRepository)
	authService    = auth.NewAuthService()
	userHandler    = handlers.NewUserHandler(userService, authService)
)

func UserRoute(r *gin.Engine) {
	r.GET("/users", userHandler.ShowAllUsers)

	user := r.Group("/users")
	user.GET("/:id", middleware.AuthMiddleware(authService, userService), userHandler.ShowUser)
	user.POST("/signup", userHandler.SignUpUser)
	user.POST("/signin", userHandler.SignInUser)
	user.POST("/email_checker", userHandler.CheckEmailAvailability)
	user.POST("/avatar", middleware.AuthMiddleware(authService, userService), userHandler.UploadAvatar)

}
