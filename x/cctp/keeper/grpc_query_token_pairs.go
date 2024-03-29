package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TokenPair(c context.Context, req *types.QueryGetTokenPairRequest) (*types.QueryGetTokenPairResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTokenPairHex(ctx, req.RemoteDomain, req.RemoteToken)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTokenPairResponse{Pair: val}, nil
}

func (k Keeper) TokenPairs(c context.Context, req *types.QueryAllTokenPairsRequest) (*types.QueryAllTokenPairsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tokenPairs []types.TokenPair
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	TokenPairsStore := prefix.NewStore(store, types.KeyPrefix(types.TokenPairKeyPrefix))

	pageRes, err := query.Paginate(TokenPairsStore, req.Pagination, func(key []byte, value []byte) error {
		var tokenPair types.TokenPair
		if err := k.cdc.Unmarshal(value, &tokenPair); err != nil {
			return err
		}

		tokenPairs = append(tokenPairs, tokenPair)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTokenPairsResponse{TokenPairs: tokenPairs, Pagination: pageRes}, nil
}
