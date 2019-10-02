package jpegquality

import (
	"encoding/base64"
	"log"
	"os"
	"testing"
)

func init() {
	SetLogger(log.New(os.Stderr, "test-", log.Ltime|log.Lshortfile))
}

func TestJpegQuality(t *testing.T) {
	buf, err := base64.StdEncoding.DecodeString(jpegData)
	if err != nil {
		t.Fatal(err)
	}

	j, err := NewWithBytes(buf)
	if err != nil {
		t.Fatal(err)
	}
	q := j.Quality()
	t.Logf("jpeg quality %d", q)
	if q != 80 {
		t.Errorf("mismatch quality %d", q)
	}
}

const (
	jpegData = `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`
)
