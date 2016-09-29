package main

import (
	"github.com/kataras/iris"
	//"github.com/iris-contrib/middleware/logger"
	//"github.com/iris-contrib/middleware/recovery"
	"github.com/itsbalamurali/parse-server/database"
	"github.com/itsbalamurali/parse-server/controllers"
	"github.com/itsbalamurali/parse-server/config"
)

func main() {
	config.InitConfig()
	api := iris.New()
	// set the global middlewares
	//api.Use(logger.New())
	//api.Use(recovery.Handler)
	//Main Database Connection
	Db := database.MgoDb{}
	Db.Init()
	// index keys
	//keys := []string{"email"}
	///Db.Index("auth", keys)


	//API Version 1
	v1 := iris.Party("/1")

	//V1 Routes
	v1.Get("/swagger.json",controllers.SwaggerJSON) //Serves Swagger JSON API Spec

	//Classes
	class := new(controllers.ClassAPI)
	v1.Post("/classes/:className", class.Create)
	v1.Get("/classes/:className/:objectId",class.Get)
	v1.Put("/classes/:className/:objectId",class.Update)
	v1.Get("/classes/:className",class.GetAll)
	v1.Delete("/classes/:className/:objectId",class.Delete)

	//Users
	users := new(controllers.UserAPI)
	v1.Post("/users",users.Signup)
	v1.Get("/login",users.Login)
	v1.Post("/logout",users.Logout)
	v1.Get("/users/me",users.Me)
	v1.Get("/users/:objectId",users.Get)
	v1.Put("/users/:objectId",users.Update)
	v1.Delete("/users/:objectId",users.Delete)
	v1.Get("/users",users.GetAll)
	v1.Post("/requestPasswordReset",users.ResetPassword)
	/*
	//Sessions
	sessions := new(controllers.SessionAPI)

	//Roles
	roles := new(controllers.RoleAPI)

	//Files
	files := new(controllers.FileAPI)

	//Analytics
	event := new(controllers.EventAPI)

	//Push Notifications
	push := new(controllers.PushAPI)

	//Installations
	installations := new(controllers.InstallAPI)

	//Cloud Functions
	functions := new(controllers.FunctionAPI)

	//Schemas
	schema := new(controllers.SchemaAPI)

	//Apps
	apps := new(controllers.AppAPI)

	v1.Post("/apps/:appId")
	v1.Get("/apps/:appId")
	v1.Put("/apps/:appId")
	v1.Get("/apps")
	*/
	//Function Hooks
	//Trigger Hooks
	api.StaticWeb("/","./public",0)
	api.Listen(":8080")
}
