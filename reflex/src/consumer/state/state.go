package state

//go:generate stategen --inputFile=state.go --inputStruct=stateImpl --outputInterface=State --outputFile=state_gen.go

import (
	"database/sql"
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/neurotempest/mq_test/reflex/src/producer"
	producer_grpc "github.com/neurotempest/mq_test/reflex/src/producer/client/grpc"
)

var dbURI = flag.String("db_uri", "user:pw@tcp(localhost:3306)/db_name", "db connection uri")

type stateImpl struct {
	db *sql.DB
	producerClient producer.Client
}

func New() *stateImpl {

	db, err := sql.Open("mysql", *dbURI)
	if err != nil {
		log.Fatal("failed to create DB connection:", err.Error())
	}

	producerClient, err := producer_grpc.New()
	if err != nil {
		log.Fatal("failed to create producer grpc client:", err.Error())
	}

	return &stateImpl{
		db: db,
		producerClient: producerClient,
	}
}

