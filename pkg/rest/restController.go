package rest

import "github.com/tabed23/hexagon-architecture/pkg/service"

type Rest struct {
	s *service.Service
}

func NewRest(s *service.Service) *Rest {
	return &Rest{s: s}
}
