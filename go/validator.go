// https://github.com/0xERR0R/blocky/blob/main/lists/parsers/hosts.go
package main

import (
	"fmt"
	"net"
	"regexp"
	"strings"
	"unicode/utf8"

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
	return len(host) >= 2 && strings.HasPrefix(host, "/") && strings.HasSuffix(host, "/")
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

// idnaProfile maps an entry to the form DNS clients put on the wire: UTS #46 lookup mapping
// (case folding, NFC, width folding), then punycode. STD3 and label validity checks are off,
// they reject entries for reasons that amount to "that domain should not be used".
var idnaProfile = idna.New(idna.MapForLookup(), idna.StrictDomainName(false), idna.ValidateLabels(false))

func normalizeHostsListEntry(host string) (string, error) {
	// remove optional start and end markers for ABP styled lists
	host = strings.TrimPrefix(host, "||")
	host = strings.TrimSuffix(host, "^")

	// URL-style IPv6 literal, as in "||[2001:db8::1]^"
	if ip, ok := unwrapIPv6Literal(host); ok {
		return ip, nil
	}

	if !isRegex(host) {
		var err error

		host, err = toASCII(host)
		if err != nil {
			return "", err
		}
	}

	if err := validateHostsListEntry(host); err != nil {
		return "", err
	}

	// Blocky strips the root dot from query names, so "example.com." would never match
	return strings.TrimSuffix(host, "."), nil
}

func unwrapIPv6Literal(s string) (string, bool) {
	if len(s) < 2 || s[0] != '[' || s[len(s)-1] != ']' {
		return "", false
	}

	ip := s[1 : len(s)-1]
	if !strings.Contains(ip, ":") || net.ParseIP(ip) == nil {
		return "", false
	}

	return ip, true
}

// toASCII only runs IDNA on entries with non-ASCII bytes or an "xn--" prefix, it is slow.
func toASCII(host string) (string, error) {
	if !needsIDNA(host) {
		return host, nil
	}

	ascii, err := idnaProfile.ToASCII(host)
	if err != nil {
		return "", fmt.Errorf("%w: %s", err, host)
	}

	// A label that vanishes in the mapping would turn "*.com.xn--" into "*.com."
	if nonEmptyLabels(ascii) < nonEmptyLabels(host) {
		return "", fmt.Errorf("label vanishes in IDNA mapping: %s", host)
	}

	return ascii, nil
}

// nonEmptyLabels counts labels with content, UTS #46 also maps the ideographic
// and fullwidth full stops to '.'
func nonEmptyLabels(s string) int {
	n, inLabel := 0, false

	for _, r := range s {
		switch r {
		case '.', '\u3002', '\uff0e', '\uff61':
			inLabel = false
		default:
			if !inLabel {
				n++
				inLabel = true
			}
		}
	}

	return n
}

func needsIDNA(host string) bool {
	return strings.ContainsFunc(host, func(r rune) bool { return r >= utf8.RuneSelf }) ||
		strings.Contains(strings.ToLower(host), "xn--")
}
