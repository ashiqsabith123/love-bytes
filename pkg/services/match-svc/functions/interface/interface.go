package interfaces

import (
	"context"
	"mime/multipart"

	"github.com/ashiqsabith123/api-gateway/pkg/models/responce"
)

type MatchFunctions interface {
	UploadPhotos(ctx context.Context, files []*multipart.FileHeader) (responce.Response, bool)
}
