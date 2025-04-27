package log

import (
	"bufio"
	"bytes"
	"context"
	"net/http"
	"time"
)

// elasticsearch日志

type EsLogger struct {
	RequestEnabled  bool `desc:"是否开启请求日志"`
	ResponseEnabled bool `desc:"是否开启响应日志"`
}

func (logger EsLogger) LogRoundTrip(request *http.Request, response *http.Response, err error, t time.Time, ti time.Duration) error {
	if logger.RequestBodyEnabled() && request != nil && request.Body != nil && request.Body != http.NoBody {
		var buf bytes.Buffer
		if request.GetBody != nil {
			b, _ := request.GetBody()
			_, err = buf.ReadFrom(b)
		} else {
			_, err = buf.ReadFrom(request.Body)
		}
		if err != nil {
			Errorf(context.TODO(), "ES日志error: %+v", err)
			return nil
		}
		scanner := bufio.NewScanner(&buf)
		for scanner.Scan() {
			s := scanner.Text()
			if s != "" {
				Infof(context.TODO(), "ES日志: 请求参数: %s; 时间: %+v; 消耗：%+v", s, t, ti)
			}
		}
	}

	return nil
}

func (logger EsLogger) RequestBodyEnabled() bool {
	return logger.RequestEnabled
}

func (logger EsLogger) ResponseBodyEnabled() bool {
	return logger.ResponseEnabled
}
