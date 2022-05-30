package tests

//package main
//
//import (
//	"fmt"
//	"github.com/nsqio/go-nsq"
//	"log"
//	"os"
//	"os/signal"
//	"syscall"
//)
//
//type msgHandler struct {}
//
//func processMsg(b []byte) error{
//	content := string(b)
//	fmt.Printf("Msg: %v \n", content)
//	return nil
//}
//
//// 必须实现该接口
//func (h *msgHandler) HandleMessage(m *nsq.Message) error {
//	if len(m.Body) == 0 {
//		return nil
//	}
//
//	err := processMsg(m.Body)
//	m.Finish()
//	return err
//}
//
//
//func main() {
//	config := nsq.NewConfig()
//	consumer, err := nsq.NewConsumer("topic_test", "channel", config)
//	if err != nil {
//		panic(err)
//	}
//
//	consumer.AddHandler(&msgHandler{})
//	//err = consumer.ConnectToNSQLookupd("127.0.0.1:4161")
//	lookupAddr := []string {
//		"127.0.0.1:4161",
//	}
//	err = consumer.ConnectToNSQLookupds(lookupAddr)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	sigChan := make(chan os.Signal, 1)
//	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
//	<- sigChan
//
//	consumer.Stop()
//}
