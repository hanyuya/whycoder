package formodel

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"

	log "github.com/sirupsen/logrus"
)

/*
MVC（Model-View-Controller）是一种经典的软件设计模式，它将应用程序分为三个部分：模型（Model）、视图（View）和控制器（Controller）。
	这种分层结构可以帮助程序员更好地组织代码，简化开发流程，实现可维护、可扩展的应用程序。

在 MVC 模式中，模型表示应用程序的数据和业务逻辑，视图负责呈现模型的数据，控制器接受用户输入并更新模型和视图。
	MVC 模式实现了关注点分离（Separation of Concerns），使得开发人员可以独立地开发和维护模型、视图和控制器。

这种按照 MVC 模式组织代码的方式可以更好地实现关注点分离，减少了各个部分的耦合度，有利于提高代码的可维护性和可重用性。
	在实际应用中，模型、视图和控制器可以更复杂，视图可以使用各种不同的模板引擎和技术，例如 React、Vue 或 Angular 等现代前端框架。
	控制器和模型也可以包含更多的功能，例如输入验证、缓存和数据库连接池等。

需要注意的是，在使用 MVC 模式时，应该尽量遵循模型、视图和控制器的职责，避免在某个部分中处理过多的业务逻辑。
	同时，应该根据应用程序的特点来选择合适的模板引擎、数据库连接池等工具，以实现最佳的性能和可维护性。
*/

// Model
type User struct {
	ID       int
	Name     string
	Age      int
	Email    string
	Password string
}

func (u *User) GetByName(name string) error {
	// 模拟从数据库中查询用户信息
	if name == "admin" {
		u.ID = 1
		u.Name = "admin"
		u.Age = 25
		u.Email = "admin@example.com"
		u.Password = "password"
		return nil
	}
	return errors.New("user not found")
}

// Controller
func UserController(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/user/"):]
	user := &User{}

	// 实例化模型并获取数据
	err := user.GetByName(name)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// 使用模型数据渲染视图
	t, err := template.ParseFiles("user.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, user)
}

func TestMVC(t *testing.T) {
	http.HandleFunc("/user/", UserController)
	req := httptest.NewRequest("GET", "/user/admin", nil)
	resp := httptest.NewRecorder()
	http.HandlerFunc(UserController).ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		log.Warnf("UserController error: %d", resp.Code)
	}
	respBody := resp.Body.String()
	fmt.Printf("respBody: %s\n", respBody)
}
