package adder

import(
	"context"
	"grpcadder/pkg/api"
	"fmt"

)

//GRPCServer ...
type GRPCServer struct{}

func (s GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {

	fmt.Println(|"Add")
	return &api.AddResponse{Result: req.GetX() + req.GetY()},nil
	
}
