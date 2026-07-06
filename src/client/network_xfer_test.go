package client

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/common/ntype"
	"github.com/noxworld-dev/opennox/v1/server"
)

func TestClient_handleCBORXfer(t *testing.T) {
	c := &Client{}
	c.Objs.Ext.init(&c.Objs)

	// Test with nil data (no object)
	m1 := &server.XferObjectSetLabel{
		Object: 12345,
		Label:  "Test",
	}
	c.handleCBORXfer(ntype.PlayerInd(0), m1)
	// Should not panic

	// Test with existing data via SetByNetCode
	data := c.Objs.Ext.SetByNetCode(12345)
	if data == nil {
		t.Fatal("expected data not nil")
	}

	m2 := &server.XferObjectSetLabel{
		Object: 12345,
		Label:  "TestLabel",
		Color:  &server.XferColor{R: 255, G: 0, B: 0, A: 255},
	}
	c.handleCBORXfer(ntype.PlayerInd(0), m2)

	if data.DisplayName != "TestLabel" {
		t.Fatalf("expected label TestLabel, got %q", data.DisplayName)
	}

	// Test with unknown message type (nil)
	c.handleCBORXfer(ntype.PlayerInd(0), nil)
}
