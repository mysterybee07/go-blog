package controller

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	// "github.com/gofiber/fiber"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"github.com/mysterybee07/blogbackend/database"
	"github.com/mysterybee07/blogbackend/models"
	"github.com/mysterybee07/blogbackend/utils"
)

func validateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse")
	}
	//check if the user password is less than 6 characters
	if len(data["password"].(string)) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be at least 8 characters long",
		})
	}

	// if data["password"].(string) != data["confirm_password"].(string) {
	// 	c.Status(400)
	// 	return c.JSON(fiber.Map{
	// 		"message": "Passwords do not match",
	// 	})
	// }

	email := strings.TrimSpace(data["email"].(string))
	if !validateEmail(email) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid email",
		})
	}

	database.DB.Where("email = ?", email).First(&userData)
	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email is already taken",
		})
	}

	// hashedPassword, err := hashPassword(data["password"].(string))
	// if err != nil {
	// 	return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	// }

	user := models.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
		// Password:  hashedPassword,
		Phone: data["phone"].(string),
	}

	// if err := database.DB.Create(&user).Error; err != nil {
	// 	return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	// }
	user.SetPassword(data["password"].(string))
	err := database.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "Account created successfully",
	})
}

//	func hashPassword(password string) ([]byte, error) {
//		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
//		if err != nil {
//			return nil, err
//		}
//		return hashedPassword, nil
//	}
//
// Controller for Login
func Login(c *fiber.Ctx) error {
	var data map[string]string
	// var userData models.User
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	var user models.User
	database.DB.Where("email=?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Email Address is not found, Kindly create an account",
		})
	}
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid password",
		})
	}
	token, err := utils.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "You have successfully login",
	})
}

type claims struct {
	jwt.StandardClaims
}
