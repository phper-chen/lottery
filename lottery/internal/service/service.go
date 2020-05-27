package service

import (
	"log"

	"golang.org/x/time/rate"

	"app/lottery/internal/dao"
)

type Svc struct {
	dao          *dao.Dao
	ReqLimiter *rate.Limiter
}

func New() *Svc {

	s := &Svc{
		dao: dao.New(),
		ReqLimiter: rate.NewLimiter(500, 1000),
	}
	return s
}

func (s *Svc) Close() {
	s.dao.Close()
}

func (s *Svc) GetSysLogger() *log.Logger {
	return s.dao.Logger
}