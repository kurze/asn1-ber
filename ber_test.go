package ber

import (
	"bytes"
	"io"
	"testing"
)

func TestEncodeDecodeInteger(t *testing.T) {
	encodeDecodeInteger(t, 0)
	encodeDecodeInteger(t, 10)
	encodeDecodeInteger(t, 128)
	encodeDecodeInteger(t, 1024)
	encodeDecodeInteger(t, -1)
	encodeDecodeInteger(t, -100)
	encodeDecodeInteger(t, -1024)
}

func encodeDecodeInteger(t *testing.T, value int64) {
	encodedInteger := EncodeSignedInteger(value)
	decodedInteger := int64(DecodeInteger(encodedInteger))

	if value != int64(decodedInteger) {
		t.Error("wrong should be equal", value, decodedInteger)
	}
}

func TestBoolean(t *testing.T) {
	var value bool = true

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

}

func TestInteger(t *testing.T) {
	var value uint64 = 10

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

}

func TestString(t *testing.T) {
	var value string = "Hic sunt dracones"

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

}

func TestSequenceAndAppendChild(t *testing.T) {

	p1 := NewString(ClassUniversal, TypePrimitive, TagOctetString, "HIC SVNT LEONES", "String")
	p2 := NewString(ClassUniversal, TypePrimitive, TagOctetString, "HIC SVNT DRACONES", "String")
	p3 := NewString(ClassUniversal, TypePrimitive, TagOctetString, "Terra Incognita", "String")

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

}

func TestPrint(t *testing.T) {
	p1 := NewString(ClassUniversal, TypePrimitive, TagOctetString, "Answer to the Ultimate Question of Life, the Universe, and Everything", "Question")
	p2 := NewInteger(ClassUniversal, TypePrimitive, TagInteger, 42, "Answer")
	p3 := NewBoolean(ClassUniversal, TypePrimitive, TagBoolean, true, "Validity")

	sequence := NewSequence("a sequence")
	sequence.AppendChild(p1)
	sequence.AppendChild(p2)
	sequence.AppendChild(p3)

	PrintPacket(sequence)

	encodedSequence := sequence.Bytes()
	PrintBytes(encodedSequence, "\t")
}

func TestReadPacket(t *testing.T) {
	packet := NewString(ClassUniversal, TypePrimitive, TagOctetString, "Ad impossibilia nemo tenetur", "string")
	var buffer io.ReadWriter
	buffer = new(bytes.Buffer)

	buffer.Write(packet.Bytes())

	newPacket, err := ReadPacket(buffer)
	if err != nil {
		t.Error("error during ReadPacket", err)
	}

	if !bytes.Equal(newPacket.ByteValue, packet.ByteValue) {
		t.Error("packets should be the same")
	}
}

func TestBinaryInteger(t *testing.T) {
	// data src : http://luca.ntop.org/Teaching/Appunti/asn1.html 5.7

	if !bytes.Equal([]byte{0x02, 0x01, 0x00}, NewSignedInteger(ClassUniversal, TypePrimitive, TagInteger, 0, "").Bytes()) {
		t.Error("wrong binary generated")
	}
	if !bytes.Equal([]byte{0x02, 0x01, 0x7F}, NewSignedInteger(ClassUniversal, TypePrimitive, TagInteger, 127, "").Bytes()) {
		t.Error("wrong binary generated")
	}
	if !bytes.Equal([]byte{0x02, 0x02, 0x00, 0x80}, NewSignedInteger(ClassUniversal, TypePrimitive, TagInteger, 128, "").Bytes()) {
		t.Error("wrong binary generated")
	}
	if !bytes.Equal([]byte{0x02, 0x02, 0x01, 0x00}, NewSignedInteger(ClassUniversal, TypePrimitive, TagInteger, 256, "").Bytes()) {
		t.Error("wrong binary generated")
	}
	if !bytes.Equal([]byte{0x02, 0x01, 0x80}, NewSignedInteger(ClassUniversal, TypePrimitive, TagInteger, -128, "").Bytes()) {
		t.Error("wrong binary generated")
	}
	if !bytes.Equal([]byte{0x02, 0x01, 0xFF, 0x7F}, NewSignedInteger(ClassUniversal, TypePrimitive, TagInteger, -129, "").Bytes()) {
		t.Error("wrong binary generated")
	}

}

