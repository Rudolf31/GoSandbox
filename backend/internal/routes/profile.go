package routes

import (
	customeerrors "interface_lesson/internal/customerrors"
	"interface_lesson/internal/middleware"
	"interface_lesson/internal/models/dto"
	"interface_lesson/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func profileRoutes(router *gin.Engine, profileService services.ProfileService, authService services.AuthService) {
	profileGroup := router.Group("/profile")
	profileGroup.Use(middleware.NewAuthMiddleware(authService))

	profileGroup.POST("/", func(c *gin.Context) { createProfile(c, profileService) })
	profileGroup.GET("/:id", func(c *gin.Context) { getProfile(c, profileService) })
	profileGroup.PUT("/:id", func(c *gin.Context) { updateProfile(c, profileService) })
	profileGroup.PATCH("/:id", func(c *gin.Context) { patchProfile(c, profileService) })
	profileGroup.DELETE("/:id", func(c *gin.Context) { deleteProfile(c, profileService) })

}

// @Summary 	Create a new profile
// @Tags 		Profile
// @Accept		json
// @Produce		json
// @Param		request	body		dto.NewProfileDTO	true	"New profile"
// @Success      200              {int}    "ok"
// @Failure		400 "Bad Request"
// @Router		/profile [post]
func createProfile(c *gin.Context, profileService services.ProfileService) {
	var profile dto.NewProfileDTO

	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := profileService.CreateProfile(profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": *result})
}

// @Summary 	Get a profile
// @Tags 		Profile
// @Accept		json
// @Produce		json
// @Param		id	path		int	true	"Profile id"
// @Success		200 {object} dto.ProfileDTO "ok"
// @Failure		400 "Bad Request"
// @Failure		404 {object} customerrors.Wrapper "Profile not found"
// @Router		/profile/{id} [get]
func getProfile(c *gin.Context, profileService services.ProfileService) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	result, serviceErr := profileService.GetProfile(int32(id))
	if serviceErr != nil {
		switch serviceErr.Error {
		case customeerrors.ErrNotFound:
			c.JSON(http.StatusNotFound, serviceErr)
		default:
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, *result)
}

// @Summary 	Update a profile
// @Tags 		Profile
// @Accept		json
// @Produce		json
// @Param		id	path		int	true	"Profile id"
// @Param		request body dto.NewProfileDTO true "New profile"
// @Success		200 {object} dto.ProfileDTO "ok"
// @Failure		400 "Bad Request"
// @Failure		404 {object} customerrors.Wrapper "Profile not found"
// @Router		/profile/{id} [put]
func updateProfile(c *gin.Context, profileService services.ProfileService) {
	var profile dto.NewProfileDTO

	idStr := c.Param("id")
	id, errId := strconv.Atoi(idStr)
	if errId != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, serviceErr := profileService.UpdateProfile(int32(id), profile)
	if serviceErr != nil {
		switch serviceErr.Error {
		case customeerrors.ErrNotFound:
			c.JSON(http.StatusNotFound, serviceErr)
		default:
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, *result)
}

// @Summary 	Patch a profile
// @Tags 		Profile
// @Accept		json
// @Produce		json
// @Param		id	path		int	true	"Profile id"
// @Param		request body dto.PatchProfileDTO true "New profile"
// @Success		200 {object} dto.ProfileDTO "ok"
// @Failure		400 "Bad Request"
// @Failure		404 {object} customerrors.Wrapper "Profile not found"
// @Router		/profile/{id} [patch]
func patchProfile(c *gin.Context, profileService services.ProfileService) {
	var profile dto.PatchProfileDTO

	idStr := c.Param("id")
	id, errId := strconv.Atoi(idStr)
	if errId != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, serviceErr := profileService.PatchProfile(int32(id), profile)
	if serviceErr != nil {
		switch serviceErr.Error {
		case customeerrors.ErrNotFound:
			c.JSON(http.StatusNotFound, serviceErr)
		default:
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, *result)
}

// @Summary 	Delete a profile
// @Tags 		Profile
// @Accept		json
// @Produce		json
// @Param		id	path		int	true	"Profile id"
// @Success		200
// @Failure		400 "Bad Request"
// @Failure		404 {object} customerrors.Wrapper "Profile not found"
// @Router		/profile/{id} [delete]
func deleteProfile(c *gin.Context, profileService services.ProfileService) {
	id, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if serviceErr := profileService.DeleteProfile(int32(id)); serviceErr != nil {
		switch serviceErr.Error {
		case customeerrors.ErrNotFound:
			c.JSON(http.StatusNotFound, serviceErr)
		default:
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.Status(http.StatusOK)
}
