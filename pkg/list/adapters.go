package list

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type AdaptableType interface {
	Class(s ...string)
	ID(id string)
	Attr(key string, val string)
	Body(items ...app.UI)
	UI() app.UI
	Aria(s string, s2 string)
	DataSet(s string, s2 string)
	TabIndex(i int)
	Href(href string)
}

type htmlUl struct{ r app.HTMLUl }

func (a *htmlUl) Class(s ...string)           { a.r.Class(s...) }
func (a *htmlUl) ID(id string)                { a.r.ID(id) }
func (a *htmlUl) Attr(key string, val string) { a.r.Attr(key, val) }
func (a *htmlUl) Aria(s string, s2 string)    { a.r.Aria(s, s2) }
func (a *htmlUl) DataSet(s string, s2 string) { a.r.DataSet(s, s2) }
func (a *htmlUl) TabIndex(i int)              { a.r.TabIndex(i) }
func (a *htmlUl) Body(items ...app.UI)        { a.r.Body(items...) }
func (a *htmlUl) UI() app.UI                  { return a.r }
func (a *htmlUl) Href(s string)               {}

type htmlNav struct{ r app.HTMLNav }

func (n *htmlNav) Class(s ...string)           { n.r.Class(s...) }
func (n *htmlNav) ID(id string)                { n.r.ID(id) }
func (n *htmlNav) Attr(key string, val string) { n.r.Attr(key, val) }
func (n *htmlNav) Aria(s string, s2 string)    { n.r.Aria(s, s2) }
func (n *htmlNav) DataSet(s string, s2 string) { n.r.DataSet(s, s2) }
func (n *htmlNav) TabIndex(i int)              { n.r.TabIndex(i) }
func (n *htmlNav) Body(items ...app.UI)        { n.r.Body(items...) }
func (n *htmlNav) UI() app.UI                  { return n.r }
func (n *htmlNav) Href(s string)               {}

func (l *List) adapt() AdaptableType {
	switch l.Type {
	case Navigation:
		return &htmlNav{app.Nav()}
	default:
		return &htmlUl{app.Ul()}
	}
}

type htmlA struct{ Root app.HTMLA }

func (n *htmlA) Class(s ...string)           { n.Root.Class(s...) }
func (n *htmlA) ID(id string)                { n.Root.ID(id) }
func (n *htmlA) Attr(key string, val string) { n.Root.Attr(key, val) }
func (n *htmlA) Aria(s string, s2 string)    { n.Root.Aria(s, s2) }
func (n *htmlA) DataSet(s string, s2 string) { n.Root.DataSet(s, s2) }
func (n *htmlA) TabIndex(i int)              { n.Root.TabIndex(i) }
func (n *htmlA) Body(items ...app.UI)        { n.Root.Body(items...) }
func (n *htmlA) UI() app.UI                  { return n.Root }
func (n *htmlA) Href(s string)               { n.Root.Href(s) }

type htmlLi struct{ Root app.HTMLLi }

func (n *htmlLi) Class(s ...string)           { n.Root.Class(s...) }
func (n *htmlLi) ID(id string)                { n.Root.ID(id) }
func (n *htmlLi) Attr(key string, val string) { n.Root.Attr(key, val) }
func (n *htmlLi) Aria(s string, s2 string)    { n.Root.Aria(s, s2) }
func (n *htmlLi) DataSet(s string, s2 string) { n.Root.DataSet(s, s2) }
func (n *htmlLi) TabIndex(i int)              { n.Root.TabIndex(i) }
func (n *htmlLi) Body(items ...app.UI)        { n.Root.Body(items...) }
func (n *htmlLi) UI() app.UI                  { return n.Root }
func (n *htmlLi) Href(s string)               {}

func (i *Item) adapt() AdaptableType {
	switch i.Type {
	case ItemTypeAnchor:
		return &htmlA{app.A()}
	default:
		return &htmlLi{app.Li()}
	}
}
