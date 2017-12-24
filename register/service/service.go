package service

import (
	"time"
)

type Register struct {
	NS      string `form:"ns" json:"ns"`
	Version string `form:"ver" json:"ver" binding:"required"`
	Address string `form:"addr" json:"addr" binding:"required"`
	//Route    string `form:"route" json:"route" binding:"required"`
	Protocol string `form:"protocol" json:"protocol" binding:"required"`
}

type Services struct {
	List []Service
}

type Service struct {
	Register
	LastUpdate time.Time
}

// NewServiceRepository new service repository.
// The data will be stored in memory
func NewServiceRepository(maxNum int) *Services {
	return &Services{}
}

func (rep *Services) Append(one Service) {
	rep.List = append(rep.List, one)
}
