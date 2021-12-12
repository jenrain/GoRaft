package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/rpc"
	"time"
)

// 等待节点访问
func (rf *Raft) getRequest(writer http.ResponseWriter, request *http.Request) {
	// 解析表单参数
	request.ParseForm()
	//http://localhost:8080/req?message=ohmygod
	fmt.Println()
	fmt.Println("开始处理请求1。。。")
	fmt.Println()
	if len(request.Form["message"]) > 0 && rf.currentLeader != "-1" {
		fmt.Println()
		fmt.Println("开始处理请求2。。。")
		fmt.Println()
		message := request.Form["message"][0]
		m := new(Message)
		m.MsgID = getRandom()
		m.Msg = message
		// 接收消息以后，直接转发到领导者
		fmt.Println("http监听到了消息，准备发送给领导者，消息id：", m.MsgID)
		port := nodeTable[rf.currentLeader]
		rp, err := rpc.DialHTTP("tcp", "127.0.0.1"+port)
		if err != nil {
			log.Panic(err)
		}
		b := false
		err = rp.Call("Raft.LeaderReceiveMessage", m, &b)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println("消息是否已经发送到领导者：", b)
		writer.Write([]byte("ok!!!"))
	}
}

func (rf *Raft) httpListen() {
	// 创建getRequest() 回调方法
	http.HandleFunc("/req", rf.getRequest)
	fmt.Println("监听9090")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Println(err)
		return
	}
}

// 返回一个十位数的随机数，作为消息的ID
func getRandom() int {
	for {
		result := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10000000000)
		if result > 1000000000 {
			return result
		}
	}
}
