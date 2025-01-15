package formodel

import (
	"errors"
	"fmt"
	"testing"
)

/*
状态模式（State Pattern）是一种行为型设计模式，它适用于一个对象在在不同的状态下有不同的行为时，
	比如说电灯的开、关、闪烁是不停的状态，状态不同时，对应的行为也不同，
	在没有状态模式的情况下，为了添加新的状态或修改现有的状态，往往需要修改已有的代码，这违背了开闭原则，
	而且如果对象的状态切换逻辑和各个状态的行为都在同一个类中实现，就可能导致该类的职责过重，不符合单一职责原则。

	而状态模式将每个状态的行为封装在一个具体状态类中，使得每个状态类相对独立，
	并将对象在不同状态下的行为进行委托，从而使得对象的状态可以在运行时动态改变，每个状态的实现也不会影响其他状态。

状态模式包括以下几个重要角色：
1. State（状态）： 定义一个接口，用于封装与Context的一个特定状态相关的行为。
2. ConcreteState（具体状态）： 负责处理Context在状态改变时的行为, 每一个具体状态子类实现一个与Context的一个状态相关的行为。
3. Context（上下文）: 维护一个具体状态子类的实例，这个实例定义当前的状态。

使用场景
1. 状态模式将每个状态的实现都封装在一个类中，每个状态类的实现相对独立，使得添加新状态或修改现有状态变得更加容易，避免了使用大量的条件语句来控制对象的行为。
	但是如果状态过多，会导致类的数量增加，可能会使得代码结构复杂。

2. 总的来说，状态模式适用于有限状态机的场景，其中对象的行为在运行时可以根据内部状态的改变而改变.

3. 在游戏开发中，Unity3D 的 Animator 控制器就是一个状态机。它允许开发人员定义不同的状态（动画状态），并通过状态转换来实现角色的动画控制和行为切换。
*/

type TVSState interface {
	On() error
	Off() error
	Mute() error
}

type OnState struct {
}

func (o *OnState) On() error {
	return errors.New("The TV is already on")
}

func (o *OnState) Off() error {
	fmt.Println("Turning off the TV")
	return nil
}

func (o *OnState) Mute() error {
	fmt.Println("Muting the TV")
	return nil
}

type OffState struct {
}

func (o *OffState) On() error {
	fmt.Println("Turning on the TV")
	return nil
}

func (o *OffState) Off() error {
	return errors.New("The TV is already off")
}

func (o *OffState) Mute() error {
	return errors.New("The TV is off, cannot mute")
}

type MutedState struct {
}

func (m *MutedState) On() error {
	fmt.Println("Turning on the TV")
	return nil
}

func (m *MutedState) Off() error {
	fmt.Println("Turning off the TV")
	return nil
}

func (m *MutedState) Mute() error {
	return errors.New("The TV is already muted")
}

type STV struct {
	state TVSState
}

func (t *STV) setState(state TVSState) {
	t.state = state
}

func (t *STV) On() error {
	if err := t.state.On(); err != nil {
		return err
	}

	t.setState(new(OnState))

	return nil
}

func (t *STV) Off() error {
	if err := t.state.Off(); err != nil {
		return err
	}

	t.setState(new(OffState))

	return nil
}

func (t *STV) Mute() error {
	if err := t.state.Mute(); err != nil {
		return err
	}

	t.setState(new(MutedState))

	return nil
}

func TestState(t *testing.T) {
	tv := STV{}
	tv.setState(new(OffState))

	if err := tv.On(); err != nil {
		fmt.Printf("Error: %v", err)
	}

	if err := tv.Mute(); err != nil {
		fmt.Printf("Error: %v", err)
	}

	if err := tv.Off(); err != nil {
		fmt.Printf("Error: %v", err)
	}

	if err := tv.Mute(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}
