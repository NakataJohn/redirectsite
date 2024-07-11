package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RedirectHandler 用于处理重定向逻辑
func RedirectHandler(c *gin.Context, redirectCount int) {
	if redirectCount > 0 {
		// 减少重定向次数，并设置重定向
		c.Redirect(http.StatusMovedPermanently, c.Request.URL.Path+"?redirect="+strconv.FormatInt(int64(redirectCount-1), 10))
	} else {
		// 最后一次重定向后，返回成功消息
		c.String(http.StatusOK, "Successfully!")
	}
}

// RedirectHandler 用于处理重定向逻辑
func RedirectHandlerW(c *gin.Context, redirectCount int) {
	if redirectCount > 0 {
		// 减少重定向次数，并设置重定向
		c.Redirect(http.StatusMovedPermanently, c.Request.URL.Path+"?redirect="+strconv.FormatInt(int64(redirectCount-1), 10))
	} else {
		// 最后一次重定向后，返回成功消息
		c.String(http.StatusForbidden, "Run along.")
	}
}

func main() {
	router := gin.Default()

	// 设置路由，这里我们假设查询参数redirect用于控制重定向次数
	router.GET("/right", func(c *gin.Context) {
		redirectCount := c.Query("redirect")
		if redirectCount == "" {
			// 如果没有提供redirect参数，默认从5开始
			redirectCount = "4"
		}

		// 将字符串转换为int
		count, err := strconv.Atoi(redirectCount)
		if err != nil {
			// 如果转换失败，返回错误
			c.String(http.StatusBadRequest, "无效的redirect参数")
			return
		}

		// 调用RedirectHandler来处理重定向逻辑
		RedirectHandler(c, count)
	})

	// 设置路由，这里我们假设查询参数redirect用于控制重定向次数
	router.GET("/wrong", func(c *gin.Context) {
		redirectCount := c.Query("redirect")
		if redirectCount == "" {
			// 如果没有提供redirect参数，默认从5开始
			redirectCount = "4"
		}

		// 将字符串转换为int
		count, err := strconv.Atoi(redirectCount)
		if err != nil {
			// 如果转换失败，返回错误
			c.String(http.StatusBadRequest, "无效的redirect参数")
			return
		}

		// 调用RedirectHandler来处理重定向逻辑
		RedirectHandlerW(c, count)
	})
	// 启动服务
	router.Run(":8081") // 监听并在 0.0.0.0:8080 上启动服务
}

// 注意：上面的代码示例中，为了简化，我省略了strconv包的导入，但在实际代码中你需要这样做：
// import "strconv"
// 并在需要的地方使用它来进行字符串到整数的转换。
