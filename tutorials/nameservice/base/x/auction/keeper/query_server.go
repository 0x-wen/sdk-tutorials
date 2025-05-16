package keeper

import (
	"context"
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	auction "github.com/cosmos/sdk-tutorials/tutorials/nameservice/base/x/auction"
)

var _ auction.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the module QueryServer.
func NewQueryServerImpl(k Keeper) auction.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}

func (qs queryServer) Name(ctx context.Context, r *auction.QueryNameRequest) (*auction.QueryNameResponse, error) {
	if r == nil {
		return nil, errors.New("empty request")
	}
	if len(r.Name) == 0 {
		return nil, auction.ErrEmptyName
	}
	goCtx := sdk.UnwrapSDKContext(ctx)

	record, err := qs.k.GetNameRecord(goCtx, r.Name)
	if err != nil {
		return nil, err
	}

	return &auction.QueryNameResponse{
		Name: &record,
	}, nil
}

func (qs queryServer) Names(ctx context.Context, r *auction.QueryNamesRequest) (*auction.QueryNamesResponse, error) {
	if r == nil {
		return nil, errors.New("empty request")
	}
	goCtx := sdk.UnwrapSDKContext(ctx)

	names, err := qs.k.GetNames(goCtx)
	if err != nil {
		return nil, err
	}

	return &auction.QueryNamesResponse{
		Names: names,
	}, nil

}
