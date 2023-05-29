package main

import jwtdemo "hi-supergirl/go-learning-gin/loginlogout/jwtDemo"

func main() {
	//logger.DisabledColor()
	//logger.EnabledColor()
	//logger.SetCustomerLogger()

	//middlewaredemo.AddFirstCustomMiddleWare()
	//middlewaredemo.PropogateVariableViaMWDmo()
	//middlewaredemo.AddMultipleMiddlewares()
	//middlewaredemo.GinRecoveryDemo()

	//bindingdemos.TestBindingViaFormAndJson()
	//validationdemo.SimpleValidator()
	//designreponse.Demo()

	//deadline.Main()

	// I think this solution is not applicable
	//loginlogout.LoginLogoutWithSessionDemo()

	jwtdemo.StartHttPServer()

}
