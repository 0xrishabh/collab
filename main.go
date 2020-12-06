package main

import (
	"fmt"
	"sync"
	"github.com/0xrishabh/collab/service"
	"github.com/alexflint/go-arg"
)



func main(){

	var wg sync.WaitGroup

	var args struct {
		Domain string
		Ipv4 string
	}

	arg.MustParse(&args)
	fmt.Println(args.Domain)
	

	wg.Add(2)

	go service.Http_run(args.Domain)
	go service.Dns_run(args.Ipv4)


	wg.Wait()
}