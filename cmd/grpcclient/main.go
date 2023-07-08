package main

import (
	"context"

	"github.com/ecshreve/jepp/internal/ent/proto/entpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	// Open a connection to the server.
	conn, err := grpc.Dial(":5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed connecting to server: %s", err)
	}
	defer conn.Close()

	// Create a User service Client on the connection.
	sc := entpb.NewSeasonServiceClient(conn)
	// cac := entpb.NewCategoryServiceClient(conn)
	// gc := entpb.NewGameServiceClient(conn)
	// clc := entpb.NewClueServiceClient(conn)
	ctx := context.Background()

	// On a separate RPC invocation, retrieve the user we saved previously.
	get, err := sc.Get(ctx, &entpb.GetSeasonRequest{
		Id: 37,
	})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed retrieving season: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("retrieved season with id=%d: %v", get.Id, get)

	// getCat, err := cac.Get(ctx, &entpb.GetCategoryRequest{
	// 	Id: 35,
	// })
	// if err != nil {
	// 	se, _ := status.FromError(err)
	// 	log.Fatalf("failed retrieving cat: status=%s message=%s", se.Code(), se.Message())
	// }
	// log.Printf("retrieved cat with id=%d: %v", getCat.Id, getCat)

	// getGame, err := gc.Get(ctx, &entpb.GetGameRequest{
	// 	Id: 4147,
	// })
	// if err != nil {
	// 	se, _ := status.FromError(err)
	// 	log.Fatalf("failed retrieving game: status=%s message=%s", se.Code(), se.Message())
	// }
	// log.Printf("retrieved game with id=%d: %v", getGame.Id, getGame)

	// getClue, err := clc.Get(ctx, &entpb.GetClueRequest{
	// 	Id: 740101024,
	// })
	// if err != nil {
	// 	se, _ := status.FromError(err)
	// 	log.Fatalf("failed retrieving clue: status=%s message=%s", se.Code(), se.Message())
	// }
	// log.Printf("retrieved clue with id=%d: %v", getClue.Id, getClue)

}
