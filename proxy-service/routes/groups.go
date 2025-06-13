package routes

import "github.com/gofiber/fiber/v2"

type RouteGroupFunc func(router fiber.Router)

var RouteGroups = []struct {
	Path    string
	Handler RouteGroupFunc
}{
	{Path: "/auth", Handler: AuthRoutes},
	{Path: "/user", Handler: UserRoutes},
}
