package main

import (
	"fmt"

	builder_facets "design_patterns/builder/facets"
	builder_parameter "design_patterns/builder/parameter"
	state_classic "design_patterns/state/classic"
	state_handmade "design_patterns/state/handmade"
	classic_visitor "design_patterns/visitor/classic"
	intrusive_visitor "design_patterns/visitor/intrusive"
	reflective_visitor "design_patterns/visitor/reflective"
)

func main() {
	fmt.Println("======BUILDER PATTERN=====")
	builder_parameter.Run()

	fmt.Println("======BUILDER FACETS=====")
	fmt.Println(builder_facets.Run())

	fmt.Println("======INTRUSIVE VISITOR=====")
	fmt.Println(intrusive_visitor.Run())

	fmt.Println("======REFLECTIVE VISITOR=====")
	fmt.Println(reflective_visitor.Run())

	fmt.Println("======CLASSIC VISITOR=====")
	fmt.Println(classic_visitor.Run())

	fmt.Println("======CLASSIC STATE=====")
	state_classic.Run()

	fmt.Println("======HANDMADE STATE=====")
	state_handmade.Run()
}
