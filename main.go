package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/LukaDegan/PS-DN05/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "vreme",
		Usage: "Upravljanje vremenskih postaj in meritev temperature",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "minTemp",
				Usage: "Najnižja dovoljena temperatura meritve v °C",
				Value: -50,
			},
			&cli.Float64Flag{
				Name:  "maxTemp",
				Usage: "Najvišja dovoljena temperatura meritve v °C",
				Value: 50,
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Usage: "Najmanjše število meritev za izračun statistike",
				Value: 3,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float64("minTemp")
			maxTemp := cmd.Float64("maxTemp")
			stMeritev := cmd.Int("stMeritev")
			zagon(minTemp, maxTemp, stMeritev)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func zagon(minTemp, maxTemp float64, stMeritev int) {
	postaje := make(map[string]vreme.Postaja)
	vreme.DodajPostajo(postaje, "KP01", "Koper Luka", "Koper")
	vreme.DodajPostajo(postaje, "KR01", "Kredarica", "Triglav")
	vreme.DodajPostajo(postaje, "LJ01", "Ljubljana Bežigrad", "Ljubljana")
	vreme.DodajPostajo(postaje, "MB01", "Maribor Tabor", "Maribor")

	for _, m := range []float64{24.5, 27, 26.1, 28.4} {
		vreme.DodajMeritev(postaje, "KP01", m, minTemp, maxTemp)
	}
	for _, m := range []float64{2, -1.5, 3.5} {
		vreme.DodajMeritev(postaje, "KR01", m, minTemp, maxTemp)
	}
	for _, m := range []float64{18.2, 21.5, 19.8, 16.5} {
		vreme.DodajMeritev(postaje, "LJ01", m, minTemp, maxTemp)
	}
	vreme.DodajMeritev(postaje, "MB01", 12.4, minTemp, maxTemp)

	fmt.Println("| Začetno stanje |")
	vreme.IzpisMeritev(postaje)

	fmt.Println("\n| Dodana meritev obstoječi postaji (KP01) |")
	vreme.DodajMeritev(postaje, "KP01", 43.21, minTemp, maxTemp)
	vreme.IzpisMeritev(postaje)

	fmt.Println("\n| Meritev pod minTemp |")
	vreme.DodajMeritev(postaje, "KP01", minTemp-0.1, minTemp, maxTemp)

	fmt.Println("\n| Meritev nad maxTemp |")
	vreme.DodajMeritev(postaje, "KP01", maxTemp+0.1, minTemp, maxTemp)

	fmt.Println("\n| Neobstoječa postaja |")
	vreme.DodajMeritev(postaje, "AB12", 25, minTemp, maxTemp)

	fmt.Println("\n| Poročilo |")
	vreme.IzpisPorocila(postaje, stMeritev)

}
