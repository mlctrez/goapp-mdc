package base

import (
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// OnEvent logs a message when the event fires.
func OnEvent(msg string) func(ctx app.Context, e app.Event) {
	return func(ctx app.Context, e app.Event) {
		log.Println(msg)
	}
}
