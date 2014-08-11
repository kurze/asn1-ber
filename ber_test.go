package ber

import "testing"

func TestEncodeDecodeInterger(t *testing.T) {
	// Pre
	var integer uint64 = 10

	// Test
	encodedInteger := EncodeInteger(integer)
	decodedInteger := DecodeInteger(encodedInteger)

	if integer != decodedInteger {
		t.Error("wrong should be equal", integer, decodedInteger)
	}

	// Post
}

func TestBoolean(t *testing.T) {
	// Pre
	var value bool = true

	// Test
	packet := NewBoolean(ClassUniversal, TypePrimitive, TagBoolean, value, "first Packet, True")

	newBoolean, ok := packet.Value.(bool)
	if !ok || newBoolean != value {
		t.Error("error during creating packet")
	}

	encodedPacket := packet.Bytes()

	newPacket := DecodePacket(encodedPacket)

	newBoolean, ok = newPacket.Value.(bool)
	if !ok || newBoolean != value {
		t.Error("error during decoding packet")
	}

	// Post
}

func TestInteger(t *testing.T) {
	// Pre
	var value uint64 = 10

	// Test
	packet := NewInteger(ClassUniversal, TypePrimitive, TagInteger, value, "Integer, 10")

	newInteger, ok := packet.Value.(uint64)
	if !ok || newInteger != value {
		t.Error("error during creating packet")
	}

	encodedPacket := packet.Bytes()

	newPacket := DecodePacket(encodedPacket)

	newInteger, ok = newPacket.Value.(uint64)
	if !ok || newInteger != value {
		t.Error("error during decoding packet")
	}

	// Post
}

func TestString(t *testing.T) {
	// Pre
	var value string = "Hic sunt dracones"

	// Test
	packet := NewString(ClassUniversal, TypePrimitive, TagOctetString, value, "String")

	newValue, ok := packet.Value.(string)
	if !ok || newValue != value {
		t.Error("error during creating packet")
	}

	encodedPacket := packet.Bytes()

	newPacket := DecodePacket(encodedPacket)

	newValue, ok = newPacket.Value.(string)
	if !ok || newValue != value {
		t.Error("error during decoding packet")
	}

	// Post
}
