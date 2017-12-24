package service

import (
	"time"

	"github.com/satori/go.uuid"
)

type Register struct {
	NS      string `form:"ns" json:"ns"`
	Name    string `form:"name" json:"name" binding:"required"`
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
	Id         uuid.UUID
	LastUpdate time.Time
}

// NewServiceRepository new service repository.
// The data will be stored in memory
func NewServiceRepository(maxNum int) *Services {
	return &Services{}
}

// Append new service to service repository
func (rep *Services) Append(one Service) {
	if !rep.Has(one) {
		one.Id = uuid.NewV4()
		rep.List = append(rep.List, one)
	}
}

// Has determine if given service is exist in the list
func (rep *Services) Has(one Service) bool {
	for _, srv := range rep.List {
		if one.NS == srv.NS {
			if one.Name == srv.Name {
				return true
			}
		}
	}
	return false
}
