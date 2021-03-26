package client

type SubscriptionInitialPosition int

const (
	// Latest position which means the start consuming position will be the last message
	SubscriptionPositionLatest SubscriptionInitialPosition = iota

	// Earliest position which means the start consuming position will be the first message
	SubscriptionPositionEarliest
)

type SubscriptionType int

const (
	// Exclusive there can be only 1 consumer on the same topic with the same subscription name
	Exclusive SubscriptionType = iota

	// Shared subscription mode, multiple consumer will be able to use the same subscription name
	// and the messages will be dispatched according to
	// a round-robin rotation between the connected consumers
	Shared

	// Failover subscription mode, multiple consumer will be able to use the same subscription name
	// but only 1 consumer will receive the messages.
	// If that consumer disconnects, one of the other connected consumers will start receiving messages.
	Failover

	// KeyShared subscription mode, multiple consumer will be able to use the same
	// subscription and all messages with the same key will be dispatched to only one consumer
	KeyShared
)

type UniqueID = int64

type ConsumerOptions struct {
	// The topic that this consumer will subscribe on
	Topic string

	// The subscription name for this consumer
	SubscriptionName string

	// InitialPosition at which the cursor will be set when subscribe
	// Default is `Latest`
	SubscriptionInitialPosition

	// Message for this consumer
	// When a message is received, it will be pushed to this channel for consumption
	MessageChannel chan ConsumerMessage

	// Set receive channel size
	BufSize int64

	// Select the subscription type to be used when subscribing to the topic.
	// Default is `Exclusive`
	Type SubscriptionType
}

type Consumer interface {
	// returns the subscription for the consumer
	Subscription() string

	// Message channel
	Chan() <-chan ConsumerMessage

	// Seek to the uniqueID position
	Seek(MessageID) error //nolint:govet

	// Make sure that msg is received. Only used in pulsar
	Ack(ConsumerMessage)

	// Close consumer
	Close()
}
