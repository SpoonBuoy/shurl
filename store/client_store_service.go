package store

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type Client struct {
	Limiter  *rate.Limiter
	LastSeen time.Time
}

type Clients struct {
	Mu    sync.Mutex
	Table map[string]*Client
}

func NewClientsTable() *Clients {
	return &Clients{
		Mu:    sync.Mutex{},
		Table: make(map[string]*Client),
	}
}
