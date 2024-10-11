package ybc

type Client interface {
	Author
	Verifier
	Decoder
}
