package handlers

import (
	"github.com/fazilnbr/learn-docker/database"
	"github.com/fazilnbr/learn-docker/models"
	"github.com/gofiber/fiber/v2"
	_ "gorm.io/gorm"
)

// @title Go + Fiber True fact API
// @version 1.0
// @description This is a sample true fact api qustion and answer. You can visit the GitHub repository at https://github.com/fazilnbr/learn-docker

// @contact.name API Support
// @contact.url https://fazilnbr.github.io/mypeosolal.web.portfolio/
// @contact.email fazilkp2000@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:3000
// @BasePath /
// @query.collection.format multi

// @Summary List True Fact 	Q&A
// @ID listTrueFact
// @Tags APIs
// @Produce json
// @Success 200 {object} []models.Fact{}
// @Router / [get]
func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

// @Summary Create True Fact 	Q&A
// @ID createTrueFact
// @Tags APIs
// @Produce json
// @Param Fact body models.Fact{} true "admin login"
// @Success 200 {object} models.Fact{}
// @Failure 422 {object} models.Fact{}
// @Router /fact [post]
func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}
