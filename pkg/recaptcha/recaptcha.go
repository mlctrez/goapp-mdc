package recaptcha

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	fetch "github.com/mlctrez/wasm-fetch"
)

const recaptchaKey = "6Ldt8sgcAAAAACwJjJMaRH3b31xDXBB6IYvBpLmc"

type recaptcha struct {
	app.Compo
	id     string
	action string
	event  string
}

func New(id, action string) *recaptcha {
	return &recaptcha{
		id:     id,
		action: action,
		event:  id + "_" + action + "_event",
	}
}

func (d *recaptcha) Render() app.UI {
	script := app.Raw(fmt.Sprintf(`<script>
grecaptcha.enterprise.ready(function () {
    grecaptcha.enterprise.execute('%s', {action: '%s'}).then(function (token) {
		window.dispatchEvent(new CustomEvent('%s', { detail: {token : token} }));		
    });
});
</script>`, recaptchaKey, d.action, d.event))
	return script
}

func (d *recaptcha) OnMount(ctx app.Context) {
	app.Window().AddEventListener(d.event, func(ctx app.Context, e app.Event) {
		token := e.Get("detail").Get("token").String()
		marshal, err := json.Marshal(map[string]interface{}{"token": token})
		if err != nil {
			panic(err)
		}

		assessmentEndpoint := *(ctx.Page().URL())
		assessmentEndpoint.Path = "/assessmentendpoint"

		opts := &fetch.Opts{Method: fetch.MethodPost, Body: bytes.NewReader(marshal), Headers: map[string]string{"Content-Type": "application/json"}}

		response, err := fetch.Fetch(assessmentEndpoint.String(), opts)
		if err != nil {
			panic(err)
		}
		if response.Status == 200 {
			fmt.Println(string(response.Body))
		}

	})
}
