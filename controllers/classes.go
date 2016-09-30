package controllers

import (
	"github.com/kataras/iris"
	"github.com/itsbalamurali/parse-server/database"
	"encoding/json"
)

type ClassAPI struct {
	*iris.Context
}

func (c *ClassAPI) Get(ctx *iris.Context) {
	//name := c.Param("name")
	Db := database.MgoDb{}
	Db.Init()

	Db.Close()
}

func (c *ClassAPI) GetAll(ctx *iris.Context) {

	Db := database.MgoDb{}
	Db.Init()

	Db.Close()
}


func (c ClassAPI) Create(ctx *iris.Context) {
	var i interface{}
	classname := ctx.Param("className")
	b := ctx.Request.Body()
	if err := json.Unmarshal(b, &i); err != nil {
		panic(err.Error())
	}
	Db := database.MgoDb{}
	Db.Init()
	// Insert
	if err := Db.C(classname).Insert(&i) ; err != nil {
		// Is a duplicate key, but we don't know which one
		//ctx.JSON(iris.StatusOK, models.Err("5"))
		if Db.IsDup(err) {
			//ctx.JSON(iris.StatusOK, models.Err("6"))
		}
	} else {
		ctx.JSON(iris.StatusOK, iris.Map{"response": true})
	}

	Db.Close()
}


func (c *ClassAPI) Update(ctx *iris.Context) {
	var i interface{}
	//classname := ctx.Param("className")
	b := ctx.Request.Body()
	if err := json.Unmarshal(b, &i); err != nil {
		panic(err.Error())
	}

	Db := database.MgoDb{}
	Db.Init()

	Db.Close()
}

func (c *ClassAPI) Delete(ctx *iris.Context)  {

	Db := database.MgoDb{}
	Db.Init()


	Db.Close()
}
