package rabbitmq

import (
	"github.com/streadway/amqp"
)

func declareQuorumQueue(ch *amqp.Channel, name string) error {
	_, err := ch.QueueDeclare(
		name,  // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		map[string]interface{}{"x-queue-type": "quorum"}, // arguments
	)
	if err != nil {
		return err
	}
	return nil
}

func declareQueueBinding(ch *amqp.Channel, exchange, queue, routingKey string) error {
	return ch.QueueBind(
		queue,      // queue name
		routingKey, // routing key
		exchange,   // exchange
		false,
		nil,
	)
}

func declareDirectExchange(ch *amqp.Channel, name string) error {
	return ch.ExchangeDeclare(
		name,     // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // argument
	)
}

func declareRetryExchange(ch *amqp.Channel, name string) error {
	return ch.ExchangeDeclare(
		name,                // name
		"x-delayed-message", // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		map[string]interface{}{"x-delayed-type": "direct"},
	)
}

// InitRabbitMQObjects initializes RabbitMQ objects
func InitRabbitMQObjects(config Config) error {
	conn, err := amqp.Dial(config.RabbitMQURI)
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	err = declareDirectExchange(ch, ExchangePaymentCMS)
	if err != nil {
		return err
	}

	err = declareRetryExchange(ch, ExchangePaymentCMSRetry)
	if err != nil {
		return err
	}

	err = declareQuorumQueue(ch, QueuePaymentCMSUnsuspendPayoutMerchant)
	if err != nil {
		return err
	}

	err = declareQuorumQueue(ch, QueuePaymentCMSUnsuspendPayoutShop)
	if err != nil {
		return err
	}

	err = declareQueueBinding(
		ch,
		ExchangePaymentCMS,
		QueuePaymentCMSUnsuspendPayoutMerchant,
		RoutingKeyUnsuspendPayoutMerchant,
	)
	if err != nil {
		return err
	}

	err = declareQueueBinding(
		ch,
		ExchangePaymentCMSRetry,
		QueuePaymentCMSUnsuspendPayoutMerchant,
		RoutingKeyUnsuspendPayoutMerchant,
	)
	if err != nil {
		return err
	}

	err = declareQueueBinding(
		ch,
		ExchangePaymentCMS,
		QueuePaymentCMSUnsuspendPayoutShop,
		RoutingKeyUnsuspendPayoutShop,
	)
	if err != nil {
		return err
	}

	err = declareQueueBinding(
		ch,
		ExchangePaymentCMSRetry,
		QueuePaymentCMSUnsuspendPayoutShop,
		RoutingKeyUnsuspendPayoutShop,
	)
	if err != nil {
		return err
	}

	return nil
}

func InitRabbitMQObjectsWithPreset(config Config, presets []Preset) error {
	conn, err := amqp.Dial(config.RabbitMQURI)
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	for _, p := range presets {
		err = declareDirectExchange(ch, p.DirectExchangeName)
		if err != nil {
			return err
		}

		err = declareRetryExchange(ch, p.RetryExchangeName)
		if err != nil {
			return err
		}
		err = declareQuorumQueue(ch, p.QueueName)
		if err != nil {
			return err
		}

		err = declareQueueBinding(
			ch,
			p.DirectExchangeName,
			p.QueueName,
			p.DirectRoutingKey,
		)
		if err != nil {
			return err
		}

		err = declareQueueBinding(
			ch,
			p.RetryExchangeName,
			p.QueueName,
			p.RetryRoutingKey,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
