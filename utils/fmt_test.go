package utils

import "testing"

func TestWireshark(t *testing.T) {
	s := Wireshark(0, []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15,
		0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23,
		0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30, 0x31,
	})
	expected := "00 01 02 03 04 05 06 07  08 09 10 11 12 13 14 15\n16 17 18 19 20 21 22 23  24 25 26 27 28 29 30 31\n"
	if s != expected {
		t.Errorf("wireshark result doesn't match.\nWant:\n%s\nGot:\n%s", expected, s)
	}
}

func TestWiresharkOffset(t *testing.T) {
	s := Wireshark(5, []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15,
		0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23,
		0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30, 0x31,
	})
	expected := "00 00 00 00 00 00 01 02  03 04 05 06 07 08 09 10\n11 12 13 14 15 16 17 18  19 20 21 22 23 24 25 26\n27 28 29 30 31\n"
	if s != expected {
		t.Errorf("wireshark result with offset doesn't match.\nWant:\n%s\nGot:\n%s", expected, s)
	}
}
