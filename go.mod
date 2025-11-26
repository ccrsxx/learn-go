module github.com/ccrsxx/learn-go

go 1.25.4

require (
	github.com/ccrsxx/learn-go/src/getting-started/greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace github.com/ccrsxx/learn-go/src/getting-started/greetings => ./src/getting-started/greetings
