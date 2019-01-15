package main

import (
    "fmt"
    "github.com/Shopify/sarama"
)

func main() {
    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Partitioner = sarama.NewRandomPartitioner
    config.Producer.Return.Successes = true

    msg := &sarama.ProducerMessage{}
    msg.Topic = "videoreco_message"
    msg.Value = sarama.StringEncoder("this is a good test, my message is good")

    client, err := sarama.NewSyncProducer([]string{"221.122.77.13:9093"}, config)
    if err != nil {
        fmt.Println("producer close, err:", err)
        return
    }

    defer client.Close()

    pid, offset, err := client.SendMessage(msg)
    if err != nil {
        fmt.Println("send message failed,", err)
        return
    }

    fmt.Printf("pid:%v offset:%v\n", pid, offset)
}