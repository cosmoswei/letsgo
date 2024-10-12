package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/exp/constraints"
	"golang.org/x/net/context"
	_ "letsgo/gin"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"sync"
	"time"
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

func Describe(i interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", i, i)
}

// 定义一个泛型类型，T 是类型参数
type Stack[T any] struct {
	elements []T
}

// 为泛型类型定义方法
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() T {
	n := len(s.elements)
	element := s.elements[n-1]
	s.elements = s.elements[:n-1]
	return element
}

// 泛型函数，T 是类型参数
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// AddTest 函数，约束 T 只能是数字类型
func Add[T constraints.Ordered](a, b T) T {
	return a + b
}

// 定义一个函数，接收两个不同的类型参数
func Swap[T, U any](a T, b U) (U, T) {
	return b, a
}

// 使用泛型函数求最小值
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// 定义泛型 List
type List[T any] struct {
	items []T
}

func (l *List[T]) Add(item T) {
	l.items = append(l.items, item)
}

func (l *List[T]) Get(index int) T {
	return l.items[index]
}

type Project struct{}

func (p *Project) deferError() {
	if err := recover(); err != nil {
		fmt.Println("recover: ", err)
	}
}
func (p *Project) exec(msgchan chan interface{}) {
	defer p.deferError()
	for msg := range msgchan {
		m := msg.(int)
		fmt.Println("msg: ", m)
	}
}
func (p *Project) run(msgchan chan interface{}) {
	for {
		go p.exec(msgchan)
		time.Sleep(time.Second * 2)
	}
}
func (p *Project) Main() {
	a := make(chan interface{}, 100)
	go p.run(a)
	go func() {
		for {
			a <- "1"
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 20)
}

func main() {
	orderPrint()
	strs := []string{"abc", "abcd", "aa", "100a"}
	res := CountArrStr(strs, 2)
	fmt.Println(res['b'])
}

type LetterFreq map[rune]int

func CountArrStr(strs []string, currency int) LetterFreq {
	res := make(map[rune]int, 7)
	ch := make(chan string, len(strs))
	cur := make(chan struct{}, currency)
	wg1 := sync.WaitGroup{}
	mu := sync.Mutex{}
	for _, str := range strs {
		wg1.Add(1)
		go func(s string) {
			defer wg1.Done()
			cur <- struct{}{}
			ch <- s
			<-cur
		}(str)
	}
	go func() {
		wg1.Wait()
		close(ch)
	}()

	for str := range ch {
		subCount(str, res, &mu)
	}

	return LetterFreq(res)
}
func subCount(str string, res map[rune]int, mu *sync.Mutex) {
	for _, char := range str {
		mu.Lock()
		res[char]++
		mu.Unlock()
	}
}

func Interview2() {
	var a uint = 0
	var b uint = 1
	c := a - b
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(c)
	ax := [3]int{1, 2, 3} // 数组
	for k, v := range ax {
		if k == 0 {
			ax[0], ax[1] = 100, 200
			fmt.Println(ax)
		}
		ax[k] = 100 + v
	}
	fmt.Println(ax)

	az := []int{1, 2, 3} // 切片
	for k, v := range az {
		if k == 0 {
			az[0], az[1] = 100, 200
			fmt.Println(az)
		}
		az[k] = 100 + v
	}
	fmt.Println(az)

	//var wg sync.WaitGroup
	//c1 := make(chan bool, 1) // 用于协调 cat 的信号
	//c2 := make(chan bool, 1) // 用于协调 dog 的信号
	//c3 := make(chan bool, 1) // 用于协调 fish 的信号
	//
	//// 启动三个 goroutine
	//wg.Add(3)
	//go printCat(c1, c2, &wg)
	//go printDog(c2, c3, &wg)
	//go printFish(c3, c1, &wg)
	//
	//// 初始化信号
	//c1 <- true // 让 cat goroutine 先开始

	//// 等待所有 goroutine 完成
	//wg.Wait()

	//ctx, cancel := context.WithCancel(context.Background())
	//go worker(ctx, "worker1 ")
	//go worker(ctx, "worker2 ")
	//time.Sleep(time.Second * 2)
	//cancel()
	//time.Sleep(time.Second * 1)

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*1)
	defer cancelFunc()
	select {
	case <-time.After(time.Second * 2):
		fmt.Println("finish work")
	case <-timeout.Done():
		fmt.Println("timeout")
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, ctxkey("userId"), "18411100111")
	go worker2(ctx)
	time.Sleep(time.Second * 1)
}

type ctxkey string

func worker2(ctx context.Context) {
	value := ctx.Value(ctxkey("userId"))
	if value != nil {
		fmt.Println("receive a value from context, the value is ", value)
	} else {
		fmt.Println("no value")
	}
}

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("收到退出信号，我退出了")
			return
		default:
			fmt.Println(name + "正在工作")
			time.Sleep(time.Second)
		}
	}
}

func printCat(c1, c2 chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-c1                              // 等待前一个信号
		fmt.Println("cat")                // 打印 cat
		time.Sleep(10 * time.Millisecond) // 模拟处理时间
		c2 <- true                        // 发送信号给下一个
	}
}

func printDog(c2, c3 chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-c2               // 等待前一个信号
		fmt.Println("dog") // 打印 dog
		time.Sleep(10 * time.Millisecond)
		c3 <- true // 发送信号给下一个
	}
}

func printFish(c3, c1 chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-c3                // 等待前一个信号
		fmt.Println("fish") // 打印 fish
		time.Sleep(10 * time.Millisecond)
		c1 <- true // 发送信号给下一个
	}
}
func Interview() {
	//p := new(Project)
	//p.Main()
	extendSliceErr()
	fmt.Println(live(), " World")
	fmt.Println(a1, b2, name1, c3)

	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]
	str2[1] = "new"
	fmt.Println(str1)
	str2 = append(str2, "z", "x", "y")
	fmt.Println(str1)
	fmt.Println(&Student{Name: "menglu"} == &Student{Name: "menglu"})
	fmt.Println(Student{Name: "menglu"} == Student{Name: "menglu"})
	fmt.Println([...]string{"1"} == [...]string{"1"})
	//fmt.Println([]string{"1"} == []string{"1"})

	chain = "main"
	////A()
	fmt.Println(chain)

	//var wg2 sync.WaitGroup
	//wg2.Add(1)
	//go func() {
	//	time.Sleep(time.Second * 2)
	//	wg2.Done()
	//	wg2.Add(1)
	//}()
	//wg2.Wait()

	//var mu MyMutex
	//mu.Lock()
	//var mu2 = mu
	//mu.count++
	//mu.Unlock()
	//mu2.Lock()
	//mu2.count++
	//mu2.Unlock()
	//fmt.Println(mu.count, mu2.count)

	//var ch chan int
	//var count int
	//go func() {
	//	ch <- 1
	//}()
	//go func() {
	//	count++
	//	time.Sleep(1 * time.Second)
	//	close(ch)
	//}()
	//<-ch
	//fmt.Println(count)
}

type MyMutex struct {
	sync.Mutex
	count int
}

var mu = sync.Mutex{}
var chain string

func A() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> A"
	B()
}
func B() {
	chain = chain + " --> B"
	C()
}
func C() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> C"
}

const (
	a1 = iota
	b2 = iota
)

const (
	name1 = "huangxuwei"
	name2 = "huangxuwei"
	c3    = iota
	d4    = iota
)

type People interface {
	Show()
}

type Student2 struct {
}

func (student *Student2) Show() {
}

func live() People {
	var stu Student2
	return &stu
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	fmt.Println(stus)
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	for s, v := range m {
		fmt.Println(s)
		fmt.Println(v)
	}
}

func gpmTest() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func defer_Call() {
	defer func() { fmt.Println("print first") }()
	defer func() { fmt.Println("print second") }()
	defer func() { fmt.Println("print third") }()
	panic("panic")
}

func generic() {
	// 使用泛型函数
	intSlice := []int{1, 2, 3}
	stringSlice := []string{"Hello", "World"}

	PrintSlice(intSlice)    // 输出: 1 2 3
	PrintSlice(stringSlice) // 输出: Hello World

	// 使用泛型类型 Stack
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	fmt.Println(intStack.Pop()) // 输出: 2
	stringStack := Stack[string]{}
	stringStack.Push("Go")
	stringStack.Push("Generics")
	fmt.Println(stringStack.Pop()) // 输出: Generics

	fmt.Println(Add(3, 4))     // 输出: 7
	fmt.Println(Add(2.5, 3.5)) // 输出: 6

	x, y := Swap(1, "Go")
	fmt.Println(x, y) // 输出: Go 1

	fmt.Println(Min(10, 20))       // 输出: 10
	fmt.Println(Min(3.5, 2.5))     // 输出: 2.5
	fmt.Println(Min("Go", "Java")) // 输出: Go

	intList := List[int]{}
	intList.Add(1)
	intList.Add(2)
	fmt.Println(intList.Get(1)) // 输出: 2

	stringList := List[string]{}
	stringList.Add("Hello")
	stringList.Add("Go")
	fmt.Println(stringList.Get(0)) // 输出: Hello
}

func InterfaceDemo() {
	Describe(42)
	Describe("hello")
	Describe(true)

	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s) // 输出: hello

	var i2 interface{} = 2
	s, ok := i2.(string)
	if ok {
		fmt.Println(s)
	} else {
		fmt.Println("类型断言失败")
	}

	var s2 Shape

	s2 = Circle{Radius: 5}
	fmt.Println("Circle Area:", s2.Area())
	fmt.Println("Circle Perimeter:", s2.Perimeter())

	s2 = Rectangle{Width: 3, Height: 4}
	fmt.Println("Rectangle Area:", s2.Area())
	fmt.Println("Rectangle Perimeter:", s2.Perimeter())
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func TypeCheck(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("i 是字符串: %s\n", v)
	case int:
		fmt.Printf("i 是整数: %d\n", v)
	default:
		fmt.Printf("未知类型: %T\n", v)
	}
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type PP struct {
	Role string
}

func (c PP) Write(p []byte) (n int, err error) {
	panic("")
}
func (c PP) Read(p []byte) (n int, err error) {
	panic("")
}

func orderPrint2() {
	var wg sync.WaitGroup

	// 创建4个channel，用来控制4个goroutine的顺序
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
	ch4 := make(chan struct{})

	// 启动4个goroutine，编号分别为1、2、3、4
	wg.Add(4)
	go printNumber(1, ch1, ch2, &wg) // goroutine 1
	go printNumber(2, ch2, ch3, &wg) // goroutine 2
	go printNumber(3, ch3, ch4, &wg) // goroutine 3
	go printNumber(4, ch4, ch1, &wg) // goroutine 4

	// 启动时，先给第一个goroutine发送启动信号
	ch1 <- struct{}{}

	// 防止主程序过早退出，等待goroutine执行
	wg.Wait()
}

func printNumber(id int, currentChan, nextChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		<-currentChan               // 等待信号
		fmt.Println(id)             // 打印自己的编号
		time.Sleep(1 * time.Second) // 每秒打印一次
		nextChan <- struct{}{}      // 通知下一个 goroutine 开始
	}
}

func orderPrint() {
	c1 := make(chan bool, 1)
	c2 := make(chan bool, 1)
	c3 := make(chan bool, 1)
	c4 := make(chan bool, 1)
	round := 3
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			<-c1
			fmt.Println("1")
			time.Sleep(time.Second)
			c2 <- true
		}
	}()
	go func() {
		for {
			<-c2
			fmt.Println("2")
			time.Sleep(time.Second)
			c3 <- true
		}
	}()
	go func() {
		for {
			<-c3
			fmt.Println("3")
			time.Sleep(time.Second)
			c4 <- true
		}
	}()
	go func() {
		for {
			<-c4
			fmt.Println("4")
			time.Sleep(time.Second)
			round--
			if round <= 0 {
				close(c1)
				wg.Done()
				return
			}
			c1 <- true
		}
	}()
	c1 <- true
	wg.Wait()
	fmt.Println("print is finished")
}

// https://learn.lianglianglee.com/%E4%B8%93%E6%A0%8F/Go%20%E8%AF%AD%E8%A8%80%E9%A1%B9%E7%9B%AE%E5%BC%80%E5%8F%91%E5%AE%9E%E6%88%98
//func project() {
//	name := "飞雪无情"
//
//	nameP := &name //取地址
//
//	fmt.Println("name变量的值为:", name)
//
//	fmt.Println("name变量的内存地址为:", nameP)
//
//	*nameP = "huangxuwei"
//
//	fmt.Println("name变量的值为:", name)
//
//	i2 := new(int)
//
//	var intP *int
//
//	intP = new(int)
//
//	*intP = 10
//
//	fmt.Println(*i2)
//
//	i := 3
//	of := reflect.ValueOf(i)
//	typeOf := reflect.TypeOf(i)
//	fmt.Println(of, typeOf)
//	p := Person{"huangxuwei", 25}
//	marshal, err := json.Marshal(p)
//	if err == nil {
//		fmt.Println(string(marshal))
//	}
//
//	respJson := "{\"Name\":\"huangxuwei\",\"Age\":25}"
//	json.Unmarshal([]byte(respJson), &p)
//	fmt.Println(p)
//
//	x := 10
//	ip := &x
//	f2 := (*float64)(unsafe.Pointer(ip))
//	i3 := *f2 * 3
//	fmt.Println(i3)
//
//	fmt.Println(unsafe.Sizeof(true))
//
//	fmt.Println(unsafe.Sizeof(int8(0)))
//
//	fmt.Println(unsafe.Sizeof(int16(10)))
//
//	fmt.Println(unsafe.Sizeof(int32(10000000)))
//
//	fmt.Println(unsafe.Sizeof(int64(10000000000000)))
//
//	fmt.Println(unsafe.Sizeof(int(10000000000000000)))
//
//	fmt.Println(unsafe.Sizeof(string("飞雪无情")))
//
//	fmt.Println(unsafe.Sizeof([]string{"飞雪u无情", "张三"}))
//
//	arr := []int{1, 2, 3, 4, 5}
//	arr = append(arr, 2)
//	arr = append(arr, 3)
//	fmt.Println(len(arr), cap(arr))
//	extendSliceErr()
//	s := "飞雪无情"
//	fmt.Printf("s的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
//	b := []byte(s)
//	fmt.Printf("b的内存地址：%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)
//	s3 := string(b)
//	fmt.Printf("s3的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s3)).Data)
//}

func init() {
	fmt.Println("init in main.go ")
}

func extendSliceErr() {
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
