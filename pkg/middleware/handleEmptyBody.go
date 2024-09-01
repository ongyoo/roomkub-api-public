package middleware

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ongyoo/roomkub-api/pkg/api"
)

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func HandleEmptyBody(c *gin.Context) {
	bw := &bodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = bw
	c.Next()
	if len(c.Errors) != 0 || c.Writer.Status() != http.StatusOK {
		return
	}

	if bw.body.String() == "" {
		c.JSON(http.StatusOK, api.APIMessage{
			Success: true,
			Message: "SUCCESSFUL",
		})
	}
	return
}
