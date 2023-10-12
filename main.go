package main

import (
	"fmt"

	builder_facets "design_patterns/builder/facets"
	builder_parameter "design_patterns/builder/parameter"
	intrusive_visitor "design_patterns/visitor/intrusive"
)

func main() {
	builder_parameter.Run()
	fmt.Println(builder_facets.Run())
	fmt.Println(intrusive_visitor.Run())
}
