package manager

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup){
	// public router
	// userRouterPublic := Router.Group("/admin") 
	// {
	// 	userRouterPublic.POST("/register")
	// 	userRouterPublic.POST("/otp")
	// }

	// private router
	userRouterPrivate := Router.Group("/admin/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.POST("/active_user")
	}
}