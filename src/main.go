package main

import (
	"otus.ru/tbw/msa-25/src/app"
	"otus.ru/tbw/msa-25/src/deps"
)

func main() {
	deps := deps.Prepare()
	app.Run(deps)
}
