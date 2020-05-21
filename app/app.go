package app

const defaultHttpServer = "localhost:7777"

// Represents simple zero-cost message that can be used
// as signal between goroutines.
type sig struct{}
