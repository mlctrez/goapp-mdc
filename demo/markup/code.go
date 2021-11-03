package markup
type CodeDetails struct {
	Name string
	Code string
}
var Code = []CodeDetails{
    CodeDetails{Name:"index.go",Code:`<pre><code class="language-go">package demo

import (
	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
)

type Index struct {
	app.Compo
}

func (i *Index) Render() app.UI {
	return PageBody(app.Div().Text(&quot;&quot;))
}
</code></pre>
`},
    CodeDetails{Name:"banner.go",Code:`<pre><code class="language-go">package demo

import (
	&quot;fmt&quot;

	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/banner&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/base&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/button&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/checkbox&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/layout&quot;
)

type BannerDemo struct {
	app.Compo
	base.JsUtil
	floating *banner.Banner
	fixed    *banner.Banner
	message  *Message
}

type Message struct {
	app.Compo
	Text string
}

func (c *Message) Render() app.UI {
	return app.Code().Text(c.Text)
}

func (c *BannerDemo) Render() app.UI {

	if c.floating == nil {
		c.floating = &amp;banner.Banner{
			Id: &quot;normalBanner&quot;, Text: &quot;This is the banner text for a normal banner&quot;,
			Buttons: []app.UI{
				&amp;button.Button{Id: c.UUID(), Label: &quot;Primary&quot;, Banner: true, BannerAction: &quot;primary&quot;},
				&amp;button.Button{Id: c.UUID(), Label: &quot;Secondary&quot;, Banner: true, BannerAction: &quot;secondary&quot;},
			},
		}
		c.fixed = &amp;banner.Banner{
			Id: &quot;fixedBanner&quot;, Text: &quot;This is the banner text for a fixed banner&quot;, Fixed: true,
			Buttons: []app.UI{
				&amp;button.Button{Id: c.UUID(), Label: &quot;Primary&quot;, Banner: true, BannerAction: &quot;primary&quot;},
				&amp;button.Button{Id: c.UUID(), Label: &quot;Secondary&quot;, Banner: true, BannerAction: &quot;secondary&quot;},
			},
		}
		c.message = &amp;Message{Text: &quot;banner events will appear here&quot;}
	}
	openFloating := &amp;button.Button{Id: c.UUID(), Label: &quot;floating&quot;, Callback: func(button app.HTMLButton) {
		button.OnClick(func(ctx app.Context, e app.Event) {
			ctx.NewActionWithValue(string(banner.Open), c.floating)
		})
	}}
	openFixed := &amp;button.Button{Id: c.UUID(), Label: &quot;fixed&quot;, Callback: func(button app.HTMLButton) {
		button.OnClick(func(ctx app.Context, e app.Event) {
			ctx.NewActionWithValue(string(banner.Open), c.fixed)
		})
	}}
	centered := &amp;checkbox.Checkbox{Id: c.UUID(), Label: &quot;centered&quot;, Callback: func(input app.HTMLInput) {
		input.OnClick(func(ctx app.Context, e app.Event) {
			centeredValue := ctx.JSSrc().Get(&quot;checked&quot;).Bool()
			c.floating.Centered = centeredValue
			c.floating.Update()
			c.fixed.Centered = centeredValue
			c.fixed.Update()
		})
	}}

	body := app.Div().Body(
		c.floating, c.fixed,
		layout.Grid().Body(layout.Inner().Body(
			layout.Cell().Body(openFloating, openFixed, centered),
			layout.CellModified(&quot;middle&quot;, 12).Body(c.message),
		)))
	return PageBody(body)

}

func (c *BannerDemo) OnMount(ctx app.Context) {
	// handle all banner events
	for _, n := range []banner.EventType{banner.Opening, banner.Opened, banner.Closing, banner.Closed} {
		ctx.Handle(string(n), c.actionHandler)
	}
}

func (c *BannerDemo) actionHandler(ctx app.Context, action app.Action) {
	if !(action.Value == c.fixed || action.Value == c.floating) {
		return
	}
	if b, ok := action.Value.(*banner.Banner); ok {
		c.message.Text = fmt.Sprintf(&quot;message from banner %q: Event=%25s Tags=%v&quot;, b.Id, action.Name, action.Tags)
		c.message.Update()
	}
}
</code></pre>
`},
    CodeDetails{Name:"button.go",Code:`<pre><code class="language-go">package demo

import (
	&quot;time&quot;

	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/button&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/checkbox&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/layout&quot;
)

type ButtonDemo struct {
	app.Compo
	button *button.Button
}

func (d *ButtonDemo) Render() app.UI {
	if d.button == nil {
		d.button = &amp;button.Button{Id: &quot;subjectButton&quot;, Label: &quot;a button&quot;}
	}
	handleCheckboxChange := func(before func(checkVal bool)) func(input app.HTMLInput) {
		return func(input app.HTMLInput) {
			input.OnChange(func(ctx app.Context, e app.Event) {
				before(ctx.JSSrc().Get(&quot;checked&quot;).Bool())
				d.button.Update()
				// attempt to re-attach ripple, trying with a delay
				ctx.After(500*time.Millisecond, func(context app.Context) {
					d.button.OnMount(context)
				})
			})
		}
	}

	body := layout.Grid().Body(layout.Inner().Body(
		layout.CellModified(&quot;middle&quot;, 12).Body(d.button),
		layout.Cell().Body(
			&amp;checkbox.Checkbox{Id: &quot;toggleIcon&quot;, Label: &quot;has icon&quot;,
				Callback: handleCheckboxChange(func(checkVal bool) {
					if checkVal {
						d.button.Icon = &quot;bookmark&quot;
					} else {
						d.button.Icon = &quot;&quot;
					}
				})},
			&amp;checkbox.Checkbox{Id: &quot;toggleTrailing&quot;, Label: &quot;trailing icon&quot;,
				Callback: handleCheckboxChange(func(checkVal bool) { d.button.TrailingIcon = checkVal })},
			&amp;checkbox.Checkbox{Id: &quot;toggleOutline&quot;, Label: &quot;outlined&quot;,
				Callback: handleCheckboxChange(func(checkVal bool) { d.button.Outlined = checkVal })},
			&amp;checkbox.Checkbox{Id: &quot;toggleRaised&quot;, Label: &quot;raised&quot;,
				Callback: handleCheckboxChange(func(checkVal bool) { d.button.Raised = checkVal })},
			&amp;checkbox.Checkbox{Id: &quot;toggleUnelevated&quot;, Label: &quot;unelevated&quot;,
				Callback: handleCheckboxChange(func(checkVal bool) { d.button.Unelevated = checkVal })}),
	))
	return PageBody(body)
}
</code></pre>
`},
    CodeDetails{Name:"card.go",Code:`<pre><code class="language-go">package demo

import (
	&quot;fmt&quot;

	&quot;github.com/google/uuid&quot;
	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/button&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/card&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/layout&quot;
)

type CardDemo struct {
	app.Compo
}

func (d CardDemo) Render() app.UI {

	buttonCallback := func(text string) func(button app.HTMLButton) {
		return func(button app.HTMLButton) {
			button.OnClick(func(ctx app.Context, e app.Event) {
				fmt.Println(&quot;you clicked on button&quot; + text)
			})
		}
	}

	body := FlexGrid(
		layout.Cell().Body(
			&amp;card.Card{Id: uuid.New().String(), Padding: 16,
				PrimaryAction: []app.UI{app.Div().Text(&quot;Primary action card no outline&quot;)}},
		),
		layout.Cell().Body(
			&amp;card.Card{Id: uuid.New().String(), Width: 200, Height: 200, Outlined: true,
				PrimaryAction: []app.UI{app.Div().Text(&quot;Primary action card 200x200px with outline&quot;)}},
		),
		layout.Cell().Body(
			&amp;card.Card{Id: uuid.New().String(), Outlined: true, Padding: 16,
				PrimaryAction: []app.UI{app.Div().Text(&quot;Primary action card card with buttons&quot;)},
				ActionButtons: []app.UI{
					&amp;button.Button{Id: uuid.New().String(), CardAction: true,
						Label: &quot;Button One&quot;, Callback: buttonCallback(&quot;one&quot;)},
					&amp;button.Button{Id: uuid.New().String(), CardAction: true,
						Label: &quot;Button Two&quot;, Callback: buttonCallback(&quot;two&quot;)},
				},
			},
		),
		layout.Cell().Body(GopherCard(&quot;Media&quot;)),
		layout.Cell().Body(GopherCard(&quot;&quot;)),
		gopherAttribution(),
	)

	return PageBody(body)
}

func gopherAttribution() app.HTMLDiv {
	return layout.CellModified(&quot;bottom&quot;, 12).Body(
		app.Text(&quot;Gopher images courtesy of &quot;),
		app.A().Href(&quot;https://github.com/golang-samples/gopher-vector&quot;).Text(&quot;gopher-vector&quot;),
		app.Br(),
		app.Text(&quot;Licensed under the Creative Commons 3.0 Attributions license.&quot;))
}

func GopherCard(title string) app.UI {
	return &amp;card.Card{Id: uuid.New().String(), Width: 202, Height: 259,
		PrimaryAction: []app.UI{&amp;card.Media{
			Width: 202, Height: 259, Image: &quot;/web/gopher-front.png&quot;, Title: title}}}
}
</code></pre>
`},
    CodeDetails{Name:"checkbox.go",Code:`<pre><code class="language-go">package demo

import (
	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/checkbox&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/layout&quot;
)

type CheckboxDemo struct {
	app.Compo
	checkboxOne *checkbox.Checkbox
}

func (d *CheckboxDemo) onChange(changed func(bool)) func(checkbox app.HTMLInput) {
	return func(checkbox app.HTMLInput) {
		checkbox.OnChange(func(ctx app.Context, e app.Event) {
			changed(ctx.JSSrc().Get(&quot;checked&quot;).Bool())
			d.checkboxOne.Update()
		})
	}
}

func (d *CheckboxDemo) Render() app.UI {

	if d.checkboxOne == nil {
		d.checkboxOne = &amp;checkbox.Checkbox{Id: &quot;checkboxOne&quot;, Label: &quot;A Checkbox&quot;,
			Callback: func(input app.HTMLInput) {
				input.OnChange(func(ctx app.Context, e app.Event) {
					d.checkboxOne.Checked = ctx.JSSrc().Get(&quot;checked&quot;).Bool()
					if d.checkboxOne.Indeterminate {
						d.checkboxOne.Indeterminate = false
					}
					d.Update()
				})
			}}
	}

	body := layout.Grid().Body(layout.Inner().Body(
		layout.CellModified(&quot;middle&quot;, 12).
			Text(&quot;Demonstration of interacting with checkbox state from other checkboxes.&quot;),
		layout.Cell().Body(d.checkboxOne),
		layout.Cell().Body(
			&amp;checkbox.Checkbox{Id: &quot;checked&quot;, Label: &quot;checked&quot;, Checked: d.checkboxOne.Checked,
				Callback: d.onChange(func(b bool) { d.checkboxOne.Checked = b })},
			&amp;checkbox.Checkbox{Id: &quot;indeterminate&quot;, Label: &quot;indeterminate&quot;, Checked: d.checkboxOne.Indeterminate,
				Callback: d.onChange(func(b bool) { d.checkboxOne.Indeterminate = b })},
			&amp;checkbox.Checkbox{Id: &quot;disabled&quot;, Label: &quot;disabled&quot;, Checked: d.checkboxOne.Disabled,
				Callback: d.onChange(func(b bool) { d.checkboxOne.Disabled = b })},
		),
	))

	return PageBody(body)

}
</code></pre>
`},
    CodeDetails{Name:"dialog.go",Code:`<pre><code class="language-go">package demo

import (
	&quot;fmt&quot;

	&quot;github.com/mlctrez/goapp-mdc/pkg/button&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/dialog&quot;

	&quot;github.com/google/uuid&quot;
	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
)

type DialogDemo struct {
	app.Compo
}

func (d *DialogDemo) Render() app.UI {
	diagId := uuid.New().String()

	diag := &amp;dialog.Dialog{Id: diagId}
	diag.Title = []app.UI{app.Div().Text(&quot;Dialog Title&quot;)}
	diag.Content = []app.UI{app.Div().Text(&quot;This is the content section of the dialog. There is quite &quot; +
		&quot;a bit of text here to demonstrate how the dialog renders with this amount of text.&quot;)}

	diag.Buttons = []app.UI{
		&amp;button.Button{Id: diagId + &quot;-cancel&quot;, Dialog: true, DialogAction: &quot;cancel&quot;, Label: &quot;cancel&quot;,
			Callback: func(button app.HTMLButton) {
				button.OnClick(func(ctx app.Context, e app.Event) {
					fmt.Println(&quot;you clicked on the cancel button&quot;)
				})
			}},
		&amp;button.Button{Id: diagId + &quot;-dismiss&quot;, Dialog: true, DialogAction: &quot;dismiss&quot;, Label: &quot;dismiss&quot;},
	}

	openDialog := &amp;button.Button{Id: &quot;openDialogButton&quot;, Label: &quot;open dialog&quot;}
	openDialog.Callback = func(b app.HTMLButton) {
		b.OnClick(func(ctx app.Context, e app.Event) {
			diag.Open()
		})
	}

	return PageBody(openDialog, diag)
}
</code></pre>
`},
    CodeDetails{Name:"drawer.go",Code:`<pre><code class="language-go">package demo

import (
	&quot;log&quot;

	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/base&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/drawer&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/icon&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/list&quot;
)

type DrawerDemo struct {
	app.Compo
	base.JsUtil
}

func (d *DrawerDemo) Render() app.UI {

	navItems := list.Items{
		&amp;list.Item{Text: &quot;Inbox&quot;, Graphic: icon.MIInbox},
		&amp;list.Item{Text: &quot;Outgoing&quot;, Graphic: icon.MISend},
		&amp;list.Item{Type: list.ItemTypeDivider},
		&amp;list.Item{Text: &quot;Drafts&quot;, Graphic: icon.MIDrafts},
		&amp;list.Item{Text: &quot;Settings&quot;, Graphic: icon.MISettings},
		&amp;list.Item{Text: &quot;Ramen&quot;, Graphic: icon.MIRamenDining},
	}
	navItems.Select(0)

	body := &amp;drawer.Drawer{Id: d.UUID(), Type: drawer.Standard,
		List: &amp;list.List{Id: &quot;navigation&quot;, Type: list.Navigation, Items: navItems.UIList()}}
	return PageBody(body)
}

func (d *DrawerDemo) OnMount(ctx app.Context) {
	ctx.Handle(string(list.Select), d.eventHandler)
}

func (d *DrawerDemo) eventHandler(ctx app.Context, action app.Action) {
	log.Println(&quot;you clicked on item&quot;, action.Tags.Get(&quot;index&quot;))
}
</code></pre>
`},
    CodeDetails{Name:"fab.go",Code:`<pre><code class="language-go">package demo

import (
	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/fab&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/layout&quot;
)

type FabDemo struct {
	app.Compo
	activeIndex int
}

func (d *FabDemo) Render() app.UI {

	return PageBody(FlexGrid(
		layout.Grid().Body(
			&amp;fab.Fab{Id: id(), Icon: &quot;favorite&quot;},
			app.Text(&quot;regular&quot;),
		),
		layout.Grid().Body(
			&amp;fab.Fab{Id: id(), Icon: &quot;favorite&quot;, Mini: true},
			app.Text(&quot;mini&quot;),
		),
		layout.Grid().Body(
			&amp;fab.Fab{Id: id(), Icon: &quot;favorite&quot;, Extended: true, Label: &quot;Favorite&quot;},
			app.Text(&quot;extended&quot;),
		),
	))

}
</code></pre>
`},
    CodeDetails{Name:"form.go",Code:`<pre><code class="language-go">package demo

import (
	&quot;github.com/google/uuid&quot;
	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/helperline&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/layout&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/textarea&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/textfield&quot;
)

type FormDemo struct {
	app.Compo
}

func id() string {
	return uuid.New().String()
}

func textAreaExample() []app.UI {
	idOne := id()
	taOne := textarea.New(idOne).Size(8, 40).Outlined(true).
		Label(&quot;outlined text area&quot;).MaxLength(240)
	helpOne := helperline.New(idOne, &quot;textarea help text&quot;, &quot;0 / 240&quot;)

	return []app.UI{app.Div().Style(&quot;display&quot;, &quot;inline-block&quot;).Body(taOne, helpOne)}
}

func (e *FormDemo) Render() app.UI {

	//body := app.Div().Style(&quot;display&quot;, &quot;block&quot;).Body(
	//	app.Div().Text(&quot;test1&quot;),
	//	app.Div().Text(&quot;test2&quot;),
	//)

	body := layout.Grid().Body(layout.Inner().Style(&quot;display&quot;, &quot;flex&quot;).Body(
		layout.Cell().Body(layout.Inner().Style(&quot;display&quot;, &quot;flex&quot;).Body(
			layout.CellWide().Body(app.H4().Text(&quot;Text Area&quot;)),
			layout.Cell().Body(&amp;textfield.TextField{Id: id(), Label: &quot;normal&quot;}),
			layout.Cell().Body(&amp;textfield.TextField{Id: id(), Label: &quot;required&quot;, Required: true}),
			layout.Cell().Body(&amp;textfield.TextField{Id: id(), Label: &quot;outlined&quot;, Outlined: true}),
			layout.Cell().Body(&amp;textfield.TextField{Id: id(), Label: &quot;outlined required&quot;,
				Outlined: true, Required: true}),
			layout.Cell().Body(&amp;textfield.TextField{Id: id(), Placeholder: &quot;placeholder&quot;}),
		)),
		layout.Cell().Body(layout.Inner().Style(&quot;display&quot;, &quot;flex&quot;).Body(
			layout.CellWide().Body(app.H4().Text(&quot;Text Field&quot;)),
			layout.Cell().Body(textAreaExample()...),
		)),
	))

	_ = body
	return PageBody(body)
}
</code></pre>
`},
    CodeDetails{Name:"icon.go",Code:`<pre><code class="language-go">package demo

import (
	&quot;fmt&quot;
	&quot;sort&quot;

	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/base&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/icon&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/layout&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/list&quot;
)

type IconDemo struct {
	app.Compo
	base.JsUtil
	counterOne    *base.Counter
	toggleOne     *icon.Button
	toggleTwo     *icon.Button
	iconGroupList *list.List
}

func iconGroupNamesSorted() (result []string) {
	for s := range icon.AllGroupFunctions() {
		result = append(result, s)
	}
	sort.Strings(result)
	return
}

func (d *IconDemo) Render() app.UI {
	if d.counterOne == nil {
		d.counterOne = &amp;base.Counter{Label: &quot;bookmark&quot;}
		d.toggleOne = &amp;icon.Button{Id: d.UUID(), Icon: icon.MIFavorite,
			IconOff: icon.MIFavoriteBorder, AriaOn: &quot;remove from favorites&quot;, AriaOff: &quot;add to favorites&quot;}
		d.toggleTwo = &amp;icon.Button{Id: d.UUID(), Icon: icon.MIFavorite,
			IconOff: icon.MIFavoriteBorder, AriaOn: &quot;remove from favorites&quot;, AriaOff: &quot;add to favorites&quot;}
		d.toggleOne.ButtonToggleChange = func(isOn bool) { d.toggleTwo.SetState(isOn) }
		d.toggleTwo.ButtonToggleChange = func(isOn bool) { d.toggleOne.SetState(isOn) }

		d.iconGroupList = &amp;list.List{Type: list.SingleSelection, Id: &quot;iconGroupList&quot;}
		groups := list.Items{}
		for _, g := range iconGroupNamesSorted() {
			groups = append(groups, &amp;list.Item{Text: g})
		}
		d.iconGroupList.Items = groups.UIList()

	}

	body := layout.Grid().Body(
		layout.Inner().Style(&quot;display&quot;, &quot;flex&quot;).Body(
			layout.Cell().Body(
				&amp;icon.Button{Id: d.UUID(), Icon: icon.MIBookmark,
					AriaOff: &quot;bookmark this&quot;, Callback: d.IconButtonClicked}, d.counterOne),
			layout.Cell().Body(d.toggleOne, d.toggleTwo),
		),
		layout.Inner().Style(&quot;display&quot;, &quot;flex&quot;).Body(
			layout.CellWide().Body(app.Text(&quot;Material Icon Groups&quot;)),
			layout.Cell().Body(d.iconGroupList),
		),
	)

	return PageBody(body)
}

func (d *IconDemo) IconButtonClicked(button app.HTMLButton) {
	button.OnClick(func(ctx app.Context, e app.Event) {
		fmt.Println(&quot;you clicked bookmark&quot;)
		d.counterOne.Count += 1
		d.counterOne.Update()
	})
}
</code></pre>
`},
    CodeDetails{Name:"list.go",Code:`<pre><code class="language-go">package demo

import (
	&quot;fmt&quot;

	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/base&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/layout&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/list&quot;
)

type ListDemo struct {
	app.Compo
	base.JsUtil
}

func (d *ListDemo) Render() app.UI {

	regularList := list.Items{&amp;list.Item{Text: &quot;item one&quot;}, &amp;list.Item{Text: &quot;item two&quot;},
		&amp;list.Item{Text: &quot;item three&quot;}}.Select(-1)
	twoLineList := list.Items{
		&amp;list.Item{Text: &quot;item one&quot;, Secondary: &quot;item one subtext&quot;},
		&amp;list.Item{Text: &quot;item two&quot;, Secondary: &quot;item two subtext&quot;},
		&amp;list.Item{Text: &quot;item three&quot;, Secondary: &quot;item three subtext&quot;}}.Select(-1)

	groupedListOne := list.Items{&amp;list.Item{Text: &quot;group 1-1&quot;}, &amp;list.Item{Text: &quot;group 1-2&quot;}}.Select(0)
	groupedListTwo := list.Items{&amp;list.Item{Text: &quot;group 2-1&quot;}, &amp;list.Item{Text: &quot;group 2-2&quot;}}.Select(1)

	singleSelectionList := list.Items{&amp;list.Item{Text: &quot;item one&quot;}, &amp;list.Item{Text: &quot;item two&quot;},
		&amp;list.Item{Text: &quot;item three&quot;}, &amp;list.Item{Text: &quot;item four&quot;}}.Select(2)

	dividedList := list.Items{
		&amp;list.Item{Text: &quot;item one&quot;}, &amp;list.Item{Text: &quot;item two before divider&quot;},
		&amp;list.Item{Type: list.ItemTypeDivider},
		&amp;list.Item{Text: &quot;item three after divider&quot;}, &amp;list.Item{Text: &quot;item four&quot;}}
	dividedList.Select(0)

	checkboxGroupList := make(list.Items, 4)
	for i := range checkboxGroupList {
		checkboxGroupList[i] = &amp;list.Item{Type: list.ItemTypeCheckbox, Text: fmt.Sprintf(&quot;checkbox %d&quot;, i)}
	}
	checkboxGroupList.Select(-1)

	body := FlexGrid(
		layout.Cell().Body(
			app.P().Text(&quot;regular&quot;), &amp;list.List{Id: &quot;regularList&quot;, Items: regularList.UIList()}),
		layout.Cell().Body(
			app.P().Text(&quot;two line&quot;), &amp;list.List{Id: &quot;twoLineList&quot;, TwoLine: true, Items: twoLineList.UIList()}),
		layout.Cell().Body(
			app.P().Text(&quot;grouped&quot;),
			&amp;list.Group{Items: []*list.GroupItem{
				{SubHeader: &quot;group 1&quot;, List: &amp;list.List{Id: &quot;groupedList1&quot;, Items: groupedListOne.UIList()}},
				{SubHeader: &quot;group 2&quot;, List: &amp;list.List{Id: &quot;groupedList2&quot;, Items: groupedListTwo.UIList()}},
			}},
		),
		layout.Cell().Body(app.P().Text(&quot;divided&quot;), &amp;list.List{Id: &quot;dividedList&quot;, Items: dividedList.UIList()}),
		layout.Cell().Body(
			app.P().Text(&quot;single select&quot;),
			&amp;list.List{Id: &quot;singleSelectionList&quot;, Type: list.SingleSelection, Items: singleSelectionList.UIList()},
		),
		layout.Cell().Body(
			app.P().Text(&quot;checkbox group&quot;),
			&amp;list.List{Id: &quot;checkboxGroupList&quot;, Type: list.CheckBox, Items: checkboxGroupList.UIList()},
		),
	)

	return PageBody(body)

}
</code></pre>
`},
    CodeDetails{Name:"tab.go",Code:`<pre><code class="language-go">package demo

import (
	&quot;fmt&quot;

	&quot;github.com/google/uuid&quot;
	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/tab&quot;
)

type TabDemo struct {
	app.Compo
}

func (d *TabDemo) Render() app.UI {

	id := uuid.New().String()

	tab := tab.NewBar(id, []*tab.Tab{
		tab.NewTab(&quot;Tab One&quot;, 0).Active(),
		tab.NewTab(&quot;Tab Two&quot;, 1).Icon(&quot;api&quot;),
		tab.NewTab(&quot;Tab Three&quot;, 2).Icon(&quot;favorite&quot;),
	})
	tab.ActivateCallback(func(index int) {
		fmt.Println(&quot;you clicked on tab index&quot;, index)
	})

	return PageBody(tab)
}
</code></pre>
`},
    CodeDetails{Name:"code.go",Code:`<pre><code class="language-go">package demo

import (
	&quot;fmt&quot;
	&quot;strconv&quot;

	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/demo/markup&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/base&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/drawer&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/list&quot;
)

type CodeDemo struct {
	app.Compo
	base.JsUtil
	Content string
	Active  int
	list    *list.List
}

func (d *CodeDemo) OnNav(ctx app.Context) {
	d.SetActive(urlFragmentToInt(ctx))
	d.list.Select(d.Active)
	d.Update()
	ctx.Defer(prismHiglightAll)
}

func (d *CodeDemo) SetActive(index int) {
	if index &lt; 0 || index &gt; len(markup.Code)-1 {
		d.Active = 0
	} else {
		d.Active = index
	}
}

func urlFragmentToInt(ctx app.Context) (result int) {
	if idx, err := strconv.Atoi(ctx.Page().URL().Fragment); err == nil {
		result = idx
	}
	return
}

func (d *CodeDemo) Render() app.UI {
	if d.list == nil {
		items := list.Items{}
		for i, c := range markup.Code {
			items = append(items, &amp;list.Item{Text: c.Name,
				Type: list.ItemTypeAnchor, Href: fmt.Sprintf(&quot;/code#%d&quot;, i)})
		}
		d.list = &amp;list.List{Id: &quot;codeNav&quot;, Type: list.Navigation, Items: items.UIList()}
	}

	body := &amp;drawer.Drawer{Id: &quot;codeNavigation&quot;, Type: drawer.Standard, List: d.list}
	d.Content = markup.Code[d.Active].Code
	return PageBody(body, app.Raw(d.Content))
}

func prismHiglightAll(ctx app.Context) {
	prism := app.Window().Get(&quot;Prism&quot;)
	if prism.Truthy() {
		prism.Call(&quot;highlightAll&quot;)
	}
}

func (d *CodeDemo) OnMount(ctx app.Context) {
	ctx.Defer(prismHiglightAll)
}

func (d *CodeDemo) eventHandler(ctx app.Context, action app.Action) {
	if selectedIndex, err := strconv.Atoi(action.Tags.Get(&quot;index&quot;)); err != nil {
		return
	} else {
		ctx.Navigate(fmt.Sprintf(&quot;/code#%d&quot;, selectedIndex))
	}
}
</code></pre>
`},
    CodeDetails{Name:"routes.go",Code:`<pre><code class="language-go">package demo

import (
	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/icon&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/list&quot;
)

func addRoute(nav *list.Item, compo app.Composer) {
	nav.Type = list.ItemTypeAnchor
	NavigationItems = append(NavigationItems, nav)
	app.Route(nav.Href, compo)
}

func Routes() {
	addRoute(&amp;list.Item{Text: &quot;Home&quot;, Graphic: icon.MIHome, Href: &quot;/&quot;}, &amp;Index{})
	addRoute(&amp;list.Item{Text: &quot;Banner&quot;, Graphic: icon.MIVoicemail, Href: &quot;/banner&quot;}, &amp;BannerDemo{})
	addRoute(&amp;list.Item{Text: &quot;Button&quot;, Graphic: icon.MISmartButton, Href: &quot;/button&quot;}, &amp;ButtonDemo{})
	addRoute(&amp;list.Item{Text: &quot;Card&quot;, Graphic: icon.MICreditCard, Href: &quot;/card&quot;}, &amp;CardDemo{})
	addRoute(&amp;list.Item{Text: &quot;Checkbox&quot;, Graphic: icon.MICheckBox, Href: &quot;/checkbox&quot;}, &amp;CheckboxDemo{})
	addRoute(&amp;list.Item{Text: &quot;Dialog&quot;, Graphic: icon.MISpeaker, Href: &quot;/dialog&quot;}, &amp;DialogDemo{})
	addRoute(&amp;list.Item{Text: &quot;Drawer&quot;, Graphic: icon.MIDashboard, Href: &quot;/drawer&quot;}, &amp;DrawerDemo{})
	addRoute(&amp;list.Item{Text: &quot;Fab&quot;, Graphic: icon.MIFavorite, Href: &quot;/fab&quot;}, &amp;FabDemo{})
	addRoute(&amp;list.Item{Text: &quot;Form&quot;, Graphic: icon.MIInput, Href: &quot;/form&quot;}, &amp;FormDemo{})
	addRoute(&amp;list.Item{Text: &quot;Icon&quot;, Graphic: icon.MIIcecream, Href: &quot;/icon&quot;}, &amp;IconDemo{})
	addRoute(&amp;list.Item{Text: &quot;List&quot;, Graphic: icon.MIList, Href: &quot;/list&quot;}, &amp;ListDemo{})
	addRoute(&amp;list.Item{Text: &quot;Tab&quot;, Graphic: icon.MITab, Href: &quot;/tab&quot;}, &amp;TabDemo{})
	NavigationItems = append(NavigationItems, &amp;list.Item{Type: list.ItemTypeDivider})
	addRoute(&amp;list.Item{Text: &quot;Code&quot;, Graphic: icon.MICode, Href: &quot;/code&quot;}, &amp;CodeDemo{})
}
</code></pre>
`},
    CodeDetails{Name:"appupdate.go",Code:`<pre><code class="language-go">package demo

import (
	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/banner&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/base&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/button&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/icon&quot;
)

// AppUpdateBanner demonstrates how to wrap banner.Banner to handle go-app OnAppUpdate.
type AppUpdateBanner struct {
	app.Compo
	base.JsUtil
	bnr *banner.Banner
}

func (d *AppUpdateBanner) Render() app.UI {
	if d.bnr == nil {
		d.bnr = &amp;banner.Banner{
			Id: &quot;appUpdateBanner&quot;, Fixed: true, Centered: true,
			Text: &quot;A new version is available, would you like to install?&quot;,
		}
		d.bnr.Buttons = d.bannerButtons()
	}
	return d.bnr
}

func (d *AppUpdateBanner) bannerButtons() []app.UI {
	primary := &amp;button.Button{Id: &quot;updateBannerYes&quot;, Label: &quot;yes&quot;,
		Icon: string(icon.MIUpdate), Banner: true, BannerAction: &quot;primary&quot;}
	secondary := &amp;button.Button{Id: &quot;updateBannerNo&quot;, Label: &quot;later&quot;,
		Icon: string(icon.MIWatchLater), Banner: true, BannerAction: &quot;secondary&quot;}
	return []app.UI{primary, secondary}
}

func (d *AppUpdateBanner) onBannerClose(ctx app.Context, reason string) {
	switch reason {
	case &quot;primary&quot;: // Yes button
		ctx.Reload()
	case &quot;secondary&quot;: // Later button
		// This could SetState for a future time to ask
	}
}

func (d *AppUpdateBanner) OnMount(ctx app.Context) {
	d.bnr.ActionClose(ctx, d.onBannerClose)
}

func (d *AppUpdateBanner) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		d.bnr.ActionOpen(ctx)
	}
}
</code></pre>
`},
    CodeDetails{Name:"navigation.go",Code:`<pre><code class="language-go">package demo

import (
	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/base&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/drawer&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/list&quot;
)

// TODO: make this NavigationItems immutable

var NavigationItems list.Items

type Navigation struct {
	app.Compo
	base.JsUtil
	items list.Items
	list  *list.List
}

func (n *Navigation) Render() app.UI {
	return &amp;drawer.Drawer{Type: drawer.Standard, Id: &quot;navigationDrawer&quot;, List: n.list}
}

func (n *Navigation) OnMount(ctx app.Context) {
	if n.items == nil {
		n.items = NavigationItems
		n.list = &amp;list.List{Type: list.Navigation, Id: &quot;navigationList&quot;, Items: n.items.UIList()}
	}
	n.items.SelectHref(ctx.Page().URL().Path)
}
</code></pre>
`},
    CodeDetails{Name:"page.go",Code:`<pre><code class="language-go">package demo

import (
	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/layout&quot;
)

// PageBody applies the navigation, update banner, and demo page layout to the provided pageContent.
func PageBody(pageContent ...app.UI) app.UI {

	content := []app.UI{&amp;Navigation{}}
	content = append(content, pageContent...)

	return app.Div().Body(
		&amp;AppUpdateBanner{},

		app.Div().Style(&quot;display&quot;, &quot;flex&quot;).Body(content...),
	)
}

func FlexGrid(cells ...app.UI) app.UI {
	return layout.Grid().Body(
		layout.Inner().Style(&quot;display&quot;, &quot;flex&quot;).Body(cells...),
	)
}
</code></pre>
`},
    CodeDetails{Name:"handler.go",Code:`<pre><code class="language-go">package demo

import &quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;

func BuildHandler() *app.Handler {
	return &amp;app.Handler{
		Author:          &quot;mlctrez&quot;,
		Description:     &quot;Material Design Components for go-app&quot;,
		Icon:            app.Icon{Default: &quot;/web/logo-192.png&quot;, Large: &quot;/web/logo-512.png&quot;},
		Name:            &quot;MDC for go-app&quot;,
		BackgroundColor: &quot;#111&quot;,
		Scripts: []string{
			&quot;https://cdnjs.cloudflare.com/ajax/libs/material-components-web/13.0.0/material-components-web.min.js&quot;,
			&quot;https://cdnjs.cloudflare.com/ajax/libs/prism/1.25.0/prism.min.js&quot;,
			&quot;https://cdnjs.cloudflare.com/ajax/libs/prism/1.25.0/components/prism-go.min.js&quot;,
			&quot;/web/app.js&quot;,
		},
		Env: map[string]string{
			&quot;RECAPTCHA_SITE_KEY&quot;: &quot;6Ldt8sgcAAAAACwJjJMaRH3b31xDXBB6IYvBpLmc&quot;,
		},
		ShortName: &quot;goapp-mdc&quot;,
		Styles: []string{
			&quot;https://fonts.googleapis.com/icon?family=Material+Icons&quot;,
			&quot;https://fonts.googleapis.com/css2?family=Roboto&amp;display=swap&quot;,
			&quot;https://cdnjs.cloudflare.com/ajax/libs/material-components-web/13.0.0/material-components-web.min.css&quot;,
			&quot;https://cdnjs.cloudflare.com/ajax/libs/prism-themes/1.9.0/prism-material-light.min.css&quot;,
			&quot;/web/style.css&quot;,
		},
		Title: &quot;Material Design Components for go-app&quot;,
	}
}
</code></pre>
`},
    CodeDetails{Name:"wasm_server.go",Code:`<pre><code class="language-go">//go:build wasm

package main

func httpServer() {}
</code></pre>
`},
    CodeDetails{Name:"main.go",Code:`<pre><code class="language-go">package main

import (
	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/demo&quot;
)

func main() {
	demo.Routes()
	app.RunWhenOnBrowser()
	httpServer()
}
</code></pre>
`},
    CodeDetails{Name:"server.go",Code:`<pre><code class="language-go">//go:build !wasm

package main

import (
	&quot;flag&quot;
	&quot;fmt&quot;
	&quot;log&quot;
	&quot;net/http&quot;

	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/demo&quot;
)

// GitCommit is set using : go build -ldflags &quot;-X main.GitCommit=$(GIT_COMMIT)&quot;
var GitCommit string

func httpServer() {
	handler := setupVersion(demo.BuildHandler())
	if err := http.ListenAndServe(&quot;:8000&quot;, handler); err != nil {
		log.Println(err)
	}
}

func setupVersion(handler *app.Handler) *app.Handler {
	flag.Parse()
	switch flag.Arg(0) {
	case &quot;dev&quot;:
		fmt.Println(&quot;using dynamic version&quot;)
	default:
		handler.Version = GitCommit
	}
	return handler
}
</code></pre>
`},
}
