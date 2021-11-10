package demo

import (
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/button"
	"github.com/mlctrez/goapp-mdc/pkg/snackbar"
)

type SnackBarDemo struct {
	app.Compo
	base.JsUtil
	snackBar *snackbar.SnackBar
}

func (s *SnackBarDemo) Render() app.UI {
	if s.snackBar == nil {
		s.snackBar = &snackbar.SnackBar{
			Leading:          false,
			Stacked:          false,
			LabelText:        "Snackbar Label",
			ActionButtonText: "ActionButton",
		}
	}

	return PageBody(
		s.snackBar,

		&button.Button{Label: "open", Outlined: true, Callback: s.buttonAction(snackbar.Open)},
		&button.Button{Label: "close", Outlined: true, Callback: s.buttonAction(snackbar.Close)},
		&button.Button{Label: "close custom", Outlined: true,
			Callback: s.buttonAction(snackbar.Close, app.T("reason", "custom"))},

		&button.Button{Label: "10sec", Outlined: true, Callback: func(button app.HTMLButton) {
			button.OnClick(func(ctx app.Context, e app.Event) {
				s.ConsoleLog(s.snackBar.TimeoutMs(10000))
			})
		}},
		&button.Button{Label: "4sec", Outlined: true, Callback: func(button app.HTMLButton) {
			button.OnClick(func(ctx app.Context, e app.Event) {
				s.ConsoleLog(s.snackBar.TimeoutMs(4000))
			})
		}},
		&button.Button{Label: "forever", Outlined: true, Callback: func(button app.HTMLButton) {
			button.OnClick(func(ctx app.Context, e app.Event) {
				s.ConsoleLog(s.snackBar.TimeoutMs(-1))
			})
		}},
		&button.Button{Label: "read", Outlined: true, Callback: func(button app.HTMLButton) {
			button.OnClick(func(ctx app.Context, e app.Event) {
				s.ConsoleLog(s.snackBar.TimeoutMs(0))
			})
		}},
	)
}

func (s *SnackBarDemo) buttonAction(evt snackbar.Event, tags ...app.Tagger) func(button app.HTMLButton) {
	return button.Click(func(ctx app.Context, e app.Event) {
		ctx.NewActionWithValue(string(evt), s.snackBar, tags...)
	})
}

var _ app.Mounter = (*SnackBarDemo)(nil)

func (s *SnackBarDemo) logAction(context app.Context, action app.Action) {
	if action.Value == s.snackBar {
		log.Println("snackbar says", action)
	}
}

func (s *SnackBarDemo) OnMount(ctx app.Context) {
	for _, action := range []snackbar.Event{snackbar.Opening, snackbar.Opened, snackbar.Closing, snackbar.Closed} {
		ctx.Handle(string(action), s.logAction)
	}
}
