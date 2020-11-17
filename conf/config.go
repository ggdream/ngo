package conf



const (
	NAME		= "ngo"
	VERSION		= "1.1.0"
	USAGE		= "Enter `ngo` in the terminal to run quickly."
)

var (
	HOST		= "0.0.0.0"
	PORT		= 54639
	PATH		= "/"
	STATIC		= "."

	Origin		= []string{"*"}

	PrintLog	= true
	Daemon		= false
)
