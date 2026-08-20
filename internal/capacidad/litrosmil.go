package capacidad

import (
	entrada "calculator-cli/internal/Entrada"
	"fmt"
)

const (
	Mililitro  float64 = 0.001
	Centilitro float64 = 0.01
	Decilitro  float64 = 0.1

	Litro      float64 = 1.0
	Decalitro  float64 = 10.0
	Hectolitro float64 = 100.0
	Kilolitro  float64 = 1000.0
)

func Convertir(valor float64, origen float64, destino float64) float64 {
	return (valor / origen) * destino
}

func ProcesarCOnversion(origenTxt string, destinoTxt string, origenConst float64, destinoConst float64) {
	pregunta := fmt.Sprintf("Ingrese la cantidad en %s: ", origenTxt)
	valor := entrada.LeerNumero(pregunta)
	resultado := Convertir(valor, origenConst, destinoConst)
	fmt.Printf("%.2f %s son %.2f %s\n", valor, origenTxt, resultado, destinoTxt)
}
func MostrarMenu() {
	fmt.Println("====Convetidor de Capacidad=====")
	fmt.Println("1. Litro a Decalitro")
	fmt.Println("2. Litro a Kilolitro")
	fmt.Println("3. Litro a Militro")
	fmt.Println("4. Regresaar")
}
func MainCapa() {
	for {
		entrada.LimpiarPantalla()
		MostrarMenu()
		opcion := entrada.LeerOpcion("Seleccione una opcion (1-4): ")
		switch opcion {
		case 1:
			ProcesarCOnversion("Litros", "Decalitros", Litro, Decalitro)
		case 2:
			ProcesarCOnversion("Litros", "Kilolitros", Litro, Kilolitro)
		case 3:
			ProcesarCOnversion("Litros", "Mililitros", Litro, Mililitro)
		case 4:
			fmt.Println("Regresando al menu principal...")
			return
		default:
			fmt.Println("Opcion no valida, Intentalo de nuevo")
		}
		entrada.Pausar()
	}
}
