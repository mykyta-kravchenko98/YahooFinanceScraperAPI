package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "yahooscraper/pkg/yahooScraper_v1"

	"yahooscraper/internal/configs"
	"yahooscraper/internal/db/repositories"
	"yahooscraper/internal/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type grpcServer struct {
	pb.UnimplementedYahooStocksServiceServer
	stocksService services.StockService
}

func (s *grpcServer) GetAllValidStocks(ctx context.Context, req *pb.GetAllValidStocksRequest) (*pb.GetAllValidStocksResponse, error) {
	stocks, err := s.stocksService.GetActuallyStocks()

	if err != nil {
		return nil, err
	}

	protoStocks := make([]*pb.Stocks, 0, len(stocks))
	for _, c := range stocks {
		protoStocks = append(protoStocks, c.ProtoStock())
	}

	return &pb.GetAllValidStocksResponse{Stocks: protoStocks}, nil
}

// Start server
func Init(db *gorm.DB) {
	config := configs.GetConfig()
	port := config.Server.GRPCPort
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		log.Fatalf("Failed when listen: #{err}")
	}

	s := grpc.NewServer()
	reflection.Register(s)

	rep := repositories.NewCurrencySnapshotDataService(db)
	serv := services.NewStockService(rep)

	pb.RegisterYahooStocksServiceServer(s, &grpcServer{stocksService: serv})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed when serve: #{err}")
	}
}