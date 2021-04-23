package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Step 1. Invoke Search Everywhere and search for Manage Targets...

// Step 2. Click on the + button and add the Run Target you need
//		   from the options available: Docker, SSH, WSL 2

// Step 3. Go to the green run triangle on the editor gutter
// 		   and select Modify Run Configuration

// Step 4. Change the Run on for the Run Configuration to the
//		   Run Target that you created in Step 2.

func main() {
	http.HandleFunc("/", homeHandler)

	err := http.ListenAndServe(":38000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func inDocker() string {
	result := "inside"
	if _, err := os.Lstat("/.dockerenv"); err != nil && os.IsNotExist(err) {
		result = "outside"
	}

	return fmt.Sprintf("I'm running %s a Docker container.", result)
}

func inWSL() string {
	result := "outside"
	cmd := exec.Command("uname", "-a")
	if output, err := cmd.Output(); err == nil {
		if strings.Contains(strings.ToLower(string(output)), "microsoft") {
			result = "inside"
		}
	}

	return fmt.Sprintf("Running %s a WSL environment.", result)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	myOS, myArch := runtime.GOOS, runtime.GOARCH

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	_, _ = fmt.Fprintf(w, "Hello, %s!\n", r.UserAgent())
	_, _ = fmt.Fprintf(w, "I'm running on %s/%s.\n", myOS, myArch)
	_, _ = fmt.Fprintf(w, "%s\n", inDocker())
	_, _ = fmt.Fprintf(w, "%s\n", inWSL())
}
