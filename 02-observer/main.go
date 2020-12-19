package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	s1 := Sensor{Id: "s1"}
	s2 := Sensor{Id: "s2"}
	s3 := Sensor{Id: "s3"}
	st := Station{}
	st.Register(&s1)
	st.Register(&s2)
	st.Register(&s3)

	println("notify 1")
	st.Notify()
	println("notify 2")
	st.Notify()

	st.Remove(&s1)
	println("notify 3")
	st.Notify()

	return nil
}

type (
	Observer interface {
		ID() string // TODO: for comparing. Need Fix
		Update()
	}
	Observabler interface {
		Register(Observer)
		Remove(Observer)
		Notify()
	}
)

type Sensor struct {
	Id string
}

func (s *Sensor) ID() string {
	return s.Id
}

func (s *Sensor) Update() {
	println("update sensor " + s.ID())
}

type Station struct {
	observes []Observer
}

func (s *Station) Register(o Observer) {
	s.observes = append(s.observes, o)
}

func (s *Station) Remove(o Observer) {
	obs := s.observes[:0]
	for _, ob := range s.observes {
		if ob.ID() != o.ID() {
			obs = append(obs, ob)
		}
	}
	s.observes = obs
}

func (s *Station) Notify() {
	for _, o := range s.observes {
		o.Update()
	}
}
