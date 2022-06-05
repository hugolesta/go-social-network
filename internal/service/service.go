package service

import (
	"database/sql"

	"github.com/hako/branca"
)

// Service contain the core logic. You can use it to back a REST, GraphQL or RPC API.

type service struct {
	db    *sql.DB
	codec *branca.Branca
}
