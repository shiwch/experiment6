package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"strconv"
	"math/rand"
	"time"
	"example/metrics"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main(){
	http.HandleFunc("/abc", index)  //收到字符abc直接调用index函数
	//收到metrics字符时会将metrics之前收集的数据发送给prometheus
	http.Handle("/metrics", promhttp.Handler()) 
	metrics.Register()
	err := http.ListenAndServe(":5565", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func index(w http.ResponseWriter, r *http.Request) {
	//调用NewNum_Produce函数
	metrics.NewNum_Produce()
	num:=os.Getenv("Num")
	if num==""{
		Fibonacci(10)
		_,err:=w.Write([]byte("there is no env Num. Computation successed\n"))
		if err!=nil{
			log.Println("err:"+err.Error()+" No\n")
		}
	}else{
		numInt,_:=strconv.Atoi(num)
		Fibonacci(numInt)
		_,err:=w.Write([]byte("there is env Num. Computation successed\n"))
		if err!=nil{
			log.Println("err:"+err.Error()+" Yes\n")
		}
	}
}

func Fibonacci(n int)int{
	if n<=2{
		return 1
	}else{
		return Fibonacci(n-1)+Fibonacci(n-2)
	}
}
