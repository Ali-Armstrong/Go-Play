package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"strconv"

	"google.golang.org/grpc"
	pb "moviesapp.com/grpc/protos"
)

const (
	port = ":50051"
)

var movies [] *pb.MovieInfo

type movieServer struct{
	pb.UnimplementedMovieServer
}

func main(){
	initMovies()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterMovieServer(s, &movieServer{})

	log.Printf("Server Listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initMovies(){
	movie1 := &pb.MovieInfo{Id: "1", Isbn: "12345",
		Title: "Spiderman No Way Home", Director: &pb.Director{ Firstname: "Sam", Lastname: "Raimi"}}
	movie2 := &pb.MovieInfo{Id: "2", Isbn: "67890",
		Title: "The-Batman", Director: &pb.Director{ Firstname: "Matt", Lastname: "Reeves"}}
	
	movies = append(movies, movie1)
	movies = append(movies, movie2)
}

func (s *movieServer) GetMovies( 
	in *pb.Empty, 
	stream pb.Movie_GetMoviesServer) error {
	
	for _, movie := range movies{
		if err:= stream.Send(movie); err != nil {
			return err
		}
	}
	return nil
}

func (s *movieServer) GetMovie(
	ctx context.Context,
	in *pb.Id) (*pb.MovieInfo, error) {
	
	res := &pb.MovieInfo{}

	for _, movie := range movies {
		if movie.GetId() == in.GetValue(){
			res = movie
			break
		}
	}

	return res, nil
}

func (s *movieServer) CreateMovie( ctx context.Context, in *pb.MovieInfo ) (*pb.Id, error){
	log.Printf("Received: %v", in)
	res := pb.Id{}
	res.Value = strconv.Itoa(rand.Intn(100000000))
	in.Id = res.GetValue()
	movies = append(movies, in)
	return &res, nil
}