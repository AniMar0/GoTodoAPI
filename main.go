package main

import (
	API "gotodoapi/API"
)

func main() {
	S := API.Server{}
	S.Run("8080")
}
