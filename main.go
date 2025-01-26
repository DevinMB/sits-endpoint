package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	// "net/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, or error loading it.")
	}

	couchURL := os.Getenv("COUCH_URL")   
	couchUser := os.Getenv("COUCH_USER") 
	couchPass := os.Getenv("COUCH_PASS") 

	err := ConnectCouchDB(couchURL, couchUser, couchPass)
	if err != nil {
		log.Fatalf("Failed to connect to CouchDB: %v", err)
	}

	count, err := GetRowCount()
	if err != nil {
		log.Fatalf("Row count error: %v", err)
	}
	log.Printf("Row count in 'sits': %d", count)

	// data, err := GetStackedHistoricalEvents()
	// log.Printf("Stacked historical events: %+v", data)

	// 	// Start Kafka consumer in a separate goroutine
	// 	go StartKafkaConsumer("your-topic", "localhost:9092")

	// 	// Initialize WebSocket server
	// 	http.HandleFunc("/ws", WebSocketHandler)

	// 	log.Println("Server is running on :8080...")
	// 	log.Fatal(http.ListenAndServe(":8080", nil))
}
