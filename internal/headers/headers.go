package headers

import (
	"bytes"
	"fmt"
	"strings"
)

type Headers map[string]string

const crlf = "\r\n"

func NewHeaders() Headers {
	return map[string]string{}
}

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
	idx := bytes.Index(data, []byte(crlf))
	if idx == -1 {
		return 0, false, nil
	}
	if idx == 0 {
		// the empty line
		// headers are done, consume the CRLF
		return 2, true, nil
	}

	headerText := string(data[:idx])
	headerParts := strings.SplitN(headerText, ":", 2)
	headerName := headerParts[0]
	if headerName != strings.TrimRight(headerName, " ") {
		return 0, false, fmt.Errorf("invalid header name: %s", headerName)
	}

	headerValue := strings.TrimSpace(headerParts[1])
	headerName = strings.TrimSpace(headerName)

	h.Set(headerName, headerValue)
	return idx + len(crlf), false, nil
}

func (h Headers) Set(key, value string) {
	h[key] = value
}
