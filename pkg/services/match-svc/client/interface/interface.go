package interfaces

import "github.com/ashiqsabith123/love-bytes-proto/match/pb"

type MatchClient interface {
	InitMatchClient()
	GetClient() pb.MatchServiceClient
}
