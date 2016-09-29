package controllers

import (
	"github.com/kataras/iris"
	"github.com/itsbalamurali/parse-server/database"
)

type UserAPI struct {
	*iris.Context
}

func (u UserAPI) Signup(ctx *iris.Context)  {

	Db := database.MgoDb{}
	Db.Init()
	Db.Close()
}

func (u UserAPI) Login(ctx *iris.Context) {

	Db := database.MgoDb{}
	Db.Init()
	Db.Close()
}

func (u UserAPI) Logout(ctx *iris.Context) {


	Db := database.MgoDb{}
	Db.Init()
	Db.Close()
}

func (u UserAPI) Get(ctx *iris.Context) {

	Db := database.MgoDb{}
	Db.Init()
	Db.Close()
}

func (u UserAPI) Me(ctx *iris.Context) {

	Db := database.MgoDb{}
	Db.Init()
	Db.Close()
}

func (u UserAPI) Update(ctx *iris.Context)  {


	Db := database.MgoDb{}
	Db.Init()
	Db.Close()
}

func (u *UserAPI) GetAll(ctx *iris.Context) {

	Db := database.MgoDb{}
	Db.Init()
	Db.Close()
}

func (u *UserAPI) Delete(ctx *iris.Context){

	Db := database.MgoDb{}
	Db.Init()
	Db.Close()
}

func (u *UserAPI) ResetPassword(ctx *iris.Context) {

	Db := database.MgoDb{}
	Db.Init()
	Db.Close()
}

/*
func (u *UserAPI) IsAuthenticated() bool {
}

func (u *UserAPI) GetSessionToken() string {
}

func (u *UserAPI) setSessionToken(token string) {
}

func (u *UserAPI) clearSession() {
}
*/