package btcclient

import (
	"fmt"
	"net"
	"os"

	"github.com/btcsuite/btclog"
	"github.com/btcsuite/btcwallet/chain"
	"github.com/lightningnetwork/lnd/blockcache"
	"github.com/lightningnetwork/lnd/chainntnfs"
	"github.com/lightningnetwork/lnd/chainntnfs/bitcoindnotify"

	"github.com/babylonlabs-io/staking-expiry-checker/internal/config"
	"github.com/babylonlabs-io/staking-expiry-checker/internal/utils"
)

type BTCNotifier struct {
	*bitcoindnotify.BitcoindNotifier
}

func NewBTCNotifier(
	cfg *config.BTCConfig,
	hintCache HintCache,
) (*BTCNotifier, error) {
	params, err := utils.GetBTCParams(cfg.NetParams)
	if err != nil {
		return nil, fmt.Errorf("invalid BTC network params: %w", err)
	}

	bitcoindCfg := &chain.BitcoindConfig{
		ChainParams:        params,
		Host:               cfg.RPCHost,
		User:               cfg.RPCUser,
		Pass:               cfg.RPCPass,
		Dialer:             BuildDialer(cfg.RPCHost),
		PrunedModeMaxPeers: cfg.PrunedNodeMaxPeers,
		PollingConfig: &chain.PollingConfig{
			BlockPollingInterval:    cfg.BlockPollingInterval,
			TxPollingInterval:       cfg.TxPollingInterval,
			TxPollingIntervalJitter: cfg.TxPollingIntervalJitter,
		},
	}

	// Setup logging for chainntnfs
	backend := btclog.NewBackend(os.Stdout)
	logger := backend.Logger("NTFN")
	logger.SetLevel(btclog.LevelDebug)
	chainntnfs.UseLogger(logger)

	bitcoindConn, err := chain.NewBitcoindConn(bitcoindCfg)
	if err != nil {
		return nil, err
	}

	if err := bitcoindConn.Start(); err != nil {
		return nil, fmt.Errorf("unable to connect to "+
			"bitcoind: %v", err)
	}

	chainNotifier := bitcoindnotify.New(
		bitcoindConn, params, hintCache,
		hintCache, blockcache.NewBlockCache(cfg.BlockCacheSize),
	)

	return &BTCNotifier{BitcoindNotifier: chainNotifier}, nil
}

func BuildDialer(rpcHost string) func(string) (net.Conn, error) {
	return func(addr string) (net.Conn, error) {
		return net.Dial("tcp", rpcHost)
	}
}

type HintCache interface {
	chainntnfs.SpendHintCache
	chainntnfs.ConfirmHintCache
}

type EmptyHintCache struct{}

var _ HintCache = (*EmptyHintCache)(nil)

func (c *EmptyHintCache) CommitSpendHint(height uint32, spendRequests ...chainntnfs.SpendRequest) error {
	return nil
}
func (c *EmptyHintCache) QuerySpendHint(spendRequest chainntnfs.SpendRequest) (uint32, error) {
	return 0, nil
}
func (c *EmptyHintCache) PurgeSpendHint(spendRequests ...chainntnfs.SpendRequest) error {
	return nil
}

func (c *EmptyHintCache) CommitConfirmHint(height uint32, confRequests ...chainntnfs.ConfRequest) error {
	return nil
}
func (c *EmptyHintCache) QueryConfirmHint(confRequest chainntnfs.ConfRequest) (uint32, error) {
	return 0, nil
}
func (c *EmptyHintCache) PurgeConfirmHint(confRequests ...chainntnfs.ConfRequest) error {
	return nil
}
