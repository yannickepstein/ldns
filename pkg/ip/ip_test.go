package ip

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	staticIP     = "192.168.0.1"
	hostNotFound = "host not found"
)

func TestIPLookupService(t *testing.T) {
	urls := append([]string{}, "github.com")

	t.Run("resolves to url: ip", func(t *testing.T) {
		service := newIPLookupService(resolve)

		res := service.Lookup(urls)
		expected := fmt.Sprintf("%s: %s", urls[0], staticIP)

		if diff := cmp.Diff(expected, res); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("resolves errors to url: failure", func(t *testing.T) {
		service := newIPLookupService(resolveToError)

		res := service.Lookup(urls)
		expected := fmt.Sprintf("%s: %s", urls[0], hostNotFound)

		if diff := cmp.Diff(expected, res); diff != "" {
			t.Error(diff)
		}
	})
}

func resolve(url string) (string, error) {
	return staticIP, nil
}

func resolveToError(url string) (string, error) {
	return "", fmt.Errorf(hostNotFound)
}
