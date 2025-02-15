package rotation

import (
	"time"

	"github.com/charitan-go/key-server/internal/key/service"
)

type RotationServer struct {
	keySvc service.KeyService
}

func NewRotationServer(keySvc service.KeyService) *RotationServer {
	return &RotationServer{keySvc}
}

func (s *RotationServer) Run() {

	for true {
		time.Sleep(30 * time.Second)

		s.keySvc.GenerateKeyPairs()
	}
}
