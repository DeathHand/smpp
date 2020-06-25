package smpp

import (
	"bytes"
	"io"
)

// Encoder encodes smpp pdu
type Encoder struct {
	w *bytes.Buffer
	h *bytes.Buffer
	b *bytes.Buffer
}

// NewEncoder constructs Encoder
func NewEncoder(w *bytes.Buffer) *Encoder {
	return &Encoder{
		w: w,
		h: &bytes.Buffer{},
		b: &bytes.Buffer{},
	}
}

// writeIntField writes smpp integer
func (e *Encoder) writeInt(v *int, b *bytes.Buffer) error {
	p := make([]byte, 4)
	p[0] = byte(*v >> 24)
	p[1] = byte(*v >> 16)
	p[2] = byte(*v >> 8)
	p[3] = byte(*v)
	n, err := b.Write(p)
	if n < len(p) {
		return io.EOF
	}
	return err
}

// writeString writes smpp octet string
func (e *Encoder) writeString(v *string, b *bytes.Buffer) error {
	n, err := b.WriteString(*v)
	if err != nil {
		return err
	}
	if n < len(*v) {
		return io.EOF
	}
	return b.WriteByte(byte(0))
}

// writeString writes smpp pdu header
func (e *Encoder) writeHeader(header *Header) error {
	header.CommandLength = e.b.Len() + PduHeaderLength
	if err := e.writeInt(&header.CommandLength, e.h); err != nil {
		return err
	}
	if err := e.writeInt(&header.CommandID, e.h); err != nil {
		return err
	}
	if err := e.writeInt(&header.CommandStatus, e.h); err != nil {
		return err
	}
	return e.writeInt(&header.SequenceNumber, e.h)
}

// writeTlvMap writes smpp tlv map
func (e *Encoder) writeTlvMap(tlvMap *TlvMap) error {
	for _, tlv := range *tlvMap {
		if err := e.writeInt(&tlv.Tag, e.b); err != nil {
			return err
		}
		if err := e.writeInt(&tlv.Length, e.b); err != nil {
			return err
		}
		if err := e.writeString(&tlv.Value, e.b); err != nil {
			return err
		}
	}
	return nil
}

// writeBindBody writes smpp bind body
func (e *Encoder) writeBindBody(body *BindBody) error {
	if err := e.writeString(&body.SystemID, e.b); err != nil {
		return err
	}
	if err := e.writeString(&body.Password, e.b); err != nil {
		return err
	}
	if err := e.writeString(&body.SystemType, e.b); err != nil {
		return err
	}
	if err := e.writeInt(&body.InterfaceVersion, e.b); err != nil {
		return err
	}
	if err := e.writeInt(&body.AddrTon, e.b); err != nil {
		return err
	}
	if err := e.writeInt(&body.AddrNpi, e.b); err != nil {
		return err
	}
	return e.writeString(&body.AddressRange, e.b)
}

// writeBindRespBody writes smpp bind resp
func (e *Encoder) writeBindRespBody(body *BindRespBody) error {
	return e.writeString(&body.SystemID, e.b)
}

// writeOutBindBody writes smpp outbind body
func (e *Encoder) writeOutBindBody(body *OutBindBody) error {
	if err := e.writeString(&body.SystemID, e.b); err != nil {
		return err
	}
	return e.writeString(&body.Password, e.b)
}

// writeSmBody writes smpp short message body
func (e *Encoder) writeSmBody(body *SmBody) error {
	if err := e.writeString(&body.ServiceType, e.b); err != nil {
		return err
	}
	if err := e.writeInt(&body.SourceAddrTon, e.b); err != nil {
		return err
	}
	if err := e.writeInt(&body.SourceAddrNpi, e.b); err != nil {
		return err
	}
	if err := e.writeString(&body.SourceAddr, e.b); err != nil {
		return err
	}
	if err := e.writeInt(&body.DestAddrTon, e.b); err != nil {
		return err
	}
	if err := e.writeInt(&body.DestAddrNpi, e.b); err != nil {
		return err
	}
	if err := e.writeString(&body.DestinationAddr, e.b); err != nil {
		return err
	}
	if err := e.writeInt(&body.EsmClass, e.b); err != nil {
		return err
	}
	if err := e.writeInt(&body.ProtocolID, e.b); err != nil {
		return err
	}
	if err := e.writeInt(&body.PriorityFlag, e.b); err != nil {
		return err
	}
	if err := e.writeString(&body.ScheduleDeliveryTime, e.b); err != nil {
		return err
	}
	if err := e.writeString(&body.ValidityPeriod, e.b); err != nil {
		return err
	}
	if err := e.writeInt(&body.RegisteredDelivery, e.b); err != nil {
		return err
	}
	if err := e.writeInt(&body.ReplaceIfPresentFlag, e.b); err != nil {
		return err
	}
	if err := e.writeInt(&body.DataCoding, e.b); err != nil {
		return err
	}
	if err := e.writeInt(&body.SmDefaultMessageID, e.b); err != nil {
		return err
	}
	if err := e.writeInt(&body.SmLength, e.b); err != nil {
		return err
	}
	return e.writeString(&body.ShortMessage, e.b)
}

// writeSmRespBody writes smpp message response
func (e *Encoder) writeSmRespBody(body *SmRespBody) error {
	return e.writeString(&body.MessageID, e.b)
}

// writeBindReceiver writes Bind Receiver smpp pdu
func (e *Encoder) writeBindReceiver(pdu *BindReceiverPdu) error {
	if err := e.writeBindBody(pdu.Body); err != nil {
		return err
	}
	return e.writeHeader(pdu.Header)
}

// writeBindReceiverResp writes Bind Receiver Resp smpp pdu
func (e *Encoder) writeBindReceiverResp(pdu *BindReceiverRespPdu) error {
	if err := e.writeBindRespBody(pdu.Body); err != nil {
		return err
	}
	return e.writeHeader(pdu.Header)
}

// writeBindTransmitter writes Bind Transmitter smpp pdu
func (e *Encoder) writeBindTransmitter(pdu *BindTransmitterPdu) error {
	if err := e.writeBindBody(pdu.Body); err != nil {
		return err
	}
	return e.writeHeader(pdu.Header)
}

// writeBindTransmitterResp writes Bind Transmitter Resp smpp pdu
func (e *Encoder) writeBindTransmitterResp(pdu *BindTransmitterRespPdu) error {
	if err := e.writeBindRespBody(pdu.Body); err != nil {
		return err
	}
	return e.writeHeader(pdu.Header)
}

// writeBindTransceiver writes Bind Transceiver smpp pdu
func (e *Encoder) writeBindTransceiver(pdu *BindTransceiverPdu) error {
	if err := e.writeBindBody(pdu.Body); err != nil {
		return err
	}
	return e.writeHeader(pdu.Header)
}

// writeBindTransceiverResp writes Bind Transceiver Resp smpp pdu
func (e *Encoder) writeBindTransceiverResp(pdu *BindTransceiverRespPdu) error {
	if err := e.writeBindRespBody(pdu.Body); err != nil {
		return err
	}
	return e.writeHeader(pdu.Header)
}

// writeUnbind writes Unbind smpp pdu
func (e *Encoder) writeUnbind(pdu *UnbindPdu) error {
	return e.writeHeader(pdu.Header)
}

// writeUnbindResp writes Unbind Resp smpp pdu
func (e *Encoder) writeUnbindResp(pdu *UnbindRespPdu) error {
	return e.writeHeader(pdu.Header)
}

// writeOutBind writes Out Bind smpp pdu
func (e *Encoder) writeOutBind(pdu *OutBindPdu) error {
	if err := e.writeOutBindBody(pdu.Body); err != nil {
		return err
	}
	return e.writeHeader(pdu.Header)
}

// writeSubmitSm writes Submit Sm smpp pdu
func (e *Encoder) writeSubmitSm(pdu *SubmitSmPdu) error {
	if err := e.writeSmBody(pdu.Body); err != nil {
		return err
	}
	return e.writeHeader(pdu.Header)
}

// writeSubmitSmResp Submit Sm Resp smpp pdu
func (e *Encoder) writeSubmitSmResp(pdu *SubmitSmRespPdu) error {
	if err := e.writeSmRespBody(pdu.Body); err != nil {
		return err
	}
	return e.writeHeader(pdu.Header)
}

// writeDeliverSm writes Deliver Sm smpp pdu
func (e *Encoder) writeDeliverSm(pdu *DeliverSmPdu) error {
	if err := e.writeSmBody(pdu.Body); err != nil {
		return err
	}
	return e.writeHeader(pdu.Header)
}

// writeDeliverSmResp writes Deliver Sm Resp smpp pdu
func (e *Encoder) writeDeliverSmResp(pdu *DeliverSmRespPdu) error {
	if err := e.writeSmRespBody(pdu.Body); err != nil {
		return err
	}
	return e.writeHeader(pdu.Header)
}

// writeEnquireLink writes Enquire Link smpp pdu
func (e *Encoder) writeEnquireLink(pdu *EnquireLinkPdu) error {
	return e.writeHeader(pdu.Header)
}

// writeEnquireLinkResp writes Enquire Link Resp smpp pdu
func (e *Encoder) writeEnquireLinkResp(pdu *EnquireLinkRespPdu) error {
	return e.writeHeader(pdu.Header)
}

// writeGenericNack writes Generic Nac smpp pdu
func (e *Encoder) writeGenericNack(pdu *GenericNackPdu) error {
	return e.writeHeader(pdu.Header)
}

// Encode encodes smpp pdu
func (e *Encoder) Encode(pdu interface{}) error {
	switch p := pdu.(type) {
	case *BindReceiverPdu:
		if err := e.writeBindBody(p.Body); err != nil {
			return err
		}
		if err := e.writeHeader(p.Header); err != nil {
			return err
		}
	case *BindReceiverRespPdu:
		if err := e.writeBindRespBody(p.Body); err != nil {
			return err
		}
		if err := e.writeHeader(p.Header); err != nil {
			return err
		}
	case *BindTransmitterPdu:
		if err := e.writeBindBody(p.Body); err != nil {
			return err
		}
		if err := e.writeHeader(p.Header); err != nil {
			return err
		}
	case *BindTransmitterRespPdu:
		if err := e.writeBindRespBody(p.Body); err != nil {
			return err
		}
		if err := e.writeHeader(p.Header); err != nil {
			return err
		}
	case *BindTransceiverPdu:
		if err := e.writeBindBody(p.Body); err != nil {
			return err
		}
		if err := e.writeHeader(p.Header); err != nil {
			return err
		}
	case *BindTransceiverRespPdu:
		if err := e.writeBindRespBody(p.Body); err != nil {
			return err
		}
		if err := e.writeHeader(p.Header); err != nil {
			return err
		}
	case *UnbindPdu:
		if err := e.writeHeader(p.Header); err != nil {
			return err
		}
	case *UnbindRespPdu:
		if err := e.writeHeader(p.Header); err != nil {
			return err
		}
	case *OutBindPdu:
		if err := e.writeOutBindBody(p.Body); err != nil {
			return err
		}
		if err := e.writeHeader(p.Header); err != nil {
			return err
		}
	case *SubmitSmPdu:
		if err := e.writeSmBody(p.Body); err != nil {
			return err
		}
		if err := e.writeHeader(p.Header); err != nil {
			return err
		}
	case *SubmitSmRespPdu:
		if err := e.writeSmRespBody(p.Body); err != nil {
			return err
		}
		if err := e.writeHeader(p.Header); err != nil {
			return err
		}
	case *DeliverSmPdu:
		if err := e.writeSmBody(p.Body); err != nil {
			return err
		}
		if err := e.writeHeader(p.Header); err != nil {
			return err
		}
	case *DeliverSmRespPdu:
		if err := e.writeSmRespBody(p.Body); err != nil {
			return err
		}
		if err := e.writeHeader(p.Header); err != nil {
			return err
		}
	case *EnquireLinkPdu:
		if err := e.writeHeader(p.Header); err != nil {
			return err
		}
	case *EnquireLinkRespPdu:
		if err := e.writeHeader(p.Header); err != nil {
			return err
		}
	case *GenericNackPdu:
		if err := e.writeHeader(p.Header); err != nil {
			return err
		}
	default:
		return ErrUnsupportedPdu
	}
	n, err := e.w.Write(e.h.Bytes())
	if err != nil {
		return err
	}
	if n < e.h.Len() {
		return io.ErrShortWrite
	}
	n, err = e.w.Write(e.b.Bytes())
	if err != nil {
		return err
	}
	if n < e.b.Len() {
		return io.ErrShortWrite
	}
	return nil
}
