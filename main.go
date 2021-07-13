package main

import (
	"strconv"

	"github.com/kzw200015/go-list/args"
	"github.com/kzw200015/go-list/routes"
)

func main() {
	r := routes.CreateRouter()

	r.Run(":" + strconv.Itoa(args.GetPort()))
}
