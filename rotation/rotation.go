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
	time.Sleep(10 * time.Second)
	for true {
		s.keySvc.GenerateKeyPairs()
		time.Sleep(2 * time.Minute)
	}
}
