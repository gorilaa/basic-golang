package main

import "fmt"

func logging()  {
	fmt.Println("Selesai menjalankan fungsi logging")
}

func runApplicaton() {
	defer logging();
	fmt.Println("App is running");
}

func main() {
	runApplicaton();
}

