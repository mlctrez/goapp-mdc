package base

import (
	"testing"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type ComponentTest struct {
	T          *testing.T
	Compo      app.UI
	Dispatcher app.ServerDispatcher
	path       []int
}

func (c *ComponentTest) Match(expected app.UI) {
	err := app.TestMatch(c.Compo, app.TestUIDescriptor{Path: c.path, Expected: expected})
	if err != nil {
		c.T.Helper()
		c.T.Errorf("path %d : %v", c.path, err)
	}
}

func (c *ComponentTest) Close() {
	c.Dispatcher.Close()
}

type Match interface {
	Match(expected app.UI)
}

func (c *ComponentTest) At(path ...int) Match {
	c.path = path
	return c
}
