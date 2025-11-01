package routes

import (
	customeerrors "interface_lesson/internal/customeErrors"
	"interface_lesson/internal/models/dto"
	"interface_lesson/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func profileRoutes(router *gin.Engine, profileService services.ProfileService) {
	profileGroup := router.Group("/profile")
	{
		profileGroup.POST("/", func(c *gin.Context) { createProfile(c, profileService) })
		profileGroup.GET("/:id", func(c *gin.Context) { getProfile(c, profileService) })
		profileGroup.PUT("/:id", func(c *gin.Context) { updateProfile(c, profileService) })
		profileGroup.PATCH("/:id", func(c *gin.Context) { patchProfile(c, profileService) })
		profileGroup.DELETE("/:id", func(c *gin.Context) { deleteProfile(c, profileService) })
	}
}

// TODO: write a doc for swagger
func createProfile(c *gin.Context, profileService services.ProfileService) {
	var profile dto.NewProfileDTO

	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := profileService.CreateProfile(profile)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": *result})
}

// TODO: write a doc for swagger
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

// TODO: write a doc for swagger
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

// TODO: write a doc for swagger
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

// TODO: write a doc for swagger
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
