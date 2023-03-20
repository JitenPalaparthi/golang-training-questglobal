package main

import (
	"flag"
	"log"
	"net"

	pb "gitlab.stackroute.in/JitenP/golang-training-questglobal/contactspb"
	"gitlab.stackroute.in/JitenP/golang-training-questglobal/database"
	"gitlab.stackroute.in/JitenP/golang-training-questglobal/grpc/handler"
	"google.golang.org/grpc"
)

var (
	port string
	dsn  string //= "host=localhost user=postgres password=postgres dbname=contactsbd port=55432 sslmode=disable TimeZone=Asia/Shanghai"
	// for mysql: --dsn="user:password@tcp(127.0.0.1:53306)/contactsdb?charset=utf8mb4&parseTime=True&loc=Local"

)

func main() {

	if port == "" {
		flag.StringVar(&port, "port", "8081", "--port=8081 or --port 8081 or -port 8081")
	}
	if dsn == "" {
		flag.StringVar(&dsn, "dsn", "host=localhost user=postgres password=postgres dbname=contactsbd port=65432 sslmode=disable TimeZone=Asia/Shanghai", "--dsn=host=localhost user=postgres password=postgres dbname=contactsbd port=55432 sslmode=disable TimeZone=Asia/Shanghai")
	}
	flag.Parse()

	db, err := database.Connect(dsn)
	if err != nil {
		log.Fatalln(err) //Fatals cannot be recoverd but panics can. Fatal exit the application with os.Exit(1)
	} else {
		log.Println(db)
	}

	lis, err := net.Listen("tcp", ":"+port) // Server is listening on port 58081 using tcp protocol
	log.Println("gRPC server started and lisening on port", port)
	if err != nil {
		log.Fatal(err)
	}
	var opts []grpc.ServerOption

	grpcserver := grpc.NewServer(opts...) // new grpc server with no options.

	cHandler := new(handler.Contact)
	cHandler.DB = db

	pb.RegisterContactServer(grpcserver, cHandler)

	grpcserver.Serve(lis)

}
