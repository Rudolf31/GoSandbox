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

	router.POST("/profile", func(c *gin.Context) {
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
	})

	router.GET("/profile/:id", func(c *gin.Context) {

		id_str := c.Param("id")
		id, err := strconv.Atoi(id_str)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		result, serviceErr := profileService.GetProfile(id)
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
	})

	router.PUT("/profile/:id", func(c *gin.Context) {

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

		result, serviceErr := profileService.UpdateProfile(id, profile)
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
	})

	router.PATCH("/profile/:id", func(c *gin.Context) {
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

		result, serviceErr := profileService.PatchProfile(id, profile)
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
	})

	router.DELETE("/profile/:id", func(c *gin.Context) {

		id, errId := strconv.Atoi(c.Param("id"))
		if errId != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		if serviceErr := profileService.DeleteProfile(id); serviceErr != nil {
			switch serviceErr.Error {
			case customeerrors.ErrNotFound:
				c.JSON(http.StatusNotFound, serviceErr)
			default:
				c.Status(http.StatusInternalServerError)
			}
			return
		}

		c.Status(http.StatusOK)
	})

}
