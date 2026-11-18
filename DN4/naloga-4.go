package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"sync"
	"time"
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

func (i izdelek) obdelaj() error {

	if i.imeIzdelka == "" {
		return fmt.Errorf("Ime izdelka %+v je prazno.\n", i)
	}
	if i.cena <= 0 {
		return fmt.Errorf("Cena izdelka %+v je neveljavna\n", i)
	}
	if i.teza <= 0 {
		return fmt.Errorf("Teza izdelka %+v je  neveljavna\n", i)
	}

	lock.Lock()

	stNarocil++
	promet += i.cena

	fmt.Printf("Številka naročila: %d\n", stNarocil)
	fmt.Printf("Ime izdelka: %s\n", i.imeIzdelka)
	fmt.Printf("Cena: %.2f\n", i.cena)
	fmt.Printf("Teža: %.2f\n", i.teza)
	fmt.Println()

	lock.Unlock()

	return nil
}

func (k eknjiga) obdelaj() error {

	if k.naslovKnjige == "" {
		return fmt.Errorf("Naslov knjige %+v je prazen.\n", k)
	}
	if k.cena <= 0 {
		return fmt.Errorf("Cena knjige %+v je neveljavna.\n", k)
	}

	lock.Lock()

	stNarocil++
	promet += k.cena

	fmt.Printf("Številka naročila: %d\n", stNarocil)
	fmt.Printf("Naslov knjige: %s\n", k.naslovKnjige)
	fmt.Printf("Cena: %.2f\n", k.cena)
	fmt.Println()

	lock.Unlock()

	return nil
}

func (st spletniTecaj) obdelaj() error {

	if st.imeTecaja == "" {
		return fmt.Errorf("Ime tečaja %+v je prazno.\n", st)
	}
	if st.cenaUre <= 0 {
		return fmt.Errorf("Postavka tečaja %+v je neveljavna.\n", st)
	}
	if st.trajanjeUre <= 0 {
		return fmt.Errorf("Trajanje tečaja %+v je neveljavno.\n", st)
	}

	lock.Lock()

	stNarocil++
	promet += st.cenaUre * float64(st.trajanjeUre)

	fmt.Printf("Številka naročila: %d\n", stNarocil)
	fmt.Printf("Ime tečaja: %s\n", st.imeTecaja)
	fmt.Printf("Cena: %.2f\n", float64(st.trajanjeUre)*st.cenaUre)
	fmt.Println()

	lock.Unlock()

	return nil
}

var promet float64
var stNarocil int

var lock sync.Mutex
var wgProizvajalci sync.WaitGroup
var wgPorabniki sync.WaitGroup

const verjetnostNapake = 0.1

func vrednost(max int) float64 {
	val := rand.Float64() * float64(max)
	if rand.Float64() < verjetnostNapake {
		return -val
	}
	return val
}

func ime(predpona string) string {
	if rand.Float64() < verjetnostNapake {
		return ""
	}
	return predpona + strconv.Itoa(rand.IntN(100)+1)
}

func main() {

	narocila := make(chan narocilo, 10)
	stop := make(chan struct{})

	for i := 0; i < 5; i++ {
		wgProizvajalci.Add(1)
		go func() {
			defer wgProizvajalci.Done()
			for {
				var n narocilo

				switch rand.IntN(3) {
				case 0:
					n = izdelek{
						imeIzdelka: ime("Izdelek "),
						cena:       vrednost(100),
						teza:       vrednost(10),
					}
				case 1:
					n = eknjiga{
						naslovKnjige: ime("Knjiga "),
						cena:         vrednost(100),
					}
				case 2:
					n = spletniTecaj{
						imeTecaja:   ime("Tečaj "),
						cenaUre:     vrednost(10),
						trajanjeUre: int(vrednost(8)) + 1,
					}
				}
				time.Sleep(time.Millisecond * 500)

				select {
				case <-stop:
					return
				case narocila <- n:
				}
			}
		}()
	}

	for i := 0; i < 2; i++ {
		wgPorabniki.Add(1)
		go func() {
			defer wgPorabniki.Done()
			for trenutnoNarocilo := range narocila {
				err := trenutnoNarocilo.obdelaj()
				if err != nil {
					fmt.Println(err)
				}
				time.Sleep(time.Millisecond * 500)
			}

		}()
	}

	fmt.Scanln()
	close(stop)
	wgProizvajalci.Wait()
	close(narocila)
	wgPorabniki.Wait()

	fmt.Printf("\nŠtevilo naročil: %d\n", stNarocil)
	fmt.Printf("Promet: %.2f €\n", promet)

}
