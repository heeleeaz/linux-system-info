module main

go 1.16

replace domain => ./domain

replace data => ./data

replace presenter => ./presenter

require (
	presenter v0.0.0-00010101000000-000000000000
)
