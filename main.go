package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "letsgo/gin"
	"log"
	"math/rand"
	"net/http"
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

func main() {
	Context()
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
