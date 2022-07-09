package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/yannickepstein/ldns/pkg/ip"
	"github.com/yannickepstein/ldns/pkg/lookup"
)

func Execute() {
	services := map[string]lookup.LookupService{
		"ip": ip.StandardIPLookupService(),
	}
	record := flag.String("record", "ip", "Type of record that you want to lookup")
	flag.Parse()
	service, ok := services[*record]
	if !ok {
		err := fmt.Errorf("unknown type of record: %s", *record)
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	urls := os.Args[1:]
	fmt.Fprintln(os.Stdout, service.Lookup(urls))
}
