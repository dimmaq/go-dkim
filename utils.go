package dkim

import (
	"bytes"
	"strings"
)

func reduceWitespaces(s string) string {
	l := len(s)
	var buf *bytes.Buffer
	p := 0
	for s != `` {
		p = strings.IndexAny(s, " \t\r\n")
		if p >= 0 {
			if buf == nil {
				buf = bytes.NewBuffer(make([]byte, 0, l))
			}
			buf.WriteString(s[:p])
			buf.WriteByte(32)
			var ch byte
			for p < l {
				ch = s[p]
				if ch == 32 || ch == 9 || ch == 13 || ch == 10 {
					p++
				} else {
					break
				}
			}
			s = s[p:]
			l = len(s)
		} else {
			if buf == nil {
				return s
			}
			buf.WriteString(s)
			break
		}
	}
	return buf.String()
}
