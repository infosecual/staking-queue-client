package client

import (
	"context"
	"testing"
	"time"

	"github.com/babylonlabs-io/staking-queue-client/config"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
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

func Fuzz_Nosy_QueueMessage_GetRetryAttempts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueueMessage
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetRetryAttempts()
	})
}

func Fuzz_Nosy_QueueMessage_IncrementRetryAttempts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueueMessage
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.IncrementRetryAttempts()
	})
}

func Fuzz_Nosy_RabbitMqClient_DeleteMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.QueueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var queueName string
		fill_err = tp.Fill(&queueName)
		if fill_err != nil {
			return
		}
		var deliveryTag string
		fill_err = tp.Fill(&deliveryTag)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		c, err := NewRabbitMqClient(config, queueName)
		if err != nil {
			return
		}
		c.DeleteMessage(deliveryTag)
	})
}

func Fuzz_Nosy_RabbitMqClient_GetQueueName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.QueueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var queueName string
		fill_err = tp.Fill(&queueName)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		c, err := NewRabbitMqClient(config, queueName)
		if err != nil {
			return
		}
		c.GetQueueName()
	})
}

func Fuzz_Nosy_RabbitMqClient_Ping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.QueueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var queueName string
		fill_err = tp.Fill(&queueName)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		c, err := NewRabbitMqClient(config, queueName)
		if err != nil {
			return
		}
		c.Ping(ctx)
	})
}

func Fuzz_Nosy_RabbitMqClient_ReQueueMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.QueueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var queueName string
		fill_err = tp.Fill(&queueName)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var message QueueMessage
		fill_err = tp.Fill(&message)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		c, err := NewRabbitMqClient(config, queueName)
		if err != nil {
			return
		}
		c.ReQueueMessage(ctx, message)
	})
}

func Fuzz_Nosy_RabbitMqClient_ReceiveMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.QueueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var queueName string
		fill_err = tp.Fill(&queueName)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		c, err := NewRabbitMqClient(config, queueName)
		if err != nil {
			return
		}
		c.ReceiveMessages()
	})
}

func Fuzz_Nosy_RabbitMqClient_SendMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.QueueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var queueName string
		fill_err = tp.Fill(&queueName)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var messageBody string
		fill_err = tp.Fill(&messageBody)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		c, err := NewRabbitMqClient(config, queueName)
		if err != nil {
			return
		}
		c.SendMessage(ctx, messageBody)
	})
}

func Fuzz_Nosy_RabbitMqClient_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.QueueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var queueName string
		fill_err = tp.Fill(&queueName)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		c, err := NewRabbitMqClient(config, queueName)
		if err != nil {
			return
		}
		c.Stop()
	})
}

func Fuzz_Nosy_RabbitMqClient_sendMessageWithAttempts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.QueueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var q2 string
		fill_err = tp.Fill(&q2)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var messageBody string
		fill_err = tp.Fill(&messageBody)
		if fill_err != nil {
			return
		}
		var q5 string
		fill_err = tp.Fill(&q5)
		if fill_err != nil {
			return
		}
		var attempts int32
		fill_err = tp.Fill(&attempts)
		if fill_err != nil {
			return
		}
		var ttl time.Duration
		fill_err = tp.Fill(&ttl)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		c, err := NewRabbitMqClient(config, q2)
		if err != nil {
			return
		}
		c.sendMessageWithAttempts(ctx, messageBody, q5, attempts, ttl)
	})
}

// skipping Fuzz_Nosy_EventMessage_GetEventType__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-queue-client/client.EventMessage

// skipping Fuzz_Nosy_EventMessage_GetStakingTxHashHex__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-queue-client/client.EventMessage

func Fuzz_Nosy_QueueClient_DeleteMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.QueueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var queueName string
		fill_err = tp.Fill(&queueName)
		if fill_err != nil {
			return
		}
		var receipt string
		fill_err = tp.Fill(&receipt)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		_x1, err := NewQueueClient(config, queueName)
		if err != nil {
			return
		}
		_x1.DeleteMessage(receipt)
	})
}

func Fuzz_Nosy_QueueClient_GetQueueName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.QueueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var queueName string
		fill_err = tp.Fill(&queueName)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		_x1, err := NewQueueClient(config, queueName)
		if err != nil {
			return
		}
		_x1.GetQueueName()
	})
}

func Fuzz_Nosy_QueueClient_Ping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.QueueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var queueName string
		fill_err = tp.Fill(&queueName)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		_x1, err := NewQueueClient(config, queueName)
		if err != nil {
			return
		}
		_x1.Ping(ctx)
	})
}

func Fuzz_Nosy_QueueClient_ReQueueMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.QueueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var queueName string
		fill_err = tp.Fill(&queueName)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var message QueueMessage
		fill_err = tp.Fill(&message)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		_x1, err := NewQueueClient(config, queueName)
		if err != nil {
			return
		}
		_x1.ReQueueMessage(ctx, message)
	})
}

func Fuzz_Nosy_QueueClient_ReceiveMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.QueueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var queueName string
		fill_err = tp.Fill(&queueName)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		_x1, err := NewQueueClient(config, queueName)
		if err != nil {
			return
		}
		_x1.ReceiveMessages()
	})
}

func Fuzz_Nosy_QueueClient_SendMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.QueueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var queueName string
		fill_err = tp.Fill(&queueName)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var messageBody string
		fill_err = tp.Fill(&messageBody)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		_x1, err := NewQueueClient(config, queueName)
		if err != nil {
			return
		}
		_x1.SendMessage(ctx, messageBody)
	})
}

func Fuzz_Nosy_QueueClient_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.QueueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var queueName string
		fill_err = tp.Fill(&queueName)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		_x1, err := NewQueueClient(config, queueName)
		if err != nil {
			return
		}
		_x1.Stop()
	})
}

func Fuzz_Nosy_StakingEvent_GetEventType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var stakerBtcPkHex string
		fill_err = tp.Fill(&stakerBtcPkHex)
		if fill_err != nil {
			return
		}
		var finalityProviderBtcPksHex []string
		fill_err = tp.Fill(&finalityProviderBtcPksHex)
		if fill_err != nil {
			return
		}
		var stakingAmount uint64
		fill_err = tp.Fill(&stakingAmount)
		if fill_err != nil {
			return
		}
		var stateHistory []string
		fill_err = tp.Fill(&stateHistory)
		if fill_err != nil {
			return
		}

		e := NewUnbondingStakingEvent(stakingTxHashHex, stakerBtcPkHex, finalityProviderBtcPksHex, stakingAmount, stateHistory)
		e.GetEventType()
	})
}

func Fuzz_Nosy_StakingEvent_GetStakingTxHashHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var stakerBtcPkHex string
		fill_err = tp.Fill(&stakerBtcPkHex)
		if fill_err != nil {
			return
		}
		var finalityProviderBtcPksHex []string
		fill_err = tp.Fill(&finalityProviderBtcPksHex)
		if fill_err != nil {
			return
		}
		var stakingAmount uint64
		fill_err = tp.Fill(&stakingAmount)
		if fill_err != nil {
			return
		}
		var stateHistory []string
		fill_err = tp.Fill(&stateHistory)
		if fill_err != nil {
			return
		}

		e := NewUnbondingStakingEvent(stakingTxHashHex, stakerBtcPkHex, finalityProviderBtcPksHex, stakingAmount, stateHistory)
		e.GetStakingTxHashHex()
	})
}
