package aritmetica

import (
	entrada "calculator-cli/internal/Entrada"
	"fmt"
	"math"
)

const (
	OpPotencia = iota + 1
	OpRaiz
	OpRegresar
)

func potencia(base float64, exponente float64) float64 {
	return math.Pow(base, exponente)
}

func mostrarMenuP() {
	fmt.Println("\n===== Potencia y Raiz ======")
	fmt.Println("1. Potencia")
	fmt.Println("2. Raiz")
	fmt.Println("3. Regresar")
}

func PotenciaRaiz() {
	for {
		entrada.LimpiarPantalla()
		mostrarMenuP()
		opcionPR := entrada.LeerOpcion("Selecionar una opcion (1-3)")
		switch opcionPR {
		case OpRegresar:
			fmt.Println("Regresando al menu prinipal...")
			return
		case OpPotencia:
			numbase := entrada.LeerNumero("Ingrese el numero base de la potencia:")
			numEx := entrada.LeerNumero("Ingrese el numero de exponente de la potencia: ")

			Potencia := potencia(numbase, numEx)
			fmt.Printf("El resultado de %.2f elevado a %.2f es: %.2f\n", numbase, numEx, Potencia)
		case OpRaiz:
			numRaiz := entrada.LeerNumero("Ingrese el numero para calcular la raiz cuadrada:")
			resultado := math.Sqrt(numRaiz)
			fmt.Printf("La raiz cuadrada de %.2f es: %.2f\n", numRaiz, resultado)
		default:
			fmt.Println("Opcion no valida, Intentalo de nuevo")
		}
		entrada.Pausar()
	}
}
