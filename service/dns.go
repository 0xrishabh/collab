package service

import (
	"net"
	"log"
	"fmt"
	"github.com/miekg/dns"
	"github.com/fatih/color"
)

var yellow = color.New(color.FgYellow).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()

type Handler struct {
	Ipv4 string
}


func (this Handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)

	fmt.Println("DNS for ",red(msg.Question[0].Name), " from ", yellow(w.RemoteAddr().String()))
	switch r.Question[0].Qtype {
	case dns.TypeA:
		msg.Authoritative = true
		domain := msg.Question[0].Name
		address := this.Ipv4
		msg.Answer = append(msg.Answer, &dns.A{
			Hdr: dns.RR_Header{ Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60 },
			A: net.ParseIP(address),
		})
	}
	w.WriteMsg(&msg)
}

func Dns_run(Ipv4 string) {
	srv := &dns.Server{Addr: ":53", Net: "udp"}
	srv.Handler = Handler{Ipv4}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}else{
		fmt.Println("Dns server running on :53")	
	}

}