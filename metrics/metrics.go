package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
	"fmt"
	"math/rand"
)


//我们使用的指标仍然时一个计数器requestCount
//以及一个requestLatency来统计数字所处区间的bucket
var (
	NumCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "Num_Total", //改变Name属性，作为稍后产生数字的个数或者数组长度
			Help: "Number of request processed by this service.",
		}, []string{},
	)
	NumScope = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:	"Num_less_than", //改变Name属性，记录数组中小于当前bucket的元素个数
			Help:	"Num less than buckets",
			//依次为数组中小于等于1.0的元素个数，小于等于2.0的元素个数......小于等于10.0的元素个数
			Buckets:   []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		}, []string{},
	)
)


//定义函数NewNum_Produce，在本函数中，我们首先使用产生一个5～10的随机数
//这个随机数我们可以作为数组长度
//然后再使用随机数产生数组元素，每一个数组元素产生之后直接进行计数器加一和
//bucket统计操作，所以数组元素就不需要进行存储了
func NewNum_Produce() {
	Array_Length := rand.Intn(5) + 5 //使用随机数产生数组长度
	fmt.Printf("\n")
	for i := 0; i < Array_Length; i++ {
		//随机生成数字，作为数组元素，由于每次产生一个数组元素就进行一次计数器加一
		//操作，以及判断数组元素属于哪一个区间的操作，所以无需定义数组变量存储数组元素
		num := rand.Intn(10) 
		//调用NumScopeIncrease函数，对计数器进行加操作
		NumScopeIncrease()  
		//在终端打印随机生产的数字
		fmt.Printf("%d ", num) 
		//将数字丢入NumScope，判断新产生的数字属于哪一个bucket
		NumScope.WithLabelValues().Observe(float64(num)) 
	}
}

//对计数器进行加操作,count自行加一
func RequestIncrease() {
	NumCount.WithLabelValues().Add(1)
}

//将NumCount和NumScope采集到的数据给prometheus
func Register() {
	prometheus.MustRegister(NumCount)
	prometheus.MustRegister(NumScope)
}
