package utils

import log "github.com/sirupsen/logrus"

const TRUNCATE_MAX = 40

// Truncate truncates a string to n characters. If pad is not nil, then the
// string is padded with pad until it is n characters long.
func Truncate(s string, n int, pad *string) string {
	if n > TRUNCATE_MAX {
		log.Warnf("Truncate: n is greater than TRUNCATE_MAX (%d), using TRUNCATE_MAX", TRUNCATE_MAX)
		n = TRUNCATE_MAX
	}

	if len(s) > n {
		return s[:n]
	}

	if pad != nil {
		if len(*pad) > 1 {
			log.Warn("Truncate: pad is longer than 1 character, ignoring")
			return s
		}

		for len(s) < n {
			s += *pad
		}
	}

	return s
}

// StringPtr returns a pointer to a string.
func StringPtr(s string) *string {
	return &s
}
