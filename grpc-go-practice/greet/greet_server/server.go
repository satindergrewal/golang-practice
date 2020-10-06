package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"golang-practice/grpc-go-practice/greet/greetpb"

	"google.golang.org/grpc"

	"google.golang.org/grpc/reflection"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	firstName := req.GetGreeting().GetFirstName()
	min := 1
	// max := 10
	for i := 0; i < 10; i++ {
		// randNum := rand.Intn(max - min)
		randNum := min
		result := "Hello " + firstName + " number " + strconv.Itoa(i) + ". Sleeping for " + strconv.Itoa(randNum) + " seconds..."
		res := &greetpb.GreetmanyTimesResponse{
			Result: result,
		}
		fmt.Println(res)
		stream.Send(res)

		time.Sleep(time.Duration(randNum) * time.Second)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet function was invoked with a streaming request\n")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			//We have finished reading the client stream
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
}

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Printf("LongGreetEveryone function was invoked with a streaming request\n")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! "

		sendErr := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", sendErr)
			return sendErr
		}
	}
}

func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
	fmt.Printf("GreetWithDeadline function was invoked with %v\n", req)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			// the client canceled the request
			fmt.Println("The client canceled the request!")
			return nil, status.Error(codes.DeadlineExceeded, "the client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetWithDeadlineResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	tls := false
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
		if sslErr != nil {
			log.Fatalf("Failed loading certificates: %v", sslErr)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	greetpb.RegisterGreetServiceServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
