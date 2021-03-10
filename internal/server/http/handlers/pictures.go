package handlers

import (
	"net/http"

	"github.com/PawlikMateusz/TWF0ZXVzeiBHb2dvQXBwcyBOQVNB/internal/models"
	"github.com/PawlikMateusz/TWF0ZXVzeiBHb2dvQXBwcyBOQVNB/pkg/imageprovider"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

type PicturesHandler struct {
	imageProvider imageprovider.ImageProvider
}

func (h *PicturesHandler) Get(c *gin.Context) {
	var req models.Request
	if err := c.ShouldBind(&req); err != nil {
		log.Errorf("Failed to parse request params: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := req.Validate(); err != nil {
		log.Debugf("Failed to validate query params: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.String(200, "Success")
}
