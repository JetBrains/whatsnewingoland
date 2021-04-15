package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Step 1. Run this application then jump to check.http file

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

	return fmt.Sprintf("I'm running running %s a Docker container.", result)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := struct {
		Message     string
		Environment string
		InDocker    string
		InWSL       string
	}{
		Message:     fmt.Sprintf("Hello, %s!", r.UserAgent()),
		Environment: fmt.Sprintf("I'm running on %s/%s.", myOS, myArch),
		InDocker:    inDocker(),
		InWSL:       inWSL(),
	}

	resp, _ := json.Marshal(response)
	resp = append(resp, '\n')
	for i := 0; i < 10; i++ {
		_, _ = w.Write(resp)
	}
}
