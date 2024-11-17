package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func main() {
	// Configure the Cassandra cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "system"
	cluster.Consistency = gocql.Quorum

	// Create a session
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Unable to connect to Cassandra: %v", err)
	}
	defer session.Close()

	fmt.Println("Connected to Cassandra!")

	// Example: Query Cassandra
	var clusterName string
	if err := session.Query(`SELECT cluster_name FROM system.local`).Scan(&clusterName); err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	fmt.Printf("Cluster name: %s\n", clusterName)
}
