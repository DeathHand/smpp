package smpp

import (
	"bytes"
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
func (d *Decoder) readInt(v *int) error {
	p := make([]byte, 4)
	n, err := d.r.Read(p)
	if err != nil {
		return err
	}
	if n < len(p) {
		return io.EOF
	}
	*v = int(p[3]) | int(p[2])<<8 | int(p[1])<<16 | int(p[0])<<24
	return nil
}

// readString reads smpp octet string
func (d *Decoder) readString(v *string, length int) error {
	w := &strings.Builder{}
	for i := 0; i < length; i++ {
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
		return err
	}
	if err := d.readInt(&header.CommandID); err != nil {
		return err
	}
	if err := d.readInt(&header.CommandStatus); err != nil {
		return err
	}
	return d.readInt(&header.SequenceNumber)
}

// readTlvMap reads smpp tlv map
func (d *Decoder) readTlvMap(tlvMap *TlvMap) error {
	for d.r.Len() > 0 {
		tlv := new(Tlv)
		if err := d.readInt(&tlv.Tag); err != nil {
			return err
		}
		if err := d.readInt(&tlv.Length); err != nil {
			return err
		}
		if err := d.readString(&tlv.Value, tlv.Length); err != nil {
			return err
		}
		(*tlvMap)[TlvName(tlv.Tag)] = *tlv
	}
	return nil
}

// readBindBody reads smpp bind body
func (d *Decoder) readBindBody(body *BindBody) error {
	if err := d.readString(&body.SystemID, 16); err != nil {
		return err
	}
	if err := d.readString(&body.Password, 9); err != nil {
		return err
	}
	if err := d.readString(&body.SystemType, 13); err != nil {
		return err
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
		return err
	}
	return d.readString(&body.Password, 9)
}

// readSmBody reads smpp short message body
func (d *Decoder) readSmBody(body *SmBody) error {
	if err := d.readString(&body.ServiceType, 6); err != nil {
		return err
	}
	if err := d.readInt(&body.SourceAddrTon); err != nil {
		return err
	}
	if err := d.readInt(&body.SourceAddrNpi); err != nil {
		return err
	}
	if err := d.readString(&body.SourceAddr, 21); err != nil {
		return err
	}
	if err := d.readInt(&body.DestAddrTon); err != nil {
		return err
	}
	if err := d.readInt(&body.DestAddrNpi); err != nil {
		return err
	}
	if err := d.readString(&body.DestinationAddr, 21); err != nil {
		return err
	}
	if err := d.readInt(&body.EsmClass); err != nil {
		return err
	}
	if err := d.readInt(&body.ProtocolID); err != nil {
		return err
	}
	if err := d.readInt(&body.PriorityFlag); err != nil {
		return err
	}
	if err := d.readInt(&body.PriorityFlag); err != nil {
		return err
	}
	if err := d.readString(&body.ScheduleDeliveryTime, 17); err != nil {
		return err
	}
	if err := d.readString(&body.ValidityPeriod, 17); err != nil {
		return err
	}
	if err := d.readInt(&body.RegisteredDelivery); err != nil {
		return err
	}
	if err := d.readInt(&body.ReplaceIfPresentFlag); err != nil {
		return err
	}
	if err := d.readInt(&body.DataCoding); err != nil {
		return err
	}
	if err := d.readInt(&body.SmDefaultMessageID); err != nil {
		return err
	}
	if err := d.readInt(&body.SmLength); err != nil {
		return err
	}
	return d.readString(&body.ShortMessage, body.SmLength)
}

// readSmRespBody reads smpp message response
func (d *Decoder) readSmRespBody(body *SmRespBody) error {
	return d.readString(&body.MessageID, 65)
}

func (d *Decoder) readBindReceiver(header *Header) (*BindReceiverPdu, error) {
	pdu := &BindReceiverPdu{Header: header}
	if err := d.readBindBody(pdu.Body); err != nil {
		return nil, err
	}
	return pdu, nil
}

func (d *Decoder) readBindReceiverResp(header *Header) (*BindReceiverRespPdu, error) {
	pdu := &BindReceiverRespPdu{Header: header}
	if err := d.readBindRespBody(pdu.Body); err != nil {
		return nil, err
	}
	if err := d.readTlvMap(pdu.Tlv); err != nil {
		return nil, err
	}
	return pdu, nil
}

func (d *Decoder) readBindTransmitter(header *Header) (*BindTransmitterPdu, error) {
	pdu := &BindTransmitterPdu{Header: header}
	if err := d.readBindBody(pdu.Body); err != nil {
		return nil, err
	}
	return pdu, nil
}

func (d *Decoder) readBindTransmitterResp(header *Header) (*BindTransmitterRespPdu, error) {
	pdu := &BindTransmitterRespPdu{Header: header}
	if err := d.readBindRespBody(pdu.Body); err != nil {
		return nil, err
	}
	if err := d.readTlvMap(pdu.Tlv); err != nil {
		return nil, err
	}
	return pdu, nil
}

func (d *Decoder) readBindTransceiver(header *Header) (*BindTransceiverPdu, error) {
	pdu := &BindTransceiverPdu{Header: header}
	if err := d.readBindBody(pdu.Body); err != nil {
		return nil, err
	}
	return pdu, nil
}

func (d *Decoder) readBindTransceiverResp(header *Header) (*BindTransceiverRespPdu, error) {
	pdu := &BindTransceiverRespPdu{Header: header}
	if err := d.readBindRespBody(pdu.Body); err != nil {
		return nil, err
	}
	if err := d.readTlvMap(pdu.Tlv); err != nil {
		return nil, err
	}
	return pdu, nil
}

func (d *Decoder) readUnbind(header *Header) (*UnbindPdu, error) {
	return &UnbindPdu{Header: header}, nil
}

func (d *Decoder) readUnbindResp(header *Header) (*UnbindRespPdu, error) {
	return &UnbindRespPdu{Header: header}, nil
}

func (d *Decoder) readOutBind(header *Header) (*OutBindPdu, error) {
	pdu := &OutBindPdu{Header: header}
	if err := d.readOutBindBody(pdu.Body); err != nil {
		return nil, err
	}
	return pdu, nil
}

func (d *Decoder) readSubmitSm(header *Header) (*SubmitSmPdu, error) {
	pdu := &SubmitSmPdu{Header: header}
	if err := d.readSmBody(pdu.Body); err != nil {
		return nil, err
	}
	return pdu, nil
}

func (d *Decoder) readSubmitSmResp(header *Header) (*SubmitSmRespPdu, error) {
	pdu := &SubmitSmRespPdu{Header: header}
	if err := d.readSmRespBody(pdu.Body); err != nil {
		return nil, err
	}
	return pdu, nil
}

func (d *Decoder) readDeliverSm(header *Header) (*DeliverSmPdu, error) {
	pdu := &DeliverSmPdu{Header: header}
	if err := d.readSmBody(pdu.Body); err != nil {
		return nil, err
	}
	if err := d.readTlvMap(pdu.Tlv); err != nil {
		return nil, err
	}
	return pdu, nil
}

func (d *Decoder) readDeliverSmResp(header *Header) (*DeliverSmRespPdu, error) {
	pdu := &DeliverSmRespPdu{Header: header}
	if err := d.readSmRespBody(pdu.Body); err != nil {
		return nil, err
	}
	return pdu, nil
}

func (d *Decoder) readEnquireLink(header *Header) (*EnquireLinkPdu, error) {
	return &EnquireLinkPdu{Header: header}, nil
}

func (d *Decoder) readEnquireLinkResp(header *Header) (*EnquireLinkRespPdu, error) {
	return &EnquireLinkRespPdu{Header: header}, nil
}

func (d *Decoder) readGenericNack(header *Header) (*GenericNackPdu, error) {
	return &GenericNackPdu{Header: header}, nil
}

// Decode decodes smpp pdu
func (d *Decoder) Decode() (interface{}, error) {
	header := new(Header)
	if err := d.readHeader(header); err != nil {
		return nil, err
	}
	switch header.CommandID {
	case BindReceiver:
		return d.readBindReceiver(header)
	case BindReceiverResp:
		return d.readBindReceiverResp(header)
	case BindTransmitter:
		return d.readBindTransmitter(header)
	case BindTransmitterResp:
		return d.readBindTransmitterResp(header)
	case BindTransceiver:
		return d.readBindTransceiver(header)
	case BindTransceiverResp:
		return d.readBindTransceiverResp(header)
	case Unbind:
		return d.readUnbind(header)
	case UnbindResp:
		return d.readUnbindResp(header)
	case OutBind:
		return d.readOutBind(header)
	case SubmitSm:
		return d.readSubmitSm(header)
	case SubmitSmResp:
		return d.readSubmitSmResp(header)
	case DeliverSm:
		return d.readDeliverSm(header)
	case DeliverSmResp:
		return d.readDeliverSmResp(header)
	case EnquireLink:
		return d.readEnquireLink(header)
	case EnquireLinkResp:
		return d.readEnquireLinkResp(header)
	case GenericNack:
		return d.readGenericNack(header)
	}
	return nil, ErrUnsupportedPdu
}
