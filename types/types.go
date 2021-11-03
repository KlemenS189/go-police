package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

type CspReport struct {
	Details CspReportDetails `json:"csp-report"`
}

func (c CspReport) Raw() string {
	return c.Details.Raw()
}

func (c CspReport) Json() string {
	return c.Details.Json()
}

type CspReportDetails struct {
	BlockedUri         string `json:"blocked-uri,omitempty"`
	Disposition        string `json:"disposition,omitempty"`
	DocumentUri        string `json:"document-uri,omitempty"`
	EffectiveDirective string `json:"effective-directive,omitempty"`
	LineNumber         int    `json:"line-number,omitempty"`
	OriginalPolicy     string `json:"original-policy,omitempty"`
	Referrer           string `json:"referrer,omitempty"`
	ScriptSample       string `json:"script-sample,omitempty"`
	SourceFile         string `json:"source-file,omitempty"`
	ViolatedDirective  string `json:"violated-directive,omitempty"`
	StatusCode         int    `json:"status-code,omitempty"`
}

func (c CspReportDetails) Raw() string {
	builder := strings.Builder{}
	if c.BlockedUri != "" {
		builder.WriteString(fmt.Sprintf("blocked-uri:%s;", c.BlockedUri))
	}
	if c.Disposition != "" {
		builder.WriteString(fmt.Sprintf("disposition:%s;", c.Disposition))
	}
	if c.DocumentUri != "" {
		builder.WriteString(fmt.Sprintf("document-uri:%s;", c.DocumentUri))
	}
	if c.EffectiveDirective != "" {
		builder.WriteString(fmt.Sprintf("effective-directive:%s;", c.EffectiveDirective))
	}
	if c.LineNumber != 0 {
		builder.WriteString(fmt.Sprintf("line-number:%d;", c.LineNumber))
	}
	if c.OriginalPolicy != "" {
		builder.WriteString(fmt.Sprintf("original-policy:%s;", c.OriginalPolicy))
	}
	if c.Referrer != "" {
		builder.WriteString(fmt.Sprintf("referrer:%s;", c.Referrer))
	}
	if c.ScriptSample != "" {
		builder.WriteString(fmt.Sprintf("script-sample:%s;", c.ScriptSample))
	}
	if c.SourceFile != "" {
		builder.WriteString(fmt.Sprintf("source-file:%s;", c.SourceFile))
	}
	if c.ViolatedDirective != "" {
		builder.WriteString(fmt.Sprintf("violated-directive:%s;", c.ViolatedDirective))
	}
	if c.StatusCode != 0 {
		builder.WriteString(fmt.Sprintf("status-code:%d;", c.StatusCode))
	}
	return builder.String()
}

func (c CspReportDetails) Json() string {
	b, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(b)
}
