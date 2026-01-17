package main

func main() {

	cfg := Config{
		addr: ":8080",
		db:   DBconfig{},
	}

	api := Application{
		config: cfg,
	}

}