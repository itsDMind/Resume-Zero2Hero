package main

import (
 "fmt"
 "net/http"
 "sync"
)

// LoadBalancer struct represents the load balancer.
type LoadBalancer struct {
 servers []string
 index   int
 mutex   sync.Mutex
}

// NewLoadBalancer creates a new LoadBalancer instance.
func NewLoadBalancer(servers []string) *LoadBalancer {
 return &LoadBalancer{
  servers: servers,
  index:   0,
 }
}

// ServeHTTP handles incoming HTTP requests and forwards them to backend servers.
func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
 lb.mutex.Lock()
 defer lb.mutex.Unlock()

 // Get the next server in a round-robin fashion
 server := lb.servers[lb.index]
 lb.index = (lb.index + 1) % len(lb.servers)

 // Forward the request to the selected server
 proxy := NewProxy(server)
 proxy.ServeHTTP(w, r)
}

// Proxy struct represents a simple reverse proxy.
type Proxy struct {
 target string
}

// NewProxy creates a new Proxy instance.
func NewProxy(target string) *Proxy {
 return &Proxy{target: target}
}

// ServeHTTP forwards the incoming HTTP request to the target server.
func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
 // Forward the request to the target server
 targetURL := fmt.Sprintf("http://%s%s", p.target, r.URL.Path)
 resp, err := http.Get(targetURL)
 if err != nil {
  http.Error(w, "Error forwarding request", http.StatusInternalServerError)
  return
 }
 defer resp.Body.Close()

 // Copy the response from the target server to the original response writer
 for key, values := range resp.Header {
  for _, value := range values {
   w.Header().Add(key, value)
  }
 }
 w.WriteHeader(resp.StatusCode)
 _, _ = w.Write([]byte("Response from: " + p.target))
}

func main() {
 // Define backend servers
 servers := []string{"localhost:8081", "localhost:8082", "localhost:8083"}

 // Create a new load balancer
 loadBalancer := NewLoadBalancer(servers)

 // Start the load balancer on port 8080
 http.ListenAndServe(":8080", loadBalancer)
}