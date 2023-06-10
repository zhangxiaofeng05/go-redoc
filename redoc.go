package go_redoc

import (
	_ "embed"
)

//go:embed assets/redoc.standalone.js
var javaScriptData string

//go:embed assets/index.html
var htmlData string
