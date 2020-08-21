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
const PduHeaderLength uint32 = 16

// Command ids - SMPP v3.4 - 5.1.2.1 page 110-111
const (
	GenericNack         uint32 = 0x80000000
	BindReceiver        uint32 = 0x00000001
	BindReceiverResp    uint32 = 0x80000001
	BindTransmitter     uint32 = 0x00000002
	BindTransmitterResp uint32 = 0x80000002
	QuerySm             uint32 = 0x00000003
	QuerySmResp         uint32 = 0x80000003
	SubmitSm            uint32 = 0x00000004
	SubmitSmResp        uint32 = 0x80000004
	DeliverSm           uint32 = 0x00000005
	DeliverSmResp       uint32 = 0x80000005
	Unbind              uint32 = 0x00000006
	UnbindResp          uint32 = 0x80000006
	ReplaceSm           uint32 = 0x00000007
	ReplaceSmResp       uint32 = 0x80000007
	CancelSm            uint32 = 0x00000008
	CancelSmResp        uint32 = 0x80000008
	BindTransceiver     uint32 = 0x00000009
	BindTransceiverResp uint32 = 0x80000009
	OutBind             uint32 = 0x0000000B
	EnquireLink         uint32 = 0x00000015
	EnquireLinkResp     uint32 = 0x80000015
)

//  Command status - SMPP v3.4 - 5.1.3 page 112-114
const (
	EsmeRok              uint32 = 0x00000000
	EsmeRinvMsgLen       uint32 = 0x00000001
	EsmeRinvCmdLen       uint32 = 0x00000002
	EsmeRinvCmdId        uint32 = 0x00000003
	EsmeRinvBndSts       uint32 = 0x00000004
	EsmeRalyBnd          uint32 = 0x00000005
	EsmeRinvPrtFlg       uint32 = 0x00000006
	EsmeRinvRegDlvFlg    uint32 = 0x00000007
	EsmeRsysErr          uint32 = 0x00000008
	EsmeRinvSrcAdr       uint32 = 0x0000000A
	EsmeRinvDstAdr       uint32 = 0x0000000B
	EsmeRinvMsgId        uint32 = 0x0000000C
	EsmeRbindFail        uint32 = 0x0000000D
	EsmeRinvPaswd        uint32 = 0x0000000E
	EsmeRinvSysId        uint32 = 0x0000000F
	EsmeRcancelFail      uint32 = 0x00000011
	EsmeRreplaceFail     uint32 = 0x00000013
	EsmeRmsgqFul         uint32 = 0x00000014
	EsmeRinvSerTyp       uint32 = 0x00000015
	EsmeRinvNumDests     uint32 = 0x00000033
	EsmeRinvDlName       uint32 = 0x00000034
	EsmeRinvDestFlag     uint32 = 0x00000040
	EsmeRinvSubRep       uint32 = 0x00000042
	EsmeRinvEsmClass     uint32 = 0x00000043
	EsmeRcntSubDl        uint32 = 0x00000044
	EsmeRsubmitFail      uint32 = 0x00000045
	EsmeRinvSrcTon       uint32 = 0x00000048
	EsmeRinvSrcNpi       uint32 = 0x00000049
	EsmeRinvDstTon       uint32 = 0x00000050
	EsmeRinvDstNpi       uint32 = 0x00000051
	EsmeRinvSysTyp       uint32 = 0x00000053
	EsmeRinvRepFlag      uint32 = 0x00000054
	EsmeRinvNumMsgs      uint32 = 0x00000055
	EsmeRthrottled       uint32 = 0x00000058
	EsmeRinvSched        uint32 = 0x00000061
	EsmeRinvExpiry       uint32 = 0x00000062
	EsmeRinvDftMsgId     uint32 = 0x00000063
	EsmeRxTAppn          uint32 = 0x00000064
	EsmeRxPAppn          uint32 = 0x00000065
	EsmeRxRAppn          uint32 = 0x00000066
	EsmeRqueryFail       uint32 = 0x00000067
	EsmeRinvoptParStream uint32 = 0x000000C0
	EsmeRoptParNotAllwd  uint32 = 0x000000C1
	EsmeRinvParLen       uint32 = 0x000000C2
	EsmeRmissingOptParam uint32 = 0x000000C3
	EsmeRinvOptParamVal  uint32 = 0x000000C4
	EsmeRdeliveryFailure uint32 = 0x000000FE
	EsmeRinvDcs          uint32 = 0x00000104
	EsmeRunknownErr      uint32 = 0x000000FF
)

var ErrEsmeRinvMsgLen = errors.New("message length is invalid")
var ErrEsmeRinvCmdLen = errors.New("command length is invalid")
var ErrEsmeRinvCmdId = errors.New("invalid command id")
var ErrEsmeRinvBndSts = errors.New("incorrect bind status for given command")
var ErrEsmeRalyBnd = errors.New("esme already in bound state")
var ErrEsmeRinvPrtFlg = errors.New("invalid priority flag")
var ErrEsmeRinvRegDlvFlg = errors.New("invalid registered delivery flag")
var ErrEsmeRsysErr = errors.New("system error")
var ErrEsmeRinvSrcAdr = errors.New("invalid source address")
var ErrEsmeRinvDstAdr = errors.New("invalid dest addr")
var ErrEsmeRinvMsgId = errors.New("message id is invalid")
var ErrEsmeRbindFail = errors.New("bind failed")
var ErrEsmeRinvPaswd = errors.New("invalid password")
var ErrEsmeRinvSysId = errors.New("invalid system id")
var ErrEsmeRcancelFail = errors.New("cancel sm failed")
var ErrEsmeRreplaceFail = errors.New("replace sm failed")
var ErrEsmeRmsgqFul = errors.New("message queue full")
var ErrEsmeRinvSerTyp = errors.New("invalid service type")
var ErrEsmeRinvNumDests = errors.New("invalid number of destinations")
var ErrEsmeRinvDlName = errors.New("invalid distribution list name")
var ErrEsmeRinvDestFlag = errors.New("invalid destination flag")
var ErrEsmeRinvSubRep = errors.New("invalid submit with replace request")
var ErrEsmeRinvEsmClass = errors.New("invalid esm class set")
var ErrEsmeRcntSubDl = errors.New("cannot submit to distribution list")
var ErrEsmeRsubmitFail = errors.New("submit_sm or submit_multi failed")
var ErrEsmeRinvSrcTon = errors.New("invalid source address ton")
var ErrEsmeRinvSrcNpi = errors.New("invalid source address npi")
var ErrEsmeRinvDstTon = errors.New("invalid destination address ton")
var ErrEsmeRinvDstNpi = errors.New("invalid destination address npi")
var ErrEsmeRinvSysTyp = errors.New("invalid system_type field")
var ErrEsmeRinvRepFlag = errors.New("invalid replace_if_present flag")
var ErrEsmeRinvNumMsgs = errors.New("invalid number of messages")
var ErrEsmeRthrottled = errors.New("throttling error (esme has exceeded allowed message limits)")
var ErrEsmeRinvSched = errors.New("invalid scheduled delivery time")
var ErrEsmeRinvExpiry = errors.New("invalid message (expiry time)")
var ErrEsmeRinvDftMsgId = errors.New("predefined message invalid or not found")
var ErrEsmeRxTAppn = errors.New("esme receiver temporary app error code")
var ErrEsmeRxPAppn = errors.New("esme receiver permanent app error code")
var ErrEsmeRxRAppn = errors.New("esme receiver reject message error code")
var ErrEsmeRqueryFail = errors.New("query_sm request failed")
var ErrEsmeRinvoptParStream = errors.New("error in the optional part of the pdu body")
var ErrEsmeRoptParNotAllwd = errors.New("optional parameter not allowed")
var ErrEsmeRinvParLen = errors.New("invalid parameter length")
var ErrEsmeRmissingOptParam = errors.New("expected optional parameter missing")
var ErrEsmeRinvOptParamVal = errors.New("invalid optional parameter value")
var ErrEsmeRdeliveryFailure = errors.New("delivery failure (data_sm_resp)")
var ErrEsmeRinvDcs = errors.New("invalid data coding scheme")
var ErrEsmeRunknownErr = errors.New("unknown error")

var ErrCodes = map[uint32]error{
	0x00000001: ErrEsmeRinvMsgLen,
	0x00000002: ErrEsmeRinvCmdLen,
	0x00000003: ErrEsmeRinvCmdId,
	0x00000004: ErrEsmeRinvBndSts,
	0x00000005: ErrEsmeRalyBnd,
	0x00000006: ErrEsmeRinvPrtFlg,
	0x00000007: ErrEsmeRinvRegDlvFlg,
	0x00000008: ErrEsmeRsysErr,
	0x0000000A: ErrEsmeRinvSrcAdr,
	0x0000000B: ErrEsmeRinvDstAdr,
	0x0000000C: ErrEsmeRinvMsgId,
	0x0000000D: ErrEsmeRbindFail,
	0x0000000E: ErrEsmeRinvPaswd,
	0x0000000F: ErrEsmeRinvSysId,
	0x00000011: ErrEsmeRcancelFail,
	0x00000013: ErrEsmeRreplaceFail,
	0x00000014: ErrEsmeRmsgqFul,
	0x00000015: ErrEsmeRinvSerTyp,
	0x00000033: ErrEsmeRinvNumDests,
	0x00000034: ErrEsmeRinvDlName,
	0x00000040: ErrEsmeRinvDestFlag,
	0x00000042: ErrEsmeRinvSubRep,
	0x00000043: ErrEsmeRinvEsmClass,
	0x00000044: ErrEsmeRcntSubDl,
	0x00000045: ErrEsmeRsubmitFail,
	0x00000048: ErrEsmeRinvSrcTon,
	0x00000049: ErrEsmeRinvSrcNpi,
	0x00000050: ErrEsmeRinvDstTon,
	0x00000051: ErrEsmeRinvDstNpi,
	0x00000053: ErrEsmeRinvSysTyp,
	0x00000054: ErrEsmeRinvRepFlag,
	0x00000055: ErrEsmeRinvNumMsgs,
	0x00000058: ErrEsmeRthrottled,
	0x00000061: ErrEsmeRinvSched,
	0x00000062: ErrEsmeRinvExpiry,
	0x00000063: ErrEsmeRinvDftMsgId,
	0x00000064: ErrEsmeRxTAppn,
	0x00000065: ErrEsmeRxPAppn,
	0x00000066: ErrEsmeRxRAppn,
	0x00000067: ErrEsmeRqueryFail,
	0x000000C0: ErrEsmeRinvoptParStream,
	0x000000C1: ErrEsmeRoptParNotAllwd,
	0x000000C2: ErrEsmeRinvParLen,
	0x000000C3: ErrEsmeRmissingOptParam,
	0x000000C4: ErrEsmeRinvOptParamVal,
	0x000000FE: ErrEsmeRdeliveryFailure,
	0x00000104: ErrEsmeRinvDcs,
	0x000000FF: ErrEsmeRunknownErr,
}

// Err returns error by code
func Err(code uint32) error {
	if err, ok := ErrCodes[code]; ok {
		return err
	}
	return ErrEsmeRunknownErr
}

// SMPP v3.4 - 5.2.5 page 117
const (
	TonUnknown          uint32 = 0x00
	TonInternational    uint32 = 0x01
	TonNational         uint32 = 0x02
	TonNetworkSpecific  uint32 = 0x03
	TonSubscriberNumber uint32 = 0x04
	TonAlphanumeric     uint32 = 0x05
	TonAbbreviated      uint32 = 0x06
)

// SMPP v3.4 - 5.2.6 page 118
const (
	NpiUnknown   uint32 = 0x00
	NpiE164      uint32 = 0x01
	NpiData      uint32 = 0x03
	NpiTelex     uint32 = 0x04
	NpiE212      uint32 = 0x06
	NpiNational  uint32 = 0x08
	NpiPrivate   uint32 = 0x09
	NpiErmes     uint32 = 0x0a
	NpiInternet  uint32 = 0x0e
	NpiWapclient uint32 = 0x12
)

// ESM bits 1-0 - SMPP v3.4 - 5.2.12 page 121-122
const (
	EsmSubmitModeDatagram        uint32 = 0x01
	EsmSubmitModeForward         uint32 = 0x02
	EsmSubmitModeStoreAndForward uint32 = 0x03
)

// ESM bits 5-2
const (
	EsmSubmitDefault      uint32 = 0x00
	EsmSubmitBinary       uint32 = 0x04
	EsmSubmitTypeEsmeDAck uint32 = 0x08
	EsmSubmitTypeEsmeUAck uint32 = 0x10
	EsmDeliverSmscReceipt uint32 = 0x04
	EsmDeliverSmeAck      uint32 = 0x08
	EsmDeliverUAck        uint32 = 0x10
	EsmDeliverConvAbort   uint32 = 0x18
)

// Intermediate delivery notification
const EsmDeliverIdn uint32 = 0x20

// ESM bits 7-6
const (
	EsmUdhiNone  uint32 = 0x00
	EsmUdhi      uint32 = 0x40
	EsmUdhiDlr   uint32 = 0x04
	EsmReplyPath uint32 = 0x80
)

// SMPP v3.4 - 5.2.13 page 123
const ProtocolId uint32 = 0x34

// SMPP v3.4 - 5.2.14 page 123
const (
	PriorityFlag0 uint32 = 0x00
	PriorityFlag1 uint32 = 0x01
	PriorityFlag2 uint32 = 0x02
	PriorityFlag3 uint32 = 0x03
)

// SMPP v3.4 - 5.2.17 page 124
const RegDeliveryNo uint32 = 0x00

// both success and failure
const (
	RegDeliverySmscBoth   uint32 = 0x01
	RegDeliverySmscFailed uint32 = 0x02
	RegDeliverySmeDAck    uint32 = 0x04
	RegDeliverySmeUAck    uint32 = 0x08
	RegDeliverySmeBoth    uint32 = 0x10
)

// Intermediate notification
const RegDeliveryIdn uint32 = 0x16

// SMPP v3.4 - 5.2.18 page 125
const (
	ReplaceNo  uint32 = 0x00
	ReplaceYes uint32 = 0x01
)

// SMPP v3.4 - 5.2.19 page 126
//UTF-8 as internal SMSC coding
const DataCodingDefault uint32 = 0

// IA5 (CCITT T.50)/ASCII (ANSI X3.4)
const (
	DataCodingIa5         uint32 = 1
	DataCodingBinaryAlias uint32 = 2
	DataCodingIso88591    uint32 = 3
	DataCodingBinary      uint32 = 4
	DataCodingJis         uint32 = 5
	DataCodingIso88595    uint32 = 6
	DataCodingIso88598    uint32 = 7
	DataCodingUcs2        uint32 = 8
	DataCodingPictogram   uint32 = 9
	DataCodingIso2022Jp   uint32 = 10
	DataCodingKanji       uint32 = 13
	DataCodingKsc5601     uint32 = 14
	DataCodingUtf16be     uint32 = 15
)

// SMPP v3.4 - 5.2.21 page 128
const NoUserDataSm uint32 = 0x00

// SMPP v3.4 - 5.2.25 page 129
const (
	DestFlagSme      uint32 = 1
	DestFlagDistlist uint32 = 2
)

// SMPP v3.4 - 5.2.28 page 130
const (
	StateEnroute       uint32 = 1
	StateDelivered     uint32 = 2
	StateExpired       uint32 = 3
	StateDeleted       uint32 = 4
	StateUndeliverable uint32 = 5
	StateAccepted      uint32 = 6
	StateUnknown       uint32 = 7
	StateRejected      uint32 = 8
)

// SMPP v3.4 - 5.2.28 page 132
const (
	DestAddrSubunitTlv          uint32 = 0x0005
	DestNetworkTypeTlv          uint32 = 0x0006
	DestBearerTypeTlv           uint32 = 0x0007
	DestTelematicsIdTlv         uint32 = 0x0008
	SourceAddrSubunitTlv        uint32 = 0x000D
	SourceNetworkTypeTlv        uint32 = 0x000E
	SourceBearerTypeTlv         uint32 = 0x000F
	SourceTelematicsIdTlv       uint32 = 0x0010
	QosTimeToLiveTlv            uint32 = 0x0017
	PayloadTypeTlv              uint32 = 0x0019
	AdditionalStatusInfoTextTlv uint32 = 0x001D
	ReceiptedMessageIdTlv       uint32 = 0x001E
	MsMsgWaitFacilitiesTlv      uint32 = 0x0030
	PrivacyIndicatorTlv         uint32 = 0x0201
	SourceSubaddressTlv         uint32 = 0x0202
	DestSubaddressTlv           uint32 = 0x0203
	UserMessageReferenceTlv     uint32 = 0x0204
	UserResponseCodeTlv         uint32 = 0x0205
	SourcePortTlv               uint32 = 0x020A
	DestinationPortTlv          uint32 = 0x020B
	SarMsgRefNumTlv             uint32 = 0x020C
	LanguageIndicatorTlv        uint32 = 0x020D
	SarTotalSegmentsTlv         uint32 = 0x020E
	SarSegmentSeqnumTlv         uint32 = 0x020F
	ScInterfaceVersionTlv       uint32 = 0x0210
	CallbackNumPresIndTlv       uint32 = 0x0302
	CallbackNumAtagTlv          uint32 = 0x0303
	NumberOfMessagesTlv         uint32 = 0x0304
	CallbackNumTlv              uint32 = 0x0381
	DpfResultTlv                uint32 = 0x0420
	SetDpfTlv                   uint32 = 0x0421
	MsAvailabilityStatusTlv     uint32 = 0x0422
	NetworkErrorCodeTlv         uint32 = 0x0423
	MessagePayloadTlv           uint32 = 0x0424
	DeliveryFailureReasonTlv    uint32 = 0x0425
	MoreMessagesToSendTlv       uint32 = 0x0426
	MessageStateTlv             uint32 = 0x0427
	UssdServiceOpTlv            uint32 = 0x0501
	DisplayTimeTlv              uint32 = 0x1201
	SmsSignalTlv                uint32 = 0x1203
	MsValidityTlv               uint32 = 0x1204
	AlertOnMessageDeliveryTlv   uint32 = 0x130C
	ItsReplyTypeTlv             uint32 = 0x1380
	ItsSessionInfoTlv           uint32 = 0x1383
)

var TlvNames = map[uint32]string{
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
func TlvName(tag uint32) string {
	if name, ok := TlvNames[tag]; ok {
		return name
	}
	return "unknown"
}
