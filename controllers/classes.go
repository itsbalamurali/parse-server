package controllers

import (
	"github.com/kataras/iris"
	"github.com/itsbalamurali/parse-server/database"
)

type ClassAPI struct {
	*iris.Context
}

// Select gets class data information
func (c *ClassAPI) Get(ctx *iris.Context) {

	Db := database.MgoDb{}
	Db.Init()

	Db.Close()
}

// Select gets class data information by custom query
func (c *ClassAPI) GetAll(ctx *iris.Context) {

	Db := database.MgoDb{}
	Db.Init()

	Db.Close()
}

// Create creates class from data
func (c *ClassAPI) Create(ctx *iris.Context) {

	Db := database.MgoDb{}
	Db.Init()

	Db.Close()
}

// Update updates class by ID
func (c *ClassAPI) Update(ctx *iris.Context) {

	Db := database.MgoDb{}
	Db.Init()

	Db.Close()
}

// Delete deletes class by ID
func (c *ClassAPI) Delete(ctx *iris.Context)  {

	Db := database.MgoDb{}
	Db.Init()

	Db.Close()
}
