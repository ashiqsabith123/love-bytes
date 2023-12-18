package functions

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/ashiqsabith123/api-gateway/pkg/helper"
	"github.com/ashiqsabith123/api-gateway/pkg/models/responce"
	client "github.com/ashiqsabith123/api-gateway/pkg/services/match-svc/client/interface"
	interfaces "github.com/ashiqsabith123/api-gateway/pkg/services/match-svc/functions/interface"
	"github.com/ashiqsabith123/love-bytes-proto/match/pb"
)

type MatchFunctions struct {
	client client.MatchClient
}

var clients pb.MatchServiceClient

func NewMatchFunctions(client client.MatchClient) interfaces.MatchFunctions {
	matchfunc := &MatchFunctions{client: client}
	clients = matchfunc.client.GetClient()
	return &MatchFunctions{}

}

func (M *MatchFunctions) UploadPhotos(ctx context.Context, files []*multipart.FileHeader) (responce.Response, bool) {

	if len(files) < 2 || len(files) > 4 {
		response := helper.CreateResponse(http.StatusBadRequest, "Rquired min 2 photos and max 4 photos", "Invalid photos", nil)
		return response, false
	}

	for _, fileHeader := range files {

		_, err := clients.UplaodPhotos(ctx)
		if err != nil {
			response := helper.CreateResponse(500, "Cant open photo upload stream", "Server error", nil)
			return response, false
		}

		file, err := fileHeader.Open()
		if err != nil {
			response := helper.CreateResponse(http.StatusBadRequest, "Photos cant be opened", "Cant open photos", nil)
			return response, false
		}
		defer file.Close()

		buffer := make([]byte, 1024)

		for {
			_, err := file.Read(buffer)
			if err != nil {
				response := helper.CreateResponse(http.StatusBadRequest, "Cant read buffer", "Cant open photos", nil)
				return response, false
			}
		}
	}

	return responce.Response{}, true

}
