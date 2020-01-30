package main

import (
	"fmt"
	"github.com/jonascavalcantineto/cloud-connector/connectors/aws"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(connectors.AWSApiStart())
	 fmt.Println(connectors.AzureApiStart())
	// fmt.Println(connectors.GCPApiStart())
	// fmt.Println(connectors.ReverseRunes("jonascavalcantineto"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
