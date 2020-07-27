package main

import (
	kafka "github.com/segmentio/kafka-go"

	namespaces "github.com/circuit-platform/namespaces/model"

	models "github.com/circuit-platform/models-utils"
)

func SyncNamespaces(DatabaseSource string, KafkaSource string) {
	index := namespaces.CreateNamespacesIndex(DatabaseSource, "private")

	models.SyncIndex(
		"raven.public",
		kafka.ReaderConfig{
			Brokers: []string{KafkaSource},
			Partition: 0,
		},
		index,
	)
}