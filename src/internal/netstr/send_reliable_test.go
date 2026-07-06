package netstr

import (
	"testing"
	"time"

	"github.com/noxworld-dev/opennox-lib/noxnet"
)

func TestNewReliableSender(t *testing.T) {
	sent := make([][]byte, 0)
	sendFn := func(data []byte) (int, error) {
		sent = append(sent, append([]byte(nil), data...))
		return len(data), nil
	}

	rs := newReliableSender(sendFn)
	if rs == nil {
		t.Fatal("newReliableSender should not return nil")
	}
	if rs.write == nil {
		t.Error("newReliableSender write func should not be nil")
	}
	if rs.Now == nil {
		t.Error("newReliableSender Now func should not be nil")
	}
}

func TestReliableSender_Send(t *testing.T) {
	sent := make([][]byte, 0)
	sendFn := func(data []byte) (int, error) {
		sent = append(sent, append([]byte(nil), data...))
		return len(data), nil
	}

	rs := newReliableSender(sendFn)

	// Test Send with empty buffer
	seq, err := rs.Send(1, []byte{})
	if err == nil {
		t.Error("Send with empty buffer should return error")
	}
	if seq != -1 {
		t.Errorf("Send with empty buffer seq = %d, want -1", seq)
	}

	// Test Send with valid buffer
	sent = sent[:0]
	seq, err = rs.Send(5, []byte{0x01, 0x02, 0x03})
	if err != nil {
		t.Errorf("Send with valid buffer should not return error: %v", err)
	}
	if seq != 0 {
		t.Errorf("First Send seq = %d, want 0", seq)
	}
	if len(sent) != 1 {
		t.Errorf("Send should call write once, called %d times", len(sent))
	}
	if len(sent[0]) != 5 { // 2 byte header + 3 byte payload
		t.Errorf("Sent data len = %d, want 5", len(sent[0]))
	}
	if sent[0][0] != (5 | reliableFlag) {
		t.Errorf("Sent data[0] = %d, want %d", sent[0][0], 5|reliableFlag)
	}
	if sent[0][1] != 0 {
		t.Errorf("Sent data[1] (seq) = %d, want 0", sent[0][1])
	}

	// Test second Send increments seq
	sent = sent[:0]
	seq, err = rs.Send(6, []byte{0x04})
	if err != nil {
		t.Errorf("Second Send should not return error: %v", err)
	}
	if seq != 1 {
		t.Errorf("Second Send seq = %d, want 1", seq)
	}
}

func TestReliableSender_InQueue(t *testing.T) {
	sent := make([][]byte, 0)
	sendFn := func(data []byte) (int, error) {
		sent = append(sent, append([]byte(nil), data...))
		return len(data), nil
	}

	rs := newReliableSender(sendFn)

	// Initially empty
	if cnt := rs.InQueue(); cnt != 0 {
		t.Errorf("InQueue on empty = %d, want 0", cnt)
	}

	// Send some packets with op byte at position 2
	// Packet format: [id|reliableFlag, seq, op, ...]
	rs.Send(1, []byte{0x01, 0x01}) // op 0x01
	rs.Send(2, []byte{0x02, 0x02}) // op 0x02
	rs.Send(3, []byte{0x01, 0x03}) // op 0x01

	// Count all
	if cnt := rs.InQueue(); cnt != 3 {
		t.Errorf("InQueue all = %d, want 3", cnt)
	}

	// Count specific op
	if cnt := rs.InQueue(noxnet.Op(0x01)); cnt != 2 {
		t.Errorf("InQueue op 0x01 = %d, want 2", cnt)
	}
	if cnt := rs.InQueue(noxnet.Op(0x02)); cnt != 1 {
		t.Errorf("InQueue op 0x02 = %d, want 1", cnt)
	}
	if cnt := rs.InQueue(noxnet.Op(0xFF)); cnt != 0 {
		t.Errorf("InQueue non-existent op = %d, want 0", cnt)
	}
}

func TestReliableSender_Clear(t *testing.T) {
	sendFn := func(data []byte) (int, error) { return len(data), nil }
	rs := newReliableSender(sendFn)

	rs.Send(1, []byte{0x01})
	rs.Send(2, []byte{0x02})

	if cnt := rs.InQueue(); cnt != 2 {
		t.Errorf("InQueue before Clear = %d, want 2", cnt)
	}

	rs.Clear()

	if cnt := rs.InQueue(); cnt != 0 {
		t.Errorf("InQueue after Clear = %d, want 0", cnt)
	}
}

func TestReliableSender_SetCurrent(t *testing.T) {
	sent := make([][]byte, 0)
	sendFn := func(data []byte) (int, error) {
		sent = append(sent, append([]byte(nil), data...))
		return len(data), nil
	}

	rs := newReliableSender(sendFn)

	// Send 3 packets
	rs.Send(1, []byte{0x01})
	rs.Send(2, []byte{0x02})
	rs.Send(3, []byte{0x03})

	sent = sent[:0]

	// SetCurrent should mark seq for retry and ack older
	rs.SetCurrent(1)

	// Schedule and resend should send the marked packet
	rs.ScheduleAndResend()

	// Should have sent packet with seq 1
	found := false
	for _, data := range sent {
		if len(data) >= 2 && data[1] == 1 {
			found = true
			break
		}
	}
	if !found {
		t.Error("SetCurrent(1) should mark seq 1 for retry")
	}
}

func TestReliableSender_ScheduleAndResend(t *testing.T) {
	sent := make([][]byte, 0)
	sendFn := func(data []byte) (int, error) {
		sent = append(sent, append([]byte(nil), data...))
		return len(data), nil
	}

	rs := newReliableSender(sendFn)
	// Override Now to control time
	now := time.Duration(0)
	rs.Now = func() time.Duration { return now }

	rs.Send(1, []byte{0x01})
	sent = sent[:0] // Clear initial send

	// Advance time past retry interval
	now = reliableRetry + time.Second

	rs.ScheduleAndResend()

	if len(sent) == 0 {
		t.Error("ScheduleAndResend should resend after retry interval")
	}
}

func TestReliableSender_Resend(t *testing.T) {
	sent := make([][]byte, 0)
	sendFn := func(data []byte) (int, error) {
		sent = append(sent, append([]byte(nil), data...))
		return len(data), nil
	}

	rs := newReliableSender(sendFn)
	rs.Send(1, []byte{0x01})
	sent = sent[:0]

	// Manually mark shouldSend
	for it := rs.queue; it != nil; it = it.next {
		it.shouldSend = true
	}

	rs.Resend()

	if len(sent) == 0 {
		t.Error("Resend should send packets with shouldSend=true")
	}
}

func TestConn_SendReliable(t *testing.T) {
	// Test with nil conn
	var ns *Conn
	seq, err := ns.SendReliable([]byte{0x01})
	if err == nil {
		t.Error("SendReliable on nil conn should return error")
	}
	if seq != -3 {
		t.Errorf("SendReliable on nil conn seq = %d, want -3", seq)
	}

	// Test with empty buffer
	ns = &Conn{}
	ns.reliable = newReliableSender(func(data []byte) (int, error) { return len(data), nil })
	seq, err = ns.SendReliable([]byte{})
	if err == nil {
		t.Error("SendReliable with empty buffer should return error")
	}
	if seq != -2 {
		t.Errorf("SendReliable with empty buffer seq = %d, want -2", seq)
	}
}

func TestConn_SendReliableMsg(t *testing.T) {
	// Test with nil conn - SendReliableMsg calls SendReliable which checks for nil
	var ns *Conn
	_, err := ns.SendReliableMsg(&noxnet.MsgConnect{})
	if err == nil {
		t.Error("SendReliableMsg on nil conn should return error")
	}
}

func TestConn_ReliableInQueue(t *testing.T) {
	ns := &Conn{}
	ns.reliable = newReliableSender(func(data []byte) (int, error) { return len(data), nil })

	// Use add directly instead of SendReliable which requires data2hdr
	ns.reliable.add(1, []byte{0x01, 0x01})

	cnt := ns.ReliableInQueue(noxnet.Op(0x01))
	if cnt != 1 {
		t.Errorf("ReliableInQueue = %d, want 1", cnt)
	}
}

func TestStreams_MaybeSendReliable(t *testing.T) {
	var g Streams
	// streams is [128]*Conn, so we need to use correct type
	g.streams = [128]*Conn{}
	g.Now = func() time.Duration { return time.Duration(0) }

	// Should not panic with empty streams
	g.MaybeSendReliable()

	// Add a conn with reliable sender (but don't call SendReliable which needs data2hdr)
	ns := &Conn{}
	ns.reliable = newReliableSender(func(data []byte) (int, error) { return len(data), nil })
	// Manually add a packet to the queue without using SendReliable
	ns.reliable.add(1, []byte{0x01})
	g.streams[0] = ns

	// First call should set lastSendReliable
	g.MaybeSendReliable()

	// Second call within 1 second should return early
	g.MaybeSendReliable()
}
