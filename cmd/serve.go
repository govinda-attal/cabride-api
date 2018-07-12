package cmd

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/govinda-attal/cabride-api/handler"
	"github.com/govinda-attal/cabride-api/internal/provider"
	"github.com/govinda-attal/cabride-api/internal/rideds"
	"github.com/govinda-attal/cabride-api/pkg/rides/pb"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launches the grpc server on localhost:10000  and webserver on localhost:9080",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func newServer() *handler.CabTrips {
	return handler.NewCabTripsHandler(rideds.NewCabRideStore(), rideds.NewCabRideCache())
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func serve() {
	// GRPC Server Configuration
	provider.Setup()
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCabTripsServer(grpcServer, newServer())
	lis, err := net.Listen("tcp", "localhost:10000")
	if err != nil {
		os.Exit(1)
	}

	// Web Server Configuration
	ctx := context.Background()
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	gwmux := runtime.NewServeMux()

	err = pb.RegisterCabTripsHandlerFromEndpoint(ctx, gwmux, "localhost:10000", dopts)
	if err != nil {
		os.Exit(1)
	}

	mux := http.NewServeMux()
	mux.Handle("/v1/", gwmux)
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("api"))))

	h := cors.Default().Handler(mux)
	srv := &http.Server{
		Addr:    "localhost:9080",
		Handler: h,
	}

	// Start GRPC Server
	go func() {
		grpcServer.Serve(lis)
	}()

	// Start Web Server
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	provider.Cleanup()
	srv.Shutdown(ctx)
	grpcServer.GracefulStop()

	log.Println("microservice shutdown ...")
	os.Exit(0)
}
