package main

import (
    "fmt"
)

const (
    scontoBambini = 50
    scontoGiovani = 30
    scontoAdulti  = 20
    etaMassima = 99
)

// calcolaSconto restituisce lo sconto in base all'età fornita.
func calcolaSconto(eta int) int {
  switch {
  case eta >= 0 && eta <= 9 || eta >= 76:
    return scontoBambini
  case eta >= 10 && eta <= 16 || eta >= 60 && eta <= 75:
    return scontoGiovani
  case eta >= 17 && eta <= 25 || eta >= 46 && eta <= 59:
    return scontoAdulti
  default:
    return scontoAdulti // Sconto predefinito
  }
}


    func main() {
        var eta int
        fmt.Printf("Inserisci l'età della persona: ")
        _, err := fmt.Scan(&eta)

        if err != nil || eta < 0 || eta > etaMassima {
            fmt.Printf("Inserimento non valido. Inserisci un'età compresa tra 0 e %d.\n", etaMassima)
            return
        }

        sconto := calcolaSconto(eta)
        fmt.Printf("Lo sconto associato all'età %d è del %d%%\n", eta, sconto)
    }
  
