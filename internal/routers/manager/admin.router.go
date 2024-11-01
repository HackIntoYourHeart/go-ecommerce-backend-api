package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (pr *UserRouter) InitAdminRouter(Router *gin.RouterGroup){
	// public router
	adminRouterPublic := Router.Group("/admin") 
	{
		adminRouterPublic.POST("/login")
	}

	// private router
	adminRouterPrivate := Router.Group("/admin/admin")
	// adminRouterPrivate.Use(limiter())
	// adminRouterPrivate.Use(Authen())
	// adminRouterPrivate.Use(Permission())
	{
		adminRouterPrivate.POST("/active_user")
	}
}