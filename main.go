package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/jonascavalcantineto/cloud-connector/connectors"
)

func main() {
	fmt.Println(connectors.aws.AWSApiStart())
	//fmt.Println(azure.AzureApiStart())
	// fmt.Println(connectors.GCPApiStart())
	// fmt.Println(connectors.ReverseRunes("jonascavalcantineto"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
