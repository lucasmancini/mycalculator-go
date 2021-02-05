package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calc struct{}

func LeerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	entrada := scanner.Text()
	return entrada
}

func (Calc) Operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])

	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
		return operador1 + operador2
	case "-":
		fmt.Println(operador1 - operador2)
		return operador1 - operador2
	case "*":
		fmt.Println(operador1 * operador2)
		return operador1 * operador2
	case "/":
		fmt.Println(operador1 / operador2)
		return operador1 / operador2
	default:
		fmt.Println(operador, "no soportado")
		return 0
	}

}

func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func main() {
	entrada := LeerEntrada()
	operador := LeerEntrada()
	c := Calc{}
	fmt.Println(c.Operate(entrada, operador))
}

// Go realmente no tiene soporte para Objetos como tal, pero si para Structs. Los structs son colecciones
// de campos tipados. Son útiles para agrupar datos y formar registros. Por ejemplo puede existir
// el struct persona con campos nombre (string) y edad (int). En la definición de struct
// solo ponemos sus atributos. Luego, podemos crear funciones (receiver functions) que hagan parte del struct
// como se hizo aquí con operate que se relacionó con calc. Manejar structs nos permite organizar
// mejor el código también.
