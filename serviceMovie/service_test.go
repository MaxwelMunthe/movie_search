package serviceMovie

import (
	"context"
	"log"
	"net"
	"net/url"
	"testing"

	"movie-search/infrastructure"
	"movie-search/proto"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func init() {
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	env := infrastructure.Environment{}
	env.SetEnvironment()
	env.LoadConfig()
	env.InitOmdbAPI()
}

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	moviepb.RegisterMovieServiceServer(server, InjectServiceMovie())

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestMovieServiceServer_GetMovie(t *testing.T) {
	expectedResult := &moviepb.GetMovieRes{Response: "True"}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := moviepb.NewMovieServiceClient(conn)
	request := &moviepb.GetMovieReq{Movie: &moviepb.MovieReq{
		Pagination: "2",
		Searchword: "batman",
	}}

	response, err := client.GetMovie(ctx, request)
	assert.Equal(t, expectedResult.Response, response.Response)
	assert.NoError(t, err)
}
