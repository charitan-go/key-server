package app

import (
	"log"

	"github.com/charitan-go/key-server/external/auth"
	"github.com/charitan-go/key-server/grpc"
	"github.com/charitan-go/key-server/internal/key"
	"github.com/charitan-go/key-server/rotation"
	"go.uber.org/fx"
)

// Run both servers concurrently
func runServers(grpcSrv *grpc.GrpcServer, rotationSrv *rotation.RotationServer) {
	log.Println("In invoke")

	// Start gRPC server
	go func() {
		log.Println("In goroutine of grpc")
		grpcSrv.Run()
	}()

	// Start key rotation
	go func() {
		log.Println("Start key rotation")
		rotationSrv.Run()
	}()
}

func Run() {

	fx.New(
		key.KeyModule,
		auth.AuthModule,
		fx.Provide(
			grpc.NewGrpcServer,
			rotation.NewRotationServer,
		),
		fx.Invoke(runServers),
	).Run()
}
