package defaults

import (
	"github.com/wlMalk/mirango/framework"
)

var (
	MimeType = framework.MIME_JSON

	MimeInAccept = true

	Indent = true

	Schemes = []string{framework.SCHEME_HTTP, framework.SCHEME_HTTPS}
	Accepts = []string{framework.MIME_JSON}
	Returns = []string{framework.MIME_JSON}
)
