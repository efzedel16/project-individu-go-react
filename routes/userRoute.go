package routes

import (
	"github.com/gin-gonic/gin"
	"silih_a3/auth"
	"silih_a3/config"
	"silih_a3/handler"
	"silih_a3/middleware"
	"silih_a3/user"
)

var (
	DB             = config.ConnectDB()
	userRepository = user.NewUserRepository(DB)
	userService    = user.NewUserService(userRepository)
	authService    = auth.NewService()
	userHandler    = handler.NewUserHandler(userService, authService)
)

func UserRoute(r *gin.Engine) {
	r.GET("/users", userHandler.ShowAllUsers)

	user := r.Group("/users")
	user.GET("/:id", middleware.AuthMiddleware(authService, userService), userHandler.GetUser)
	user.POST("/register", userHandler.SignUpUser)
	user.POST("/login", middleware.AuthMiddleware(authService, userService), userHandler.SignInUser)
	user.POST("/email_checker", userHandler.CheckEmailAvailability)
	r.POST("/avatar", middleware.AuthMiddleware(authService, userService), userHandler.UploadAvatar)

}
