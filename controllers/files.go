package controllers

import (
	"github.com/goamz/goamz/s3"
	"github.com/itsbalamurali/parse-server/models"
	"github.com/kataras/iris"
	"io/ioutil"
)

type FileAPI struct {
	*iris.Context
	bucket *s3.Bucket
}

/*
InvalidFileName	122	An invalid filename was used for Parse File. A valid file name contains only a-zA-Z0-9_. characters and is between 1 and 128 characters.
MissingContentType	126	Missing content type.
MissingContentLength	127	Missing content length.
InvalidContentLength	128	Invalid content length.
FileTooLarge	129	File size exceeds maximum allowed.
FileSaveError	130	Error saving a file.
FileDeleteError	131	File could not be deleted.
*/

func (f *FileAPI) Upload(ctx *iris.Context) {
	if ctx.PostBody() == nil {
		ctx.JSON(iris.StatusBadRequest, models.Error{Message: "InvalidContentLength", Code: 128})
	}

	if ctx.Request.Header.ContentLength() == nil {
		ctx.JSON(iris.StatusBadRequest, models.Error{Message: "MissingContentLength", Code: 128})
	}

	if ctx.Request.Header.ContentType() == nil {
		ctx.JSON(iris.StatusBadRequest, models.Error{Message: "MissingContentType", Code: 128})
	}

	filename := ctx.Param("name")

	content, err := ioutil.ReadAll(ctx.PostBody())

	if err != nil {
		ctx.JSON(iris.StatusBadRequest, models.Error{Message: "FileSaveError", Code: 130})
	}

	err = f.bucket.Put(filename, content, ctx.Request.Header.ContentType(), s3.PublicRead)

	if err != nil {
		ctx.JSON(iris.StatusBadRequest, models.Error{Message: "FileSaveError", Code: 130})
	}

	url := f.bucket.URL(filename)
	ctx.Response.Header.Set("Location", url)
	ctx.JSON(iris.StatusCreated, iris.Map{"url": url, "name": filename})
}

func (f *FileAPI) Delete(ctx *iris.Context) {
	filename := ctx.Param("name")
	f.bucket.Del(filename)
	ctx.Response.Header.StatusCode(204)
}
