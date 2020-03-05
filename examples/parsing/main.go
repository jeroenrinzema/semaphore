package main

import (
	"fmt"
	"strings"

	"github.com/jexia/maestro/specs/intermediate"
)

const definition = `
flow "user" {
	input {
		name = "<string>"
		message "address" {
			city = "<string>"
			state = "<string>"
			country = "<string>"
		}
	}

	call "logging" "logger.Log" {
		request {
			options {
				method = "GET"
			}

			header {
			}

			message = "{{ input:name }}"
		}
	}
}
`

func main() {
	manifest, err := intermediate.UnmarshalHCL("flows.hcl", strings.NewReader(definition))
	if err != nil {
		panic(err)
	}

	fmt.Println(manifest)
}
