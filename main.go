// @APIVersion 1.0.0
// @APITitle Parse Server
// @APIDescription Parse Mobile Backend API.
// @Contact balamurali@live.com
// @License BSD
// @LicenseUrl http://opensource.org/licenses/BSD-2-Clause
// @BasePath http://0.0.0.0:8080/1/
package main

import (
	"github.com/kataras/iris"
	//"github.com/iris-contrib/middleware/logger"
	//"github.com/iris-contrib/middleware/recovery"
	"github.com/itsbalamurali/parse-server/database"
	"github.com/itsbalamurali/parse-server/controllers"
	"github.com/itsbalamurali/parse-server/config"
	"os"
)

func main() {
	config.InitConfig()
	api := iris.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// set the global middlewares
	//api.Use(logger.New())
	//api.Use(recovery.Handler)

	//Main Database Connection
	Db := database.MgoDb{}
	Db.Init()


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

	//Sessions
	sessions := new(controllers.SessionAPI)
	v1.Post("/sessions", sessions.Create)
	v1.Get("/sessions",sessions.GetAll)
	v1.Get("/sessions/me",sessions.Get)
	v1.Get("/sessions/:objectId",sessions.Get)
	v1.Put("/sessions/me",sessions.Update)
	v1.Put("/sessions/:objectId",sessions.Update)
	v1.Delete("/sessions/:objectId",sessions.Delete)

	//Roles
	roles := new(controllers.RoleAPI)
	v1.Post("/roles",roles.Create)
	v1.Get("/roles/:objectId",roles.Get)
	v1.Put("/roles/:objectId",roles.Update)
	v1.Delete("/roles/:objectId",roles.Delete)

	//Files
	files := new(controllers.FileAPI)
	v1.Post("/push",files.Upload)
	v1.Delete("/push",files.Delete)

	//Analytics
	event := new(controllers.EventAPI)
	v1.Post("/events/AppOpened",event.AppOpened)
	v1.Post("/events/:eventName",event.CustomEvent)

	//Push Notifications
	push := new(controllers.PushAPI)
	v1.Post("/push",push.Create)

	//Installations
	installations := new(controllers.InstallAPI)
	v1.Post("/installations/:objectId",installations.Create)
	v1.Get("/installations/:objectId",installations.GetByID)
	v1.Put("/installations/:objectId",installations.Update)
	v1.Get("/installations",installations.GetAll)
	v1.Post("/installations",installations.Create)

	//Cloud Functions
	functions := new(controllers.FunctionAPI)
	v1.Post("/functions/:name",functions.ExecuteFunc)
	v1.Post("/jobs/:name",functions.TriggerJob)

	//Schemas
	schema := new(controllers.SchemaAPI)
	v1.Post("/schemas/:className",schema.Create)
	v1.Get("/schemas/:className",schema.GetByID)
	v1.Put("/schemas/:className",schema.Update)
	v1.Get("/schemas",schema.GetAll)
	v1.Delete("/className/:className",schema.Delete)

	//Apps
	apps := new(controllers.AppAPI)
	v1.Post("/apps/:appId",apps.Create)
	v1.Get("/apps/:appId",apps.GetByID)
	v1.Put("/apps/:appId",apps.Update)
	v1.Get("/apps",apps.Get)

	//Config
	config := new(controllers.ConfigAPI)
	v1.Get("/config",config.Get)
	v1.Put("/config",config.Update)

	//Function Hooks
	hookfuncs := new(controllers.HFuncsAPI)
	v1.Post("/hooks/functions",hookfuncs.Create)
	v1.Get("/hooks/functions/:funcName",hookfuncs.GetByID)
	v1.Put("/hooks/functions/:funcName",hookfuncs.Update)
	v1.Delete("/hooks/functions/:funcName",hookfuncs.Delete)

	//Trigger Hooks
	hooktriggers := new(controllers.HTriggerAPI)
	v1.Post("/hooks/triggers",hooktriggers.Create)
	v1.Get("/hooks/triggers/:className/:triggerName",hookfuncs.GetByID)
	v1.Put("/hooks/triggers/:funcName/:triggerName",hookfuncs.Update)
	v1.Delete("/hooks/triggers/:funcName/:triggerName",hookfuncs.Delete)


	//Serve Swagger UI
	api.StaticWeb("/swagger","./public",0)
	api.Listen(":"+port)
}
