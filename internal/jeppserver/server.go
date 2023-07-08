package jeppserver

import (
	"context"
	"database/sql"
	"net"

	"github.com/ecshreve/jepp/internal/ent"
	"github.com/ecshreve/jepp/internal/ent/proto/entpb"
	"github.com/ecshreve/jepp/internal/utils"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

type JeppServer struct {
	*grpc.Server
	cl *ent.Client
	db *sql.DB
}

func NewServer() *JeppServer {
	logger := log.New()
	logrusEntry := log.NewEntry(logger)
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	cl, db := utils.InitDB()
	grpc.EnableTracing = true

	season_svc := entpb.NewSeasonService(cl)
	category_svc := entpb.NewCategoryService(cl)
	game_svc := entpb.NewGameService(cl)
	clue_svc := entpb.NewClueService(cl)

	alwaysLoggingDeciderServer := func(ctx context.Context, fullMethodName string, servingObject interface{}) bool { return true }

	// Create a new gRPC server (you can wire multiple services to a single server).
	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_logrus.StreamServerInterceptor(logrusEntry),
			grpc_logrus.PayloadStreamServerInterceptor(log.NewEntry(logger), alwaysLoggingDeciderServer)),
		),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
			grpc_logrus.PayloadUnaryServerInterceptor(log.NewEntry(logger), alwaysLoggingDeciderServer)),
		),
	)

	// Register the services with the server.
	entpb.RegisterSeasonServiceServer(server, season_svc)
	entpb.RegisterCategoryServiceServer(server, category_svc)
	entpb.RegisterGameServiceServer(server, game_svc)
	entpb.RegisterClueServiceServer(server, clue_svc)

	return &JeppServer{
		server,
		cl,
		db,
	}
}

func (s *JeppServer) Start() {
	// Open port 5000 for listening to traffic.
	log.Info("listening on :5000")
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}

	// Listen for traffic indefinitely.
	log.Info("serving")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}
