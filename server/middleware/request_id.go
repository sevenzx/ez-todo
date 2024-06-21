package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
	"github.com/hertz-contrib/requestid"
	"strings"
	"time"
)

func RequestId() app.HandlerFunc {
	return requestid.New(
		requestid.WithGenerator(func(ctx context.Context, c *app.RequestContext) string {
			s := uuid.New().String()
			s = strings.ReplaceAll(s, "-", "")
			s = strings.ToUpper(s)
			return fmt.Sprintf("%s%s", time.Now().Format("20060102150405"), s)
		}),
	)
}
