package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type responseData struct {
	status int
	body   map[string]interface{}
}

func TimeoutMiddleware(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer func() {

			if ctx.Err() == context.DeadlineExceeded {
				c.Abort()
			}

			cancel()
		}()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func TimedHandler(duration time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {

		ctx := c.Request.Context()
		doneChan := make(chan responseData)

		//這邊就是之後移去別的地方當固定的值
		timeoutRes := responseData{}
		timeoutRes.status = http.StatusGatewayTimeout
		timeoutRes.body = gin.H{"message": "timeout"}

		go doSomeThing(duration, doneChan)

		select {
		case <-ctx.Done():
			c.JSON(timeoutRes.status, timeoutRes.body)
			return

		case res := <-doneChan:
			c.JSON(res.status, res.body)
		}
	}
}

func MyHandler(c *gin.Context) {

	ctx := c.Request.Context()
	var param QureyParam
	if c.Bind(&param) != nil {
		c.String(http.StatusBadRequest, "fail")
		return
	}
	doneChan := make(chan responseData)
	duration := time.Second * time.Duration(param.Second)

	timeoutRes := responseData{}
	timeoutRes.status = http.StatusGatewayTimeout
	timeoutRes.body = gin.H{"message": "timeout"}

	go doSomeThing(duration, doneChan)

	select {
	case <-ctx.Done():
		c.JSON(timeoutRes.status, timeoutRes.body)
		return

	case res := <-doneChan:
		c.JSON(res.status, res.body)
	}

}

func doSomeThing(duration time.Duration, doneChan chan responseData) {

	time.Sleep(duration)
	doneChan <- responseData{
		status: http.StatusOK,
		body:   gin.H{"message": "hello world"},
	}

}

type QureyParam struct {
	Second int
}
