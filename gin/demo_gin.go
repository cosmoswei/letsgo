package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func ServerStart() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	r.GET("/index", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/user/:name", func(c *gin.Context) {
		param := c.Param("name")
		c.String(http.StatusOK, "Hello %s", param)
	})

	r.GET("/users", func(c *gin.Context) {
		value := c.Query("name")
		role := c.DefaultQuery("role", "teacher")
		c.String(http.StatusOK, "Hello %s is a %s", value, role)
	})

	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000")
		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})

	r.POST("/post/map", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")
		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	r.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})

	defaultHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}

	v1 := r.Group("/v1")
	{
		v1.GET("/post", defaultHandler)
		v1.GET("/series", defaultHandler)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/post", defaultHandler)
		v2.GET("/series", defaultHandler)
	}

	r.POST("/upload1", func(c *gin.Context) {
		file, err := c.FormFile("file")
		fmt.Println(err)
		if err != nil {
			fmt.Println(file.Filename)
			c.String(http.StatusBadRequest, file.Filename)
		}
	})

	r.POST("/upload2", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			fmt.Println(file.Filename)
		}

		c.String(http.StatusOK, "ok, %d files uploaded!", len(files))
	})

	r.LoadHTMLGlob("gin/templates/*")
	stu1 := &student{
		"huangxuwei", 25,
	}
	stu2 := &student{
		"lisi", 52,
	}

	r.GET("/arr", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arr.html", gin.H{
			"title": "Gin", "stuArr": [2]*student{stu1, stu2},
		})
	})

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/benchmark", Logger())

	authorized := r.Group("/v3")
	authorized.Use(Logger())

	{
		authorized.POST("/login", func(c *gin.Context) {})
		authorized.POST("/submit", func(c *gin.Context) {})
	}

	r.POST("/json", func(c *gin.Context) {
		var stu3 student
		err := c.ShouldBind(&stu3)
		if err != nil {
			fmt.Println("出错了", stu3)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("received stu is %v\n", stu3)
		fmt.Println()
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"name":    stu3.Name,
			"age":     stu3.Age,
		})
	})

	r.Run(":8090")
}

type student struct {
	Name string
	Age  int
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("session", "5s1ZXsc7kkoyha7niIMIhiqm14mkjjgF45dfaegacs3eghw3")
		c.Next()
		since := time.Since(t)
		log.Print(since)
	}
}
