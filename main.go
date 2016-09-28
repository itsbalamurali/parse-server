package main

import "github.com/kataras/iris"

func main() {
	api := iris.New()
	//API Version 1
	v1 := iris.Party("/1")

	//V1 Routes
	v1.Get("/")

	//Classes
	v1.Post("/classes/:className")
	v1.Get("/classes/:className/:objectId")
	v1.Put("/classes/:className/:objectId")
	v1.Get("/classes/:className")
	v1.Delete("/classes/:className/:objectId")

	//Users
	v1.Post("/users")
	v1.Get("/login")
	v1.Post("/logout")
	v1.Get("/users/me")
	v1.Get("/users/:objectId")
	v1.Put("/users/:objectId")
	v1.Delete("/users/:objectId")
	v1.Get("/users")
	v1.Post("/requestPasswordReset")

	//Sessions
	//Roles
	//Files
	//Analytics
	//Push Notifications
	//Installations
	//Cloud Functions
	//Schemas
	//Apps
	v1.Post("/apps/:appId")
	v1.Get("/apps/:appId")
	v1.Put("/apps/:appId")
	v1.Get("/apps")

	//Function Hooks
	//Trigger Hooks



	api.Get("/hi")
	api.Listen(":8080")
}
