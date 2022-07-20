package robotsHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rpparas/flight_log/database"
	"github.com/rpparas/flight_log/internals/model"
)

// GetFlight func one robot by ID
// @Description Get one robot by ID
// @Tags Flight
// @Accept json
// @Produce json
// @Success 200 {object} model.Flight
// @router /api/v1/robot/{id} [get]
func GetRobot(c *fiber.Ctx) error {
	db := database.DB
	var robot model.Robot

	// Read the param flightsId
	id := c.Params("robotId")

	// Find the flights with the given Id
	db.Find(&robot, "id = ?", id)

	// If no such robot present return an error
	if robot.ID == uuid.Nil {
		// TODO: Determine appropriate error code, 204 or 404 instead of 200
		return c.Status(200).JSON(fiber.Map{"status": "no content", "message": "No robot found", "data": nil})
	}

	// Return the flights with the Id
	return c.JSON(fiber.Map{"status": "success", "message": "Robot Found", "data": robot})
}

// GetRobot func one robot by ID
// @Description Get one robot by ID
// @Tags Robot
// @Accept json
// @Produce json
// @Success 200 {object} model.Robot
// @router /api/robots/ [get]
func GetRobots(c *fiber.Ctx) error {
	db := database.DB
	var robots []model.Robot

	// TODO: support params for filtering robots

	db.Find(&robots)

	if len(robots) == 0 {
		// TODO: Determine appropriate error code, 204 or 404 instead of 200
		return c.Status(200).JSON(fiber.Map{"status": "no content", "message": "No robots present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Robots Found", "data": robots})
}
