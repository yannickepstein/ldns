package ip

import (
	"fmt"
	"net"
	"strings"
)

type ipResolver func(url string) (string, error)

type IPLookupService struct {
	resolve ipResolver
}

func (ips IPLookupService) Lookup(urls []string) string {
	return resolverFactory(ips.resolve)(urls)
}

func StandardIPLookupService() IPLookupService {
	return newIPLookupService(ipNetResolver)
}

func newIPLookupService(resolve ipResolver) IPLookupService {
	return IPLookupService{
		resolve: resolve,
	}
}

func ipNetResolver(url string) (string, error) {
	ips, err := net.LookupIP(url)
	if err != nil {
		return "", err
	}
	if len(ips) <= 0 {
		return "", fmt.Errorf("did not find any IP address")
	}
	return ips[0].String(), nil
}

type resolver func(urls []string) string

func resolverFactory(resolve ipResolver) resolver {
	return func(urls []string) string {
		var results []string
		for _, url := range urls {
			ip, err := resolve(url)
			if err != nil {
				res := fmt.Sprintf("%s: %s", url, err)
				results = append(results, res)
			} else {
				res := fmt.Sprintf("%s: %s", url, ip)
				results = append(results, res)
			}
		}
		return strings.Join(results, "\n")
	}
}
