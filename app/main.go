package main

import (
	"github.com/sugar-hit/banana/app/web"
)

func main() {
	route := web.Reg()
	web.Start(route)
}
