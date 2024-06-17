package timeseries

import (
	"net/http"

	"github.com/VictoriaMetrics/VictoriaMetrics/app/vminsert"
	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmselect"
	"github.com/pingcap/log"
	"go.uber.org/zap"
)

var _ http.HandlerFunc = InsertHandler
var _ http.HandlerFunc = SelectHandler

func InsertHandler(writer http.ResponseWriter, request *http.Request) {
	vminsert.RequestHandler(writer, request)
}

func SelectHandler(writer http.ResponseWriter, request *http.Request) {
	log.Info("select handler", zap.String("request", request.URL.String()), zap.String("form", request.Form.Encode()))
	vmselect.RequestHandler(writer, request)
}
