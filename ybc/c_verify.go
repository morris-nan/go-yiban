package ybc

type Verifier interface {
	Verify(verifyRequest string) (bool, error)
}
