package main

import (
	"fmt"
	"github.com/miekg/dns"
	"net"
	"os"
	"strings"
)

func splitRR(rr dns.RR) []string {
	rrary := strings.SplitN(rr.String(), "\t", 5)
	return rrary
}

func checkOpenResolve(ns string) bool {
	dst := "techack.net"
	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(dst), dns.TypeNS)
	m.RecursionDesired = true
	r, _, err := c.Exchange(m, net.JoinHostPort(ns, "53"))

	if err != nil {
		return false
	}

	if len(r.Answer) != 0 {
		for _, rr := range r.Answer {
			rrs := splitRR(rr)
			if strings.Compare(rrs[4], "49.212.146.45") == 0 {
				return true
			}
		}
	}
	return false
}

func main() {
	dst := os.Args[1]
	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(dst), dns.TypeNS)
	m.RecursionDesired = true
	ns := "8.8.8.8"
	r, _, err := c.Exchange(m, net.JoinHostPort(ns, "53"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, rr := range r.Answer {
		rrs := splitRR(rr)
		check := checkOpenResolve(rrs[4])
		if check == true {
			fmt.Println(rrs[4] + " is OpenResolve")
		}
	}
}
