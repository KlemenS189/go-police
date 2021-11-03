package types

import (
	"log"
	"testing"
)

func TestCspReport_Raw(t *testing.T) {
	report := CspReport{
		CspReportDetails{
			BlockedUri:         "/test",
			Disposition:        "test",
			DocumentUri:        "test",
			EffectiveDirective: "test",
			LineNumber:         100,
			OriginalPolicy:     "test",
			Referrer:           "test",
			ScriptSample:       "test",
			SourceFile:         "test",
			ViolatedDirective:  "test",
			StatusCode:         100,
		},
	}
	expected := "blocked-uri:/test;disposition:test;document-uri:test;effective-directive:test;line-number:100;" +
		"original-policy:test;referrer:test;script-sample:test;source-file:test;violated-directive:test;status-code:100;"
	if report.Raw() != expected {
		log.Println(report.Raw(), expected)
		t.Error("Wrong raw output")
	}
}

func TestCspReport_Json(t *testing.T) {
	report := CspReport{
		CspReportDetails{
			BlockedUri:         "/test",
			Disposition:        "test",
			DocumentUri:        "test",
			EffectiveDirective: "test",
			LineNumber:         100,
			OriginalPolicy:     "test",
			Referrer:           "test",
			ScriptSample:       "test",
			SourceFile:         "test",
			StatusCode:         100,
		},
	}
	j := report.Json()
	excpected := "{\"blocked-uri\":\"/test\",\"disposition\":\"test\",\"document-uri\":\"test\",\"effective-directive" +
		"\":\"test\",\"line-number\":100,\"original-policy\":\"test\",\"referrer\":\"test\",\"script-sample\":\"test" +
		"\",\"source-file\":\"test\",\"status-code\":100}"
	if j != excpected {
		t.Error("Wrong json output")
	}
}
