package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "letsgo/gin"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"sync"
	"time"
	"unsafe"
)

var name string = " lisi"

const (
	a = iota
	b
	c
	d = 100
	e
	f = iota
	g
)

const (
	h = iota
	i
	j
	k
	l
)

const (
	m = 2
	n
	o
	p
	q
)

const xxxx = "zsdsaczxc"

func demo() {
	fmt.Println("Hello World")

	var name = "huangxuwei"

	var (
		age    int  = 18
		gogogo bool = false
	)

	gogogo = true

	fmt.Println(&age)
	fmt.Println(xxxx)
	fmt.Println(gogogo)
	fmt.Printf("%T", gogogo)
	fmt.Println(name, "xxxx")

	_, x := zkx()

	fmt.Println("a", x)
	fmt.Println(a, b, c, d, e, f, g)
	fmt.Println(h, i, j, k)

	var bb byte = 24

	fmt.Println(bb)
	//var z float64
	//var y float64
	//fmt.Println("请输入2个数：")
	//fmt.Scanln(&z, &y)
	//fmt.Println("z：", z)
	//fmt.Println("y：", y)
	//fmt.Println("y * z = ", z*y)

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d \t", i, j, i*j)
		}
		fmt.Println()
	}

	fmt.Println(swap(2, 1))
	fmt.Println(sum(2, 1, 3, 4, 5, 6, 7, 8))

	//arr := [4]int{1, 2, 3, 4}
	//fmt.Println(arr)
	//updateArr(arr)
	//fmt.Println(arr)

	slice := []int{1, 2, 3, 4}
	fmt.Println(slice)
	updateSlice(slice)
	fmt.Println(slice)

}

func swap(a, b int) (int, int) {
	return b, a
}

func updateArr(arr [4]int) {
	fmt.Println("arr is ", arr)
	arr[1] = 100
	fmt.Println("arr is ", arr)
}

func updateSlice(arr []int) {
	fmt.Println("arr is ", arr)
	arr[1] = 100
	fmt.Println("arr is ", arr)
}

func sum(nums ...int) int {
	sum := 0
	for _, i1 := range nums {
		sum += i1
	}
	return sum
}

func zkx() (int, int) {
	return 1, 2
}

func main() {
	name := "飞雪无情"

	nameP := &name //取地址

	fmt.Println("name变量的值为:", name)

	fmt.Println("name变量的内存地址为:", nameP)

	*nameP = "huangxuwei"

	fmt.Println("name变量的值为:", name)

	i2 := new(int)

	var intP *int

	intP = new(int)

	*intP = 10

	fmt.Println(*i2)

	i := 3
	of := reflect.ValueOf(i)
	typeOf := reflect.TypeOf(i)
	fmt.Println(of, typeOf)
	p := Person{"huangxuwei", 25}
	marshal, err := json.Marshal(p)
	if err == nil {
		fmt.Println(string(marshal))
	}

	respJson := "{\"Name\":\"huangxuwei\",\"Age\":25}"
	json.Unmarshal([]byte(respJson), &p)
	fmt.Println(p)

	x := 10
	ip := &x
	f2 := (*float64)(unsafe.Pointer(ip))
	i3 := *f2 * 3
	fmt.Println(i3)

	fmt.Println(unsafe.Sizeof(true))

	fmt.Println(unsafe.Sizeof(int8(0)))

	fmt.Println(unsafe.Sizeof(int16(10)))

	fmt.Println(unsafe.Sizeof(int32(10000000)))

	fmt.Println(unsafe.Sizeof(int64(10000000000000)))

	fmt.Println(unsafe.Sizeof(int(10000000000000000)))

	fmt.Println(unsafe.Sizeof(string("飞雪无情")))

	fmt.Println(unsafe.Sizeof([]string{"飞雪u无情", "张三"}))

	arr := []int{1, 2, 3, 4, 5}
	arr = append(arr, 2)
	arr = append(arr, 3)
	fmt.Println(len(arr), cap(arr))
	extendSlice()
	s := "飞雪无情"
	fmt.Printf("s的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	b := []byte(s)
	fmt.Printf("b的内存地址：%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)
	s3 := string(b)
	fmt.Printf("s3的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s3)).Data)
}

func init() {

	fmt.Println("init in main.go ")

}

func extendSlice() {
	arr2 := make([]float64, 3, 5)
	arr2 = append(arr2, 1, 2, 3, 4)
	fmt.Println(arr2, len(arr2), cap(arr2)) // 5 10

	slice2 := make([]float32, 3, 5)               // [0 0 0] 长度为3容量为5的切片
	slice2 = append(slice2, 1, 2, 3, 4)           // [0, 0, 0, 1, 2, 3, 4]
	fmt.Println(slice2, len(slice2), cap(slice2)) // 7 12
}

func doOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}

	//用于等待协程执行完毕
	done := make(chan bool)

	//启动10个协程执行once.Do(onceBody)
	for i := 0; i < 10; i++ {
		go func() {
			//把要执行的函数(方法)作为参数传给once.Do方法即可
			once.Do(onceBody)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}

// 10个人赛跑，1个裁判发号施令
func race() {

	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)

	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num, "号已经就位")
			cond.L.Lock()
			cond.Wait() //等待发令枪响
			fmt.Println(num, "号开始跑……")
			cond.L.Unlock()
		}(i)
	}

	//等待所有goroutine都进入wait状态
	time.Sleep(2 * time.Second)

	go func() {
		defer wg.Done()
		fmt.Println("裁判已经就位，准备发令枪")
		fmt.Println("比赛开始，大家准备跑")
		cond.Broadcast() //发令枪响
	}()
	//防止函数提前返回退出
	wg.Wait()
}
func channel() {

	//声明三个存放结果的channel
	firstCh := make(chan string)
	secondCh := make(chan string)
	threeCh := make(chan string)

	//同时开启3个goroutine下载
	go func() {
		firstCh <- downloadFile("firstCh")
	}()

	go func() {
		secondCh <- downloadFile("secondCh")
	}()

	go func() {
		threeCh <- downloadFile("threeCh")
	}()

	//开始select多路复用，哪个channel能获取到值，
	//就说明哪个最先下载好，就用哪个。
	select {
	case filePath := <-firstCh:
		fmt.Println(filePath)
	case filePath := <-secondCh:
		fmt.Println(filePath)
	case filePath := <-threeCh:
		fmt.Println(filePath)
	}
}

func downloadFile(chanName string) string {

	//模拟下载文件,可以自己随机time.Sleep点时间试试
	time.Sleep(time.Second)
	return chanName + ":filePath"
}

func BaseExample() {
	m1 := make(map[string]int)
	m2 := m1
	m1["ko"] = 100
	m2["ok"] = 200

	m3 := map[string]int{
		"ss": 11,
		"bb": 22,
	}
	fmt.Println(m1, m2)

	str := "golang"
	var p *string = &str
	*p = "hello"
	fmt.Println(str)

	for k, v := range m3 {
		fmt.Println(k, v)
	}

	nums := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	for i := 0; i <= 20; i++ {
		func(i int) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("捕获到 panic:", r, i)
				}
			}()
			err, _ := maybeError()
			fmt.Println(err)
		}(i)
	}

	stu := &Student{
		Name: "huangxuwei", Age: 25,
	}
	fmt.Println(stu)
	fmt.Println(stu.hello("ni da ye"))

	stu2 := new(Student)
	stu2.Name = "huangxuwei"
	stu2.Age = 25
	fmt.Println(stu2.hello("bob"))

	var stu3 Human = &Student{
		Name: "zhansan", Age: 25,
	}

	fmt.Println(stu3.getName())

	m := make(map[string]interface{})
	m["name"] = "huangxuwei"
	m["age"] = 25
	m["subject"] = [3]string{"english", "math", "chinese"}
	m["sroce"] = [3]int{99, 100, 99}
	fmt.Println(m)

	for i := 0; i < 6; i++ {
		downWg.Add(1)
		go download("a.com/" + strconv.Itoa(i))
	}
	downWg.Wait()
	fmt.Println("done")

	for i := 0; i < 6; i++ {
		go downloadFromCh("a.com/" + strconv.Itoa(i))
	}

	for i := 0; i < 6; i++ {
		msg := <-ch
		fmt.Println(msg)
	}
}

var downWg = sync.WaitGroup{}

var ch = make(chan string)

func download(url string) {
	rand.Seed(time.Now().UnixNano())
	second := rand.Intn(5)
	fmt.Println(" start to download ", url)
	time.Sleep(time.Second * time.Duration(second))
	downWg.Done()
}

func downloadFromCh(url string) {
	rand.Seed(time.Now().UnixNano())
	second := rand.Intn(5)
	time.Sleep(time.Second * time.Duration(second))
	ch <- url
}

func (stu *Student) hello(person string) string {
	return fmt.Sprintf("hello, %s! i am %s  and i am %d", person, stu.Name, stu.Age)
}

type Human interface {
	getName() string
}

type Worker struct {
	name   string
	gender string
}

func (w Worker) getName() string {
	return w.name
}

type Student struct {
	Name string
	Age  int
}

func (stu *Student) getName() string {
	return stu.Name
}

func maybeError() (error, string) {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	// 生成一个 0 到 99 的随机整数
	randomInt := rand.Intn(99)

	if randomInt%2 == 0 {
		panic("is even ")
	} else {
		return errors.New("is old"), ""
	}
}

func httpServer(port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func operaDateBase() {
	db, err := sql.Open("mySQl", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	{
		query := `
            CREATE TABLE users (
                id INT AUTO_INCREMENT,
                username TEXT NOT NULL,
                Password TEXT NOT NULL,
                created_at DATETIME,
                PRIMARY KEY (id)
            );`
		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}

		{ // Insert a new user
			username := "johndoe"
			password := "secret"
			createdAt := time.Now()

			result, err := db.Exec(`INSERT INTO users (username, Password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
			if err != nil {
				log.Fatal(err)
			}

			id, err := result.LastInsertId()
			fmt.Println(id)
		}

	}
}
