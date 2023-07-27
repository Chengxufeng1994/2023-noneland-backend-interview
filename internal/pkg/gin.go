package pkg

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"noneland/backend/interview/internal/api"
)

type TokenBucket struct {
	capacity  int64
	rate      float64
	tokens    float64
	lastToken time.Time
	mtx       sync.Mutex
}

func (tb *TokenBucket) Allow() bool {
	tb.mtx.Lock()
	defer tb.mtx.Unlock()
	now := time.Now()

	tb.tokens = tb.tokens + tb.rate*now.Sub(tb.lastToken).Seconds()
	if tb.tokens > float64(tb.capacity) {
		tb.tokens = float64(tb.capacity)
	}

	if tb.tokens >= 1 {
		tb.tokens--
		tb.lastToken = now
		return true
	} else {
		return false
	}
}

func LimitHandler(tb *TokenBucket) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(&tb)
		if !tb.Allow() {
			c.String(http.StatusConflict, "Too many request")
			c.Abort()
			return
		}
		c.Next()
	}
}

// InitHttpHandler 為了分成測試用與正式用，所以把 gin 的初始化抽出來
func InitHttpHandler() (h http.Handler) {
	return h2c.NewHandler(setupGin(), &http2.Server{})
}

func setupGin() http.Handler {
	gin.SetMode(gin.DebugMode)

	r := gin.New()
	r.Use(api.ErrorHandler())

	tb := &TokenBucket{
		capacity:  61,
		rate:      20.333,
		tokens:    0,
		lastToken: time.Now(),
	}
	apiGroup := r.Group("/api")
	apiGroup.GET("hello", api.HelloHandler)

	apiGroup.GET("exchange/balance/:userId", LimitHandler(tb), api.GetBalance)
	apiGroup.GET("exchange/transfer/records/:userId", LimitHandler(tb), api.GetTxRecords)

	return r
}
