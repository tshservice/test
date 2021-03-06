package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
        "net"
)

func main() {
        port := "8080"
        if fromEnv := os.Getenv("PORT"); fromEnv != "" {
                port = fromEnv
        }

        server := http.NewServeMux()
        server.HandleFunc("/", hello)
        log.Printf("Server listening on port %s", port)
        log.Fatal(http.ListenAndServe(":"+port, server))
}

func hello(w http.ResponseWriter, r *http.Request) {
        log.Printf("Serving request: %s", r.URL.Path)
        host, _ := os.Hostname()
	resp := os.Getenv("RESP")
	if len(resp) == 0 {
		resp = "Hello, world!"
	}
	fmt.Println("Servicing request.")
	fmt.Fprintf(w, "%s\n", resp)
        fmt.Fprintf(w, "Version: 1.1.\n")
        fmt.Fprintf(w, "Hostname: %s\n", host)
        addrs, _ := net.LookupIP(host)
        for _, addr := range addrs {
                if ipv4 := addr.To4(); ipv4 != nil {
                        fmt.Fprintf(w, "IPv4: %s\n", ipv4)
                }
        }

}
