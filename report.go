package smpp

import (
	"regexp"
	"strings"
)

// DeliveryReport is a deliver_sm message representation
// SMPP v3.4 - Appendix B page 167
type DeliveryReport struct {
	ID         string
	Sub        string
	Dlvrd      string
	SubmitDate string
	DoneDate   string
	Stat       string
	Err        string
	Text       string
}

// DeliveryReportParser is a DeliveryReport parser
type DeliveryReportParser struct {
	idRegex          *regexp.Regexp
	subRegexp        *regexp.Regexp
	dlvrdRegexp      *regexp.Regexp
	submitDateRegexp *regexp.Regexp
	doneDateRegexp   *regexp.Regexp
	statRegexp       *regexp.Regexp
	errRegexp        *regexp.Regexp
	textRegexp       *regexp.Regexp
}

// NewDeliveryReportParser parser constructor
func NewDeliveryReportParser() *DeliveryReportParser {
	p := new(DeliveryReportParser)
	p.idRegex, _ = regexp.Compile("id:(\\w+)")
	p.subRegexp, _ = regexp.Compile("sub:(\\w+)")
	p.dlvrdRegexp, _ = regexp.Compile("dlvrd:(\\w+)")
	p.submitDateRegexp, _ = regexp.Compile("submit date:(\\w+)")
	p.doneDateRegexp, _ = regexp.Compile("done date:(\\w+)")
	p.statRegexp, _ = regexp.Compile("stat:(\\w+)")
	p.errRegexp, _ = regexp.Compile("err:(\\w+)")
	p.textRegexp, _ = regexp.Compile("Text:(.+)")
	return p
}

func (p *DeliveryReportParser) splitAttr(attr string) string {
	parts := strings.Split(attr, ":")
	if len(parts) != 2 {
		return ""
	}
	return strings.TrimSpace(parts[1])
}

func (p *DeliveryReportParser) parseAttr(name string, message string) string {
	switch name {
	case "id":
		return p.splitAttr(p.idRegex.FindString(message))
	case "sub":
		return p.splitAttr(p.subRegexp.FindString(message))
	case "dlvrd":
		return p.splitAttr(p.dlvrdRegexp.FindString(message))
	case "submit date":
		return p.splitAttr(p.submitDateRegexp.FindString(message))
	case "done date":
		return p.splitAttr(p.doneDateRegexp.FindString(message))
	case "stat":
		return p.splitAttr(p.subRegexp.FindString(message))
	case "err":
		return p.splitAttr(p.errRegexp.FindString(message))
	case "text":
		return p.splitAttr(p.textRegexp.FindString(message))
	}
	return ""
}

// Parse reads delivery report attributes
func (p *DeliveryReportParser) Parse(message string) *DeliveryReport {
	return &DeliveryReport{
		ID:         p.parseAttr("id", message),
		Sub:        p.parseAttr("sub", message),
		Dlvrd:      p.parseAttr("dlvrd", message),
		SubmitDate: p.parseAttr("submit date", message),
		DoneDate:   p.parseAttr("done date", message),
		Stat:       p.parseAttr("stat", message),
		Err:        p.parseAttr("err", message),
		Text:       p.parseAttr("text", message),
	}
}
