package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-macaron/auth"
	"gopkg.in/macaron.v1"
)

type hello struct {
	d int
	s string
}

func Hello(u, p string) bool {
	return true
}

func main() {

	m := macaron.Classic()
	m.Use(macaron.Static("public"))
	m.Use(macaron.Renderer())
	m.Use(auth.BasicFunc(Hello))
	h := hello{
		d: 4,
		s: "fjhdkjf",
	}

	//every struct that we will receive from myhandler will have same value as h here
	m.Map(h)

	m.Use(func() {
		fmt.Println("u are safe")
	})

	m.Get("/", myHandler, myHandler1)

	//serve
	//m.Run()

	m.Get("/set", func(ctx *macaron.Context) {
		ctx.SetCookie("user", "Unknwon", 1000)
	})

	m.Get("/get", func(ctx *macaron.Context) string {
		ctx.Data["hello"] = hello{
			d: 5,
			s: "hello erm",
		}
		return ctx.GetCookie("user")
	})
	m.Get("/hello/*", func(ctx *macaron.Context) string {
		return "Hello " + ctx.Params("*")
	})
	m.Get("/date/*/*/*/events", func(ctx *macaron.Context) string {
		return fmt.Sprintf("Date: %s/%s/%s", ctx.Params("*0"), ctx.Params("*1"), ctx.Params("*2"))
	})
	m.Get("/user/:id([0-9]+)", func(ctx *macaron.Context) string {
		return fmt.Sprintf("User ID: %s", ctx.Params(":id"))
	})
	m.NotFound(func() string {
		// Custom handle for 404
		return "error hoise"
	})

	log.Println("server is running .....")
	log.Println(http.ListenAndServe("0.0.0.0:4000", m))
}

func myHandler(ctx *macaron.Context, hs hello, h hello) {

	fmt.Println(hs)
	fmt.Println(h)

	fmt.Println("hi")

}

func myHandler1(ctx *macaron.Context) {
	fmt.Println("hello")
	//return " path is : " + ctx.Req.RequestURI
	// override ResponseWriter with our wrapper ResponseWriter
	//ctx.MapTo(rw, (*http.ResponseWriter)(nil))

	ctx.JSON(http.StatusCreated, " path is : "+ctx.Req.RequestURI)
}
