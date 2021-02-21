package main
import "fmt"

func showArea(len int,breadth int)
{
	fmt.Println(len*breadth)
}

func main()
{
	l := 10
	b := 25

	showArea(l,b)
}
