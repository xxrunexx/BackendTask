package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id    int    `json:"id" form:"id"`
	Age   int    `json:"age" form:"age"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

// type UserTemp struct {
// 	Name  string `json:"name" form:"name"`
// 	Email string `json:"email" form:"email"`
// }

func main() {
	e := echo.New()

	// Routing
	// If user request url "/user", then run HelloController
	e.GET("/user/:id/:age", UserController)
	e.POST("/user", RegisterController)
	e.Start(":8000")
}

// Controller need 1 param echo.Context
// e echo.Context can recieved data from client
// Client will send the data (url path || query || form) to echo.Context
// echo.Context also can be use to return response to client.
func UserController(e echo.Context) error {
	// Url Param
	id, _ := strconv.Atoi(e.Param("id"))   // Store input to variable Id using e.Param
	age, _ := strconv.Atoi(e.Param("age")) // Param always return string, AtoI to conv to Int
	// Query Param
	search := e.QueryParam("search")
	sort := e.QueryParam("sort")

	HarunArRasyid := User{
		Id:    id,
		Name:  "Harun Ar Rasyid",
		Email: "rasyid.id3@gmail.com",
		Age:   age}
	// String() return 2 values, code & text
	return e.JSON(http.StatusOK, map[string]interface{}{
		"user":   HarunArRasyid,
		"search": search,
		"sort":   sort,
	})
}

// Footnote :
// Query Param : Using "?" symbol, (format using key & value)
// Url Param : single URL

func RegisterController(e echo.Context) error {
	// get data from value
	// name := e.FormValue("name") Commented to add binding data
	// email := e.FormValue("email") Commented to add binding data

	user := User{}
	e.Bind(&user)

	return e.JSON(http.StatusOK, map[string]interface{}{
		// "email": email,
		// "name":  name,
		"messages": "Successfully Creating New User",
		"user":     user,
	})
}
