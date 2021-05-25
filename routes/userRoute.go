package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"silih_a3/auth"
	"silih_a3/config"
	"silih_a3/user"
)

var (
	DB             *gorm.DB = config.ConnectDB()
	userRepository          = user.NewUserRepository(DB)
	userService             = user.NewUserService(userRepository)
	authService             = auth.NewService()
	userHandler             = user.NewUserHandler(userService, authService)
)

func UserRoute(r *gin.Engine) {
	//r.GET("/users", auth.AuthMiddleware(userService, authService), u)

	//user := r.Group("/users")
	{
		r.GET("/users/:id", userHandler.GetUser)
		r.POST("/users/signup", userHandler.SignUpUser)
		r.POST("/users/signin", userHandler.SignInUser)
		r.POST("/users/email_checker", userHandler.CheckEmailAvailability)
		r.POST("/users/avatar", auth.AuthMiddleware(authService, userService), userHandler.UploadAvatar)
	}
}
