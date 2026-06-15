package lura

import (
	"context"
	"net/http"

	kotelhttpserver "github.com/velonetics/velonetics-otel/http/server"
	"github.com/velonetics/velonetics-otel/state"
	luraconfig "github.com/velonetics/lura/v2/config"
	"github.com/velonetics/lura/v2/logging"
	luragin "github.com/velonetics/lura/v2/router/gin"
)

func GlobalRunServer(_ logging.Logger, next luragin.RunServerFunc) luragin.RunServerFunc {
	otelCfg := state.GlobalConfig()
	if otelCfg == nil {
		return next
	}

	return func(ctx context.Context, cfg luraconfig.ServiceConfig, h http.Handler) error {
		var trustedProxies []string
		if v, ok := cfg.ExtraConfig[luragin.Namespace].(map[string]interface{}); ok {
			if tpxs, ok := v["trusted_proxies"].([]string); ok {
				trustedProxies = tpxs
			}
		}
		wrappedH := kotelhttpserver.NewTrackingHandlerWithTrustedProxies(h, trustedProxies)
		return next(ctx, cfg, wrappedH)
	}
}
