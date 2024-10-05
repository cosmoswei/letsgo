package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
)

func TestAdd(t *testing.T) {
	if ans := AddTest(1, 2); ans != 3 {
		t.Error("add(1, 2) should be equal to 3")
	}
}

func TestMul(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		if Mul(2, 3) != 6 {
			t.Error("Mul(2,3) should be equal to 6")
		}
	})
	t.Run("neg", func(t *testing.T) {
		if Mul(-1, 2) != -2 {
			t.Error("Mul(-1, 2) should be equal to -2")
		}
	})
}

func TestMulBest(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 2, 3, 6},
		{"neg", -2, 3, -6},
		{"zero", 0, 3, 0},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := Mul(c.A, c.B); ans != c.Expected {
				t.Fatalf("%d * %d expected %d, but %d got",
					c.A, c.B, c.Expected, ans)
			}
		})
	}
}

type calcCase struct{ A, B, Expected int }

func createMulTestCase(t *testing.T, c *calcCase) {
	t.Helper()
	t.Run("", func(t *testing.T) {
		if ans := Mul(c.A, c.B); ans != c.Expected {
			t.Fatalf("%d * %d expected %d, but %d got", c.A, c.B, c.Expected, ans)
		}
	},
	)
}

func TestMul3(t *testing.T) {
	setUp()
	createMulTestCase(t, &calcCase{1, 2, 2})
	createMulTestCase(t, &calcCase{2, 6, 12})
	createMulTestCase(t, &calcCase{3, 3, 9})
	createMulTestCase(t, &calcCase{0, 0, 0})
	tarDown()
}

func setUp() {
	fmt.Printf("Before all tests\n")
}

func tarDown() {
	fmt.Printf("after all tests\n")
}

func handlerError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("got error: %s", err)
	}
}

func TestConn(t *testing.T) {
	ln, err := net.Listen("tcp", ":0")
	handlerError(t, err)
	defer ln.Close()
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

	})
	http.Serve(ln, nil)

	resp, err := http.Get("http://" + ln.Addr().String() + "/hello")
	handlerError(t, err)

	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	handlerError(t, err)

	if string(all) != "hello world" {
		t.Fatal(string(all))
	}
}

func TestConn3(t *testing.T) {
	//req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8080/hello", nil)
	//w := httptest.NewRecorder()
	//handlerError(w, req)
	//bytes, _ := ioutil.ReadAll(w.Result().Body)
	//if string(bytes) != "hello world" {
	//	t.Fatal(string(bytes))
	//}
}

func BenchmarkHello(b *testing.B) {
	// go test -benchmem -bench .
	for i := 0; i < b.N; i++ {
		fmt.Printf("hello")
	}
}

func BenchmarkParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}
