package main

import (
	entrada "calculator-cli/internal/Entrada"
	calculo "calculator-cli/internal/aritmetica"
	potencia "calculator-cli/internal/aritmetica"
	capacidad "calculator-cli/internal/capacidad"
	convertir "calculator-cli/internal/convertidor"
	longitud "calculator-cli/internal/longitud"
	masa "calculator-cli/internal/masa"

	"fmt"
	"time"
)

const (
	OpCalculo = iota + 1
	OpConvertir
	OpLongitud
	OpCapacidad
	OpFechaHoras
	OpPoteRaiz
	OpMasa
	OpSalir
)

func mostrarMenu() {
	fmt.Println("\n======Calculadro CLI======")
	fmt.Println("1. Calculadora")
	fmt.Println("2. Convertidor Grados")
	fmt.Println("3. Longitud")
	fmt.Println("4. Capacidad")
	fmt.Println("5. Fecha y Hora")
	fmt.Println("6. Potencia Raiz")
	fmt.Println("7. Masa")
	fmt.Println("8. Salir")
}
func mostrarFechaHora() {
	ahora := time.Now()
	fmt.Println("Fecha y hora actual:", ahora.Format("2006-01-02 15:04:05"))
}
func main() {
	for {
		entrada.LimpiarPantalla()

		mostrarMenu()
		opcionMenu := entrada.LeerNumero("Seleccionar una opcion (1-8):")
		switch opcionMenu {
		case OpSalir:
			fmt.Println("¡Hasta Luego!")
			return
		case OpCalculo:
			calculo.Calculadora()
		case OpConvertir:
			convertir.Convertidor()
		case OpLongitud:
			longitud.ConverLogitud()
		case OpCapacidad:
			capacidad.ConvertLitros()
		case OpPoteRaiz:
			potencia.PotenciaRaiz()
		case OpFechaHoras:
			mostrarFechaHora()
		case OpMasa:
			masa.ConverMasa()
			continue
		}

	}

}
