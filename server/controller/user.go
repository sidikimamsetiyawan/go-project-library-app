package controller

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sidikimamsetiyawan/go-project-library-app/database"
	"github.com/sidikimamsetiyawan/go-project-library-app/model"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()
var jwtSecret = []byte("your_jwt_secret_key")

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"` // Role field for registration
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Register(c *fiber.Ctx) error {
	var request RegisterRequest

	db := database.DBConn

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Validate request fields
	if err := validate.Struct(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Validate password requirements
	if !isValidPassword(request.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Password must be 8 characters, alphanumeric, contain at least 1 uppercase letter, and no special characters"})
	}

	// Check if email already exists
	var user model.Users
	if err := db.Where("email = ?", request.Email).First(&user).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Email already registered"})
	}

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)

	// Default role if not specified
	if request.Role == "" {
		request.Role = "user"
	}

	// Create new user
	newUser := model.Users{
		Email:        request.Email,
		Password:     string(hashedPassword),
		Role:         request.Role,
		CreatedBy:    "system", // Default to system if not provided
		CreatedDate:  time.Now(),
		ModifiedBy:   "system", // Default to system if not provided
		ModifiedDate: time.Now(),
	}

	if err := db.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not register user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

func Login(c *fiber.Ctx) error {
	var request LoginRequest

	db := database.DBConn

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Validate request fields
	if err := validate.Struct(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Check if user exists
	var user model.Users
	if err := db.Where("email = ?", request.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Check if password matches
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid password"})
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.UserID
	claims["email"] = user.Email
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not login"})
	}

	return c.JSON(fiber.Map{"token": tokenString})
}

func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return false
	}
	if !regexp.MustCompile(`[a-zA-Z0-9]+`).MatchString(password) {
		return false
	}
	if regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(password) {
		return false
	}
	return true
}
