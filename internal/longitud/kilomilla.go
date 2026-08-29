package longitud

import (
	entrada "calculator-cli/internal/Entrada"
	"fmt"
)

type unidadLongitud float64

const (
	Metro      unidadLongitud = 1.0
	Kilometro  unidadLongitud = 1000.0
	Centimetro unidadLongitud = 0.01
	Milla      unidadLongitud = 1609.344
	Pie        unidadLongitud = 0.3048
	Pulgada    unidadLongitud = 0.0254
)

func Convertir(valor float64, de, a unidadLongitud) float64 {
	metro := valor * float64(de)
	return metro / float64(a)
}

func mostrarMenu() {
	fmt.Println("\n===== Convercion de Loongitud ======")
	fmt.Println("1. Metro")
	fmt.Println("2 Kilometro")
	fmt.Println("3. Centimetro")
	fmt.Println("4. Milla")
	fmt.Println("5. Pies ")
	fmt.Println("6. Pulgadas")
	fmt.Println("7. Regresar")
}

func obtenerUnidades(opcion int) (unidadLongitud, string, bool) {
	switch opcion {
	case 1:
		return Metro, "metros", true
	case 2:
		return Kilometro, "kilometros", true
	case 3:
		return Centimetro, "centimetros", true
	case 4:
		return Milla, "millas", true
	case 5:
		return Pie, "pies", true
	case 6:
		return Pulgada, "pulgadas", true
	default:
		return 0, "", false
	}
}

func ConverLogitud() {
	for {
		entrada.LimpiarPantalla()
		mostrarMenu()
		// 1. Seleccion de la unidad de origen
		opcionL := entrada.LeerOpcion("Selecciona una opcion (1-8)")
		if opcionL == 7 {
			fmt.Println("Regresando al menu principal...")
			return
		}
		uOrigen, nombreOrigen, oKorigen := obtenerUnidades(opcionL)
		if !oKorigen {
			fmt.Printf("Opcion de origen no validad")
			entrada.Pausar()
			continue
		}

		// 2. Seleccion de la unidad de destino
		opcDestino := entrada.LeerOpcion("Selecciona la unida de Destino")
		uDestino, nombreDestino, oKdestino := obtenerUnidades(opcDestino)
		if !oKdestino {
			fmt.Printf("Opcion de destino no valida")
			entrada.Pausar()
			continue
		}

		// 3. Captura del valor a convertir
		valor := entrada.LeerNumero(fmt.Sprintf("Ingresa el valor en %s: ", nombreOrigen))
		resultado := Convertir(valor, uOrigen, uDestino)
		fmt.Printf("%.2f %s = %.2f %s\n", valor, nombreOrigen, resultado, nombreDestino)

		entrada.Pausar()
	}
}
