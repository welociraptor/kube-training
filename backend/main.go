package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("unable to connect to database!", err)
	}
	defer conn.Close(context.Background())
	s := Server{
		conn: conn,
	}
	http.HandleFunc("/hello", s.serve)
	http.ListenAndServe(":8080", nil)
}

type Server struct {
	conn *pgx.Conn
}

func (s *Server) getLastGreeting() (string, error) {
	var name string
	var message string
	err := s.conn.QueryRow(context.Background(), "SELECT name, message FROM greetings ORDER BY id ASC LIMIT 1").Scan(&name, &message)
	if err != nil {
		return "", err
	}
	return name + " says " + message, nil
}

func (s *Server) serve(w http.ResponseWriter, r *http.Request) {
	greeting, err := s.getLastGreeting()
	if err != nil {
		log.Fatalln("unable to get last greeting: %w", err)
	}
	fmt.Fprintf(w, greeting)
}
