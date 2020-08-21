package smpp

type Header struct {
	CommandLength  uint32
	CommandID      uint32
	CommandStatus  uint32
	SequenceNumber uint32
}

type Tlv struct {
	Tag    uint32
	Length uint32
	Value  string
}

type TlvMap map[string]Tlv

type BindBody struct {
	SystemID         string
	Password         string
	SystemType       string
	InterfaceVersion uint32
	AddrTon          uint32
	AddrNpi          uint32
	AddressRange     string
}

type OutBindBody struct {
	SystemID string
	Password string
}

type BindRespBody struct {
	SystemID string
}

type SmBody struct {
	ServiceType          string
	SourceAddrTon        uint32
	SourceAddrNpi        uint32
	SourceAddr           string
	DestAddrTon          uint32
	DestAddrNpi          uint32
	DestinationAddr      string
	EsmClass             uint32
	ProtocolID           uint32
	PriorityFlag         uint32
	ScheduleDeliveryTime string
	ValidityPeriod       string
	RegisteredDelivery   uint32
	ReplaceIfPresentFlag uint32
	DataCoding           uint32
	SmDefaultMessageID   uint32
	SmLength             uint32
	ShortMessage         string
}

type SmRespBody struct {
	MessageID string
}

type BindReceiverPdu struct {
	Header *Header
	Body   *BindBody
}

type BindReceiverRespPdu struct {
	Header *Header
	Body   *BindRespBody
	Tlv    TlvMap
}

type BindTransmitterPdu struct {
	Header *Header
	Body   *BindBody
}

type BindTransmitterRespPdu struct {
	Header *Header
	Body   *BindRespBody
	Tlv    TlvMap
}

type BindTransceiverPdu struct {
	Header *Header
	Body   *BindBody
}

type BindTransceiverRespPdu struct {
	Header *Header
	Body   *BindRespBody
	Tlv    TlvMap
}

type SubmitSmPdu struct {
	Header *Header
	Body   *SmBody
	Tlv    TlvMap
}

type SubmitSmRespPdu struct {
	Header *Header
	Body   *SmRespBody
}

type DeliverSmPdu struct {
	Header *Header
	Body   *SmBody
	Tlv    TlvMap
}

type DeliverSmRespPdu struct {
	Header *Header
	Body   *SmRespBody
}

type EnquireLinkPdu struct {
	Header *Header
}

type EnquireLinkRespPdu struct {
	Header *Header
}

type GenericNackPdu struct {
	Header *Header
}

type UnbindPdu struct {
	Header *Header
}

type UnbindRespPdu struct {
	Header *Header
}

type OutBindPdu struct {
	Header *Header
	Body   *OutBindBody
}
