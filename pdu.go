package smpp

type Header struct {
	CommandLength  int
	CommandID      int
	CommandStatus  int
	SequenceNumber int
}

type Tlv struct {
	Tag    int
	Length int
	Value  string
}

type TlvMap map[string]Tlv

type BindBody struct {
	SystemID         string
	Password         string
	SystemType       string
	InterfaceVersion int
	AddrTon          int
	AddrNpi          int
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
	SourceAddrTon        int
	SourceAddrNpi        int
	SourceAddr           string
	DestAddrTon          int
	DestAddrNpi          int
	DestinationAddr      string
	EsmClass             int
	ProtocolID           int
	PriorityFlag         int
	ScheduleDeliveryTime string
	ValidityPeriod       string
	RegisteredDelivery   int
	ReplaceIfPresentFlag int
	DataCoding           int
	SmDefaultMessageID   int
	SmLength             int
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
