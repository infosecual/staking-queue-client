package queuemngr

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/staking-queue-client/client"
	"github.com/babylonlabs-io/staking-queue-client/config"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"go.uber.org/zap"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_Nosy_QueueManager_Ping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.QueueConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if cfg == nil || logger == nil {
			return
		}

		qc, err := NewQueueManager(cfg, logger)
		if err != nil {
			return
		}
		qc.Ping()
	})
}

func Fuzz_Nosy_QueueManager_PushActiveStakingEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.QueueConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var ev *client.StakingEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if cfg == nil || logger == nil || ev == nil {
			return
		}

		qc, err := NewQueueManager(cfg, logger)
		if err != nil {
			return
		}
		qc.PushActiveStakingEvent(ctx, ev)
	})
}

func Fuzz_Nosy_QueueManager_PushUnbondingStakingEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.QueueConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var ev *client.StakingEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if cfg == nil || logger == nil || ev == nil {
			return
		}

		qc, err := NewQueueManager(cfg, logger)
		if err != nil {
			return
		}
		qc.PushUnbondingStakingEvent(ctx, ev)
	})
}

func Fuzz_Nosy_QueueManager_PushWithdrawableStakingEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.QueueConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var ev *client.StakingEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if cfg == nil || logger == nil || ev == nil {
			return
		}

		qc, err := NewQueueManager(cfg, logger)
		if err != nil {
			return
		}
		qc.PushWithdrawableStakingEvent(ctx, ev)
	})
}

func Fuzz_Nosy_QueueManager_PushWithdrawnStakingEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.QueueConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var ev *client.StakingEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if cfg == nil || logger == nil || ev == nil {
			return
		}

		qc, err := NewQueueManager(cfg, logger)
		if err != nil {
			return
		}
		qc.PushWithdrawnStakingEvent(ctx, ev)
	})
}

func Fuzz_Nosy_QueueManager_ReQueueMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.QueueConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var message client.QueueMessage
		fill_err = tp.Fill(&message)
		if fill_err != nil {
			return
		}
		var queueName string
		fill_err = tp.Fill(&queueName)
		if fill_err != nil {
			return
		}
		if cfg == nil || logger == nil {
			return
		}

		qc, err := NewQueueManager(cfg, logger)
		if err != nil {
			return
		}
		qc.ReQueueMessage(ctx, message, queueName)
	})
}

func Fuzz_Nosy_QueueManager_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.QueueConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if cfg == nil || logger == nil {
			return
		}

		qc, err := NewQueueManager(cfg, logger)
		if err != nil {
			return
		}
		qc.Start()
	})
}

func Fuzz_Nosy_QueueManager_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.QueueConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if cfg == nil || logger == nil {
			return
		}

		qc, err := NewQueueManager(cfg, logger)
		if err != nil {
			return
		}
		qc.Stop()
	})
}
