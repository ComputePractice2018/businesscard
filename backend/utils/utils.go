package utils

import "fmt"

//GetHelloWorldString возвращает строку с приветствием
func GetHelloWorldString(name string) string {
	return fmt.Sprintf("Привет, %s", name)
}
