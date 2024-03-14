package grpc

import (
	"context"
	"errors"
	"github.com/yourscarent/cars_manager/internal/log"
	"github.com/yourscarent/cars_manager/internal/usecase"
	"github.com/yourscarent/cars_manager/internal/utils"
	cm "github.com/yourscarent/cars_manager/pkg/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Params struct {
	Usecase    *usecase.Manager
	GRPCServer *grpc.Server
}

func NewServer(p Params) *Server {
	return &Server{
		gRPCServer: p.GRPCServer,
		ucase:      p.Usecase,
	}
}

func (s Server) Register() {
	cm.RegisterCarsManagerServer(s.gRPCServer, &s)
}

type Server struct {
	log log.Logger

	ucase      *usecase.Manager
	gRPCServer *grpc.Server

	valid   utils.Validator
	convert utils.Converter

	cm.UnimplementedCarsManagerServer
}

func (s Server) CreateCar(ctx context.Context, car *cm.Car) (*emptypb.Empty, error) {
	if err := s.valid.ValidateCar(car); err != nil {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err := s.ucase.CreateCar(ctx, s.convert.NewCar(car))
	if err != nil {
		msg, internal := utils.FromError(err)
		if internal {
			s.log.Error("failed to create car", msg)
			return nil, errors.New(utils.InternalError)
		}
		return nil, errors.New(msg)
	}

	return &emptypb.Empty{}, nil
}

func (s Server) DeleteCar(ctx context.Context, id *cm.ID) (*emptypb.Empty, error) {
	if err := s.valid.ValidateId(id); err != nil {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err := s.ucase.DeleteCar(ctx, id.Id)
	if err != nil {
		msg, internal := utils.FromError(err)
		if internal {
			s.log.Error("failed to delete car", msg)
			return nil, errors.New(utils.InternalError)
		}
		return nil, errors.New(msg)
	}

	return &emptypb.Empty{}, nil
}

func (s Server) UpdateCar(ctx context.Context, update *cm.Update) (*emptypb.Empty, error) {
	if err := s.valid.ValidateUpdate(update); err != nil {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err := s.ucase.UpdateCar(ctx, s.convert.NewUpdate(update))
	if err != nil {
		msg, internal := utils.FromError(err)
		if internal {
			s.log.Error("failed to create car", msg)
			return nil, errors.New(utils.InternalError)
		}
		return nil, errors.New(msg)
	}

	return &emptypb.Empty{}, nil
}
