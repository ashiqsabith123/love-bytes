package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ashiqsabith123/api-gateway/pkg/models/request"
	responce "github.com/ashiqsabith123/api-gateway/pkg/models/responce"
	match "github.com/ashiqsabith123/api-gateway/pkg/services/match-svc/functions/interface"
	Const "github.com/ashiqsabith123/love-bytes-proto/constants"
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
// @Security		BearerTokenAuth
// @Param photos formData file true "Multiple photos to upload"
// @Success 200 {object} responce.Response "Photos uploaded successfully"
// @Failure 400 {object} responce.Response "Invalid request or contains other files"
// @Failure 401 {object} responce.Response "Unauthorized - User id not found"
// @Router /upload/photos [post]
func (M *MatchHandler) UploadPhotos(C *gin.Context) {

	_, ok := C.Get("userID")

	if !ok {
		resp := responce.ErrorReposonce(http.StatusBadRequest, Const.INVALID_REQUEST, Const.USER_ID_NOT_FOUND)
		C.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	form, err := C.MultipartForm()

	if err != nil {

		resp := responce.ErrorReposonce(http.StatusBadRequest, Const.INVALID_REQUEST, "No photos found")
		C.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	files := form.File["photos"]

	for _, fileHeader := range files {

		contentType := fileHeader.Header.Get("Content-Type")
		if contentType != "image/jpeg" {
			resp := responce.ErrorReposonce(http.StatusBadRequest, Const.INVALID_REQUEST, "Contains other files")
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

func (M *MatchHandler) SaveUserPrefrences(C *gin.Context) {
	_, ok := C.Get("userID")

	if !ok {
		resp := responce.ErrorReposonce(http.StatusBadRequest, Const.INVALID_REQUEST, Const.USER_ID_NOT_FOUND)
		C.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	var userPrefReq request.UserPreferences

	if err := C.ShouldBindJSON(&userPrefReq); err != nil {
		resp := responce.ErrorReposonce(http.StatusBadRequest, Const.INVALID_REQUEST, err.Error())
		C.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	resp, ok := M.functions.SaveUserPrefrences(C, userPrefReq)

	if !ok {
		C.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	C.JSON(resp.Code, resp)

}

func (M *MatchHandler) GetMatches(C *gin.Context) {
	_, ok := C.Get("userID")

	if !ok {
		resp := responce.ErrorReposonce(http.StatusBadRequest, Const.INVALID_REQUEST, Const.USER_ID_NOT_FOUND)
		C.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	resp, ok := M.functions.GetMatches(C)

	if !ok {
		C.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	

	C.JSON(resp.Code, resp)

	// C.JSON(500, gin.H{
	// 	"error": "server error",
	// })

}

func (M *MatchHandler) CreateIntrest(C *gin.Context) {
	_, ok := C.Get("userID")

	if !ok {
		resp := responce.ErrorReposonce(http.StatusBadRequest, Const.INVALID_REQUEST, Const.USER_ID_NOT_FOUND)
		C.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	reciverId := C.Param("recieverId")

	if reciverId == "" {
		if !ok {
			resp := responce.ErrorReposonce(http.StatusBadRequest, Const.INVALID_REQUEST, "Reviver ID not found")
			C.AbortWithStatusJSON(http.StatusBadRequest, resp)
			return
		}
	}

	var intrestReq request.IntrestReq

	ID, err := strconv.Atoi(reciverId)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	intrestReq.RecieverId = uint(ID)

	resp, ok := M.functions.CreateIntrest(C, intrestReq)

	if !ok {
		C.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	C.JSON(resp.Code, resp)

}

func (M *MatchHandler) GetAllIntrestRequests(C *gin.Context) {
	_, ok := C.Get("userID")

	if !ok {
		resp := responce.ErrorReposonce(http.StatusBadRequest, Const.INVALID_REQUEST, Const.USER_ID_NOT_FOUND)
		C.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	resp, ok := M.functions.GetAllIntrestRequests(C)

	if !ok {
		C.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	C.JSON(resp.Code, resp)
}
