package main

import "fmt"

func isValid(s string) bool {
	var store []string
	if len(s)%2 != 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "(":
			store = append([]string{")"}, store...)
		case "[":
			store = append([]string{"]"}, store...)
		case "{":
			store = append([]string{"}"}, store...)
		default:
			if len(store) == 0 {
				return false
			}
			target, storeAfterRemove := store[0], store[1:]
			if target == string(s[i]) {
				store = storeAfterRemove
				continue
			} else {
				return false
			}
		}
	}
	return len(store) == 0
}

func main() {
	fmt.Println(isValid("()"))
}
