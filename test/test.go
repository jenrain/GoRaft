package main

import (
	"fmt"
	"time"
)

func main() {
	/*resp, _ := http.Get("http://127.0.0.1:9090/req?message=hello")
	defer resp.Body.Close()
	//200 ok
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)

	buf := make([]byte, 1024)
	for {
		//接收服务端信息
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			fmt.Println("读取完毕")
			res := string(buf[:n])
			fmt.Println(res)
			break
		}
	}*/
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
