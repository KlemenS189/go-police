package main

import (
	"log"
	"net/http"
)

const html = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
<h1>Testing site</h1>
<p>See console for details</p>
<!--Try to execute an inline script-->

<button onclick=clickButton>Button</button>

<script>
    const today = new Date()
    console.log(today)

	function clickButton() {
		alert("bad button")
	}
</script>

<!--Try to load axios-->
<script src="https://cdnjs.cloudflare.com/ajax/libs/axios/0.24.0/axios.min.js"
        integrity="sha512-u9akINsQsAkG9xjc1cnGF4zw5TFDwkxuc9vUp5dltDWYCSmyd0meygbvgXrlc/z7/o4a19Fb5V0OUE58J7dcyw=="
        crossorigin="anonymous" referrerpolicy="no-referrer"></script>
</body>
</html>`

// This is used as an e2e test for simulating csp violations
func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
		}
		writer.Header().Set("Content-Security-Policy-Report-Only", "default-src 'self'; report-uri http://localhost:8000/csp/violation-report/")
		_, _ = writer.Write([]byte(html))
	})

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
