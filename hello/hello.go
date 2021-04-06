package hello

import "fmt"

const prefixHelloBr = "Hello, "
const prefixHelloEs = "Hola, "
const prefixHelloFr = "Bonjour, "

func Hello(name string, language string) string {
    if (name == "") {
        name = "World"
    }
    

    return prefixSalutation(language) + name
}

func prefixSalutation(language string) (prefix string) {
    switch language {
    case "spanish":
        prefix = prefixHelloEs
    case "french":
        prefix = prefixHelloFr
    default:
        prefix = prefixHelloBr
    }
    return
}

func main() {
    fmt.Println(Hello("World", ""))
}