package loginlogout

import (
	jwtdemo "hi-supergirl/go-learning-gin/loginlogout/jwtDemo"
	sessiondemo "hi-supergirl/go-learning-gin/loginlogout/sessionDemo"
)

func LoginLogoutWithSessionDemo() {
	sessiondemo.LoginLogoutWithSessionDemo()
}

func JwtDemo() {
	jwtdemo.StartHttPServer()
}
