package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
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
	MutliChannel()
}

func httpServer(port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func operDateBase() {
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
