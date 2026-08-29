package masa

import (
	entrada "calculator-cli/internal/Entrada"
	"fmt"
)

type unidadMasa float64

const (
	Kilogramo  unidadMasa = 1000.0
	Hectogramo unidadMasa = 100.0
	Decagramo  unidadMasa = 10.0
	Gramo      unidadMasa = 1
	Decigramo  unidadMasa = 0.1
	Centigramo unidadMasa = 0.01
	Miligramo  unidadMasa = 0.001
)

func ConvertirMasa(valor float64, de, a unidadMasa) float64 {
	masa := valor * float64(de)
	return masa / float64(a)
}

func mostrarMenuMasa() {
	fmt.Println("\n===== Convercion de Masa ======")
	fmt.Println("1. Gramo ")
	fmt.Println("2. Kilogramo ")
	fmt.Println("3. Hectogramo")
	fmt.Println("4. Decagramo")
	fmt.Println("5. Decigramo")
	fmt.Println("6. Centigramo")
	fmt.Println("7. Miligramo")
	fmt.Println("8. Regresar")
}

func obtenerUnidadesM(opcion int) (unidadMasa, string, bool) {
	switch opcion {
	case 1:
		return Gramo, "gramo", true
	case 2:
		return Kilogramo, "kilogramo", true
	case 3:
		return Hectogramo, "hectogramo", true
	case 4:
		return Decagramo, "decagramo", true
	case 5:
		return Decigramo, "decigramo", true
	case 6:
		return Centigramo, "centigramo", true
	case 7:
		return Miligramo, "miligramo", true
	default:
		return 0, "", false
	}
}

func ConverMasa() {
	for {
		entrada.LimpiarPantalla()
		mostrarMenuMasa()
		OpMasa := entrada.LeerOpcion("Selecioa una opcion (1-6)")
		if OpMasa == 8 {
			fmt.Println("Regresando al menu principal...")
			return
		}

		uOrigen, nombreOrigen, okOrigen := obtenerUnidadesM(OpMasa)
		if !okOrigen {
			fmt.Printf("Opcion de origen no valida")
			entrada.Pausar()
			continue
		}

		opcDestino := entrada.LeerOpcion("Seleciona la unidad del destino")
		uDestino, nombreDestino, okDestino := obtenerUnidadesM(opcDestino)
		if !okDestino {
			fmt.Printf("Opcion de destino no vvalida")
			entrada.Pausar()
			continue
		}

		valorM := entrada.LeerNumero(fmt.Sprintf("ingrese el valor en %s:", nombreOrigen))
		resultado := ConvertirMasa(valorM, uOrigen, uDestino)
		fmt.Printf("%.2f %s = %.2f %s\n", valorM, nombreOrigen, resultado, nombreDestino)
		entrada.Pausar()
	}
}
