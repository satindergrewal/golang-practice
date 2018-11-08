//regex.go
package main

import(
    "fmt"
    "regexp"
)

func ExampleRegexp_Expand() {

	content := []byte(`

	# comment line

	option1= value1

	option2= value2


	# another comment line

	option3= value3

`)


	// Regex pattern captures "key: value" pair from the content.

	pattern := regexp.MustCompile(`(?m)(?P<key>\w+)=\s+(?P<value>\w+)$`)


	// Template to convert "key: value" to "key=value" by

	// referencing the values captured by the regex pattern.

	template := []byte("$key:$value\n")


	result := []byte{}


	// For each match of the regex in the content.

	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {

		// Apply the captured submatches to the template and append the output

		// to the result.

		result = pattern.Expand(result, template, content, submatches)

	}

	fmt.Println(string(result))

	// Output:

	// option1=value1

	// option2=value2

	// option3=value3

}

func ExampleRegexp_FindIndex() {

	content := []byte(`

	# comment line

	option1: value1

	option2: value2

`)

	// Regex pattern captures "key: value" pair from the content.

	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)


	loc := pattern.FindIndex(content)

	fmt.Println(loc)

	fmt.Println(string(content[loc[0]:loc[1]]))

	fmt.Println(string(content))
	//fmt.Println(string(content[loc[1]))

	// Output:

	// [18 33]

	// option1: value1

}

func main(){
    
    ExampleRegexp_Expand()

    //ExampleRegexp_FindIndex()
    
}