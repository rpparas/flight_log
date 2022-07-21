package robotsHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rpparas/flight_log/database"
	"github.com/rpparas/flight_log/internals/model"
)

// CreateRobot func create a robot
// @Description Create a robot
// @Tags Robot
// @Accept json
// @Produce json
// @Success 200 {object} model.Robot
// @router /api/robot [post]
func CreateRobot(c *fiber.Ctx) error {
	db := database.DB
	robot := new(model.Robot)

	// Store the body in the robots and return error if encountered
	err := c.BodyParser(robot)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	robot.ID = uuid.New()
	// Create the Robot and return error if encountered
	err = db.Create(&robot).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create robots", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created Robot", "data": robot})
}
