package markup
var Code = map[string]string{
"demo/appupdate.go":`<pre><code class="language-go">package demo

import (
	&quot;log&quot;

	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/banner&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/base&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/button&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/icon&quot;
)

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
	log.Println(&quot;banner was closed with reason&quot;, reason)
	switch reason {
	case &quot;primary&quot;:
		ctx.Reload()
	case &quot;secondary&quot;:
		// set a timer to open again in X hours?
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
`,
"demo/demo_banner.go":`<pre><code class="language-go">package demo

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
`,
"demo/demo_button.go":`<pre><code class="language-go">package demo

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
`,
"demo/demo_card.go":`<pre><code class="language-go">package demo

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

	body := layout.Grid().Body(
		layout.Inner().Body(
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
						&amp;button.Button{Id: uuid.New().String(), CardAction: true, Label: &quot;Button One&quot;, Callback: buttonCallback(&quot;one&quot;)},
						&amp;button.Button{Id: uuid.New().String(), CardAction: true, Label: &quot;Button Two&quot;, Callback: buttonCallback(&quot;two&quot;)},
					},
				},
			),
			layout.Cell().Body(
				// an example media card with title
				&amp;card.Card{Id: uuid.New().String(), Height: 100, Width: 100,
					PrimaryAction: []app.UI{
						&amp;card.Media{Width: 100, Height: 100, Image: &quot;/web/logo-192.png&quot;, Title: &quot;Media&quot;},
					}},
			),
			layout.Cell().Body(
				// an example media card with title
				&amp;card.Card{Id: uuid.New().String(), Height: 100, Width: 100,
					PrimaryAction: []app.UI{
						&amp;card.Media{Width: 100, Height: 100, Image: &quot;/web/logo-192.png&quot;},
					}},
			),
		),
	)
	return PageBody(body)
}
</code></pre>
`,
"demo/demo_checkbox.go":`<pre><code class="language-go">package demo

import (
	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/checkbox&quot;
)

type CheckboxDemo struct {
	app.Compo
	checkboxOne *checkbox.Checkbox
}

func (d *CheckboxDemo) Render() app.UI {

	if d.checkboxOne == nil {
		d.checkboxOne = &amp;checkbox.Checkbox{Id: &quot;checkboxOne&quot;, Label: &quot;Checkbox Label&quot;,
			Callback: func(input app.HTMLInput) {
				input.OnChange(func(ctx app.Context, e app.Event) {
					d.checkboxOne.Checked = ctx.JSSrc().Get(&quot;checked&quot;).Bool()
					// changing the state of the checkbox should clear this flag
					if d.checkboxOne.Indeterminate {
						d.checkboxOne.Indeterminate = false
					}
					d.Update()
				})
			}}
	}

	body := app.Div().Body(
		d.checkboxOne,
		app.Hr(),
		&amp;checkbox.Checkbox{Id: &quot;checked&quot;, Label: &quot;checked&quot;, Checked: d.checkboxOne.Checked,
			Callback: func(checkbox app.HTMLInput) {
				checkbox.OnChange(func(ctx app.Context, e app.Event) {
					d.checkboxOne.Checked = ctx.JSSrc().Get(&quot;checked&quot;).Bool()
					d.checkboxOne.Update()
				})
			}},
		&amp;checkbox.Checkbox{Id: &quot;indeterminate&quot;, Label: &quot;indeterminate&quot;, Checked: d.checkboxOne.Indeterminate,
			Callback: func(checkbox app.HTMLInput) {
				checkbox.OnChange(func(ctx app.Context, e app.Event) {
					d.checkboxOne.Indeterminate = ctx.JSSrc().Get(&quot;checked&quot;).Bool()
					d.checkboxOne.Update()
				})
			}},
		&amp;checkbox.Checkbox{Id: &quot;disabled&quot;, Label: &quot;disabled&quot;, Checked: d.checkboxOne.Disabled,
			Callback: func(checkbox app.HTMLInput) {
				checkbox.OnChange(func(ctx app.Context, e app.Event) {
					d.checkboxOne.Disabled = ctx.JSSrc().Get(&quot;checked&quot;).Bool()
					d.checkboxOne.Update()
				})
			}},
	)
	return PageBody(body)

}
</code></pre>
`,
"demo/demo_code.go":`<pre><code class="language-go">package demo

import (
	&quot;fmt&quot;
	&quot;sort&quot;
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
	list *list.List
}

func sortedNames() []string {
	var sortedNames []string
	for n := range markup.Code {
		sortedNames = append(sortedNames, n)
	}
	// reverse sort here
	sort.Slice(sortedNames, func(i, j int) bool { return sortedNames[i] &gt; sortedNames[j] })
	return sortedNames
}

func (d *CodeDemo) OnNav(ctx app.Context) {
	d.LogWithP(d, &quot;OnNav&quot;)
	url := ctx.Page().URL()
	idx, err := strconv.Atoi(url.Fragment)
	if err != nil {
		idx = 0
	}
	d.Active = idx
	d.Content = markup.Code[sortedNames()[d.Active]]
	d.list.Select(d.Active)
	d.Update()
	ctx.Defer(func(context app.Context) {
		app.Window().Get(&quot;Prism&quot;).Call(&quot;highlightAll&quot;)
	})
}

func (d *CodeDemo) Render() app.UI {
	d.LogWithPf(d, &quot;Render active=%d&quot;, d.Active)
	if d.list == nil {
		d.LogWithP(d, &quot;new list&quot;)
		items := list.Items{}
		for i, name := range sortedNames() {
			items = append(items, &amp;list.Item{Text: name,
				Type: list.ItemTypeAnchor, Href: fmt.Sprintf(&quot;/code#%d&quot;, i)})
		}
		d.list = &amp;list.List{Id: &quot;codeNav&quot;, Type: list.Navigation, Items: items.UIList()}
	}

	body := &amp;drawer.Drawer{Id: &quot;codeNavigation&quot;, Type: drawer.Standard, List: d.list}

	if d.Content == &quot;&quot; {
		d.Content = markup.Code[sortedNames()[0]]
	}

	return PageBody(body, app.Raw(d.Content))
}

func (d *CodeDemo) OnMount(ctx app.Context) {
	d.LogWithPf(d, &quot;OnMount active=%d&quot;, d.Active)
	ctx.Handle(string(list.Select), d.eventHandler)
	app.Window().Get(&quot;Prism&quot;).Call(&quot;highlightAll&quot;)
}

func (d *CodeDemo) eventHandler(ctx app.Context, action app.Action) {
	if selectedIndex, err := strconv.Atoi(action.Tags.Get(&quot;index&quot;)); err != nil {
		return
	} else {
		ctx.Navigate(fmt.Sprintf(&quot;/code#%d&quot;, selectedIndex))
	}
}
</code></pre>
`,
"demo/demo_dialog.go":`<pre><code class="language-go">package demo

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
`,
"demo/demo_drawer.go":`<pre><code class="language-go">package demo

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
`,
"demo/demo_fab.go":`<pre><code class="language-go">package demo

import (
	&quot;sort&quot;

	&quot;github.com/mlctrez/goapp-mdc/pkg/fab&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/icon&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/tab&quot;

	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
)

type FabDemo struct {
	app.Compo
	activeIndex int
}

func (d *FabDemo) Render() app.UI {
	var tabs []*tab.Tab

	sortedGroupNames := sortedGroupNames()
	for index, groupName := range sortedGroupNames {
		newTab := tab.NewTab(groupName, index)
		if d.activeIndex == index {
			newTab.Active()
		}
		tabs = append(tabs, newTab)
	}

	bar := tab.NewBar(&quot;fabDemoTabBar&quot;, tabs)
	bar.ActivateCallback(func(index int) {
		d.activeIndex = index
		d.Update()
	})

	var content []app.UI
	content = append(content, bar)

	groupName := sortedGroupNames[d.activeIndex]
	iconNames := icon.AllGroupFunctions()[groupName]()
	for _, n := range iconNames {
		content = append(content, &amp;fab.Fab{Id: &quot;fab_&quot; + string(n), Icon: n})
	}

	return PageBody(app.Div().Body(content...))
}

func sortedGroupNames() []string {
	var groupNamesSorted []string
	for groupName := range icon.AllGroupFunctions() {
		groupNamesSorted = append(groupNamesSorted, groupName)
	}
	sort.Strings(groupNamesSorted)
	return groupNamesSorted
}
</code></pre>
`,
"demo/demo_form.go":`<pre><code class="language-go">package demo

import (
	&quot;github.com/mlctrez/goapp-mdc/pkg/fab&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/helperline&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/layout&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/textarea&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/textfield&quot;

	&quot;github.com/google/uuid&quot;
	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
)

type FormDemo struct {
	app.Compo
}

func id() string {
	return uuid.New().String()
}

func cell(name string, contents ...app.UI) app.UI {
	return layout.CellModified(&quot;&quot;, 12).Body(app.H4().Text(name),
		layout.Inner().Body(
			func() []app.UI {
				var result []app.UI
				for _, content := range contents {
					result = append(result, layout.Cell().Body(content))
				}
				return result
			}()...,
		),
	)
}

func fabExamples() []app.UI {
	return []app.UI{
		&amp;fab.Fab{Id: id(), Icon: &quot;favorite&quot;},
		&amp;fab.Fab{Id: id(), Icon: &quot;favorite&quot;, Mini: true},
		&amp;fab.Fab{Id: id(), Icon: &quot;favorite&quot;, Extended: true, Label: &quot;Favorite&quot;},
	}
}

func textFieldExamples() []app.UI {
	return []app.UI{
		&amp;textfield.TextField{Id: id(), Label: &quot;normal&quot;},
		&amp;textfield.TextField{Id: id(), Label: &quot;required&quot;, Required: true},
		&amp;textfield.TextField{Id: id(), Label: &quot;outlined&quot;, Outlined: true},
		&amp;textfield.TextField{Id: id(), Label: &quot;outlined required&quot;, Outlined: true, Required: true},
		&amp;textfield.TextField{Id: id(), Placeholder: &quot;placeholder&quot;},
	}
}

func textAreaExample() []app.UI {
	idOne := id()
	taOne := textarea.New(idOne).Size(8, 40).Outlined(true).Label(&quot;outlined text area&quot;).MaxLength(240)
	helpOne := helperline.New(idOne, &quot;textarea help text&quot;, &quot;0 / 240&quot;)

	return []app.UI{app.Div().Style(&quot;display&quot;, &quot;inline-block&quot;).Body(taOne, helpOne)}

}

func (e *FormDemo) Render() app.UI {

	body := layout.Grid().Body(layout.Inner().Body(
		cell(&quot;Fab&quot;, fabExamples()...),
		cell(&quot;Text Field&quot;, textFieldExamples()...),
		cell(&quot;Text Area&quot;, textAreaExample()...),
	))
	return PageBody(body)

}
</code></pre>
`,
"demo/demo_icon.go":`<pre><code class="language-go">package demo

import (
	&quot;fmt&quot;

	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/base&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/icon&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/layout&quot;
)

type IconDemo struct {
	app.Compo
	base.JsUtil
	counterOne *base.Counter
	toggleOne  *icon.Button
	toggleTwo  *icon.Button
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
	}

	body := layout.Grid().Body(
		layout.Inner().Body(
			layout.Cell().Body(
				&amp;icon.Button{Id: d.UUID(), Icon: icon.MIBookmark,
					AriaOff: &quot;bookmark this&quot;, Callback: d.IconButtonClicked}, d.counterOne),
			layout.Cell().Body(d.toggleOne, d.toggleTwo),
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
`,
"demo/demo_list.go":`<pre><code class="language-go">package demo

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

	regularList := list.Items{&amp;list.Item{Text: &quot;item one&quot;}, &amp;list.Item{Text: &quot;item two&quot;}, &amp;list.Item{Text: &quot;item three&quot;}}.Select(-1)
	twoLineList := list.Items{
		&amp;list.Item{Text: &quot;item one&quot;, Secondary: &quot;item one subtext&quot;},
		&amp;list.Item{Text: &quot;item two&quot;, Secondary: &quot;item two subtext&quot;},
		&amp;list.Item{Text: &quot;item three&quot;, Secondary: &quot;item three subtext&quot;}}.Select(-1)

	groupedListOne := list.Items{&amp;list.Item{Text: &quot;group 1-1&quot;}, &amp;list.Item{Text: &quot;group 2-2&quot;}}.Select(0)
	groupedListTwo := list.Items{&amp;list.Item{Text: &quot;group 2-1&quot;}, &amp;list.Item{Text: &quot;group 2-2&quot;}}.Select(1)

	singleSelectionList := list.Items{&amp;list.Item{Text: &quot;item one&quot;}, &amp;list.Item{Text: &quot;item two&quot;},
		&amp;list.Item{Text: &quot;item three&quot;}, &amp;list.Item{Text: &quot;item four&quot;}}.Select(2)

	dividedList := list.Items{&amp;list.Item{Text: &quot;item one&quot;}, &amp;list.Item{Text: &quot;item two before divider&quot;},
		&amp;list.Item{Type: list.ItemTypeDivider}, &amp;list.Item{Text: &quot;item three after divider&quot;}, &amp;list.Item{Text: &quot;item four&quot;}}
	dividedList.Select(0)

	// TODO: build out radio button component first
	//radioGroupList := Items{&amp;Item{Text: &quot;item one&quot;}, &amp;Item{Text: &quot;item two&quot;},
	//	&amp;Item{Text: &quot;item three&quot;}, &amp;Item{Text: &quot;item four&quot;}}.Select(2)

	checkboxGroupList := make(list.Items, 4)
	for i := range checkboxGroupList {
		checkboxGroupList[i] = &amp;list.Item{Type: list.ItemTypeCheckbox, Text: fmt.Sprintf(&quot;checkbox %d&quot;, i)}
	}
	checkboxGroupList.Select(-1)

	body := layout.Grid().Body(layout.Inner().Body(
		layout.Cell().Body(
			app.P().Text(&quot;regular list&quot;), &amp;list.List{Id: &quot;regularList&quot;, Items: regularList.UIList()}),
		layout.Cell().Body(
			app.P().Text(&quot;two line list&quot;), &amp;list.List{Id: &quot;twoLineList&quot;, TwoLine: true, Items: twoLineList.UIList()}),
		layout.Cell().Body(
			app.P().Text(&quot;grouped List&quot;),
			&amp;list.Group{Items: []*list.GroupItem{
				{SubHeader: &quot;group 1&quot;, List: &amp;list.List{Id: &quot;groupedList1&quot;, Items: groupedListOne.UIList()}},
				{SubHeader: &quot;group 2&quot;, List: &amp;list.List{Id: &quot;groupedList2&quot;, Items: groupedListTwo.UIList()}},
			}},
		),
		layout.Cell().Body(app.P().Text(&quot;divided List&quot;), &amp;list.List{Id: &quot;dividedList&quot;, Items: dividedList.UIList()}),
		layout.Cell().Body(
			app.P().Text(&quot;single selection&quot;),
			&amp;list.List{Id: &quot;singleSelectionList&quot;, Type: list.SingleSelection, Items: singleSelectionList.UIList()},
		),
		//layout.Cell().Body(
		//	app.P().Text(&quot;radio group&quot;),
		//	&amp;List{Id: &quot;radioGroupList&quot;, Type: RadioGroup, Items: radioGroupList.UIList()},
		//),
		layout.Cell().Body(
			app.P().Text(&quot;checkbox group&quot;),
			&amp;list.List{Id: &quot;checkboxGroupList&quot;, Type: list.CheckBox, Items: checkboxGroupList.UIList()},
		),
	))
	return PageBody(body)

}
</code></pre>
`,
"demo/demo_tab.go":`<pre><code class="language-go">package demo

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
`,
"demo/handler.go":`<pre><code class="language-go">package demo

import &quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;

func BuildHandler() *app.Handler {
	return &amp;app.Handler{
		Author:          &quot;mlctrez&quot;,
		Description:     &quot;Material Design Components for go-app&quot;,
		Icon:            app.Icon{Default: &quot;/web/logo-192.png&quot;, Large: &quot;/web/logo-512.png&quot;},
		Name:            &quot;MDC for go-app&quot;,
		BackgroundColor: &quot;#111&quot;,
		Scripts: []string{
			//&quot;https://unpkg.com/material-components-web@latest/dist/material-components-web.min.js&quot;,
			&quot;/web/material-components-web.min.js&quot;,
			&quot;/web/prism.js&quot;,
		},
		ShortName: &quot;goapp-mdc&quot;,
		Styles: []string{
			&quot;https://fonts.googleapis.com/icon?family=Material+Icons&quot;,
			//&quot;https://unpkg.com/material-components-web@latest/dist/material-components-web.min.css&quot;,
			&quot;/web/material-components-web.min.css&quot;,
			&quot;/web/prism.css&quot;,
			&quot;/web/style.css&quot;,
		},
		Title: &quot;Material Design Components for go-app&quot;,
	}
}
</code></pre>
`,
"demo/index.go":`<pre><code class="language-go">package demo

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
`,
"demo/navigation.go":`<pre><code class="language-go">package demo

import (
	&quot;log&quot;

	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/base&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/drawer&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/list&quot;
)

var navigationItems list.Items

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
		n.items = navigationItems
		n.items.SelectHref(ctx.Page().URL().Path)
		n.list = &amp;list.List{Type: list.Navigation, Id: &quot;navigationList&quot;, Items: n.items.UIList()}
	}
	ctx.Handle(string(list.Select), n.eventHandler)
}

func (n *Navigation) eventHandler(ctx app.Context, action app.Action) {
	if action.Value != n.list {
		return
	}
	log.Println(&quot;you clicked on item&quot;, action.Tags.Get(&quot;index&quot;))
}
</code></pre>
`,
"demo/page.go":`<pre><code class="language-go">package demo

import (
	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
)

func PageBody(elems ...app.UI) app.UI {

	content := []app.UI{&amp;Navigation{}}
	content = append(content, elems...)

	return app.Div().Body(
		&amp;AppUpdateBanner{},

		app.Div().Style(&quot;display&quot;, &quot;flex&quot;).Body(content...),
	)
}
</code></pre>
`,
"demo/routes.go":`<pre><code class="language-go">package demo

import (
	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/icon&quot;
	&quot;github.com/mlctrez/goapp-mdc/pkg/list&quot;
)

func addRoute(nav *list.Item, compo app.Composer) {
	nav.Type = list.ItemTypeAnchor
	navigationItems = append(navigationItems, nav)
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
	navigationItems = append(navigationItems, &amp;list.Item{Type: list.ItemTypeDivider})
	addRoute(&amp;list.Item{Text: &quot;Code&quot;, Graphic: icon.MICode, Href: &quot;/code&quot;}, &amp;CodeDemo{})
}
</code></pre>
`,
"main.go":`<pre><code class="language-go">package main

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
`,
"server.go":`<pre><code class="language-go">//go:build !wasm

package main

import (
	&quot;flag&quot;
	&quot;fmt&quot;
	&quot;log&quot;
	&quot;net/http&quot;

	&quot;github.com/maxence-charriere/go-app/v9/pkg/app&quot;
	&quot;github.com/mlctrez/goapp-mdc/demo&quot;
)

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
`,
"wasm_server.go":`<pre><code class="language-go">//go:build wasm

package main

func httpServer() {}
</code></pre>
`,
}
