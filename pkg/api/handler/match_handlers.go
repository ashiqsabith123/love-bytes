package handler

import (
	"fmt"
	"net/http"

	responce "github.com/ashiqsabith123/api-gateway/pkg/models/responce"
	match "github.com/ashiqsabith123/api-gateway/pkg/services/match-svc/functions/interface"
	"github.com/gin-gonic/gin"
)

type MatchHandler struct {
	functions match.MatchFunctions
}

func NewMatchHandler(mathcFunc match.MatchFunctions) *MatchHandler {
	return &MatchHandler{functions: mathcFunc}
}

// @Summary Upload photos with additional fields
// @Description Upload photos with additional fields using a multipart form
// @ID upload-photos
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "Bearer {token}" default("your_access_token") // Include this if authorization is required
// @Param field_name formData string true "Value for the 'field_name' field"
// @Param photos formData file true "Multiple photos to upload"
// @Success 200 {object} ApiResponse "Photos uploaded successfully"
// @Failure 400 {object} ApiResponse "Invalid request or contains other files"
// @Failure 401 {object} ApiResponse "Unauthorized - User id not found"
// @Router /upload-photos [post]
func (M *MatchHandler) UploadPhotos(C *gin.Context) {

	_, ok := C.Get("userID")

	if !ok {

		resp := responce.ErrorReposonce(http.StatusBadRequest, "Invalid request", "User id not found")
		C.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	form, err := C.MultipartForm()

	if err != nil {

		resp := responce.ErrorReposonce(http.StatusBadRequest, "Invalid request", "No photos found")
		C.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	files := form.File["photos"]

	for _, fileHeader := range files {

		contentType := fileHeader.Header.Get("Content-Type")
		if contentType != "image/jpeg" {

			resp := responce.ErrorReposonce(http.StatusBadRequest, "Invalid request", "Contains other files")
			fmt.Println("ree", resp)
			C.AbortWithStatusJSON(http.StatusBadRequest, resp)
			return
		}

	}

	resp, ok := M.functions.UploadPhotos(C, files)

	if !ok {
		C.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	C.JSON(resp.Code, resp)

}
