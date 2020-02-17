package main

import "fmt"

func main() {
	// normal or interpreted string
	var name = "Steve Jobs"

	// raw string (can span multiple lines. escape characters are not interpreted.)
	var bio = `Steve Jobs was an American entrepreneur and inventor.
	He was the CEO and co-founder of Apple Inc.`

	fmt.Println(name)
	fmt.Println(bio)

	fmt.Println("")

	var website = "\thttps://www.callicoder.com\t\n"

	var siteDescription = `\t\tCalliCoder is a programming blog where you can find
                           practical guides and tutorials on programming languages, 
                           web development, and desktop app development.\t\n`

	fmt.Println(website, siteDescription)

}
