package smpp

import "errors"

// ErrUnsupportedPdu throws when pdu passed to Encode() is not a pointer to supported pdu structure
var ErrUnsupportedPdu = errors.New("pdu unsupported or not a pointer")

// SMPP v3.4 - 2.1 page 13
const (
	ConnectionModeTransmitter string = "TX"
	ConnectionModeReceiver    string = "RX"
	ConnectionModeTransceiver string = "TRX"
)

// SMPP v3.4 - 2.2 page 14
const (
	SessionOpenState    string = "OPEN"
	SessionBondTxState  string = "BOUND_TX"
	SessionBondRxState  string = "BOUND_RX"
	SessionBondTrxState string = "BOUND_TRX"
	SessionClosedState  string = "CLOSED"
)

// SMPP v3.4 - 3.2 page 38
const PduHeaderLength int = 16

// Command ids - SMPP v3.4 - 5.1.2.1 page 110-111
const (
	GenericNack         = 0x80000000
	BindReceiver        = 0x00000001
	BindReceiverResp    = 0x80000001
	BindTransmitter     = 0x00000002
	BindTransmitterResp = 0x80000002
	QuerySm             = 0x00000003
	QuerySmResp         = 0x80000003
	SubmitSm            = 0x00000004
	SubmitSmResp        = 0x80000004
	DeliverSm           = 0x00000005
	DeliverSmResp       = 0x80000005
	Unbind              = 0x00000006
	UnbindResp          = 0x80000006
	ReplaceSm           = 0x00000007
	ReplaceSmResp       = 0x80000007
	CancelSm            = 0x00000008
	CancelSmResp        = 0x80000008
	BindTransceiver     = 0x00000009
	BindTransceiverResp = 0x80000009
	OutBind             = 0x0000000B
	EnquireLink         = 0x00000015
	EnquireLinkResp     = 0x80000015
)

//  Command status - SMPP v3.4 - 5.1.3 page 112-114
const (
	// No Error
	EsmeRok = 0x00000000
	// Message Length is invalid
	EsmeRinvMsgLen = 0x00000001
	// Command Length is invalid
	EsmeRinvCmdLen = 0x00000002
	// Invalid Command ID
	EsmeRinvCmdId = 0x00000003
	// Incorrect BIND Status for given command
	EsmeRinvBndSts = 0x00000004
	// ESME Already in Bound State
	EsmeRalyBnd = 0x00000005
	// Invalid Priority Flag
	EsmeRinvPrtFlg = 0x00000006
	// Invalid Registered Delivery Flag
	EsmeRinvRegDlvFlg = 0x00000007
	// System Error
	EsmeRsysErr = 0x00000008
	// Invalid Source Address
	EsmeRinvSrcAdr = 0x0000000A
	// Invalid Dest Addr
	EsmeRinvDstAdr = 0x0000000B
	// Message ID is invalid
	EsmeRinvMsgId = 0x0000000C
	// Bind Failed
	EsmeRbindFail = 0x0000000D
	// Invalid Password
	EsmeRinvPaswd = 0x0000000E
	// Invalid System ID
	EsmeRinvSysId = 0x0000000F
	// Cancel SM Failed
	EsmeRcancelfail = 0x00000011
	// Replace SM Failed
	EsmeRreplaceFail = 0x00000013
	// Message Queue Full
	EsmeRmsgqFul = 0x00000014
	// Invalid Service Type
	EsmeRinvSerTyp = 0x00000015
	// Invalid number of destinations
	EsmeRinvNumDests = 0x00000033
	// Invalid Distribution List name
	EsmeRinvDlName = 0x00000034
	// Destination flag (submit_multi)
	EsmeRinvDestFlag = 0x00000040
	// Invalid ‘submit with replace’ request (i.e. submit_sm with replace_if_present_flag set)
	EsmeRinvSubRep = 0x00000042
	// Invalid ESM_SUBMIT field data
	EsmeRinvEsmSubmit = 0x00000043
	// Cannot Submit to Distribution List
	EsmeRcntSubDl = 0x00000044
	// submit_sm or submit_multi failed
	EsmeRsubmitFail = 0x00000045
	// Invalid Source address TON
	EsmeRinvSrcTon = 0x00000048
	// Invalid Source address NPI
	EsmeRinvSrcNpi = 0x00000049
	// Invalid Destination address TON
	EsmeRinvDstTon = 0x00000050
	// Invalid Destination address NPI
	EsmeRinvDstNpi = 0x00000051
	// Invalid system_type field
	EsmeRinvSysTyp = 0x00000053
	// Invalid replace_if_present flag
	EsmeRinvRepFlag = 0x00000054
	// Invalid number of messages
	EsmeRinvNumMsgs = 0x00000055
	// Throttling error (ESME has exceeded allowed message limits)
	EsmeRthrottled = 0x00000058
	// Invalid Scheduled Delivery Time
	EsmeRinvSched = 0x00000061
	// Invalid message (Expiry time)
	EsmeRinvExpiry = 0x00000062
	// Predefined Message Invalid or Not Found
	EsmeRinvDftMsgId = 0x00000063
	// ESME Receiver Temporary App Error Code
	EsmeRxTAppn = 0x00000064
	// ESME Receiver Permanent App Error Code
	EsmeRxPAppn = 0x00000065
	// ESME Receiver Reject Message Error Code
	EsmeRxRAppn = 0x00000066
	// query_sm request failed
	EsmeRqueryFail = 0x00000067
	// Error in the optional part of the PDU Body.
	EsmeRinvoptParStream = 0x000000C0
	// Optional Parameter not allowed
	EsmeRoptParNotAllwd = 0x000000C1
	// Invalid Parameter Length.
	EsmeRinvParLen = 0x000000C2
	// Expected Optional Parameter missing
	EsmeRmissingOptParam = 0x000000C3
	// Invalid Optional Parameter Value
	EsmeRinvOptParamVal = 0x000000C4
	// Delivery Failure (data_sm_resp)
	EsmeRdeliveryFailure = 0x000000FE
	// Unknown Error
	EsmeRunknownErr = 0x000000FF
)

// SMPP v3.4 - 5.2.5 page 117
const (
	TonUnknown          = 0x00
	TonInternational    = 0x01
	TonNational         = 0x02
	TonNetworkSpecific  = 0x03
	TonSubscriberNumber = 0x04
	TonAlphanumeric     = 0x05
	TonAbbreviated      = 0x06
)

// SMPP v3.4 - 5.2.6 page 118
const (
	NpiUnknown   = 0x00
	NpiE164      = 0x01
	NpiData      = 0x03
	NpiTelex     = 0x04
	NpiE212      = 0x06
	NpiNational  = 0x08
	NpiPrivate   = 0x09
	NpiErmes     = 0x0a
	NpiInternet  = 0x0e
	NpiWapclient = 0x12
)

// ESM bits 1-0 - SMPP v3.4 - 5.2.12 page 121-122
const (
	EsmSubmitModeDatagram        = 0x01
	EsmSubmitModeForward         = 0x02
	EsmSubmitModeStoreAndForward = 0x03
)

// ESM bits 5-2
const (
	EsmSubmitDefault      = 0x00
	EsmSubmitBinary       = 0x04
	EsmSubmitTypeEsmeDAck = 0x08
	EsmSubmitTypeEsmeUAck = 0x10
	EsmDeliverSmscReceipt = 0x04
	EsmDeliverSmeAck      = 0x08
	EsmDeliverUAck        = 0x10
	EsmDeliverConvAbort   = 0x18
)

// Intermediate delivery notification
const EsmDeliverIdn = 0x20

// ESM bits 7-6
const (
	EsmUdhiNone  = 0x00
	EsmUdhi      = 0x40
	EsmUdhiDlr   = 0x04
	EsmReplyPath = 0x80
)

// SMPP v3.4 - 5.2.13 page 123
const ProtocolId = 0x34

// SMPP v3.4 - 5.2.14 page 123
const (
	PriorityFlag0 = 0x00
	PriorityFlag1 = 0x01
	PriorityFlag2 = 0x02
	PriorityFlag3 = 0x03
)

// SMPP v3.4 - 5.2.17 page 124
const RegDeliveryNo = 0x00

// both success and failure
const (
	RegDeliverySmscBoth   = 0x01
	RegDeliverySmscFailed = 0x02
	RegDeliverySmeDAck    = 0x04
	RegDeliverySmeUAck    = 0x08
	RegDeliverySmeBoth    = 0x10
)

// Intermediate notification
const RegDeliveryIdn = 0x16

// SMPP v3.4 - 5.2.18 page 125
const (
	ReplaceNo  = 0x00
	ReplaceYes = 0x01
)

// SMPP v3.4 - 5.2.19 page 126
//UTF-8 as internal SMSC coding
const DataCodingDefault = 0

// IA5 (CCITT T.50)/ASCII (ANSI X3.4)
const (
	DataCodingIa5         = 1
	DataCodingBinaryAlias = 2
	// Latin 1
	DataCodingIso88591 = 3
	DataCodingBinary   = 4
	DataCodingJis      = 5
	// Cyrllic
	DataCodingIso88595 = 6
	// Latin/Hebrew
	DataCodingIso88598 = 7
	// UCS-2BE (Big Endian)
	DataCodingUcs2      = 8
	DataCodingPictogram = 9
	// Music codes
	DataCodingIso2022Jp = 10
	// Extended Kanji JIS
	DataCodingKanji   = 13
	DataCodingKsc5601 = 14
	DataCodingUtf16be = 15
)

// SMPP v3.4 - 5.2.21 page 128
const NoUserDataSm = 0x00

// SMPP v3.4 - 5.2.25 page 129
const (
	DestFlagSme      = 1
	DestFlagDistlist = 2
)

// SMPP v3.4 - 5.2.28 page 130
const (
	StateEnroute       = 1
	StateDelivered     = 2
	StateExpired       = 3
	StateDeleted       = 4
	StateUndeliverable = 5
	StateAccepted      = 6
	StateUnknown       = 7
	StateRejected      = 8
)

// SMPP v3.4 - 5.2.28 page 132
const (
	DestAddrSubunitTlv          = 0x0005
	DestNetworkTypeTlv          = 0x0006
	DestBearerTypeTlv           = 0x0007
	DestTelematicsIdTlv         = 0x0008
	SourceAddrSubunitTlv        = 0x000D
	SourceNetworkTypeTlv        = 0x000E
	SourceBearerTypeTlv         = 0x000F
	SourceTelematicsIdTlv       = 0x0010
	QosTimeToLiveTlv            = 0x0017
	PayloadTypeTlv              = 0x0019
	AdditionalStatusInfoTextTlv = 0x001D
	ReceiptedMessageIdTlv       = 0x001E
	MsMsgWaitFacilitiesTlv      = 0x0030
	PrivacyIndicatorTlv         = 0x0201
	SourceSubaddressTlv         = 0x0202
	DestSubaddressTlv           = 0x0203
	UserMessageReferenceTlv     = 0x0204
	UserResponseCodeTlv         = 0x0205
	SourcePortTlv               = 0x020A
	DestinationPortTlv          = 0x020B
	SarMsgRefNumTlv             = 0x020C
	LanguageIndicatorTlv        = 0x020D
	SarTotalSegmentsTlv         = 0x020E
	SarSegmentSeqnumTlv         = 0x020F
	ScInterfaceVersionTlv       = 0x0210
	CallbackNumPresIndTlv       = 0x0302
	CallbackNumAtagTlv          = 0x0303
	NumberOfMessagesTlv         = 0x0304
	CallbackNumTlv              = 0x0381
	DpfResultTlv                = 0x0420
	SetDpfTlv                   = 0x0421
	MsAvailabilityStatusTlv     = 0x0422
	NetworkErrorCodeTlv         = 0x0423
	MessagePayloadTlv           = 0x0424
	DeliveryFailureReasonTlv    = 0x0425
	MoreMessagesToSendTlv       = 0x0426
	MessageStateTlv             = 0x0427
	UssdServiceOpTlv            = 0x0501
	DisplayTimeTlv              = 0x1201
	SmsSignalTlv                = 0x1203
	MsValidityTlv               = 0x1204
	AlertOnMessageDeliveryTlv   = 0x130C
	ItsReplyTypeTlv             = 0x1380
	ItsSessionInfoTlv           = 0x1383
)

var TlvNames = map[int]string{
	DestAddrSubunitTlv:          "dest_addr_subunit",
	DestNetworkTypeTlv:          "dest_network_type",
	DestBearerTypeTlv:           "dest_bearer_type",
	DestTelematicsIdTlv:         "dest_telematics_id",
	SourceAddrSubunitTlv:        "source_addr_subunit",
	SourceNetworkTypeTlv:        "source_network_type",
	SourceBearerTypeTlv:         "source_bearer_type",
	SourceTelematicsIdTlv:       "source_telematics_id",
	QosTimeToLiveTlv:            "qos_time_to_live",
	PayloadTypeTlv:              "payload_type",
	AdditionalStatusInfoTextTlv: "additional_status_info_text",
	ReceiptedMessageIdTlv:       "receipted_message_id",
	MsMsgWaitFacilitiesTlv:      "ms_msg_wait_facilities",
	PrivacyIndicatorTlv:         "privacy_indicator",
	SourceSubaddressTlv:         "source_subaddress",
	DestSubaddressTlv:           "dest_subaddress",
	UserMessageReferenceTlv:     "user_message_reference",
	UserResponseCodeTlv:         "user_response_code",
	SourcePortTlv:               "source_port",
	DestinationPortTlv:          "destination_port",
	SarMsgRefNumTlv:             "sar_msg_ref_num",
	LanguageIndicatorTlv:        "language_indicator",
	SarTotalSegmentsTlv:         "sar_total_segments",
	SarSegmentSeqnumTlv:         "sar_segment_seqnum",
	ScInterfaceVersionTlv:       "sc_interface_version",
	CallbackNumPresIndTlv:       "callback_num_pres_ind",
	CallbackNumAtagTlv:          "callback_num_atag",
	NumberOfMessagesTlv:         "number_of_messages",
	CallbackNumTlv:              "callback_num",
	DpfResultTlv:                "dpf_result",
	SetDpfTlv:                   "set_dpf",
	MsAvailabilityStatusTlv:     "ms_availability_status",
	NetworkErrorCodeTlv:         "network_error_code",
	MessagePayloadTlv:           "message_payload",
	DeliveryFailureReasonTlv:    "delivery_failure_reason",
	MoreMessagesToSendTlv:       "more_messages_to_send",
	MessageStateTlv:             "message_state",
	UssdServiceOpTlv:            "ussd_service_op",
	DisplayTimeTlv:              "display_time",
	SmsSignalTlv:                "sms_signal",
	MsValidityTlv:               "ms_validity",
	AlertOnMessageDeliveryTlv:   "alert_on_message_delivery",
	ItsReplyTypeTlv:             "its_reply_type",
	ItsSessionInfoTlv:           "its_session_info",
}

// TlvName returns tlv tag name by tag id
func TlvName(tag int) string {
	if name, ok := TlvNames[tag]; ok {
		return name
	}
	return "unknown"
}
