package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hakuuww/go-gin/models"
	"sync"
)

var g *id_generator = &id_generator{}

type VideoController interface {
	GetAll(context *gin.Context)
	Update(context *gin.Context)
	Create(context *gin.Context)
	Delete(context *gin.Context)
}

type controller struct {
	videos []models.Video
}

type id_generator struct {
	counterId int
	mtx       sync.Mutex
}

// The use of defer for unlocking the mutex is actually a smart and easy way to ensure that we are always releasing this resource
// before the function exits.
// If we have a scenario where there are multiple return statements and and complex flow control inside a function
// then having a defer statement will make sure that we always unlock the mutex no matter where we are returning.
// otherwise we might forget to unlock the mutex before the different return statements
func (g *id_generator) getNextVideoId() int {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	g.counterId++
	return g.counterId
}

func NewVideoController() VideoController {
	return &controller{make([]models.Video,0 )}
}

func (c* controller) GetAll(context *gin.Context) {
	print("inside getall function")
	context.JSON(200, c.videos)
	return
}

func deleteElement(slice []models.Video, index int) []models.Video {

	if index < 0 || index >= len(slice) {
		return slice // Return the original slice if the index is out of bounds
	}

	return append(slice[:index], slice[index+1:]...)
}

func (c *controller) Update(context *gin.Context) {
	var newVideo models.Video

	//ShouldBind checks the Method and Content-Type to select a binding engine automatically, Depending on the "Content-Type" header different bindings are used, for example:
	// "application/json" --> JSON binding
	// "application/xml"  --> XML binding
	// It parses the request's body as JSON if Content-Type == "application/json" using JSON or XML as a JSON input.
	// It decodes the json payload into the struct specified as a pointer.
	// Like c.Bind() but this method does not set the response status code to 400 or abort if input is not valid.
	if err := context.ShouldBindUri(&newVideo); err != nil {
		context.String(400, "bad request, format does not match")
		return
	}

	if err := context.ShouldBindJSON(&newVideo); err != nil {
		context.String(400, "bad request, format does not match")
		return
	}

	//In a Go for range loop, the values returned in each iteration are passed by value, not by pointer.
	//This means that the video variable in your example is a copy of the value in the c.videos slice for each iteration.
	//Therefore in order to modify a specific value in the slice, we need to use the index to access it
	for idx, video := range c.videos {
		if newVideo.Id == video.Id {
			c.videos[idx] = newVideo
			context.String(200, "video with id %d has been updated to the db", newVideo.Id)
			return
		}
	}

	context.String(400, "bad request, cannot find video with id %d", newVideo.Id)
	return

	//c.videos = append(c.videos,newVideo)
}

func (c *controller) Create(context *gin.Context) {
	newVideo := models.Video{Id: g.getNextVideoId()}

	//BindJSON binds the passed struct pointer using the JSON binding engine.
	//It will abort the request with HTTP 400 if any error occurs.
	if err := context.BindJSON(&newVideo); err != nil {
		context.String(400, "Bad request %v, error in parsing data", err)
		return
	}

	c.videos = append(c.videos, newVideo)
	context.String(200, "Successfully added new video, new video id is %d", newVideo.Id)
	return
}

func (c *controller) Delete(context *gin.Context) {
	var videoToDelete models.Video

	if err := context.ShouldBindUri(&videoToDelete); err != nil {
		context.String(400, "bad request, format does not match. ERROR: %v", err)
		return
	}

	for idx, video := range c.videos {
		if videoToDelete.Id == video.Id {
			c.videos = deleteElement(c.videos, idx)
			context.String(200, "video with id %d has been removed from the db", videoToDelete.Id)
			return
		}
	}

	context.String(400, "Cannot find the video to delete, video id:%d", videoToDelete.Id)

	return

}
