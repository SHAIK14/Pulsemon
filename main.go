package main

func main() {
	data, err := config.Load("config.json")
	if err != nil {
		return err
	}
	for _, d := range data {
		fmt.Pri
	}

}
