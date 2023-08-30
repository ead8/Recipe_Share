package controllers

import (
	"gin-cloudinary-api/dtos"
	"gin-cloudinary-api/models"
	"gin-cloudinary-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MultipleFileUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["files"] // Assuming your HTML form field name is "files"

		var uploadUrls []string
		for _, file := range files {
			uploadedFile, err := file.Open()
			if err != nil {
				c.JSON(
					http.StatusInternalServerError,
					dtos.MediaDto{
						StatusCode: http.StatusInternalServerError,
						Message:    "error",
						Data:       map[string]interface{}{"data": "Error opening file"},
					})
				return
			}
			defer uploadedFile.Close()

			uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: uploadedFile})
			if err != nil {
				c.JSON(
					http.StatusInternalServerError,
					dtos.MediaDto{
						StatusCode: http.StatusInternalServerError,
						Message:    "error",
						Data:       map[string]interface{}{"data": "Error uploading file"},
					})
				return
			}
			uploadUrls = append(uploadUrls, uploadUrl)
		}

		c.JSON(
			http.StatusOK,
			dtos.MediaDto{
				StatusCode: http.StatusOK,
				Message:    "success",
				Data:       map[string]interface{}{"data": uploadUrls},
			})
	}
}


func RemoteUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		var url models.Url

		//validate the request body
		if err := c.BindJSON(&url); err != nil {
			c.JSON(
				http.StatusBadRequest,
				dtos.MediaDto{
					StatusCode: http.StatusBadRequest,
					Message:    "error",
					Data:       map[string]interface{}{"data": err.Error()},
				})
			return
		}

		uploadUrl, err := services.NewMediaUpload().RemoteUpload(url)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				dtos.MediaDto{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Error uploading file"},
				})
			return
		}

		c.JSON(
			http.StatusOK,
			dtos.MediaDto{
				StatusCode: http.StatusOK,
				Message:    "success",
				Data:       map[string]interface{}{"data": uploadUrl},
			})
	}
}
