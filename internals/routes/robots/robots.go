package robots

import (
	"github.com/gofiber/fiber/v2"
	robotsHandler "github.com/rpparas/flight_log/internals/handlers/robots"
)

func SetupApiRoutes(router fiber.Router) {
	// Endpoints for robots
	robots := router.Group("/robots")
	robots.Post("/", robotsHandler.CreateRobot)
	robots.Get("/", robotsHandler.GetRobots)
	robots.Get("/:robotId", robotsHandler.GetRobots)

}
