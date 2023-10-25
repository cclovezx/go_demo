package main

import (
	"fmt"
	interfaced "packagestudy/imterface/interface_01"
)

func main() {
	phone := interfaced.Phone{}
	phone.Start()
	fmt.Println()
	phone.Stop()
}
