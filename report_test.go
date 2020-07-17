package smpp

import (
	"testing"
)

var fixture = "id:IIIIIIIIII sub:SSS dlvrd:DDD submit date:YYMMDDhhmm done date:YYMMDDhhmm stat:DDDDDDD err:E Text: . . . . . . . . ."

func TestNewDeliveryReportParser(t *testing.T) {
	parser := NewDeliveryReportParser()
	report := parser.Parse(fixture)
	t.Log(report.DoneDate)
}

func BenchmarkDeliveryReportParser_Parse(b *testing.B) {
	b.ReportAllocs()
	parser := NewDeliveryReportParser()
	for i := 0; i < b.N; i++ {
		parser.Parse(fixture)
	}
}
