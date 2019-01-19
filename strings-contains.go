package main
import ( "fmt"; "strings"; )
func main() {
    fmt.Println(strings.Contains("Australia", "Aus"))   // Any part of string
    fmt.Println(strings.Contains("Aus", "Australia"))   // Any part of string
    fmt.Println()
    fmt.Println(strings.Contains("Australia", "Australian"))
    fmt.Println(strings.Contains("Australian", "Australia"))
    fmt.Println()
    fmt.Println(strings.Contains("Japan", "JAPAN")) // Case sensitive
    fmt.Println(strings.Contains("Japan", "JAP")) // Case sensitive
    fmt.Println()
    fmt.Println(strings.Contains("Become inspired to travel to Australia.", "Australia"))
    fmt.Println(strings.Contains("Australia", "Become inspired to travel to Australia.")) 
    fmt.Println()
    fmt.Println(strings.Contains("", ""))
    fmt.Println(strings.Contains("  ", " ")) // space also consider as string
    fmt.Println()
    fmt.Println(strings.Contains("12554", "1"))
    fmt.Println(strings.Contains("1", "12554"))
}
