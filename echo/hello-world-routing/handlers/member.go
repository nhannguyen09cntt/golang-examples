package handlers

import (
	"encoding/json"

	"github.com/labstack/echo"
)

type MemberHandler struct {
	BaseHandler
}

func NewMemberHandler() *MemberHandler {
	return &MemberHandler{}
}

func (h *MemberHandler) PostMember() echo.HandlerFunc {
	return func(c echo.Context) error {
		msg := Message{Message: "POST MEMBER"}
		c.Response().Header().Set("Server", "4Rum")
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		return json.NewEncoder(c.Response()).Encode(msg)
	}
}

func (h *MemberHandler) GetMembers() echo.HandlerFunc {
	return func(c echo.Context) error {
		msg := Message{Message: "GetMembers"}
		c.Response().Header().Set("Server", "4Rum")
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		return json.NewEncoder(c.Response()).Encode(msg)
	}
}

func (h *MemberHandler) GetMember() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		msg := Message{Message: "Member: " + id}
		c.Response().Header().Set("Server", "4Rum")
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		return json.NewEncoder(c.Response()).Encode(msg)
	}
}
