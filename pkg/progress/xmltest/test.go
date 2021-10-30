package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

type Svg struct {
	XMLName   struct{} `xml:"svg"`
	Namespace string   `xml:"namespace,attr"`
	Class     string   `xml:"class,attr"`
	ViewBox   string   `xml:"viewBox,attr"`
}

/*

<svg class="mdc-circular-progress__determinate-circle-graphic" viewBox="0 0 48 48" xmlns="http://www.w3.org/2000/svg">
<circle class="mdc-circular-progress__determinate-track" cx="24" cy="24" r="18.0" stroke-width="4.0"/>
<circle class="mdc-circular-progress__determinate-circle" cx="24" cy="24" r="18.0" stroke-dasharray="113.1" stroke-dashoffset="113.1" stroke-width="4.0"/>
</svg>

*/

func main() {

	output := &bytes.Buffer{}

	encoder := xml.NewEncoder(output)
	encoder.Indent("", "  ")
	err := encoder.Encode(&Svg{
		Class:     "mdc-circular-progress__determinate-circle-graphic",
		ViewBox:   "0 0 48 48",
		Namespace: "http://www.w3.org/2000/svg",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(output.String())

}
