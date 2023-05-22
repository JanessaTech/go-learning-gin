package main

import (
	middlewaredemo "hi-supergirl/go-learning-gin/MiddlewareDemo"
)

func main() {
	//logger.DisabledColor()
	//logger.EnabledColor()
	//logger.SetCustomerLogger()
	middlewaredemo.AddFirstCustomMiddleWare()
}
