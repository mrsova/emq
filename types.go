package emq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Message struct {
	UID     string `json:"uid"`
	Payload any    `json:"payload"`
}

type DoneReport struct {
	Status int `json:"status"`
}

type RejectReport struct {
	Status int    `json:"status"`
	Reason string `json:"reason"`
}

type FailReport struct {
	Status int    `json:"status"`
	Reason string `json:"reason"`
}

type Report struct {
	UID    string
	Done   *DoneReport   `json:"done,omitempty"`
	Reject *RejectReport `json:"reject,omitempty"`
	Fail   *FailReport   `json:"fail,omitempty"`
}

type QueueWorkerPool struct {
	workersCount int
	queue        string
	deliveries   <-chan amqp.Delivery
	results      chan ResultDelivery
}

type ResultDelivery struct {
	WorkerId int
	Delivery *amqp.Delivery
}

type ChannelPoolItemKey struct {
	Name       string
	Type       string
	Queue      string
	Consumer   string
	Exchange   string
	RoutingKey string
}

type ConsumerConfig struct {
	Exchange     string
	ExchangeKind string
	RoutingKey   string
	Routines     int
	Queue        string
}

type PublisherConfig struct {
	Exchange     string
	ExchangeKind string
	RoutingKey   string
}
