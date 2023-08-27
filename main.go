package main

func main() {
    apiServer := NewServer(":3000")
    apiServer.Run()
}
