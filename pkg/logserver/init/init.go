package init

import "github.com/boeing-ai-gateway/cmd/pkg/logserver"

func init() {
	go logserver.StartServerWithDefaults()
}
