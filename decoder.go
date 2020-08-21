package smpp

import (
	"bytes"
	"encoding/binary"
	"io"
	"strings"
)

// Decoder decodes smpp pdu
type Decoder struct {
	r *bytes.Buffer
}

// NewDecoder constructs Decoder
func NewDecoder(r *bytes.Buffer) *Decoder {
	return &Decoder{r: r}
}

// readInt reads smpp integer
func (d *Decoder) readInt(v *uint32) error {
	b := make([]byte, 4)
	n, err := d.r.Read(b)
	if err != nil {
		return err
	}
	if n < len(b) {
		return io.EOF
	}
	*v = binary.BigEndian.Uint32(b)
	return nil
}

// readString reads smpp octet string
func (d *Decoder) readString(v *string, length uint32) error {
	w := &strings.Builder{}
	for i := uint32(0); i < length; i++ {
		b, err := d.r.ReadByte()
		if err != nil {
			return err
		}
		if b == byte(0) {
			break
		}
		w.WriteByte(b)
	}
	*v = w.String()
	return nil
}

// readHeader reads smpp pdu header
func (d *Decoder) readHeader(header *Header) error {
	if err := d.readInt(&header.CommandLength); err != nil {
		return ErrEsmeRinvCmdLen
	}
	if err := d.readInt(&header.CommandID); err != nil {
		return ErrEsmeRinvCmdId
	}
	if err := d.readInt(&header.CommandStatus); err != nil {
		return err
	}
	return d.readInt(&header.SequenceNumber)
}

// readTlvMap reads smpp tlv map
func (d *Decoder) readTlvMap(tlvMap TlvMap) error {
	for d.r.Len() > 0 {
		tlv := new(Tlv)
		if err := d.readInt(&tlv.Tag); err != nil {
			return ErrEsmeRoptParNotAllwd
		}
		if err := d.readInt(&tlv.Length); err != nil {
			return ErrEsmeRinvParLen
		}
		if err := d.readString(&tlv.Value, tlv.Length); err != nil {
			return ErrEsmeRinvOptParamVal
		}
		tlvMap[TlvName(tlv.Tag)] = *tlv
	}
	return nil
}

// readBindBody reads smpp bind body
func (d *Decoder) readBindBody(body *BindBody) error {
	if err := d.readString(&body.SystemID, 16); err != nil {
		return ErrEsmeRinvSysId
	}
	if err := d.readString(&body.Password, 9); err != nil {
		return ErrEsmeRinvPaswd
	}
	if err := d.readString(&body.SystemType, 13); err != nil {
		return ErrEsmeRinvSysTyp
	}
	if err := d.readInt(&body.InterfaceVersion); err != nil {
		return err
	}
	if err := d.readInt(&body.AddrTon); err != nil {
		return err
	}
	if err := d.readInt(&body.AddrNpi); err != nil {
		return err
	}
	return d.readString(&body.AddressRange, 41)
}

// readBindRespBody reads smpp bind resp
func (d *Decoder) readBindRespBody(body *BindRespBody) error {
	return d.readString(&body.SystemID, 16)
}

// readOutBindBody reads smpp outbind body
func (d *Decoder) readOutBindBody(body *OutBindBody) error {
	if err := d.readString(&body.SystemID, 16); err != nil {
		return ErrEsmeRinvSysId
	}
	if err := d.readString(&body.Password, 9); err != nil {
		return ErrEsmeRinvPaswd
	}
	return nil
}

// readSmBody reads smpp short message body
func (d *Decoder) readSmBody(body *SmBody) error {
	if err := d.readString(&body.ServiceType, 6); err != nil {
		return ErrEsmeRinvSerTyp
	}
	if err := d.readInt(&body.SourceAddrTon); err != nil {
		return ErrEsmeRinvSrcTon
	}
	if err := d.readInt(&body.SourceAddrNpi); err != nil {
		return ErrEsmeRinvSrcNpi
	}
	if err := d.readString(&body.SourceAddr, 21); err != nil {
		return ErrEsmeRinvSrcAdr
	}
	if err := d.readInt(&body.DestAddrTon); err != nil {
		return ErrEsmeRinvDstTon
	}
	if err := d.readInt(&body.DestAddrNpi); err != nil {
		return ErrEsmeRinvDstNpi
	}
	if err := d.readString(&body.DestinationAddr, 21); err != nil {
		return ErrEsmeRinvDstAdr
	}
	if err := d.readInt(&body.EsmClass); err != nil {
		return ErrEsmeRinvEsmClass
	}
	if err := d.readInt(&body.ProtocolID); err != nil {
		return err
	}
	if err := d.readInt(&body.PriorityFlag); err != nil {
		return ErrEsmeRinvPrtFlg
	}
	if err := d.readString(&body.ScheduleDeliveryTime, 17); err != nil {
		return ErrEsmeRinvSched
	}
	if err := d.readString(&body.ValidityPeriod, 17); err != nil {
		return ErrEsmeRinvExpiry
	}
	if err := d.readInt(&body.RegisteredDelivery); err != nil {
		return ErrEsmeRinvRegDlvFlg
	}
	if err := d.readInt(&body.ReplaceIfPresentFlag); err != nil {
		return ErrEsmeRinvRepFlag
	}
	if err := d.readInt(&body.DataCoding); err != nil {
		return ErrEsmeRinvDcs
	}
	if err := d.readInt(&body.SmDefaultMessageID); err != nil {
		return ErrEsmeRinvMsgId
	}
	if err := d.readInt(&body.SmLength); err != nil {
		return ErrEsmeRinvMsgLen
	}
	return d.readString(&body.ShortMessage, body.SmLength)
}

// readSmRespBody reads smpp message response
func (d *Decoder) readSmRespBody(body *SmRespBody) error {
	return d.readString(&body.MessageID, 65)
}

// Decode decodes smpp pdu
func (d *Decoder) Decode() (interface{}, error) {
	header := new(Header)
	if err := d.readHeader(header); err != nil {
		return nil, err
	}
	switch header.CommandID {
	case BindReceiver:
		p := &BindReceiverPdu{Header: header, Body: &BindBody{}}
		if err := d.readBindBody(p.Body); err != nil {
			return nil, err
		}
		return p, nil
	case BindReceiverResp:
		p := &BindReceiverRespPdu{
			Header: header,
			Body:   &BindRespBody{},
			Tlv:    TlvMap{},
		}
		if err := d.readBindRespBody(p.Body); err != nil {
			return nil, err
		}
		if d.r.Len() > 0 {
			if err := d.readTlvMap(p.Tlv); err != nil {
				return nil, err
			}
		}
		return p, nil
	case BindTransmitter:
		p := &BindTransmitterPdu{Header: header, Body: &BindBody{}}
		if err := d.readBindBody(p.Body); err != nil {
			return nil, err
		}
		return p, nil
	case BindTransmitterResp:
		p := &BindTransceiverRespPdu{
			Header: header,
			Body:   &BindRespBody{},
			Tlv:    TlvMap{},
		}
		if err := d.readBindRespBody(p.Body); err != nil {
			return nil, err
		}
		if d.r.Len() > 0 {
			if err := d.readTlvMap(p.Tlv); err != nil {
				return nil, err
			}
		}
		return p, nil
	case BindTransceiver:
		p := BindTransceiverPdu{Header: header, Body: &BindBody{}}
		if err := d.readBindBody(p.Body); err != nil {
			return nil, err
		}
		return p, nil
	case BindTransceiverResp:
		p := BindTransceiverRespPdu{
			Header: header,
			Body:   &BindRespBody{},
			Tlv:    TlvMap{},
		}
		if err := d.readBindRespBody(p.Body); err != nil {
			return nil, err
		}
		if d.r.Len() > 0 {
			if err := d.readTlvMap(p.Tlv); err != nil {
				return nil, err
			}
		}
		return p, nil
	case Unbind:
		return &UnbindPdu{Header: header}, nil
	case UnbindResp:
		return &UnbindRespPdu{Header: header}, nil
	case OutBind:
		p := &OutBindPdu{Header: header}
		if err := d.readOutBindBody(p.Body); err != nil {
			return nil, err
		}
		return p, nil
	case SubmitSm:
		p := &SubmitSmPdu{
			Header: header,
			Body:   &SmBody{},
			Tlv:    TlvMap{},
		}
		if err := d.readSmBody(p.Body); err != nil {
			return nil, err
		}
		if d.r.Len() > 0 {
			if err := d.readTlvMap(p.Tlv); err != nil {
				return nil, err
			}
		}
		return p, nil
	case SubmitSmResp:
		p := &SubmitSmRespPdu{Header: header, Body: &SmRespBody{}}
		if err := d.readSmRespBody(p.Body); err != nil {
			return nil, err
		}
		return p, nil
	case DeliverSm:
		p := &DeliverSmPdu{
			Header: header,
			Body:   &SmBody{},
			Tlv:    TlvMap{},
		}
		if err := d.readSmBody(p.Body); err != nil {
			return nil, err
		}
		if d.r.Len() > 0 {
			if err := d.readTlvMap(p.Tlv); err != nil {
				return nil, err
			}
		}
		return p, nil
	case DeliverSmResp:
		p := &DeliverSmRespPdu{Header: header, Body: &SmRespBody{}}
		if err := d.readSmRespBody(p.Body); err != nil {
			return nil, err
		}
		return p, nil
	case EnquireLink:
		return &EnquireLinkPdu{Header: header}, nil
	case EnquireLinkResp:
		return &EnquireLinkRespPdu{Header: header}, nil
	case GenericNack:
		return &GenericNackPdu{Header: header}, nil
	}
	return nil, ErrUnsupportedPdu
}
