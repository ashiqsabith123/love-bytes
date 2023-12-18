package functions

import (
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"sync"

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

// func (M *MatchFunctions) UploadPhotos(ctx context.Context, files []*multipart.FileHeader) (responce.Response, bool) {

// 	if len(files) < 2 || len(files) > 4 {
// 		response := helper.CreateResponse(http.StatusBadRequest, "Rquired min 2 photos and max 4 photos", "Invalid photos", nil)
// 		return response, false
// 	}

// 	tim := time.Now()

// 	for _, fileHeader := range files {

// 		stream, err := clients.UplaodPhotos(ctx)
// 		if err != nil {
// 			response := helper.CreateResponse(500, "Cant open photo upload stream", "Server error", nil)
// 			return response, false
// 		}

// 		file, err := fileHeader.Open()
// 		if err != nil {
// 			response := helper.CreateResponse(http.StatusBadRequest, "Photos cant be opened", "Cant open photos", nil)
// 			return response, false
// 		}
// 		defer file.Close()

// 		buffer := make([]byte, 1024)

// 		for {
// 			n, err := file.Read(buffer)

// 			if err != nil {
// 				if err != io.EOF {
// 					response := helper.CreateResponse(http.StatusBadRequest, "Can't read buffer", "Can't open photos", nil)
// 					return response, false
// 				}
// 				break
// 			}

// 			chunk := pb.PhotoRequest{
// 				UserID:    1,
// 				ImageData: buffer[:n],
// 			}

// 			if n < len(buffer) {
// 				chunk.LastChunk = true
// 			}

// 			if err := stream.Send(&chunk); err != nil {
// 				response := helper.CreateResponse(http.StatusBadRequest, "Can't send to stream", err.Error(), nil)
// 				return response, false
// 			}
// 		}

// 	}

// 	fmt.Println("sinc", time.Since(tim))

// 	return responce.Response{}, true

// }

func (M *MatchFunctions) UploadPhotos(ctx context.Context, files []*multipart.FileHeader) (responce.Response, bool) {

	type chResp struct {
		resp responce.Response
		ok   bool
	}
	var wg sync.WaitGroup

	ch := make(chan chResp, len(files))

	if len(files) < 2 || len(files) > 4 {
		response := helper.CreateResponse(http.StatusBadRequest, "Rquired min 2 photos and max 4 photos", "Invalid photos", nil)
		return response, false
	}

	wg.Add(len(files))

	for _, fileHeader := range files {
		go func(ctx context.Context, fileHeader *multipart.FileHeader, wg *sync.WaitGroup, ch chan chResp) {
			stream, err := clients.UplaodPhotos(ctx)
			if err != nil {
				response := helper.CreateResponse(500, "Cant open photo upload stream", "Server error", nil)
				ch <- chResp{
					resp: response,
					ok:   false,
				}
			}

			file, err := fileHeader.Open()
			if err != nil {
				response := helper.CreateResponse(http.StatusBadRequest, "Photos cant be opened", "Cant open photos", nil)
				ch <- chResp{
					resp: response,
					ok:   false,
				}
			}
			defer file.Close()

			buffer := make([]byte, 8024)

			for {
				n, err := file.Read(buffer)

				if err != nil {
					if err != io.EOF {
						response := helper.CreateResponse(http.StatusBadRequest, "Can't read buffer", "Can't open photos", nil)
						ch <- chResp{
							resp: response,
							ok:   false,
						}
					}
					break
				}

				chunk := pb.PhotoRequest{
					UserID:    1,
					ImageData: buffer[:n],
				}

				if n < len(buffer) {
					chunk.LastChunk = true
				}

				if err := stream.Send(&chunk); err != nil {
					response := helper.CreateResponse(http.StatusBadRequest, "Can't send to stream", err.Error(), nil)
					ch <- chResp{
						resp: response,
						ok:   false,
					}
				}
			}

			wg.Done()
		}(ctx, fileHeader, &wg, ch)
	}

	wg.Wait()

	return responce.Response{}, true

}

// func UploadPic(ctx context.Context, fileHeader *multipart.FileHeader, wg *sync.WaitGroup) (responce.Response, bool) {

// }
