package routers

import (
	"hello-world-routing/handlers"

	"github.com/labstack/echo"
)

func InitRoutes(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	{
		v1.POST("/members", handlers.NewMemberHandler().PostMember())
		v1.GET("/members", handlers.NewMemberHandler().GetMembers())
		v1.GET("/members/:id", handlers.NewMemberHandler().GetMember())
	}
}
