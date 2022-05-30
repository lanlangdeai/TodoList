package tests

//package main
//
//import (
//	"github.com/nsqio/go-nsq"
//	"log"
//)
//
//func main() {
//	//初始化生产者
//	config := nsq.NewConfig()
//	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
//	if err != nil{
//		panic(err)
//	}
//	msgBody := []byte("hello")
//	topName := "topic_test"
//	err = producer.Publish(topName, msgBody)
//	if err != nil {
//		log.Fatal(err)
//	}
//	producer.Stop()
//
//}
