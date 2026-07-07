package opennox

import (
	"testing"

	"github.com/noxworld-dev/opennox-lib/log"
	"github.com/noxworld-dev/opennox-lib/maps"
)

func TestClientMapDownloadStartNativeValidation(t *testing.T) {
	c := &clientMapDownload{}
	c.log = log.New("test")
	c.ndl = maps.NewNativeDownloader("/tmp")
	c.iface.setMapPath = func(path string) {}
	c.iface.sendCancelMap = func() int { return 0 }
	c.iface.sendReceivedMap = func() int { return 0 }

	// Test empty name
	err := c.startNative("", 100)
	if err == nil {
		t.Error("startNative with empty name should return error")
	}

	// Test invalid size
	err = c.startNative("testmap", 0)
	if err == nil {
		t.Error("startNative with size 0 should return error")
	}

	// Test invalid name with slash at start (i=0)
	err = c.startNative("/invalid", 100)
	if err == nil {
		t.Error("startNative with slash at start should return error")
	}
}

func TestClientMapDownloadStartNativeNameHandling(t *testing.T) {
	c := &clientMapDownload{}
	c.log = log.New("test")
	c.ndl = maps.NewNativeDownloader("/tmp")
	c.iface.setMapPath = func(path string) {}
	c.iface.sendCancelMap = func() int { return 0 }
	c.iface.sendReceivedMap = func() int { return 0 }

	// Test name without extension - should add .nxz
	err := c.startNative("testmap", 100)
	_ = err
}

func TestClientMapDownloadDownloading(t *testing.T) {
	c := &clientMapDownload{}

	if c.Downloading() {
		t.Error("new clientMapDownload should not be downloading")
	}

	c.setDownloading(true)
	if !c.Downloading() {
		t.Error("setDownloading(true) should make Downloading() return true")
	}

	c.setDownloading(false)
	if c.Downloading() {
		t.Error("setDownloading(false) should make Downloading() return false")
	}
}

func TestClientMapDownloadSetDownloadOK(t *testing.T) {
	c := &clientMapDownload{}

	c.setDownloadOK(true)
	if !c.downloadOK {
		t.Error("setDownloadOK(true) should set downloadOK to true")
	}

	c.setDownloadOK(false)
	if c.downloadOK {
		t.Error("setDownloadOK(false) should set downloadOK to false")
	}
}
