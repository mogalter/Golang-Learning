package while

import (
	"fmt"
	"while"
)

func main() {
	x := 7
	fmt.Println("Repeated:", x)
	while.LoopWhile(x)
}
