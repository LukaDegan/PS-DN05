// Package vreme omogoča delo z vremenskimi postajami: dodajanje postaj in meritev
// ter izpis meritev in poročil o temperaturah.
package vreme

import "fmt"

// Postaja predstavlja vremensko postajo z imenom, krajem in seznamom meritev temperature.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo v slovar postaj doda novo postajo z dano oznako, imenom in krajem.
// Če postaja s to oznako že obstaja, jo prepiše.
func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	postaje[oznaka] = Postaja{ime: ime, kraj: kraj}
}

// IzpisMeritev izpiše meritve vseh postaj v slovarju.
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Printf("MERITVE:\n")
	for oznaka, postaja := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", oznaka, postaja.ime, postaja.kraj, postaja.meritve)
	}
}

// DodajMeritev postaji z dano oznako doda meritev temperature.
// Meritev je veljavna le, če leži med minTemp in maxTemp (vključno),
// sicer se izpiše opozorilo in meritev ni dodana.
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura, minTemp, maxTemp float64) {
	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Println("Meritev ni veljavna")
		return
	}
	postaja, obstaja := postaje[oznaka]
	if !obstaja {
		fmt.Println("Postaja ne obstaja")
		return
	}
	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja
}

// statistika vrne najnižjo, najvišjo in povprečno temperaturo postaje.
// Če postaja ne obstaja ali ima manj kot stMeritev meritev, vrne ok = false.
func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	postaja, obstaja := postaje[oznaka]
	if !obstaja || len(postaja.meritve) < stMeritev || len(postaja.meritve) == 0 {
		return 0, 0, 0, false
	}

	najnizja = postaja.meritve[0]
	najvisja = postaja.meritve[0]
	for _, meritev := range postaja.meritve {
		najvisja = max(najvisja, meritev)
		najnizja = min(najnizja, meritev)
		povprecna += meritev
	}
	povprecna /= float64(len(postaja.meritve))

	return najnizja, najvisja, povprecna, true
}

// IzpisPorocila za vsako postajo izpiše najnižjo, najvišjo in povprečno temperaturo
// ter opis vremena. Postaje z manj kot stMeritev meritvami so označene kot premalo meritev.
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	for oznaka, postaja := range postaje {
		najnizja, najvisja, povp, uredu := statistika(postaje, oznaka, stMeritev)
		if !uredu {
			fmt.Printf("%s: premalo meritev\n", postaja.ime)
			continue
		}
		var opis string
		switch {
		case povp >= 25:
			opis = "Vroče"
		case povp >= 15:
			opis = "Toplo"
		case povp >= 5:
			opis = "Hladno"
		default:
			opis = "Mrzlo"
		}
		fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n", postaja.ime, najnizja, najvisja, povp, opis)
	}
}
