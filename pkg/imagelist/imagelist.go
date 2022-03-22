package imagelist

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type ImageList struct {
	app.Compo
	Class   string
	Protection bool
	Masonry bool
	Items   Items
}

func (i *ImageList) Render() app.UI {
	ul := app.Ul().Class("mdc-image-list")
	ul.Class(i.Class)
	if i.Protection {
		ul.Class("mdc-image-list--with-text-protection")
	}
	if i.Masonry {
		ul.Class("mdc-image-list--masonry")
	}
	ul.Body(i.Items.UIList()...)
	return ul
}

type Items []*Item

func (i Items) UIList() (result []app.UI) {
	for _, item := range i {
		result = append(result, item)
	}
	return
}

type Item struct {
	app.Compo
	ImageContent app.UI
	Supporting   *SupportingLabel
}

func (i *Item) Render() app.UI {
	li := app.Li().Class("mdc-image-list__item").Body(
		i.ImageContent, i.Supporting,
	)
	return li
}

type SupportingLabel struct {
	app.Compo
	Text string
}

func (i *SupportingLabel) Render() app.UI {
	return app.Div().Class("mdc-image-list__supporting").Body(
		app.Span().Class("mdc-image-list__label").Text(i.Text),
	)
}

type ImageAspectContainer struct {
	app.Compo
	Image app.HTMLImg
}

func (i *ImageAspectContainer) Render() app.UI {
	container := app.Div().Class("mdc-image-list__image-aspect-container")
	container.Body(i.Image.Class("mdc-image-list__image"))
	return container
}
