package main

import (
	"fmt"
	"sync"
)

type narocilo interface {
	obdelaj() error
}

type izdelek struct {
	imeIzdelka string
	cena       float64
	teza       float64
}

type eknjiga struct {
	naslovKnjige string
	cena         float64
}

type spletniTecaj struct {
	imeTecaja   string
	trajanjeUre int
	cenaUre     float64
}

var promet float64 = 0
var stNarocil int = 0

var lock sync.Mutex

/*
Številka naročila: 1
Ime izdelka: Prenosni računalnik
Cena: 2000 €
Teža: 2.5 kg
*/

func (i izdelek) obdelaj() error {

	if i.imeIzdelka == "" {
		return fmt.Errorf("ime izdelka %+v ne sme biti prazno ", i)
	}
	if i.cena <= 0 {
		return fmt.Errorf("cena izdelka %q mora biti pozitivna, podano: %v", i.imeIzdelka, i.cena)
	}
	if i.teza <= 0 {
		return fmt.Errorf("teza izdelka %q mora biti pozitivna, podano: %v", i.imeIzdelka, i.teza)
	}

	lock.Lock()
	stNarocil++
	promet += i.cena

	fmt.Printf("Številka naročila: %d\n", stNarocil)
	fmt.Printf("Ime izdelka: %s\n", i.imeIzdelka)
	fmt.Printf("Cena: %v €\n", i.cena)
	fmt.Printf("Teža: %v kg\n", i.teza)
	fmt.Println()

	lock.Unlock()

	return nil
}

func (k eknjiga) obdelaj() error {

	if k.naslovKnjige == "" {
		return fmt.Errorf("naslov knjige %+v ne sme biti prazen", k)
	}
	if k.cena <= 0 {
		return fmt.Errorf("cena knjige %q mora biti pozitivna, podano %v", k.naslovKnjige, k.cena)
	}

	lock.Lock()
	stNarocil++
	promet += k.cena

	fmt.Printf("Številka naročila: %d\n", stNarocil)
	fmt.Printf("Naslov knjige: %s\n", k.naslovKnjige)
	fmt.Printf("Cena: %v €\n", k.cena)
	fmt.Println()

	lock.Unlock()

	return nil
}

func (st spletniTecaj) obdelaj() error {
	if st.imeTecaja == "" {
		return fmt.Errorf("naslov spletnega tečaja %+v ne sme biti prazen", st)
	}
	if st.cenaUre <= 0 {
		return fmt.Errorf("postavka spletnega tečaja %q mora biti pozitivna, podano: %v", st.imeTecaja, st.cenaUre)
	}
	if st.trajanjeUre <= 0 {
		return fmt.Errorf("trajanje spletnega tečaja %q mora biti pozitivno, podano: %v", st.imeTecaja, st.trajanjeUre)
	}

	lock.Lock()
	stNarocil++
	promet += st.cenaUre * float64(st.trajanjeUre)

	fmt.Printf("Številka naročila: %d\n", stNarocil)
	fmt.Printf("Ime tečaja: %s\n", st.imeTecaja)
	fmt.Printf("Cena tečaja: %v €\n", st.cenaUre*float64(st.trajanjeUre))
	fmt.Printf("Trajanje tečaja: %v ur\n", st.trajanjeUre)
	fmt.Println()

	lock.Unlock()

	return nil
}

func main() {

	izdelek1 := izdelek{"Lonec", 20.5, 1.5}
	izdelek2 := izdelek{"Krpa", 3, 0.2}
	izdelek3 := izdelek{"Kladivo", 37, 2.5}
	izdelek4 := izdelek{"", 45.5, 0.9}

	eknjiga1 := eknjiga{"Naučite se kuhati!", 24}
	eknjiga2 := eknjiga{"Uspavanke", -12}
	eknjiga3 := eknjiga{"Slikarstvo v renesansi", 35.5}

	spletniTecaj1 := spletniTecaj{"Kuharski tečaj", 8, 10}
	spletniTecaj2 := spletniTecaj{"Kako spečemo kruh", 3, 8}
	spletniTecaj3 := spletniTecaj{"Programiranje v GO", 99, 9.99}

	narocila := []narocilo{izdelek1, izdelek2, izdelek3, izdelek4, eknjiga1, eknjiga2, eknjiga3, spletniTecaj1, spletniTecaj2, spletniTecaj3}

	var wg sync.WaitGroup
	napake := make(chan error, len(narocila))

	for _, v := range narocila {
		wg.Add(1)
		go func() {
			defer wg.Done()
			e := v.obdelaj()
			if e != nil {
				napake <- e
			}
		}()
	}

	wg.Wait()
	close(napake)

	fmt.Println("Napake:")

	for n := range napake {
		fmt.Println(n)
	}

	fmt.Printf("\nŠtevilo naročil: %d\n", stNarocil)
	fmt.Printf("Promet: %.2f €\n", promet)

}
