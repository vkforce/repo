package consumer_group

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func resetOffsets(adminClient *sarama.KafkaAdminClient, groupId string, topic string, offset int64) error {
	// Reset offsets for all partitions
	err := adminClient.ResetOffsets(&sarama.OffsetResetRequest{
		GroupId: groupId,
		TopicPartitions: []sarama.TopicPartition{
			{
				Topic:     &topic,
				Partition: nil, // All partitions
			},
		},
		ToOffset: offset,
	})
	if err != nil {
		return fmt.Errorf("failed to reset offsets: %w", err)
	}
	return nil
}
