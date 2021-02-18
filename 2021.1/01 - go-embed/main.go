package main

import "embed"

//go:embed index.html public/favicon.ico src
var webApp embed.FS

func main() {
	_ = webApp
}
