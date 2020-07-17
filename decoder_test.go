package smpp

import (
	"bytes"
	"testing"
)

func TestDecoder_Decode(t *testing.T) {
	req := &BindReceiverPdu{
		Header: &Header{
			CommandID:      BindReceiver,
			CommandStatus:  EsmeRok,
			SequenceNumber: 1,
		},
		Body: &BindBody{
			SystemID:         "test",
			Password:         "pass",
			SystemType:       "",
			InterfaceVersion: ProtocolId,
			AddrTon:          0,
			AddrNpi:          1,
			AddressRange:     "",
		},
	}
	buffer := new(bytes.Buffer)
	if err := NewEncoder(buffer).Encode(req); err != nil {
		t.Fatal(err)
	}
	rep, err := NewDecoder(buffer).Decode()
	if err != nil {
		t.Fatal(err)
	}
	if pdu, ok := rep.(*BindReceiverPdu); ok {
		t.Log(pdu.Header)
		t.Log(pdu.Body)
	} else {
		t.Fail()
	}
	req1 := &BindReceiverRespPdu{
		Header: &Header{
			CommandID:      BindReceiverResp,
			CommandStatus:  EsmeRok,
			SequenceNumber: 2,
		},
		Body: &BindRespBody{
			SystemID: "test",
		},
		Tlv: TlvMap{},
	}
	buffer1 := new(bytes.Buffer)
	if err := NewEncoder(buffer1).Encode(req1); err != nil {
		t.Fatal(err)
	}
	rep1, err := NewDecoder(buffer1).Decode()
	if err != nil {
		t.Fatal(err)
	}
	if pdu1, ok := rep1.(*BindReceiverRespPdu); ok {
		t.Log(pdu1.Header)
		t.Log(pdu1.Body)
	} else {
		t.Fail()
	}
}
