package defaults

import (
	"github.com/mirango/framework"
)

var (
	MimeType = framework.MIME_JSON

	MimeInAccept = true

	Indent = true

	MaxMemory int64 = 32 << 20 // 32 MB

	Schemes = []string{framework.SCHEME_HTTP, framework.SCHEME_HTTPS}
	Accepts = []string{framework.MIME_JSON}
	Returns = []string{framework.MIME_JSON}
)
