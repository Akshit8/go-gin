// Package api defines endpoints for api
// impls endpoint documentation
package api

import (
	"net/http"

	"github.com/Akshit8/go-gin/controller"
	"github.com/Akshit8/go-gin/entity"
	"github.com/gin-gonic/gin"
)

// VideoAPI holds all app controllers
type VideoAPI struct {
	loginController controller.LoginController
	videoController controller.VideoController
}

// NewVideoAPI inits new video api instance
func NewVideoAPI(
	loginController controller.LoginController,
	videoController controller.VideoController,
) *VideoAPI {
	return &VideoAPI{
		loginController: loginController,
		videoController: videoController,
	}
}

// Paths Information

// Authenticate godoc
// @Summary Provides a JSON Web Token
// @Description Authenticates a user and provides a JWT to Authorize API calls
// @ID Authentication
// @Consume application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} entity.JWT
// @Failure 401 {object} entity.Response
// @Router /auth/token [post]
func (api *VideoAPI) Authenticate(ctx *gin.Context) {
	token := api.loginController.Login(ctx)
	if token != "" {
		ctx.JSON(http.StatusOK, &entity.JWT{
			Token: token,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, &entity.Response{
			Message: "Not Authorized",
		})
	}
}

// AllVideos godoc
// @Security bearerAuth
// @Summary List existing videos
// @Description Get all the existing videos
// @Tags videos,list
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Video
// @Failure 401 {object} entity.Response
// @Router /videos [get]
func (api *VideoAPI) AllVideos(ctx *gin.Context) {
	ctx.JSON(200, api.videoController.FindAll())
}

// GetVideo godoc
// @Security bearerAuth
// @Summary return a single videos
// @Description Get a video for a id
// @Tags videos
// @Accept  json
// @Produce  json
// @Param  id path int true "Video ID"
// @Success 200 {object} entity.Video
// @Failure 401 {object} entity.Response
// @Router /videos/{id} [get]
func (api *VideoAPI) GetVideo(ctx *gin.Context) {
	video, err := api.videoController.Get(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, &entity.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"video": video,
		})
	}
}

// CreateVideo godoc
// @Security bearerAuth
// @Summary Create new videos
// @Description Create a new video
// @Tags videos,create
// @Accept  json
// @Produce  json
// @Param video body entity.Video true "Create video"
// @Success 200 {object} entity.Response
// @Failure 400 {object} entity.Response
// @Failure 401 {object} entity.Response
// @Router /videos/ [post]
func (api *VideoAPI) CreateVideo(ctx *gin.Context) {
	err := api.videoController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "saved video successfully",
		})
	}
}

// UpdateVideo godoc
// @Security bearerAuth
// @Summary Update videos
// @Description Update a single video
// @Security bearerAuth
// @Tags videos
// @Accept  json
// @Produce  json
// @Param  id path int true "Video ID"
// @Param video body entity.Video true "Update video"
// @Success 200 {object} entity.Response
// @Failure 400 {object} entity.Response
// @Failure 401 {object} entity.Response
// @Router /videos/{id} [put]
func (api *VideoAPI) UpdateVideo(ctx *gin.Context) {
	err := api.videoController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, &entity.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"video": "video updated successfully",
		})
	}
}

// DeleteVideo godoc
// @Security bearerAuth
// @Summary Remove videos
// @Description Delete a single video
// @Security bearerAuth
// @Tags videos
// @Accept  json
// @Produce  json
// @Param  id path int true "Video ID"
// @Success 200 {object} entity.Response
// @Failure 400 {object} entity.Response
// @Failure 401 {object} entity.Response
// @Router /videos/{id} [delete]
func (api *VideoAPI) DeleteVideo(ctx *gin.Context) {
	err := api.videoController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, &entity.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"video": "video deleted successfully",
		})
	}
}
