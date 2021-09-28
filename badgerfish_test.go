package badgerfish

import (
	"testing"
)

func TestNothing(t *testing.T) {
	doc := &Document{}
	if doc == nil {
		t.Error("Expected nil document")
	}
}

func TestGetProlog(t *testing.T) {
	testCases := []struct {
		in     string
		expect string
	}{
		{
			"",
			"",
		},
		{
			`<?xml version="1.0" encoding="UTF-8"?>`,
			`<?xml version="1.0" encoding="UTF-8"?>`,
		},
		{
			`<?xml version="1.0" encoding="UTF-8"?> <hello>world</hello> `,
			`<?xml version="1.0" encoding="UTF-8"?>`,
		},
		{
			`<?xml version="1.0" encoding="UTF-8"?>
			<hello>world</hello>`,
			`<?xml version="1.0" encoding="UTF-8"?>`,
		},
	}

	for _, tc := range testCases {
		res := getProlog(tc.in)
		if res != tc.expect {
			t.Errorf("Expected %q, got %q", tc.expect, res)
		}
	}
}
