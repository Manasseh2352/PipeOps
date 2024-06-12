package server

import (
	"github.com/gin-gonic/gin"
	db "github.com//db/sqlc"
)


type Server struct {
	store db.store
}