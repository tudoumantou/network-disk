package mq

import (
	"github.com/streadway/amqp"
	"network-disk/config"
	"log"
)

var conn *amqp.Connection
var channel *amqp.Channel

func initChannel() bool {
 	// 1.判断channel是否已经创建过
	if channel != nil{
		return true
	}
	// 2.获得rabbitmq的一个连接
	conn, err := amqp.Dial(config.RabbitURL)
	if err != nil{
		log.Println(err.Error())
		return false
	}
	// 3.打开一个channel，用于消息的发布和接收等
	channel, err = conn.Channel()
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

// 发布消息
func Publish(exchange, routingKey string, msg []byte) bool {
	// 1. 判断channel是否正常
	if !initChannel(){
		return false
	}

	// 2.执行消息发布动作
	err := channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: msg,
		})

	if err != nil{
		log.Println(err.Error())
		return false
	}
	return true
}