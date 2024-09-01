package rabbitmq

type Preset struct {
	QueueName          string
	DirectExchangeName string
	RetryExchangeName  string
	DirectRoutingKey   string
	RetryRoutingKey    string
}
