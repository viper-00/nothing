package api

import "context"

type Server struct{}

func (s *Server) HandlePing(ctx context.Context, info *ServerInfo) (*Message, error) {
	return nil, nil
}

func (s *Server) IsUp(ctx context.Context, info *ServerInfo) (*IsActive, error) {
	return nil, nil
}

func (s *Server) InitAgent(ctx context.Context, info *ServerInfo) (*Message, error) {
	return nil, nil
}

func (s *Server) HandleMonitorData(ctx context.Context, data *MonitorData) (*Message, error) {
	return nil, nil
}

func (s *Server) HandleCustomMonitorData(ctx context.Context, data *MonitorData) (*Message, error) {
	return nil, nil
}

func (s *Server) HandleMonitorDataRequest(ctx context.Context, data *MonitorDataRequest) (*MonitorData, error) {
	return nil, nil
}

func (s *Server) HandleCustomMetricNameRequest(ctx context.Context, info *ServerInfo) (*Message, error) {
	return nil, nil
}

func (s *Server) HandleAgentIdsRequest(ctx context.Context, void *Void) (*Message, error) {
	return nil, nil
}

func (s *Server) mustEmbedUnimplementedMonitorDataServiceServer() {}
