package opennox

import (
	"context"
	"image"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestScreenshotHandlerCancelledContext(t *testing.T) {
	// Create a client with screenshots channel
	cli := &Client{}
	cli.screenshots.req = make(chan chan<- image.Image)
	noxClient = cli
	defer func() { noxClient = nil }()

	// Create a request with cancelled context
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	req := httptest.NewRequest("GET", "/nox/screenshot", nil).WithContext(ctx)
	w := httptest.NewRecorder()

	screenshotHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
}

func TestScreenshotHandlerNilImage(t *testing.T) {
	cli := &Client{}
	cli.screenshots.req = make(chan chan<- image.Image)
	noxClient = cli
	defer func() { noxClient = nil }()

	// Start a goroutine to handle screenshot request and return nil image
	go func() {
		out := <-cli.screenshots.req
		// Send nil image and close
		close(out)
	}()

	req := httptest.NewRequest("GET", "/nox/screenshot", nil)
	w := httptest.NewRecorder()

	screenshotHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusNotImplemented {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusNotImplemented)
	}
}

func TestScreenshotHandlerPNG(t *testing.T) {
	cli := &Client{}
	cli.screenshots.req = make(chan chan<- image.Image)
	noxClient = cli
	defer func() { noxClient = nil }()

	// Create a simple 1x1 image
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))

	go func() {
		out := <-cli.screenshots.req
		out <- img
		close(out)
	}()

	req := httptest.NewRequest("GET", "/nox/screenshot?f=png", nil)
	w := httptest.NewRecorder()

	screenshotHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	// PNG magic bytes: 89 50 4E 47
	body := w.Body.Bytes()
	if len(body) < 4 || body[0] != 0x89 || body[1] != 0x50 {
		t.Error("Response body does not look like a PNG")
	}
}

func TestScreenshotHandlerJPEG(t *testing.T) {
	cli := &Client{}
	cli.screenshots.req = make(chan chan<- image.Image)
	noxClient = cli
	defer func() { noxClient = nil }()

	img := image.NewRGBA(image.Rect(0, 0, 1, 1))

	go func() {
		out := <-cli.screenshots.req
		out <- img
		close(out)
	}()

	req := httptest.NewRequest("GET", "/nox/screenshot?f=jpg&q=90", nil)
	w := httptest.NewRecorder()

	screenshotHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	// JPEG magic bytes: FF D8
	body := w.Body.Bytes()
	if len(body) < 2 || body[0] != 0xFF || body[1] != 0xD8 {
		t.Error("Response body does not look like a JPEG")
	}
}

func TestScreenshotHandlerDefaultFormat(t *testing.T) {
	cli := &Client{}
	cli.screenshots.req = make(chan chan<- image.Image)
	noxClient = cli
	defer func() { noxClient = nil }()

	img := image.NewRGBA(image.Rect(0, 0, 1, 1))

	go func() {
		out := <-cli.screenshots.req
		out <- img
		close(out)
	}()

	// No format specified, should default to JPEG
	req := httptest.NewRequest("GET", "/nox/screenshot", nil)
	w := httptest.NewRecorder()

	screenshotHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
}

func TestClientScreenshotCancelled(t *testing.T) {
	cli := &Client{}
	cli.screenshots.req = make(chan chan<- image.Image)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	img, err := cli.Screenshot(ctx)
	if err == nil {
		t.Error("expected error for cancelled context")
	}
	if img != nil {
		t.Error("expected nil image for cancelled context")
	}
}

func TestClientMaybeScreenshot(t *testing.T) {
	cli := &Client{}
	cli.screenshots.req = make(chan chan<- image.Image)

	// maybeScreenshot should return quickly if no request
	done := make(chan bool)
	go func() {
		cli.maybeScreenshot()
		done <- true
	}()

	select {
	case <-done:
		// OK, returned quickly
	case <-time.After(time.Second):
		t.Error("maybeScreenshot did not return quickly")
	}
}
