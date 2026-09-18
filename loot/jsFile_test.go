package loot

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLooksLikeJS(t *testing.T) {
	cases := []struct {
		contentType string
		want        bool
	}{
		{"text/javascript", true},
		{"application/javascript; charset=utf-8", true},
		{"application/ecmascript", true},
		{"TEXT/JAVASCRIPT", true},
		{"text/html", false},
		{"application/json", false},
		{"", false},
	}
	for _, c := range cases {
		if got := looksLikeJS(c.contentType); got != c.want {
			t.Errorf("looksLikeJS(%q) = %v, want %v", c.contentType, got, c.want)
		}
	}
}

func TestStoreJSFile_SkipsBeautifyForHTML(t *testing.T) {
	dir := t.TempDir()
	htmlContent := "<html><script>var x=1;function f(){return x}</script></html>"
	f := JSFile{
		URL:         "https://example.com/page",
		Content:     htmlContent,
		ContentType: "text/html",
	}

	path, err := StoreJSFile(f, dir, true)
	if err != nil {
		t.Fatalf("StoreJSFile returned error: %v", err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read stored file: %v", err)
	}
	if string(got) != htmlContent {
		t.Errorf("HTML content was modified despite beautify=true; got %q, want %q", got, htmlContent)
	}
}

func TestStoreJSFile_BeautifiesUnknownOrJSContentType(t *testing.T) {
	// Empty ContentType preserves legacy behaviour: beautify is attempted (only skipped for
	// non-JS content types). This only checks storage succeeds, not the beautifier's actual
	// formatting output, to stay independent of whether js-beautify is installed.
	dir := t.TempDir()
	f := JSFile{
		URL:     "https://example.com/app.js",
		Content: "var x=1;",
	}

	path, err := StoreJSFile(f, dir, true)
	if err != nil {
		t.Fatalf("StoreJSFile returned error: %v", err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("expected file to exist at %s: %v", path, err)
	}
	if filepath.Base(path) != "app.js" {
		t.Errorf("unexpected file name: %s", filepath.Base(path))
	}
}
