package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com/search?q=iit+madras+qs+rank&sca_esv=e04e555e27c90511&rlz=1C5CHFA_enIN990IN990&biw=1440&bih=812&sxsrf=AHTn8zo6xHskC3X1UJHrqv9LoydI9ghXsQ%3A1739283939421&ei=412rZ4KzGea64-EPhqf2mA8&ved=0ahUKEwjCsP3K6buLAxVm3TgGHYaTHfMQ4dUDCBA&uact=5&oq=iit+madras+qs+rank&gs_lp=Egxnd3Mtd2l6LXNlcnAiEmlpdCBtYWRyYXMgcXMgcmFuazILEAAYgAQYkQIYigUyBRAAGIAEMgUQABiABDIFEAAYgAQyBBAAGB4yBBAAGB4yBhAAGAgYHjIGEAAYCBgeMgYQABgIGB4yBhAAGAgYHki-G1CtBFiUGXACeAGQAQGYAbIDoAGbF6oBCTAuMi45LjAuMbgBA8gBAPgBAZgCCKACrAvCAgoQABiwAxjWBBhHwgIKECMYgAQYJxiKBcICChAAGIAEGEMYigXCAg4QABiABBiRAhixAxiKBcICCBAAGAcYChgewgIGEAAYBxgewgIHEAAYgAQYDcICCxAAGIAEGIYDGIoFwgIFEAAY7wXCAggQABiABBiiBMICCBAAGKIEGIkFmAMAiAYBkAYIkgcFMi4wLjagB9hd&sclient=gws-wiz-serp"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response if of type: %T\n", response)
	defer response.Body.Close()
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)

}
