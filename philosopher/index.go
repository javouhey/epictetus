package philosopher

type Name string

func Salutation(name Name) string {
    return "Hello " + string(name)
}
