// Template generated by reactGen

package main

import (
	"crypto/sha256"
	"fmt"
	"net/url"
	"strings"

	"github.com/gopherjs/gopherjs/js"

	"honnef.co/go/js/dom"
	"myitcv.io/react"
)

type AppDef struct {
	react.ComponentDef
}

type AppState struct {
	handle string
	key    string
}

func App() *AppElem {
	return buildAppElem()
}

func (a AppDef) Render() react.Element {
	errStr := ""
	buttState := ""
	buttHref := ""
	if a.State().key == "" || a.State().handle == "" {
		buttState = " disabled"
	} else {
		buttHref, errStr = a.buildURL()
	}
	return react.Div(nil,
		react.Form(&react.FormProps{ClassName: ""},
			react.Div(
				&react.DivProps{ClassName: "form-group"},
				react.Label(&react.LabelProps{ClassName: "sr-only", For: "handle"}, react.S("Twitter handle")),
				react.Div(&react.DivProps{ClassName: "input-group"},
					react.Div(&react.DivProps{ClassName: "input-group-addon"},
						react.I(&react.IProps{ClassName: "fas fa-at"}),
					),
					react.Input(&react.InputProps{
						Type:        "text",
						ClassName:   "form-control",
						ID:          "handle",
						Placeholder: "Twitter handle",
						Value:       a.State().handle,
						OnChange:    handleChange{a},
					}),
				),
				react.Label(&react.LabelProps{ClassName: "sr-only", For: "key"}, react.S("Raffle key")),
				react.Div(&react.DivProps{ClassName: "input-group"},
					react.Div(&react.DivProps{ClassName: "input-group-addon"},
						react.I(&react.IProps{ClassName: "fas fa-key"}),
					),
					react.Input(&react.InputProps{
						Type:        "text",
						ClassName:   "form-control",
						ID:          "key",
						Placeholder: "Raffle key",
						Value:       a.State().key,
						OnChange:    keyChange{a},
					}),
				),
				react.A(&react.AProps{
					ClassName: "btn btn-primary" + buttState,
					Href:      buttHref,
					Role:      "button",
					Target:    "_blank",
				}, react.S("Prepare tweet!")),
			),
		),
		react.H5(&react.H5Props{ClassName: "text-danger", Style: &react.CSS{FontWeight: "bold"}},
			react.S(errStr),
		),
	)
}

type handleChange struct{ a AppDef }

func (i handleChange) OnChange(se *react.SyntheticEvent) {
	target := se.Target().(*dom.HTMLInputElement)
	ns := i.a.State()
	ns.handle = target.Value
	i.a.SetState(ns)
}

type keyChange struct{ a AppDef }

func (i keyChange) OnChange(se *react.SyntheticEvent) {
	target := se.Target().(*dom.HTMLInputElement)
	ns := i.a.State()
	ns.key = target.Value
	i.a.SetState(ns)
}

func (a AppDef) buildURL() (res string, errStr string) {
	var err error
	defer func() {
		if err != nil {
			errStr = fmt.Sprintf("failed to build Tweet URL: %v", err)
		}
	}()

	curr, err := url.Parse(js.Global.Get("window").Get("location").String())
	if err != nil {
		return
	}

	target, err := url.Parse("https://twitter.com/intent/tweet")
	if err != nil {
		return
	}

	ns := a.State()

	hash := sha256.New()
	fmt.Fprintf(hash, "Handle: %v\n", strings.ToLower(strings.TrimSpace(ns.handle)))
	fmt.Fprintf(hash, "Key: %v\n", strings.ToLower(strings.TrimSpace(ns.key)))

	q := target.Query()
	q.Set("text", fmt.Sprintf("%v %x", curr.Query().Get("greeting"), hash.Sum(nil)))
	q.Set("hashtags", curr.Query().Get("hashtags"))

	target.RawQuery = q.Encode()

	return target.String(), ""
}
