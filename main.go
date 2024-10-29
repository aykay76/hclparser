package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <tfvars-file>", os.Args[0])
	}

	tfvarsFile := os.Args[1]
	variables, err := parseTFVars(tfvarsFile)
	if err != nil {
		log.Fatalf("Error parsing TFVars file: %s", err)
	}

	fmt.Println("Terraform Variables:")
	for name, _ := range variables {
		fmt.Printf("%s\n", name)
	}
}

// parse tfvars file
func parseTFVars(tfvarsFile string) (map[string]cty.Value, error) {
	parser := hclparse.NewParser()
	file, diags := parser.ParseHCLFile(tfvarsFile)
	if diags.HasErrors() {
		return nil, diags
	}

	var variables map[string]cty.Value
	diags = gohcl.DecodeBody(file.Body, nil, &variables)
	if diags.HasErrors() {
		return nil, diags
	}

	return variables, nil
}
