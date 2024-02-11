package api

// Initialize and Run App against Default Db
func main() {
	a := App{}
	a.Default()

	a.Run(":8010")
}
