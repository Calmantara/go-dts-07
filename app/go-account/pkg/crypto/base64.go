package crypto

import (
	"context"
	"encoding/base64"

	"github.com/Calmantara/go-common/pkg/logger"
)

func DecodeBase64(ctx context.Context, token string) (payload string, err error) {
	logger.Info(ctx, "DecodeBase64 invoked")
	bpayload, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return
	}
	payload = string(bpayload)
	return
}
