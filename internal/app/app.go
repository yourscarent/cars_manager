package app

import (
	"fmt"
	"github.com/yourscarent/cars_manager/internal/config"
	"github.com/yourscarent/cars_manager/internal/db/postgres"
	"github.com/yourscarent/cars_manager/internal/log"
	server "github.com/yourscarent/cars_manager/internal/server/grpc"
	"github.com/yourscarent/cars_manager/internal/usecase"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func MustStart(cfg *config.Config) {
	l := log.MustSetup(cfg.Env)

	//defer func() {
	//	if err := recover(); err != nil {
	//		l.Error("panic recovery: ", err)
	//	}
	//}()

	l.Debug("cfg", cfg)

	// repo
	db := postgres.MustConnect(cfg.DB)
	defer db.Close()
	repo := postgres.NewRepository(db)

	// usecase
	ucase := usecase.NewManager(usecase.Params{
		Repo: repo,
	})

	// adapter
	gRPCServer := grpc.NewServer()
	server.NewServer(server.Params{
		GRPCServer: gRPCServer,
		Usecase:    ucase,
	}).Register()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	chStop := make(chan os.Signal, 1)
	signal.Notify(chStop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		l.Info("server is running")
		if err = gRPCServer.Serve(lis); err != nil {
			panic("failed to serve: " + err.Error())
		}
	}()

	<-chStop
	l.Info("server stopped")
}
