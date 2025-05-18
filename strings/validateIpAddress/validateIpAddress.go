// #468
// https://leetcode.com/problems/validate-ip-address/description/
package leetcode

import (
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	// Validate IPv4
	if validIPv4(queryIP) {
		return "IPv4"
	}
	if validIPv6(queryIP) {
		return "IPv6"
	}
	return "Neither"
}

func validIPv4(queryIP string) bool {
	chunks := strings.Split(queryIP, ".")
	if len(chunks) != 4 {
		return false
	}
	for _, c := range chunks {
		if len(c) > 1 && c[0] == '0' {
			return false
		}
		n, err := strconv.Atoi(c)
		if err != nil {
			return false
		}
		if n < 0 || n > 255 {
			return false
		}
	}
	return true
}

func validIPv6(queryIP string) bool {
	chunks := strings.Split(queryIP, ":")
	if len(chunks) != 8 {
		return false
	}
	for _, c := range chunks {
		if len(c) > 4 {
			return false
		}
		if _, err := strconv.ParseInt(c, 16, 32); err != nil {
			return false
		}
	}
	return true
}
