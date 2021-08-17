module example.com/user

go 1.16

replace example.com/math => ../math

replace example.com/greetings => ../greetings

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	example.com/math v0.0.0-00010101000000-000000000000
)
