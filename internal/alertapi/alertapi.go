package alertapi

import (
	"context"

	"github.com/viper-00/nothing/pkg/memdb"
)

type Server struct {
	Database *memdb.Database
}

func (s *Server) HandleAlerts(ctx context.Context, alert *Alert) (*Response, error) {
	return nil, nil
}

func (s *Server) AlertRequest(ctx context.Context, request *Request) (*AlertArray, error) {
	return nil, nil
}

func (s *Server) mustEmbedUnimplementedAlertServiceServer() {}
