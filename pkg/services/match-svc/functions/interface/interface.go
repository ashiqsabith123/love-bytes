package interfaces

import (
	"context"
	"mime/multipart"

	"github.com/ashiqsabith123/api-gateway/pkg/models/request"
	"github.com/ashiqsabith123/api-gateway/pkg/models/responce"
)

type MatchFunctions interface {
	UploadPhotos(ctx context.Context, files []*multipart.FileHeader) (responce.Response, bool)
	SaveUserPrefrences(ctx context.Context, userPref request.UserPreferences) (responce.Response, bool)
	GetMatches(ctx context.Context) (responce.Response, bool)
	CreateIntrest(ctx context.Context, intrest request.IntrestReq) (responce.Response, bool)
}
