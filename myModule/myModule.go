package myModule

import (
	_ "fmt"

	_ "github.com/gorilla/websocket"
	_ "github.com/valyala/fasthttp"
)

func myModule(a int32, b int32) int32 {
	return a * b
}
