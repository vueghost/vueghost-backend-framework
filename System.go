/*
 * Vueghost Tech, FZE © 2018- 2020. Unauthorized copying of this file,
 * via any medium is strictly prohibited proprietary and confidential,
 * can not be copied and/or distributed without the express permission of Vueghost
 */
package System

import (
	"net/http"
	"reflect"
)

const (
	Version = 1.0
)

const ACIILogo string = `
██╗   ██╗██╗   ██╗███████╗ ██████╗ ██╗  ██╗ ██████╗ ███████╗████████╗
██║   ██║██║   ██║██╔════╝██╔════╝ ██║  ██║██╔═══██╗██╔════╝╚══██╔══╝
██║   ██║██║   ██║█████╗  ██║  ███╗███████║██║   ██║███████╗   ██║   
╚██╗ ██╔╝██║   ██║██╔══╝  ██║   ██║██╔══██║██║   ██║╚════██║   ██║   
 ╚████╔╝ ╚██████╔╝███████╗╚██████╔╝██║  ██║╚██████╔╝███████║   ██║   
  ╚═══╝   ╚═════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚══════╝   ╚═╝ Framework v.1.1
  ╚═══╝   ╚═════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚══════╝   ╚═╝ GHOST ➕✖️

BACKEND SERVER GOLANG - CopyRight © 2018-2020 https://www.vueghost.com.
DEVELOPED BY: Ahmed.M.Yassin 
********************************************************************************************
`

//Run
func Run(Struct interface{},
	method string,
	httpResponseWriter http.ResponseWriter,
	httpRequest *http.Request) {
	structType := reflect.TypeOf(Struct)
	structPointer := reflect.New(structType)
	structPointer.MethodByName("Context").Call([]reflect.Value{
		reflect.ValueOf(httpResponseWriter),
		reflect.ValueOf(httpRequest),
	})
	structPointer.MethodByName("Constructor").Call([]reflect.Value{})
	structPointer.MethodByName(method).Call([]reflect.Value{})
	structPointer.MethodByName("Finalize").Call([]reflect.Value{})
}

//Startup
func Startup(configure Configure) {

}
