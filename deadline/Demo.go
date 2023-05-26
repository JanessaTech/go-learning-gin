package deadline

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func setDeadlineMw() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("enter setDeadlineMw ...")
		ctx := c.Request.Context()
		deadline := time.Now().Add(5 * time.Second)
		newCtx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		c.Request = c.Request.WithContext(newCtx)

		c.Next()
		cancel()
		fmt.Println("leave setDeadlineMw ...")
	}
}

// http://127.0.0.1:8080/test
func Main() {
	r := gin.New()
	r.Use(setDeadlineMw(), gin.Recovery())
	r.GET("/test", testFunc)
	r.Run(":8080")
}

func doTask() string {
	time.Sleep(7 * time.Second)
	return "task is finished"
}

func testFunc(c *gin.Context) {
	fmt.Println(time.Now(), "Get started with task..")
	ctx := c.Request.Context()
	if deadlineTime, ok := ctx.Deadline(); ok {
		fmt.Println("deadlineTime = ", deadlineTime)
		doneChan := make(chan string)

		// start a go routine to execute the task
		go func() {
			res := doTask()
			doneChan <- res
		}()

		// waiting for result: either the task is done or task isn't finished due to deadline exceeded
		select {
		case <-ctx.Done():
			err := ctx.Err()
			fmt.Println("task isn't finished due to:", err)
		case res := <-doneChan:
			fmt.Println(time.Now(), res)
		}
	}
}
