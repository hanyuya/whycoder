package formodel

import (
	"net/http"
	"net/http/httptest"
	"testing"

	log "github.com/sirupsen/logrus"
)

/*
中间件模式（Middleware Pattern）是一种常见的软件设计模式，常用于处理系统中的请求和响应，尤其是在 Web 开发领域中。

在 Go 语言中，中间件模式通常用于处理 HTTP 请求和响应，可以帮助开发人员在请求处理过程中添加额外的逻辑。
	具体来说，中间件是一种函数，可以访问请求和响应，并在请求处理期间对它们进行修改。

总之，中间件模式是一种灵活且可扩展的模式，可以帮助开发人员轻松添加和组合逻辑，从而更好地处理请求和响应。

下面是一个简单的示例，其中我们创建了两个中间件函数，一个用于记录请求信息，另一个用于检查请求头中的 API 密钥。
*/

/*
type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

适配器：这个类型的主要目的是允许普通的函数作为 HTTP 处理器（Handler）使用。
	如果一个函数 f 具有与 HandlerFunc 相同的签名（即接受 ResponseWriter 和 *Request 作为参数），
	那么 HandlerFunc(f) 就是一个可以调用 f 的 Handler。

*/
// 定义一个记录请求信息的中间件
func loggingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// 记录请求信息
		log.Infof("Received request: %s %s", r.Method, r.URL.Path)
		// 调用下一个中间件或最终处理器
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// 定义一个检查 API 密钥的中间件
func apiKeyMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// 从请求头中获取 API 密钥
		apiKey := r.Header.Get("X-API-Key")
		// 检查 API 密钥是否有效
		if apiKey == "mysecretkey" {
			// 调用下一个中间件或最终处理器
			next.ServeHTTP(w, r)
		} else {
			// 密钥无效，返回 401 错误
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	}
	return http.HandlerFunc(fn)
}

// 定义一个最终处理器
func handler(w http.ResponseWriter, r *http.Request) {
	// 处理请求
	w.Write([]byte("Hello, world!"))
}

func TestMiddleware(t *testing.T) {
	// 创建一个处理器链，其中包含两个中间件和一个最终处理器
	handler := apiKeyMiddleware(loggingMiddleware(http.HandlerFunc(handler)))
	// 调用最终处理器
	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	req.Header.Set("X-API-Key", "mysecretkey")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	// 检查结果
	if w.Code != http.StatusOK {
		log.Warnf("w.Code = %d, want %d", w.Code, http.StatusOK)
	}
}
