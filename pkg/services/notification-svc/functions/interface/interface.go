package interfaces

import (
	"context"

	"github.com/ashiqsabith123/api-gateway/pkg/models/responce"
)

type NotificationFunctions interface {
	GetAllNotifications(ctx context.Context, filter ...string) (responce.Response, bool)
}
