package main

import "pb3/db"

func main() {
	db.InitDatabase("postgres://postgres:mysecretpassword@localhost:5432/problem3?sslmode=disable", 1)
}
