module github.com/johandry/klient-examples/simple

go 1.14

require (
	github.com/johandry/klient v0.0.0
	k8s.io/apimachinery v0.17.3
)

replace github.com/johandry/klient => ../../klient
