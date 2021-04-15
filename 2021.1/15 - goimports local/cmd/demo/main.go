package main

import (
	myLog "loc.al/module/pkg/log"
	"loc.al/module/pkg/log/logrus"
	"loc.al/module/pkg/log/zap"
)

// Step 1. To enable this support, head over to
// 		   Preferences/Settings | Editor | Code Style | Go | Imports
// 		   and enable "Import grouping".
// 		   Then you can specify which imports should be grouped into separate blocks.
// E.g. you can specify the loc.al package

// Step 2. Unfold the imports above.

func main() {
	// Step 3. Uncomment the lines below
	// Shortcut: Ctrl + / on Windows/Linux
	//			 Cmd + / on macOS

	/*now := time.Now()
	_ = fmt.Sprintf("Now it's %v\n", now)*/

	/*err := errors.New("this comes from pkg/errors not errors")
	log.Println(err)*/
}

func _() {
	message := "Hello, World!"
	process(&logrus.Adapter{}, message)
	process(&zap.Adapter{}, message)
}

func process(a myLog.Do, message string) {
	_ = a.Info(message)
}
