package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/proxy"
)

func main() {
	content, err := os.ReadFile("targets.yaml")
	if err != nil {
		fmt.Println("Dosya Bulunamadı:", err)
		return
	}

	targets := strings.Split(string(content), "\n")

	torproxy := "127.0.0.1:9050"
	dialer, err := proxy.SOCKS5("tcp", torproxy, nil, proxy.Direct)
	if err != nil {
		fmt.Println("Tor Bağlantı Hatası", err)
		return
	}

	transportayari := &http.Transport{
		Dial: dialer.Dial,
	}

	istemci := &http.Client{
		Transport: transportayari,
		Timeout:   20 * time.Second,
	}

	fmt.Println("---------------------------------------------------")

	for _, line := range targets {
		url := strings.TrimSpace(line)
		if url == "" {
			continue
		}
		tarama(istemci, url)
	}
}

func tarama(client *http.Client, url string) {
	resp, err := client.Get(url)
	if err != nil {
		errStr := err.Error()
		if strings.Contains(errStr, "Timeout") || strings.Contains(errStr, "deadline exceeded") {
			fmt.Printf("[ERR]  Scanning: %s -> TIMEOUT\n", url)
		} else {
			fmt.Printf("[ERR]  Scanning: %s -> FAILED (Connection Error)\n", url)
		}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("[INFO] Scanning: %s -> SUCCESS\n", url)
	} else {
		fmt.Printf("[INFO] Scanning: %s -> SUCCESS (Status: %d)\n", url, resp.StatusCode)
	}
}
