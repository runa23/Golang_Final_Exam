package main

import (
	"Final-Exam/entity"
	"log"
	"net"

	grpchandler "Final-Exam/user/handler"
	pb "Final-Exam/user/proto/user_service/v1"
	usergorm "Final-Exam/user/repository"
	"Final-Exam/user/service"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "postgresql://postgres:password@pg-db_user:5432/go_db_user"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		log.Fatalln(err)
	}

	//run migration
	if err = gormDB.AutoMigrate(&entity.User{}); err != nil {
		log.Fatalln("Failed to migrate:", err)
	}

	//setup service
	userRepo := usergorm.NewUserRepository(gormDB)
	userService := service.NewUserService(userRepo)
	userHandler := grpchandler.NewUserHandler(userService)

	//run the grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userHandler)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Server is running on port 50051")
	_ = grpcServer.Serve(lis)
}
