package abci

import (
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/cosmos/sdk-tutorials/tutorials/nameservice/base/x/auction/mempool"
	"github.com/cosmos/sdk-tutorials/tutorials/nameservice/base/x/auction/provider"
)

// PrepareProposalHandler 处理共识过程中的提案的准备
type PrepareProposalHandler struct {
	logger      log.Logger
	txConfig    client.TxConfig
	cdc         codec.Codec
	mempool     *mempool.ThresholdMempool
	txProvider  provider.TxProvider
	runProvider bool
}

// ProcessProposalHandler 处理共识过程中的提案,验证该提案和包括的投票扩展
type ProcessProposalHandler struct {
	TxConfig client.TxConfig
	Codec    codec.Codec
	Logger   log.Logger
}

// VoteExtHandler 该结构用于处理投票扩展
type VoteExtHandler struct {
	logger  log.Logger
	mempool *mempool.ThresholdMempool
	cdc     codec.Codec
}

type InjectedVoteExt struct {
	VoteExtSigner []byte
	Bids          [][]byte
}

type InjectedVotes struct {
	Votes []InjectedVoteExt
}

type AppVoteExtension struct {
	Height int64
	Bids   [][]byte
}

type SpecialTransaction struct {
	Height int
	Bids   [][]byte
}
