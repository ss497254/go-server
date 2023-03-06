package routes

import (
	"server/database"
	"server/entities"
	"server/lib"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	json := new(LoginRequest)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "invalid data received",
		})
	}

	result := entities.User{}
	query := entities.User{Email: json.Email}

	err := database.DB.First(&result, &query).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "user not found",
			"error":   err,
		})
	}

	if !lib.ComparePasswords(result.Password, []byte(json.Password)) {
		return c.JSON(fiber.Map{
			"code":    401,
			"message": "Invalid Password",
		})
	}

	lib.SendAccessToken(jwt.MapClaims{"userId": result.UserId}, c)

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
		"userId":  result.UserId,
	})
}

func Register(c *fiber.Ctx) error {
	type CreateUserRequest struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	json := new(CreateUserRequest)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	password := lib.HashAndSalt([]byte(json.Password))

	new := entities.User{
		// UserId:    guuid.New(),
		Email:     json.Email,
		FirstName: json.FirstName,
		LastName:  json.LastName,
		Password:  password,
	}

	found := entities.User{}
	query := entities.User{Email: json.Email}

	err := database.DB.First(&found, &query).Error
	if err != gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "user already exists",
		})
	}
	database.DB.Create(&new)

	lib.SendAccessToken(jwt.MapClaims{"userId": new.UserId}, c)

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
		"userId":  new.UserId,
	})
}
