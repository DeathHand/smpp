package smpp

import (
	"bytes"
	"testing"
)

func TestEncoder_Encode(t *testing.T) {
	pdu := &EnquireLinkPdu{
		Header: &Header{
			CommandID:      EnquireLink,
			CommandStatus:  EsmeRbindFail,
			SequenceNumber: 2,
		},
	}
	buffer := &bytes.Buffer{}
	if err := NewEncoder(buffer).Encode(pdu); err != nil {
		t.Fatal(err)
	}
	t.Log(buffer.Bytes())

	result, err := NewDecoder(buffer).Decode()
	if err != nil {
		t.Fatal(err)
	}
	if p, ok := result.(*EnquireLinkPdu); ok {
		t.Log(p.Header.CommandLength)
		t.Log(p.Header.CommandID)
		t.Log(p.Header.CommandStatus)
		t.Log(p.Header.SequenceNumber)
		return
	}
	t.Fatal()
}

func BenchmarkEncoder_Encode(b *testing.B) {
	b.ReportAllocs()
	pdu := &SubmitSmPdu{
		Header: &Header{
			CommandID:      EnquireLink,
			CommandStatus:  EsmeRbindFail,
			SequenceNumber: 2,
		},
		Body: &SmBody{
			ServiceType:          "",
			SourceAddrTon:        0,
			SourceAddrNpi:        0,
			SourceAddr:           "",
			DestAddrTon:          0,
			DestAddrNpi:          0,
			DestinationAddr:      "",
			EsmClass:             0,
			ProtocolID:           0,
			PriorityFlag:         0,
			ScheduleDeliveryTime: "",
			ValidityPeriod:       "",
			RegisteredDelivery:   0,
			ReplaceIfPresentFlag: 0,
			DataCoding:           0,
			SmDefaultMessageID:   0,
			SmLength:             0,
			ShortMessage:         "",
		},
	}
	buffer := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		if err := NewEncoder(buffer).Encode(pdu); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDecoder_Decode(b *testing.B) {
	b.ReportAllocs()
	pdu := &SubmitSmPdu{
		Header: &Header{
			CommandID:      EnquireLink,
			CommandStatus:  EsmeRbindFail,
			SequenceNumber: 2,
		},
		Body: &SmBody{
			ServiceType:          "",
			SourceAddrTon:        0,
			SourceAddrNpi:        0,
			SourceAddr:           "",
			DestAddrTon:          0,
			DestAddrNpi:          0,
			DestinationAddr:      "",
			EsmClass:             0,
			ProtocolID:           0,
			PriorityFlag:         0,
			ScheduleDeliveryTime: "",
			ValidityPeriod:       "",
			RegisteredDelivery:   0,
			ReplaceIfPresentFlag: 0,
			DataCoding:           0,
			SmDefaultMessageID:   0,
			SmLength:             0,
			ShortMessage:         "",
		},
	}
	buffer := &bytes.Buffer{}
	if err := NewEncoder(buffer).Encode(pdu); err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		tmp := bytes.NewBuffer(buffer.Bytes())
		_, err := NewDecoder(tmp).Decode()
		if err != nil {
			b.Fatal(err)
		}
	}
}
