package main

import (
	"ecommerce_api/app"
	"ecommerce_api/controller"
	"ecommerce_api/helper"
	"ecommerce_api/model"
	"ecommerce_api/repository"
	"ecommerce_api/service"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := app.DBConnection()
	r := echo.New()

	tokenUseCase := helper.NewTokenUseCase()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, tokenUseCase)
	userController := controller.NewUserController(userService)

	r.Debug = true
	r.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
	r.HTTPErrorHandler = helper.BindAndValidate

	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	// Grup Route dengan middleware role "seller"
	buyerGroup := r.Group("/buyer")
	buyerGroup.Use(helper.RoleMiddleware(tokenUseCase, "buyer"))
	buyerGroup.GET("/:id", userController.GetUser)

	r.POST("/register", userController.SaveUser)

	r.GET("/users", userController.GetUserList)
	r.PUT("/user/:id", userController.UpdateUser, JWTProtection())
	r.DELETE("/user/:id", userController.DeleteUser, JWTProtection())
	r.GET("/user/deleted/:id", userController.GetUserDeleted)

	r.POST("/login", userController.LoginUser)
	r.Logger.Fatal(r.Start(":8080"))

}

func JWTProtection() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helper.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, "anda harus login", nil))
		},
	})
}
