package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
	"github.com/hertz-contrib/requestid"
	"time"
)

func RequestId() app.HandlerFunc {
	return requestid.New(
		requestid.WithGenerator(func(ctx context.Context, c *app.RequestContext) string {
			return fmt.Sprintf("%s-%v", time.Now().Format("20060102"), uuid.New().String())
		}),
	)
}
