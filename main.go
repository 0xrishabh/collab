package main

import (
	"my/collaborator/service"
	_ "fmt"
)



func main(){
	go func(){service.Http_run()}()
	service.Dns_run()
}