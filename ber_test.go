package ber

import (
	"bytes"
	"io"
	"testing"
)

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

func TestSequenceAndAppendChild(t *testing.T) {
	// Pre
	var v1, v2, v3 string
	v1 = "HIC SVNT LEONES"
	v2 = "HIC SVNT DRACONES"
	v3 = "Terra Incognita"

	// Test
	p1 := NewString(ClassUniversal, TypePrimitive, TagOctetString, v1, "String")
	p2 := NewString(ClassUniversal, TypePrimitive, TagOctetString, v2, "String")
	p3 := NewString(ClassUniversal, TypePrimitive, TagOctetString, v3, "String")

	sequence := NewSequence("a sequence")
	sequence.AppendChild(p1)
	sequence.AppendChild(p2)
	sequence.AppendChild(p3)

	if len(sequence.Children) != 3 {
		t.Error("wrong length for children array should be three =>", len(sequence.Children))
	}

	encodedSequence := sequence.Bytes()

	decodedSequence := DecodePacket(encodedSequence)
	if len(decodedSequence.Children) != 3 {
		t.Error("wrong length for children array should be three =>", len(decodedSequence.Children))
	}

	// Post
}

func TestPrint(t *testing.T) {
	// Pre
	var v1 string = "Answer to the Ultimate Question of Life, the Universe, and Everything"
	var v2 uint64 = 42
	var v3 bool = true
	// Test
	p1 := NewString(ClassUniversal, TypePrimitive, TagOctetString, v1, "Question")
	p2 := NewInteger(ClassUniversal, TypePrimitive, TagInteger, v2, "Answer")
	p3 := NewBoolean(ClassUniversal, TypePrimitive, TagBoolean, v3, "Validity")

	sequence := NewSequence("a sequence")
	sequence.AppendChild(p1)
	sequence.AppendChild(p2)
	sequence.AppendChild(p3)

	PrintPacket(sequence)

	encodedSequence := sequence.Bytes()
	PrintBytes(encodedSequence, "\t")

	// Post
}

func TestReadPacket(t *testing.T) {
	// Pre
	var value string = "Ad impossibilia nemo tenetur"
	packet := NewString(ClassUniversal, TypePrimitive, TagOctetString, value, "string")
	var buffer io.ReadWriter
	buffer = new(bytes.Buffer)

	// Test
	buffer.Write(packet.Bytes())

	newPacket, err := ReadPacket(buffer)
	if err != nil {
		t.Error("error during ReadPacket", err)
	}

	if !bytes.Equal(newPacket.ByteValue, packet.ByteValue) {
		t.Error("packets should be the same")
	}
}
