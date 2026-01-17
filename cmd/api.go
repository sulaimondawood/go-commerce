package main

type Application struct {
	config Config
}

type Config struct {
	db   DBconfig
	addr string
}

type DBconfig struct {
	dbUrl string
}