package postgresql

import (
	"crypto/tls"

	"github.com/go-pg/pg/v10"
)

type Config struct {
	// TCP host:port or Unix socket depending on Network.
	Addr string

	User     string
	Password string
	Database string

	// TLS config for secure connections.
	TLSConfig *tls.Config
}

func InitDB(cfg Config) {

	db = pg.Connect(&pg.Options{
		Addr:      cfg.Addr,
		User:      cfg.User,
		Password:  cfg.Password,
		Database:  cfg.Database,
		TLSConfig: cfg.TLSConfig,
	})
}
