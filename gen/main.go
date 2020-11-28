package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const template = `package main

var icon = []byte("%s")
`

func main() {
	data, err := ioutil.ReadFile("./gen/headset.png")
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.NewBuffer(nil)
	for _, v := range data {
		fmt.Fprintf(buf, `\x%02X`, v)
	}
	out := fmt.Sprintf(template, buf.String())
	os.Remove("data.go")
	err = ioutil.WriteFile("data.go", []byte(out), 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = exec.Command("gofmt", "-w", "data.go").Run()
	if err != nil {
		log.Fatal(err)
	}
}
