package ybc

type Decoder interface {
	Decode(verifyRequest string) (map[string]any, error)
}
