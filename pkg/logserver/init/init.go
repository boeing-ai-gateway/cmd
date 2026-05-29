package init

import "github.com/obot-platform/cmd/pkg/logserver"

func init() {
	go logserver.StartServerWithDefaults()
}
