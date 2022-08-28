package state

//go:generate stategen --inputFile=state.go --inputStruct=stateImpl --outputInterface=State --outputFile=state_gen.go

import (
	"database/sql"
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/neurotempest/mq_test/reflex/src/consumer"
	consumer_grpc "github.com/neurotempest/mq_test/reflex/src/consumer/client/grpc"
)

var dbURI = flag.String("db_uri", "user:pw@tcp(localhost:3306)/db_name", "db connection uri")

type stateImpl struct {
	db *sql.DB
	consumerClient consumer.Client
}

func New() *stateImpl {

	db, err := sql.Open("mysql", *dbURI)
	if err != nil {
		log.Fatal("failed to create DB connection:", err.Error())
	}

	consumerClient, err := consumer_grpc.New()
	if err != nil {
		log.Fatal("failed to create consumer grpc client:", err.Error())
	}

	return &stateImpl{
		consumerClient: consumerClient,
		db: db,
	}
}

