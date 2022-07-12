package ip

import (
	"fmt"
	"net"
	"strings"
)

const (
	branchingFactor = 10
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
		queue := make(chan string, branchingFactor)
		out := make(chan string, branchingFactor)
		for _, url := range urls {
			go func(url string) {
				queue <- url
			}(url)
		}
		for i := 0; i < branchingFactor; i++ {
			go func() {
				for url := range queue {
					ip, err := resolve(url)
					if err != nil {
						res := fmt.Sprintf("%s: %s", url, err)
						out <- res
					} else {
						res := fmt.Sprintf("%s: %s", url, ip)
						out <- res
					}
				}
				close(out)
			}()
		}

		var results []string
		for i := 0; i < len(urls); i++ {
			results = append(results, <-out)
		}
		return strings.Join(results, "\n")
	}
}
