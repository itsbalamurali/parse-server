package middleware

import "github.com/kataras/iris"

const (
	MasterKey = "X-Parse-Master-Key"
	RestAPIKey = "X-Parse-REST-API-Key"
	JSKey = "X-Parse-Javascript-Key"
	AppID = "X-Parse-Application-Id"
)

func APIAuth(ctx *iris.Context){
	masterkey := ctx.RequestHeader(MasterKey)
	if masterkey != "" {
		ctx.Next()
	}else {
		ctx.JSON(403,iris.Map{"error":"unauthorized"})
	}
}
