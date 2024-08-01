# Homework 1 : Command-line form

Design a simple command-line application that asks users for their name
and year of birth. Once a user has entered the information, the
application should respond with the following message.

    Hello, <name>! You are currently <age> years old.

The [fmt][] package implements functions for dealing with formatted
input and output. The `fmt.Print` function, for instance, writes the
given values to the standard output (i.e., the terminal). The
`fmt.Scanln` function scans text read from standard input, and stores
the value into the variable at to the given address.

[fmt]: https://pkg.go.dev/fmt

The user's age should be calculated based on the current year. The
`time.Now` function, which is part of the standard library's [time][]
package, returns a value of type `Time` that represents the current
local time. The `(t) Year` method can be called on `Time` values to get
the corresponding year.

[time]: https://pkg.go.dev/time

