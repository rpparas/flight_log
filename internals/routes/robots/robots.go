package robots

import (
	"github.com/gofiber/fiber/v2"
	robotsHandler "github.com/rpparas/flight_log/internals/handlers/robots"
)

func SetupApiRoutes(router fiber.Router) {
	// TODO: use singular vs plural
	robot := router.Group("/robots")
	robot.Get("/", robotsHandler.GetRobots)
	// TODO: need post request for multiple robots

	robots := router.Group("/robot")
	robots.Get("/:robotId", robotsHandler.GetRobots)
	robot.Post("/", robotsHandler.CreateRobot)
}
