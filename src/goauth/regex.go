//regex.go
package main

import(
    "fmt"
    "regexp"
    "io/ioutil"
    "os/user"
    "log"
    "strings"
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

func BytesToString(data []byte) string {
    return string(data[:])
}

func main(){
    
    //ExampleRegexp_Expand()

    //ExampleRegexp_FindIndex()


    usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
    }
    fmt.Println( usr.HomeDir )

    kmdconf := usr.HomeDir+`/.komodo/komodo.conf`
    fmt.Println(kmdconf)

    confdata, err := ioutil.ReadFile(kmdconf)
    if err != nil {
        log.Fatal( err )
    }
    //fmt.Printf("%s\n",confdata)

    var rpcu = regexp.MustCompile("(?m)^rpcuser=.+$")
    var rpcpass = regexp.MustCompile("(?m)^rpcpassword=.+$")
    //fmt.Println(rpcu)

    //fmt.Println(rpcu.Match(confdata))

    bytestr := BytesToString(confdata)
    //fmt.Println("BytesToString: "+bytestr)

    rpcuser_line := rpcu.FindString(bytestr)
    rpcpass_line := rpcpass.FindString(bytestr)
    fmt.Printf("%q\n", rpcuser_line)
    fmt.Printf("%q\n", rpcpass_line)
    fmt.Printf("\n\n")
    //fmt.Printf(strings.TrimLeft(strings.TrimLeft(rpcuser_line,`rpcuser`),`=`)+"\n")

    kmd_rpcuser := strings.TrimLeft(strings.TrimLeft(rpcuser_line,`rpcuser`),`=`)
    kmd_rpcpass := strings.TrimLeft(strings.TrimLeft(rpcpass_line,`rpcpassword`),`=`)

    fmt.Printf("RPC User: %s\nRPC Password: %s\n", kmd_rpcuser, kmd_rpcpass)

    //re := regexp.MustCompile("rpcuser=.?")
    //fmt.Printf("%q\n", re.FindString(bytestr))

    //fmt.Println(re.FindStringSubmatch(rpcuser_line))
    
}