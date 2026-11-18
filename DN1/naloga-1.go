package main

import (
	"fmt"
	"slices"
)

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

/*
	dodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64), ki postaji z dano oznako doda novo meritev.
	Pred dodajanjem naj funkcija preveri, ali je temperatura v dovoljenem območju (med -50 °C in 50 °C).
	Če postaja ne obstaja ali meritev ni veljavna, naj funkcija izpiše ustrezno obvestilo.
*/

func dodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64) {
	postaja, ok := postaje[oznaka]

	if ok {
		if -50 < temperatura && temperatura < 50 {
			postaja.meritve = append(postaja.meritve, temperatura)
			postaje[oznaka] = postaja
		} else {
			fmt.Println("Temperatura v nedovoljenem območju.")
			return
		}
	} else {
		fmt.Println("Postaja ne obstaja.")
		return
	}
}

/*
statistika(postaje map[string]Postaja, oznaka string)
(najnizja, najvisja, povprecna float64, ok bool) vrne najnižjo, najvišjo in povprečno temperaturo postaje.
Če postaja ne obstaja ali ima manj kot 3 meritve, naj funkcija vrne ok == false.
*/

func statistika(postaje map[string]Postaja, oznaka string) (najnizja, najvisja, povprecna float64, ok bool) {
	postaja, ok := postaje[oznaka]

	if len(postaja.meritve) < 3 || !ok {
		ok = false
		return
	}

	najnizja = slices.Min(postaja.meritve)
	najvisja = slices.Max(postaja.meritve)

	st_meritev := len(postaja.meritve)
	sum := 0.0
	for i := 0; i < st_meritev; i++ {
		sum += postaja.meritve[i]
	}
	povprecna = sum / float64(st_meritev)

	return

}

/*
MERITVE:
KP01 - Koper Luka (Koper): [24.5 27 26.1 28.4]
KR01 - Kredarica (Triglav): [2 -1.5 3.5]
LJ01 - Ljubljana Bežigrad (Ljubljana): [18.2 21.5 19.8 16.5]
MB01 - Maribor Tabor (Maribor): [12.4]
*/

func izpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")

	for oznaka, postaja := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", oznaka, postaja.ime, postaja.kraj, postaja.meritve)
	}
}

/*
Koper Luka: min 24.5 °C, max 28.4 °C, povprečje 26.5 °C -> Vroče
Kredarica: min -1.5 °C, max 3.5 °C, povprečje 1.3 °C -> Mrzlo
Ljubljana Bežigrad: min 16.5 °C, max 21.5 °C, povprečje 19.0 °C -> Toplo
Maribor Tabor: premalo meritev

če je povprečje >= 25 °C, naj izpiše Vroče;
če je povprečje med 15 °C in 25 °C, naj izpiše Toplo;
če je povprečje med 5 °C in 15 °C, naj izpiše Hladno;
če je povprečje < 5 °C, naj izpiše Mrzlo.
*/

func izpisPorocila(postaje map[string]Postaja) {
	for naziv, postaja := range postaje {
		najnizja, najvisja, povprecna, ok := statistika(postaje, naziv)

		var t string

		switch {
		case povprecna >= 25:
			t = "Vroče"
		case povprecna >= 15:
			t = "Toplo"
		case povprecna >= 5:
			t = "Hladno"
		default:
			t = "Mrzlo"
		}

		if ok {
			fmt.Printf("%s: min %.1f ˚C, max %.1f ˚C, povprečje %.1f ˚C -> %s\n", postaja.ime, najnizja, najvisja, povprecna, t)
		} else {
			fmt.Printf("%s: premalo meritev\n", postaja.ime)
		}

	}
}

func main() {

	var postaje map[string]Postaja
	postaje = make(map[string]Postaja)

	/*
		MERITVE:
		KP01 - Koper Luka (Koper): [24.5 27 26.1 28.4]
		KR01 - Kredarica (Triglav): [2 -1.5 3.5]
		LJ01 - Ljubljana Bežigrad (Ljubljana): [18.2 21.5 19.8 16.5]
		MB01 - Maribor Tabor (Maribor): [12.4]
	*/

	var KR01 Postaja
	KR01.ime = "Kredarica"
	KR01.kraj = "Triglav"
	KR01.meritve = []float64{2, -1.5, 3.5}

	postaje["KR01"] = KR01

	postaje["KP01"] = Postaja{"Koper Luka", "Koper", []float64{24.5, 27, 26.1, 28.4}}
	postaje["LJ01"] = Postaja{"Ljubljana Bežigrad", "Ljubljana", []float64{18.2, 21.5, 19.8, 16.5}}
	postaje["MB01"] = Postaja{"Maribor Tabor", "Maribor", []float64{12.4}}

	fmt.Println(postaje["LJ01"])

	dodajMeritev(postaje, "LJ01", 14)

	fmt.Println(postaje["LJ01"])

	// neveljavne
	dodajMeritev(postaje, "LJ02", 14)
	dodajMeritev(postaje, "LJ01", 52)

	//fmt.Println(postaje)

	fmt.Println(statistika(postaje, "LJ01"))
	fmt.Println(statistika(postaje, "LJ02"))

	izpisMeritev(postaje)

	fmt.Println()
	izpisPorocila(postaje)

}
