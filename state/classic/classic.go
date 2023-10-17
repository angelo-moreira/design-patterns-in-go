package state_classic

import "fmt"

type Switch struct {
	State State
}

func (s *Switch) On() {
	s.State.On(s)
}

func (s *Switch) Off() {
	s.State.Off(s)
}

func NewSwitch() *Switch {
	return &Switch{NewOffState()}
}

type State interface {
	On(sw *Switch)
	Off(sw *Switch)
}

type BaseState struct{}

func (b *BaseState) On(sw *Switch) {
	fmt.Println("the light is already on")
}

func (b *BaseState) Off(sw *Switch) {
	fmt.Println("the light is already off")
}

type OnState struct {
	BaseState
}

func (o *OnState) Off(sw *Switch) {
	fmt.Println("turning the light off")
	sw.State = NewOffState()
}

func NewOnState() *OnState {
	fmt.Println("light turned on")
	return &OnState{BaseState{}}
}

type OffState struct {
	BaseState
}

func (o *OffState) On(sw *Switch) {
	fmt.Println("turning the light on")
	sw.State = NewOnState()
}

func NewOffState() *OffState {
	fmt.Println("light turned off")
	return &OffState{BaseState{}}
}

func Run() {
	sw := NewSwitch()
	sw.On()
	sw.Off()
	sw.Off()
}
