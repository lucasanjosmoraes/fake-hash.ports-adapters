package errorhandler

// HTTPResponder define o que um erro precisa implementar para ser retornado
// numa requisição HTTP.
type HTTPResponder interface {
	error
	Status() int
	Response() []byte
}
