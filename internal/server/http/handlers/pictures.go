package handlers

import (
	"fmt"
	"net/http"

	"github.com/PawlikMateusz/TWF0ZXVzeiBHb2dvQXBwcyBOQVNB/internal/models"
	"github.com/PawlikMateusz/TWF0ZXVzeiBHb2dvQXBwcyBOQVNB/pkg/imageprovider"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

type PicturesHandler struct {
	ImageProvider imageprovider.ImageProvider
}

func (h *PicturesHandler) Get(c *gin.Context) {
	var req models.Request
	if err := c.ShouldBind(&req); err != nil {
		log.Errorf("Failed to parse request params: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Failed to parse request params: %s", err),
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

	urls, err := h.ImageProvider.GetImagesURLs(*req.StartDate, *req.EndDate)
	if err != nil {
		log.Debugf("Failed to fetch urls: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to fetch images: %s", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"urls": urls,
	})
}
