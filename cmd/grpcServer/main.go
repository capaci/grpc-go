package main

import (
	"database/sql"
	"net"

	"github.com/capaci/fc-grpc-go/internal/database"
	"github.com/capaci/fc-grpc-go/internal/pb"
	"github.com/capaci/fc-grpc-go/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}

	if err := createTables(db); err != nil {
		panic(err)
	}

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}

func createTables(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS categories (id string primary key, name string, description string)")
	if err != nil {
		return err
	}

	return nil
}
