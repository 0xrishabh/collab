package main

import (
	"github.com/0xishabh/collab/service"
	_ "fmt"
)



func main(){
	go func(){service.Http_run()}()
	service.Dns_run()
}