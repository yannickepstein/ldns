package ip

import (
	"fmt"
	"net"
	"strings"
)

func Query(urls []string) string {
	var results []string
	for _, url := range urls {
		ip, err := lookup(url)
		if err != nil {
			res := fmt.Sprintf("%s: %s", url, err)
			results = append(results, res)
		} else {
			res := fmt.Sprintf("%s: %s", url, ip.String())
			results = append(results, res)
		}
	}
	return strings.Join(results, "\n")
}

func lookup(url string) (net.IP, error) {
	ips, err := net.LookupIP(url)
	if err != nil {
		return nil, err
	}
	if len(ips) <= 0 {
		return nil, fmt.Errorf("did not find any IP address")
	}
	return ips[0], nil
}
