package chapter13

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// BroadcastServer manages multiple client connections
type BroadcastServer struct {
	clients map[chan string]bool
	mu      sync.RWMutex
}

// NewBroadcastServer creates a new broadcast server
func NewBroadcastServer() *BroadcastServer {
	return &BroadcastServer{
		clients: make(map[chan string]bool),
	}
}

// Register adds a new client
func (bs *BroadcastServer) Register(client chan string) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	bs.clients[client] = true
	log.Printf("Client registered. Total clients: %d", len(bs.clients))
}

// Unregister removes a client
func (bs *BroadcastServer) Unregister(client chan string) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	if _, ok := bs.clients[client]; ok {
		delete(bs.clients, client)
		close(client)
		log.Printf("Client unregistered. Total clients: %d", len(bs.clients))
	}
}

// Broadcast sends a message to all connected clients
func (bs *BroadcastServer) Broadcast(message string) {
	bs.mu.RLock()
	defer bs.mu.RUnlock()

	for client := range bs.clients {
		select {
		case client <- message:
		default:
			// Client channel is full or blocked, skip
		}
	}
}

// StreamHandler handles SSE (Server-Sent Events) connections
func (bs *BroadcastServer) StreamHandler(c *gin.Context) {
	// Set headers for SSE
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")

	// Create a channel for this client
	clientChan := make(chan string, 10)

	// Register the client
	bs.Register(clientChan)
	defer bs.Unregister(clientChan)

	// Get the response writer
	w := c.Writer

	// Send initial connection message
	fmt.Fprintf(w, "data: Connected to broadcast server\n\n")
	w.Flush()

	// Listen for messages and client disconnect
	for {
		select {
		case msg := <-clientChan:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			w.Flush()
		case <-c.Request.Context().Done():
			return
		}
	}
}

func MakeBroadcastServer() {
	server := NewBroadcastServer()

	// Start a goroutine to simulate broadcasting messages
	go func() {
		counter := 0
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			counter++
			message := fmt.Sprintf("Broadcast message #%d at %s", counter, time.Now().Format("15:04:05"))
			log.Println("Broadcasting:", message)
			server.Broadcast(message)
		}
	}()

	// Set Gin to release mode (optional)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// SSE stream endpoint
	router.GET("/stream", server.StreamHandler)

	// Home page with client
	router.GET("/", func(c *gin.Context) {
		html := templ()
		c.Header("Content-Type", "text/html")
		c.String(200, html)
	})

	log.Println("Broadcast server starting on :8080")
	log.Println("Open http://localhost:8080 in multiple browser windows to see broadcasting in action")
	router.Run(":8080")
}

func templ() string {
	return `<!DOCTYPE html>
<html>
<head><title>Broadcast Client</title></head>
<body>
	<h1>Broadcast Server Demo (Gin Framework)</h1>
	<div id="messages"></div>
	<script>
		const eventSource = new EventSource('/stream');
		const messagesDiv = document.getElementById('messages');
		
		eventSource.onmessage = function(event) {
			const p = document.createElement('p');
			p.textContent = event.data;
			messagesDiv.appendChild(p);
		};
		
		eventSource.onerror = function(err) {
			console.error('EventSource error:', err);
		};
	</script>
</body>
</html>`
}
