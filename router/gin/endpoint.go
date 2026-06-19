package gin

import (
	"github.com/gin-gonic/gin"
	luraconfig "github.com/pucora/lura/v2/config"
	"github.com/pucora/lura/v2/proxy"
	veloneticsgin "github.com/pucora/lura/v2/router/gin"
	"go.opentelemetry.io/otel/attribute"

	kotelconfig "github.com/pucora/velonetics-otel/config"
	kotelserver "github.com/pucora/velonetics-otel/http/server"
	otelstate "github.com/pucora/velonetics-otel/state"
)

// New wraps a handler factory adding some simple instrumentation to the generated handlers
func New(hf veloneticsgin.HandlerFactory) veloneticsgin.HandlerFactory {
	return func(cfg *luraconfig.EndpointConfig, p proxy.Proxy) gin.HandlerFunc {
		otelCfg := otelstate.GlobalConfig()
		if otelCfg == nil {
			return hf(cfg, p)
		}
		if otelCfg.SkipEndpoint(cfg.Endpoint) {
			return hf(cfg, p)
		}
		urlPattern := kotelconfig.NormalizeURLPattern(cfg.Endpoint)
		next := hf(cfg, p)
		var metricsAttrs []attribute.KeyValue
		var tracesAttrs []attribute.KeyValue

		cfgExtra, err := kotelconfig.LuraLayerExtraCfg(cfg.ExtraConfig)
		if err == nil && cfgExtra.Global != nil {
			for _, kv := range cfgExtra.Global.MetricsStaticAttributes {
				if len(kv.Key) > 0 && len(kv.Value) > 0 {
					metricsAttrs = append(metricsAttrs, attribute.String(kv.Key, kv.Value))
				}
			}

			for _, kv := range cfgExtra.Global.TracesStaticAttributes {
				if len(kv.Key) > 0 && len(kv.Value) > 0 {
					tracesAttrs = append(tracesAttrs, attribute.String(kv.Key, kv.Value))
				}
			}
		}

		return func(c *gin.Context) {
			// we set the matched route to a data struct stored in the
			// context by the outer http layer, so it can be reported
			// in metrics and traces.
			kotelserver.SetEndpointPattern(c.Request.Context(), urlPattern)
			kotelserver.SetStaticAttributtes(c.Request.Context(), metricsAttrs, tracesAttrs)
			next(c)
		}
	}
}
