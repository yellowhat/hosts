// https://github.com/0xERR0R/blocky/blob/main/lists/parsers/hosts.go
package main

import (
	"fmt"
	"net"
	"regexp"
	"strings"

	"golang.org/x/net/idna"
)

const (
	maxDomainNameLength = 255 // https://www.rfc-editor.org/rfc/rfc1034#section-3.1
	dnsLabelPattern     = `[a-zA-Z0-9_-]{1,63}`
)

// Validate a domain name, but with extra flexibility:
// - no restriction on the start or end of labels
// https://www.rfc-editor.org/rfc/rfc1034#section-3.5
var domainNameRegex = regexp.MustCompile(`^` + dnsLabelPattern + `(\.` + dnsLabelPattern + `)*[\._]?$`)

func validateDomainName(host string) error {
	if len(host) > maxDomainNameLength {
		return fmt.Errorf("domain name is too long: %s", host)
	}

	if domainNameRegex.MatchString(host) {
		return nil
	}

	return fmt.Errorf("invalid domain name: %s", host)
}

func isRegex(host string) bool {
	return strings.HasPrefix(host, "/") && strings.HasSuffix(host, "/")
}

func validateHostsListEntry(host string) error {
	if net.ParseIP(host) != nil {
		return nil
	}

	if isRegex(host) {
		_, err := regexp.Compile(host)

		return err
	}

	return validateDomainName(host)
}

func normalizeHostsListEntry(host string) (string, error) {
	var err error
	// Lookup is the profile preferred for DNS queries, we use Punycode here as it does less validation.
	// That avoids rejecting domains in a list for reasons that amount to "that domain should not be used"
	// since the goal of the list is to determine whether the domain should be used or not, we leave
	// that decision to it.
	idnaProfile := idna.Punycode

	if !isRegex(host) {
		host, err = idnaProfile.ToASCII(host)
		if err != nil {
			return "", fmt.Errorf("%w: %s", err, host)
		}
	}

	// remove optional start and end markers for ABP styled lists
	host = strings.TrimPrefix(host, "||")
	host = strings.TrimSuffix(host, "^")

	if err := validateHostsListEntry(host); err != nil {
		return "", err
	}

	return host, nil
}
