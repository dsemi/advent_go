package problems

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	"unicode"
)

const rateLimit = 5 * time.Second

var last = new(time.Time)

var (
	_, b, _, _ = runtime.Caller(0)
	Basepath   = filepath.Join(filepath.Dir(b), "..")
)

func GetInput(year, day int, download bool) string {
	inputFile := filepath.Join(Basepath, fmt.Sprintf("inputs/%d/input%d.txt", year, day))
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) && download {
			fmt.Printf("Downloading input for Year %d Day %d\n", year, day)
			if last != nil {
				time.Sleep(time.Until(last.Add(rateLimit)))
			}
			*last = time.Now()
			url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
			client := &http.Client{}
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("Cookie", os.Getenv("AOC_SESSION"))
			resp, err := client.Do(req)
			if err != nil {
				log.Fatalf("Problem input fetch failed: %v", err)
			}
			if resp.StatusCode < 200 || resp.StatusCode > 299 {
				log.Fatalf("Bad HTTP response: %v", resp)
			}
			b := new(bytes.Buffer)
			if _, err = b.ReadFrom(resp.Body); err != nil {
				log.Fatalf("Error reading HTTP response body: %v", err)
			}
			if err = resp.Body.Close(); err != nil {
				log.Fatalf("Error closing HTTP response body: %v", err)
			}
			buf = b.Bytes()
			if err = ioutil.WriteFile(inputFile, buf, 0644); err != nil {
				log.Fatalf("Unable to write to output file: %v", err)
			}
		} else {
			log.Fatalf("Error reading problem input file: %v", err)
		}
	}
	return strings.TrimRightFunc(string(buf), unicode.IsSpace)
}
