package client

// TODO: add logging
// https://nathanleclaire.com/blog/2015/10/10/interfaces-and-composition-for-effective-unit-testing-in-golang/
type TVMazeclient interface {
	Search(string) (string, )
}
