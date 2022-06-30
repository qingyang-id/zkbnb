package globalrpc

import (
	"context"

	"github.com/zecrey-labs/zecrey-legend/common/model/account"
	"github.com/zecrey-labs/zecrey-legend/common/model/mempool"
	"github.com/zecrey-labs/zecrey-legend/service/api/app/internal/svc"
	"github.com/zecrey-labs/zecrey-legend/service/rpc/globalRPC/globalRPCProto"
	"github.com/zecrey-labs/zecrey-legend/service/rpc/globalRPC/globalrpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type GlobalRPC interface {
	SendTx(txType uint32, txInfo string) (string, error)
	GetLpValue(pairIndex uint32, lpAmount string) (*globalRPCProto.RespGetLpValue, error)
	GetPairInfo(pairIndex uint32) (*globalRPCProto.RespGetLatestPairInfo, error)
	GetSwapAmount(ctx context.Context, pairIndex, assetId uint64, assetAmount string, isFrom bool) (string, uint32, error)
	GetNextNonce(accountIndex uint32) (uint64, error)
	GetLatestAssetsListByAccountIndex(accountIndex uint32) ([]*globalrpc.AssetResult, error)
	GetLatestAccountInfoByAccountIndex(accountIndex uint32) (*globalrpc.RespGetLatestAccountInfoByAccountIndex, error)
	GetMaxOfferId(accountIndex uint32) (uint64, error)
	SendMintNftTx(txInfo string) (int64, error)
	SendCreateCollectionTx(txInfo string) (int64, error)

	SendAddLiquidityTx(txInfo string) (string, error)
	SendAtomicMatchTx(txInfo string) (string, error)
	SendCancelOfferTx(txInfo string) (string, error)
	SendRemoveLiquidityTx(txInfo string) (string, error)
	SendSwapTx(txInfo string) (string, error)
	SendTransferNftTx(txInfo string) (string, error)
	SendTransferTx(txInfo string) (string, error)
	SendWithdrawNftTx(txInfo string) (string, error)
	SendWithdrawTx(txInfo string) (string, error)
}

func New(svcCtx *svc.ServiceContext, ctx context.Context) GlobalRPC {
	return &globalRPC{
		AccountModel:        account.NewAccountModel(svcCtx.Conn, svcCtx.Config.CacheRedis, svcCtx.GormPointer),
		AccountHistoryModel: account.NewAccountHistoryModel(svcCtx.Conn, svcCtx.Config.CacheRedis, svcCtx.GormPointer),
		MempoolModel:        mempool.NewMempoolModel(svcCtx.Conn, svcCtx.Config.CacheRedis, svcCtx.GormPointer),
		MempoolDetailModel:  mempool.NewMempoolDetailModel(svcCtx.Conn, svcCtx.Config.CacheRedis, svcCtx.GormPointer),
		RedisConnection:     svcCtx.RedisConn,
		globalRPC:           globalrpc.NewGlobalRPC(zrpc.MustNewClient(svcCtx.Config.GlobalRpc)),
		ctx:                 ctx,
	}
}
