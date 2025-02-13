package key

import (
	"github.com/charitan-go/key-server/internal/key/service"
	"go.uber.org/fx"
)

var KeyModule = fx.Module("donor",
	fx.Provide(
		service.NewKeyService,
	),
)
