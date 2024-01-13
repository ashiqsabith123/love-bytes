package functions

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/ashiqsabith123/api-gateway/pkg/helper"
	"github.com/ashiqsabith123/api-gateway/pkg/models/request"
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

	stream, err := clients.UplaodPhotos(ctx)

	if err != nil {
		response := helper.CreateResponse(500, "Cant open photo upload stream", "Server error", nil)
		return response, false
	}

	for _, fileHeader := range files {

		file, err := fileHeader.Open()
		if err != nil {
			response := helper.CreateResponse(http.StatusBadRequest, "Photos can't be opened", "Can't open photos", nil)

			return response, false
		}

		defer file.Close() // Close the file here

		buffer := make([]byte, 8024)

		for {
			bufferLen, err := file.Read(buffer)
			if err != nil {
				if err == io.EOF {
					break
				}

				response := helper.CreateResponse(http.StatusBadRequest, "Can't read buffer", "Can't open photos", nil)

				return response, false
			}

			chunk := pb.PhotoRequest{
				UserID:    helper.GetUserID(ctx),
				ImageData: buffer[:bufferLen],
			}

			if bufferLen < len(buffer) {
				chunk.LastChunk = true
			}

			if err := stream.Send(&chunk); err != nil {
				response := helper.CreateResponse(http.StatusBadRequest, "Can't send to stream", err.Error(), nil)

				return response, false
			}
		}

	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println("err", err)
	}

	if resp != nil {
		if resp.Error != nil {
			response := helper.CreateResponse(resp.Code, resp.Message, string(resp.Error.Value), nil)
			return response, false
		}

		response := helper.CreateResponse(resp.Code, resp.Message, nil, nil)

		return response, true
	}

	response := helper.CreateResponse(500, "Server error", "The service is not responding ", nil)
	return response, false

}

func (M *MatchFunctions) SaveUserPrefrences(ctx context.Context, userPref request.UserPreferences) (responce.Response, bool) {

	err := helper.Validator(userPref)

	if err != nil {
		response := helper.CreateResponse(400, "Data is not in proper format", "Invalid fields", nil)
		return response, false
	}

	resp, _ := clients.SaveUserPrefrences(ctx, &pb.UserPrefrencesRequest{
		UserId:        helper.GetUserID(ctx),
		Height:        userPref.Height,
		MaritalStatus: userPref.MaritalStatus,
		Faith:         userPref.Faith,
		MotherTounge:  userPref.MotherTongue,
		SmokeStatus:   userPref.SmokeStatus,
		AlcoholStatus: userPref.AlcoholStatus,
		SettleStatus:  userPref.SettleStatus,
		Hobbies:       userPref.Hobbies,
		TeaPerson:     userPref.TeaPerson,
		LoveLanguage:  userPref.LoveLanguage,
	})

	if resp.Error != nil {
		response := helper.CreateResponse(resp.Code, resp.Message, string(resp.Error.Value), nil)
		return response, false
	}

	response := helper.CreateResponse(resp.Code, resp.Message, nil, nil)

	return response, true
}

func (M *MatchFunctions) GetMatches() {
	clients.GetMatchedUsers(context.TODO(), &pb.UserIdRequest{
		UserID: 1,
	})
}
