package main

import (
	"fmt"
)

/*
“11:-)))((“ -> “11((“
“:-):-(“ -> “”
“:------------))(“ -> “:------------))(“
“:::-)” -> “::”
“(((())):):((“ -> “(((())):):((“
"11::-))(" -> "11:("

    :-)
    :-(
*/

func removeSmileFaces(s string) string {
    var result string 
    
    runes := []rune(s)

    for i := 0; i < len(runes); {
        if i < len(runes) - 2 && runes[i] == ':' && runes[i+1] == '-' && (runes[i+2] == ')' || runes[i+2] == '(') {
            i += 3
            for ; i < len(runes) && runes[i] == runes[i-1]; i++ {}
        } else {
            result += string(runes[i])
            i++
        }
    }
   return result 
}

func main() {
    fmt.Println(removeSmileFaces("11:-)))((") == "11((")
    fmt.Println(removeSmileFaces(":-):-(") == "")
    fmt.Println(removeSmileFaces(":-----))(") == ":-----))(")
    fmt.Println(removeSmileFaces(":::-)") == "::")
    fmt.Println(removeSmileFaces("(((())):):((") == "(((())):):((")
    fmt.Println(removeSmileFaces("11::-))(") == "11:(")
}
