package translit_test

import (
	"fmt"
	"github.com/mdigger/translit"
)

func ExampleRuTranslit() {
	tests := []string{
		"Проверочная СТРОКА для транслитерации",
		"ЧАЩА",
		"ЧаЩа",
		"Чаща",
		"чаЩА",
	}
	for _, text := range tests {
		fmt.Println(translit.RuTranslit(text))
	}
	// Output:
	// Proverochnaja STROKA dlja transliteracii
	// CHASCHA
	// ChaScha
	// Chascha
}
