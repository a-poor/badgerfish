package badgerfish

import (
	"errors"
	"strings"
)

var (
	EscapeCharMap = map[string]string{
		"<":  "&lt;",
		">":  "&gt;",
		"&":  "&amp;",
		"'":  "&apos;",
		"\"": "&quot;",
	}
	UnEscapeCharMap = map[string]string{
		"&lt;":   "<",
		"&gt;":   ">",
		"&amp;":  "&",
		"&apos;": "'",
		"&quot;": "\"",
	}
)

var (
	ErrMalformedXML = errors.New("malformed XML")
)

type Document struct {
	Prolog string
	Root   *Element
}

func getProlog(data string) string {
	if len(data) < 2 {
		return ""
	}

	if data[0:2] != "<?" {
		return ""
	}

	i := strings.Index(data, "?>") + 2

	return data[0:i]
}

func getRootElement(data string) string {
	return ""
}

func ParseDocument(data string) (*Document, error) {
	clean := strings.TrimSpace(data)
	prolog := getProlog(clean)

	clean = strings.TrimSpace(clean[len(prolog):])

	doc := Document{Prolog: prolog}

	return &doc, nil
}

// type Prolog map[string]string

type Element struct {
	Name           string
	NameSpace      string
	Attributes     map[string]string
	StringContent  []string
	NestedElements []Element
	IsEmpty        bool
}
