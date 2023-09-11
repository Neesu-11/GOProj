package controller

import (
	"fmt"
	database "goproj/databse"
	"goproj/models"
	"log"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

//validate email

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)

}

//do the user registration

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	//check if password is right

	if len(data["password"].(string)) <= 6 {

		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password must be greater than 6 character",
		})
	}

	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "invalid email address",
		})
	}

	//check if email already exist in db

	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)

	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "email already exist",
		})
	}

	user := models.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string), 
		Phone:     data["phone"].(string),
		EmailId:     strings.TrimSpace(data["email"].(string)),
	}

	user.SetPassword(data["password"].(string))

	err:=database.DB.Create(&user) 
	if err != nil {
		log.Println(err)
		
	} 
		c.Status(200)
		return c.JSON(fiber.Map{

			"user": user,
			"message": "Account Created Successfully",
		})
	
}
