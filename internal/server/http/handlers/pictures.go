package handlers

import (
	"fmt"
	"net/http"

	"github.com/PawlikMateusz/TWF0ZXVzeiBHb2dvQXBwcyBOQVNB/pkg/imageproviders"
)

type PicturesHandler struct {
	imageProvider imageproviders.ImageProvider
}

func (h *PicturesHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO impl
	fmt.Println(h.imageProvider)
	fmt.Println("Pictures get handler")
}
