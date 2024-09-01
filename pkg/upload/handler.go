package upload

import (
	"fmt"
	"net/http"
	"os"

	b64 "encoding/base64"

	"github.com/gin-gonic/gin"
	"github.com/ongyoo/roomkub-api/pkg/api"
	"github.com/ongyoo/roomkub-api/pkg/crypto"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

func (h Handler) Upload(c *gin.Context) {
	key := c.Param("key")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	resUpload, err := h.service.Upload(c, key, *file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIResponse[*UploadResponse]{
		Success: true,
		Result:  resUpload,
	})
}

func (h Handler) GetPrivate(c *gin.Context) {
	key := c.Param("key")
	id := c.Param("id")
	businessID := c.Param("business_id")
	s3Key := fmt.Sprintf("%s/%s/%s", key, businessID, id)
	fileName := id
	body, contentType, err := h.service.GetObject(c, s3Key, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer os.Remove(fileName)
	c.Header("content-disposition", fmt.Sprintf("inline; filename=\"%s\"", fileName))
	c.Header("content-length", fmt.Sprintf("%d", len(body)))
	c.Data(http.StatusOK, contentType, body)
}

func (h Handler) GetPubilc(c *gin.Context) {
	//key := c.Param("key")
	id := c.Param("id")
	sDec, err := b64.StdEncoding.DecodeString(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	//businessID := c.Param("business_id")
	//s3Key := fmt.Sprintf("%s/%s/%s", key, businessID, id)
	s3Key := string(sDec)
	decodeS3Key, err := crypto.DecryptAes256(s3Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	s3Key = decodeS3Key
	fileName := id
	body, contentType, err := h.service.GetObject(c, s3Key, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer os.Remove(fileName)
	c.Header("content-disposition", fmt.Sprintf("inline; filename=\"%s\"", fileName))
	c.Header("content-length", fmt.Sprintf("%d", len(body)))
	c.Data(http.StatusOK, contentType, body)
}

func (h Handler) Download(c *gin.Context) {
	key := c.Param("key")
	id := c.Param("id")
	s3Key := key + "/" + id
	fileName := id
	body, contentType, err := h.service.GetObject(c, s3Key, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer os.Remove(fileName)
	c.Header("content-disposition", fmt.Sprintf("inline; filename=\"%s\"", fileName))
	c.Header("content-length", fmt.Sprintf("%d", len(body)))
	c.Data(http.StatusOK, contentType, body)

	/*
		// Set the headers for the file download
			c.Header("Content-Disposition", "attachment; filename=myfile.txt")
			c.Header("Content-Type", "application/octet-stream")
			c.Header("Content-Length", fmt.Sprintf("%d", len(fileContent)))

			// Write the []byte content to the response body
			c.Data(http.StatusOK, "application/octet-stream", fileContent)
	*/
}
