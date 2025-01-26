package formodel

import (
	"fmt"
	"testing"
)

/*
策略模式
策略模式（Strategy Pattern）是一种行为型设计模式，它定义了一系列算法（这些算法完成的是相同的工作，只是实现不同），并将每个算法封装起来，使它们可以相互替换，而且算法的变化不会影响使用算法的客户。

举个例子:
	电商网站对于商品的折扣策略有不同的算法，比如新用户满减优惠，不同等级会员的打折情况不同，
		这种情况下会产生大量的if-else语句, 并且如果优惠政策修改时，还需要修改原来的代码，不符合开闭原则。
	这就可以将不同的优惠算法封装成独立的类来避免大量的条件语句，如果新增优惠算法，可以添加新的策略类来实现，
		客户端在运行时选择不同的具体策略，而不必修改客户端代码改变优惠策略。

策略模式包含下面几个结构：
	策略类Strategy: 定义所有支持的算法的公共接口。
	具体策略类ConcreteStrategy: 实现了策略接口，提供具体的算法实现。
	上下文类Context: 包含一个策略实例，并在需要时调用策略对象的方法。

使用场景
	当一个系统根据业务场景需要动态地在几种算法中选择一种时，可以使用策略模式。例如，根据用户的行为选择不同的计费策略。
	当代码中存在大量条件判断，条件判断的区别仅仅在于行为，也可以通过策略模式来消除这些条件语句。
	在已有的工具库中，Java 标准库中的 Comparator 接口就使用了策略模式，通过实现这个接口，可以创建不同的比较器（指定不同的排序策略）来满足不同的排序需求。
*/

// 策略类Strategy
type Strategy interface {
	Pay(amount float64) error
}

// 具体策略类ConcreteStrategy
type CreditCardStrategy struct {
	name       string
	cardNumber string
	cvv        string
	date       string
}

func (c *CreditCardStrategy) Pay(amount float64) error {
	// 具体的算法实现
	fmt.Printf("Paying %0.2f using credit card %s\n", amount, c.cardNumber)
	return nil
}

// 具体策略类ConcreteStrategy
type PayPalStrategy struct {
	email    string
	password string
}

func (p *PayPalStrategy) Pay(amount float64) error {
	// 具体的算法实现
	fmt.Printf("Paying %0.2f using PayPal\n", amount)
	return nil
}

// 上下文类Context
type PaymentContext struct {
	amount   float64
	strategy Strategy
}

func NewPaymentContext(amount float64, strategy Strategy) *PaymentContext {
	return &PaymentContext{
		amount:   amount,
		strategy: strategy,
	}
}

func (p *PaymentContext) Pay() error {
	return p.strategy.Pay(p.amount)
}

// 客户端代码
func TestStrategy(t *testing.T) {
	// 选择具体的策略
	creditCard := &CreditCardStrategy{
		name:       "John Doe",
		cardNumber: "1234567890123456",
		cvv:        "123",
		date:       "01/2022",
	}

	paypal := &PayPalStrategy{
		email:    "johndoe@gmail.com",
		password: "123456",
	}

	// 选择具体的策略
	paymentContext := NewPaymentContext(100.0, creditCard)
	paymentContext.Pay()

	paymentContext = NewPaymentContext(200.0, paypal)
	paymentContext.Pay()
}
