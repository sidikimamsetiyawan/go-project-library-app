package controller

import (
	"fmt"
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
	UserName string `json:"user_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"` // Role field for registration
}

type LoginRequest struct {
	Identifier string `json:"identifier" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

type UpdateRequest struct {
	Email    string `json:"email" validate:"omitempty,email"` // Optional, must be a valid email
	Password string `json:"password" validate:"omitempty"`    // Optional
}

func Register(c *fiber.Ctx) error {
	var request RegisterRequest

	fmt.Println("request", request)

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
		UserName:     request.UserName,
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
	// Determine if identifier is an email or user_name
	isEmail := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(request.Identifier)

	var user model.Users
	if isEmail {
		// If identifier is an email
		if err := db.Where("email = ?", request.Identifier).First(&user).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
	} else {
		// If identifier is a user_name
		if err := db.Where("user_name = ?", request.Identifier).First(&user).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
	}

	// Check if password matches
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid password"})
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.UserID
	claims["user_name"] = user.UserName
	claims["email"] = user.Email
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not login"})
	}

	return c.JSON(fiber.Map{"token": tokenString})
}

func UpdateUser(c *fiber.Ctx) error {
	// Extract user ID from the URL or context (depending on your auth system)
	userID := c.Params("id")

	db := database.DBConn

	// Get the existing user
	var user model.Users
	if err := db.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	var request UpdateRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Validate the incoming request fields
	if err := validate.Struct(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Update email if provided and validate uniqueness
	if request.Email != "" && request.Email != user.Email {
		// Check if the new email is already taken by another user
		var existingUser model.Users
		if err := db.Where("email = ?", request.Email).First(&existingUser).Error; err == nil {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Email is already taken"})
		}
		user.Email = request.Email
	}

	// Update password if provided
	if request.Password != "" {
		if !isValidPassword(request.Password) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Password must be 8 characters, alphanumeric, contain at least 1 uppercase letter, and no special characters"})
		}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
		user.Password = string(hashedPassword)
	}

	// Update modified date and modified_by
	user.ModifiedDate = time.Now()
	user.ModifiedBy = "system" // Replace with the current user's info if available

	// Save changes to the database
	if err := db.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not update user"})
	}

	return c.JSON(fiber.Map{"message": "User updated successfully"})
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
