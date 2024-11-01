package routers

import (
	"example.com/go-ecommerce-backend-api/internal/routers/manager"
	"example.com/go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User user.UserRouterGroup
	Manager manager.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)