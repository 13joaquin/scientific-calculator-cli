package capacidad

import (
	entrada "calculator-cli/internal/Entrada"
	"fmt"
)

type unidadLitros float64

const (
	Mililitro  unidadLitros = 0.001
	Centilitro unidadLitros = 0.01
	Decilitro  unidadLitros = 0.1
	Litro      unidadLitros = 1.0
	Decalitro  unidadLitros = 10.0
	Hectolitro unidadLitros = 100.0
	Kilolitro  unidadLitros = 1000.0
)

func ConvertirL(valor float64, de, a unidadLitros) float64 {
	litros := valor * float64(de)
	return litros / float64(a)
}

func mostrarMenuL() {
	fmt.Println("\n==== Convecion de Capacidad =====")
	fmt.Println("1. Litros")
	fmt.Println("2. Militro")
	fmt.Println("3. Centilitro")
	fmt.Println("4. Decilitro")
	fmt.Println("5. Decalitro")
	fmt.Println("6. Hectolitro")
	fmt.Println("7. Kilolitro")
	fmt.Println("8. Regresar")
}

func obtenerUnidadesL(opcion int) (unidadLitros, string, bool) {
	switch opcion {
	case 1:
		return Litro, "litros", true
	case 2:
		return Mililitro, "militro", true
	case 3:
		return Centilitro, "centilitro", true
	case 4:
		return Decilitro, "decilitros", true
	case 5:
		return Decalitro, "decalitro", true
	case 6:
		return Hectolitro, "hectolitro", true
	case 7:
		return Kilolitro, "kilolitro", true
	default:
		return 0, "", false
	}
}

func ConvertLitros() {
	for {
		entrada.LimpiarPantalla()
		mostrarMenuL()
		opcionLitros := entrada.LeerOpcion("Seleciona una opcion()")
		if opcionLitros == 8 {
			fmt.Println("Regresando al menu principal...")
			return
		}

		uOrigen, nombreOrigen, okOrigen := obtenerUnidadesL(opcionLitros)
		if !okOrigen {
			fmt.Printf("Opcion de origen no valida")
			entrada.Pausar()
			continue
		}
		opcDestino := entrada.LeerOpcion("Selecionar la unidad")
		uDestino, nombreDestino, okDestio := obtenerUnidadesL(opcDestino)
		if !okDestio {
			fmt.Printf("Opion de destino no valida")
			entrada.Pausar()
			continue
		}

		valorL := entrada.LeerNumero(fmt.Sprintf("Ingresa el valor en %s:", nombreOrigen))
		resultado := ConvertirL(valorL, uOrigen, uDestino)
		fmt.Printf("%.2f %s = %.2f %s\n", valorL, nombreOrigen, resultado, nombreDestino)

		entrada.Pausar()
	}
}

/*
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
}*/
