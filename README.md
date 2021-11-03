# Go Police

Package for logging csp violation reports from browsers.

## Features

- Logging to file or stdout
- Logging in JSON or raw format
- If used for a high traffic website, a pass through rate can be specified (0 - 1)

## Usage

Build with

``
go build go_police
``

Flags which can be used to control the behaviour

    -endpoint string
        Endpoint for reporting csp violations (default "/csp/violation-report/")
    -filename string
        File for saving csp violations. Ignored if to-file is set to false. 
        Datetime is appended to beginning (default "csp_violation_report.txt")
    -output-mode string (json, raw)
        Write as JSON or plain text
    -passThroughPercent float
        Percent of requests to be logged (default 1)
        Example: If set to 0.1 only 10% of requests will be logged
    -port string
        Port for violation server (default "8000")
    -to-file
        Switch if output should be written to a file

## Testing

Package e2e contains a simple app which will serve a html with *Content-Security-Policy-Report-Only* header and *
report-uri* pointing to default violation endpoint. It has some scripts which will produce reports to be sent to the
endpoint.