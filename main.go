package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/jonascavalcantineto/cloud-connector/connectors/aws"
	"github.com/jonascavalcantineto/cloud-connector/connectors/azure"
	"github.com/jonascavalcantineto/cloud-connector/connectors/gcp"
)

func main() {
	fmt.Println(aws.AWSApiStart())
	fmt.Println(azure.AzureApiStart())
	fmt.Println(gcp.GCPApiStart())
	fmt.Println(aws.ReverseRunes("jonascavalcantineto"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
