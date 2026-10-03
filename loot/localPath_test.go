package loot

import (
	"path/filepath"
	"testing"
)

func TestResolveLocalFile_FileName(t *testing.T) {
	cases := []struct {
		name     string
		url      string
		wantHost string
		wantFile string
	}{
		{"plain file", "https://example.com/js/app.js", "example.com", "app.js"},
		{"root path", "http://154.57.164.64:30399/", "154.57.164.64:30399", "index.html"},
		{"no path at all", "http://example.com", "example.com", "index.html"},
		{"trailing slash directory", "https://example.com/js/", "example.com", "index.html"},
		{"query string is ignored", "https://example.com/main.js?v=2", "example.com", "main.js"},
		{"root path with query", "https://example.com/?id=1", "example.com", "index.html"},
		{"scheme gets added", "example.com/app.js", "example.com", "app.js"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			dir := t.TempDir()
			parsed, absFileName, err := resolveLocalFile(c.url, dir)
			if err != nil {
				t.Fatalf("resolveLocalFile(%q) returned error: %v", c.url, err)
			}
			if parsed.Host != c.wantHost {
				t.Errorf("host = %q, want %q", parsed.Host, c.wantHost)
			}
			if got := filepath.Base(absFileName); got != c.wantFile {
				t.Errorf("file name = %q, want %q", got, c.wantFile)
			}
			if got := filepath.Base(filepath.Dir(absFileName)); got != c.wantHost {
				t.Errorf("parent directory = %q, want %q", got, c.wantHost)
			}
		})
	}
}

// Regression test: a URL whose path is just "/" used to resolve to the host directory
// itself, so storing it failed with "is a directory" and nothing was looted.
func TestStoreJSFile_RootURL(t *testing.T) {
	dir := t.TempDir()
	htmlContent := "<html><script>var a=1;</script></html>"
	f := JSFile{
		URL:         "http://154.57.164.64:30399/",
		Content:     htmlContent,
		ContentType: "text/html; charset=UTF-8",
	}

	path, err := StoreJSFile(f, dir, false)
	if err != nil {
		t.Fatalf("StoreJSFile returned error: %v", err)
	}
	if got := filepath.Base(path); got != "index.html" {
		t.Errorf("file name = %q, want %q", got, "index.html")
	}
}
