package robotsHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rpparas/flight_log/database"
	"github.com/rpparas/flight_log/internals/model"
)

// GetRobots func gets all existing robots
// @Description Get all existing robots
// @Tags Robots
// @Accept json
// @Produce json
// @Success 200 {array} model.Robot
// @router /api/robots [get]
func GetRobots(c *fiber.Ctx) error {
	db := database.DB
	var robots []model.Robot

	// find all robots in the database
	db.Find(&robots)

	// If no robots is present return an error
	if len(robots) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No robots present", "data": nil})
	}

	// Else return robots
	return c.JSON(fiber.Map{"status": "success", "message": "Robots Found", "data": robots})
}

// CreateRobot func create a robot
// @Description Create a robot
// @Tags Robots
// @Accept json
// @Produce json
// @Success 200 {object} model.Robot
// @router /api/robots [post]
func CreateRobot(c *fiber.Ctx) error {
	db := database.DB
	robots := new(model.Robot)

	// Store the body in the robots and return error if encountered
	err := c.BodyParser(robots)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Add a uuid to the robots
	robots.ID = uuid.New()
	// Create the Robot and return error if encountered
	err = db.Create(&robots).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create robots", "data": err})
	}

	// Return the created robots
	return c.JSON(fiber.Map{"status": "success", "message": "Created Robot", "data": robots})
}

// GetRobot func one robot by ID
// @Description Get one robots by ID
// @Tags Robot
// @Accept json
// @Produce json
// @Success 200 {object} model.Robot
// @router /api/robots/{id} [get]
func GetRobot(c *fiber.Ctx) error {
	db := database.DB
	var robots model.Robot

	// Read the param robotsId
	id := c.Params("robotsId")

	// Find the robots with the given Id
	db.Find(&robots, "id = ?", id)

	// If no such robots present return an error
	if robots.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No robots present", "data": nil})
	}

	// Return the robots with the Id
	return c.JSON(fiber.Map{"status": "success", "message": "Robots Found", "data": robots})
}
