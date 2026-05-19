package context

import (
	"github.com/alex-305/tuido/internal/config"
	"github.com/alex-305/tuido/internal/storage"
)

type AppContext struct {
	Config *config.Config
	Store  *storage.Store
}
