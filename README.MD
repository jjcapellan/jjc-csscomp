# CSSCOMP  
Golang module to minify css strings.  

## Usage
Import csscomp module to your project with:
```
go get github.com/jjcapellan/jjc-csscomp
```

Then you can use its unique method in your code
```go
package main

import(
    "fmt"
    "github.com/jjcapellan/jjc-csscomp"
)

var myCSS string = `  h1, h2 {
    color: rgb(102, 255, 170);
    }  `

func main(){
    minCSS := csscomp.Get(mycss) // "Get" method returns the minified css string
    fmt.Println(minCSS)
}
// Output:
// h1,h2{color:#6fa;}
```