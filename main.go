package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
)

type Options struct {
	Url  string `short:"u" long:"url" description:"url of the jpg"`
	Rows int    `short:"r" long:"rows" description:"maximum number of rows"`
}

func PrintHeader(size int64, height int) {
	fmt.Printf("\033]1337")
	fmt.Printf(";File=")
	fmt.Printf(";inline=1")
	fmt.Printf(";size=%d;", size)
	if height > 0 {
		fmt.Printf(";height=%d", height)
	}
	fmt.Printf(":")
}

func PrintFooter() {
	fmt.Println("\a")
}

func PrintImg(img io.ReadCloser) {
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	defer img.Close()
	defer encoder.Close()
	buf := make([]byte, 8192)
	for {
		n, err := img.Read(buf)
		if err == io.EOF {
			break
		}
		encoder.Write(buf[:n])
	}
}

func main() {
	var options Options
	if _, err := flags.Parse(&options); err != nil {
		os.Exit(1)
	}

	if !(strings.HasPrefix(options.Url, "http://") || strings.HasPrefix(options.Url, "https://")) {
		log.Fatalln("url doesn't start with http")
	}

	resp, err := http.Get(options.Url)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != 200 {
		log.Fatalln(fmt.Sprintf("Issue to get img, status = %d", resp.StatusCode))
	}

	PrintHeader(resp.ContentLength, options.Rows)
	PrintImg(resp.Body)

	PrintFooter()
}
