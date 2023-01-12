package main

import (
	"runtime/debug"
)

func (app *application) backgroundTask(fn func()) {
	app.wg.Add(1)

	go func() {
		defer app.wg.Done()

		defer func() {
			err := recover()
			if err != nil {
				app.logger.Printf("%s\n%s", err, debug.Stack())
			}
		}()

		fn()
	}()
}
