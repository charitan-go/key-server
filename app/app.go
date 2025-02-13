package app

import (
	"log"

	"github.com/charitan-go/key-server/external/auth"
	"github.com/charitan-go/key-server/grpc"
	"github.com/charitan-go/key-server/internal/key"
	"go.uber.org/fx"
)

// Run both servers concurrently
func runServers(grpcSrv *grpc.GrpcServer) {
	log.Println("In invoke")

	// Start gRPC server
	go func() {
		log.Println("In goroutine of grpc")
		grpcSrv.Run()
	}()
}

func Run() {

	fx.New(
		key.KeyModule,
		auth.AuthModule,
		fx.Provide(
			grpc.NewGrpcServer,
		),
		fx.Invoke(runServers),
	).Run()
}
