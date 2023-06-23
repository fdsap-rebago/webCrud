package controller

import (
	"fmt"
	"web_crud/middleware"
	"web_crud/model"

	"github.com/gofiber/fiber/v2"
)

func ViewRegistration(c *fiber.Ctx) error {
	return c.Render("registration", fiber.Map{
		"Title":         "Registration",
		"HeaderMessage": "Welcome to Fiber Template",
	})
}

func ViewLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title":         "Login",
		"HeaderMessage": "Welcome to Fiber Template",
	})
}

func CreateUser(c *fiber.Ctx) error {
	userCredentials := &model.Users{}
	c.BodyParser(userCredentials)
	middleware.DBConn.Debug().Exec("INSERT INTO users (fullname, username, password) VALUES (?,?,?)", userCredentials.Fullname, userCredentials.Username, userCredentials.Password).Find(userCredentials)
	return c.Render("login", fiber.Map{
		"Title":         "Login",
		"HeaderMessage": "Welcome to Fiber Template",
		"Status":        true,
	})
}

func VerifyAccount(c *fiber.Ctx) error {
	fmt.Println("VERIFY ACCOUNT")
	userCredentials := &model.Users{}
	users := []model.Users{}
	if parsErr := c.BodyParser(userCredentials); parsErr != nil {
		fmt.Println("Error", parsErr)
	}

	isFound := middleware.DBConn.Debug().Raw("SELECT * FROM users WHERE username=? AND password=?", userCredentials.Username, userCredentials.Password).Find(&userCredentials).RowsAffected
	if isFound > 0 {
		middleware.DBConn.Debug().Raw("SELECT * FROM users ").Find(&users)
		return c.Render("dashboard", fiber.Map{
			"Data":          users,
			"Title":         "Dashboard",
			"HeaderMessage": "Welcome to Fiber Template",
			"SweetMessage":  "Login successful.",
			"Status":        true,
		})
	} else {
		return c.Render("login", fiber.Map{
			"Title":         "Login",
			"HeaderMessage": "Welcome to Fiber Template",
			"SweetMessage":  "Invalid account, please try again.",
			"Status":        false,
		})
	}
}

func ViewDashboard(c *fiber.Ctx) error {

	return c.Render("dashboard", fiber.Map{
		"Title":         "Dashboard",
		"HeaderMessage": "Welcome to Fiber Template",
	})
}

func GetUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	fmt.Println("ID", userId)
	return nil
}

func GetUsers(c *fiber.Ctx) error {
	users := &model.Users{}
	middleware.DBConn.Debug().Raw("SELECT * FROM users").Find(users)
	return c.JSON(users)
}
