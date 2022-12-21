package clients

import (
	"github.com/uacademy/e_commerce/api_gateway/config"
	ecom "github.com/uacademy/e_commerce/api_gateway/proto-gen/e_commerce"

	"google.golang.org/grpc"
)

type GrpcClients struct {
	Category ecom.CategoryServiceClient
	Product  ecom.ProductServiceClient
	Order    ecom.OrderServiceClient
	Auth     ecom.AuthServiceClient

	conns []*grpc.ClientConn
}

func NewGrpcClients(cfg config.Config) (*GrpcClients, error) {
	connCategory, err := grpc.Dial(cfg.CatalogServiceGrpcHost+cfg.CatalogServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	category := ecom.NewCategoryServiceClient(connCategory)

	connProduct, err := grpc.Dial(cfg.CatalogServiceGrpcHost+cfg.CatalogServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	product := ecom.NewProductServiceClient(connProduct)

	connOrder, err := grpc.Dial(cfg.OrderServiceGrpcHost+cfg.OrderServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	order := ecom.NewOrderServiceClient(connOrder)

	connAuth, err := grpc.Dial(cfg.AuthServiceGrpcHost+cfg.AuthServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	auth := ecom.NewAuthServiceClient(connAuth)

	conns := make([]*grpc.ClientConn, 0)

	return &GrpcClients{
		Category: category,
		Product:  product,
		Order:    order,
		Auth:     auth,
		conns:    append(conns, connCategory, connProduct, connOrder, connAuth),
	}, nil
}

// Close ...
func (c *GrpcClients) Close() {
	for _, v := range c.conns {
		v.Close()
	}
}
