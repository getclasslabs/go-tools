package db

import (
	"database/sql"
	"github.com/getclasslabs/go-tools/pkg/tracer"
)

type Database interface {
	Connect(Config)
	Insert(*tracer.Infos, string, ...interface{}) (sql.Result, error)
	Update(*tracer.Infos, string, ...interface{}) (sql.Result, error)
	Get(*tracer.Infos, string, ...interface{}) (map[string]interface{}, error)
	Fetch(*tracer.Infos, string, ...interface{}) ([]map[string]interface{}, error)
}


type Config interface {
	GetUser() string
	GetPassword() string
	GetHost() string
	GetPort() string
	GetDatabase() string
}