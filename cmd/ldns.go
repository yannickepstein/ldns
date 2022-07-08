package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/yannickepstein/ldns/pkg/ip"
	"github.com/yannickepstein/ldns/pkg/lookup"
)

func Execute() {
	queries := map[string]lookup.Query{
		"ip": ip.Query,
	}
	record := flag.String("record", "ip", "Type of record that you want to lookup")
	flag.Parse()
	query, ok := queries[*record]
	if !ok {
		err := fmt.Errorf("unknown type of record: %s", *record)
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	urls := os.Args[1:]
	fmt.Fprintln(os.Stdout, query(urls))
}
