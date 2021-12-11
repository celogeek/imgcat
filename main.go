package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Options struct {
	url    string
	height int
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
	var opts = &Options{}
	flag.StringVar(&opts.url, "url", "", "url of the jpg")
	flag.IntVar(&opts.height, "height", 0, "maximum height in lines")
	flag.Parse()

	if !(strings.HasPrefix(opts.url, "http://") || strings.HasPrefix(opts.url, "https://")) {
		log.Fatalln("url doesn't start with http")
	}

	resp, err := http.Get(opts.url)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != 200 {
		log.Fatalln(fmt.Sprintf("Issue to get img, status = %d", resp.StatusCode))
	}

	PrintHeader(resp.ContentLength, opts.height)
	PrintImg(resp.Body)
	PrintFooter()
}
