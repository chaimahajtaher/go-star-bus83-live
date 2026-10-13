package main

import (
        "crypto/tls"
        "encoding/json"
        "fmt"
        "io"
        "net/http"
)

///////////////////////////////////////////////////////
// STRUCTURES JSON
///////////////////////////////////////////////////////

type APIResponse struct {
        Results []Passage `json:"results"`
}

type Passage struct {

        // Identifiant interne de ligne
        IDLigne string `json:"idligne"`

        // Numéro visible de la ligne
        NomCourtLigne string `json:"nomcourtligne"`

        // Destination
        Destination string `json:"destination"`

        // Arrêt
        NomArret string `json:"nomarret"`

        // Heure réelle d'arrivée
        Arrivee string `json:"arrivee"`
}

///////////////////////////////////////////////////////
// MAIN
///////////////////////////////////////////////////////

func main() {

        fmt.Println("===================================")
        fmt.Println(" BUS 83 EN TEMPS REEL ")
        fmt.Println("===================================")

        url := "https://data.rennesmetropole.fr/api/explore/v2.1/catalog/datasets/prochains-passages-des-lignes-de-bus-du-reseau-star-en-temps-reel/records?limit=20&refine=nomcourtligne%3A%2283%22"

        ///////////////////////////////////////////////////
        // CLIENT TLS
        ///////////////////////////////////////////////////

        client := &http.Client{
                Transport: &http.Transport{
                        TLSClientConfig: &tls.Config{
                                InsecureSkipVerify: true,
                        },
                },
        }

        ///////////////////////////////////////////////////
        // APPEL API
        ///////////////////////////////////////////////////

        resp, err := client.Get(url)

        if err != nil {
                fmt.Println("Erreur API :", err)
                return
        }

        defer resp.Body.Close()

        ///////////////////////////////////////////////////
        // LECTURE REPONSE
        ///////////////////////////////////////////////////

        body, err := io.ReadAll(resp.Body)

        if err != nil {
                fmt.Println("Erreur lecture :", err)
                return
        }

        ///////////////////////////////////////////////////
        // CONVERSION JSON
        ///////////////////////////////////////////////////

        var data APIResponse

        err = json.Unmarshal(body, &data)

        if err != nil {
                fmt.Println("Erreur JSON :", err)
                return
        }

        ///////////////////////////////////////////////////
        // RECHERCHE BUS 83
        ///////////////////////////////////////////////////

        nbPassages := 0

        for _, passage := range data.Results {

                if passage.NomCourtLigne == "83" {

                        nbPassages++

                        fmt.Println("-----------------------------------")
                        fmt.Println("Ligne       :", passage.NomCourtLigne)
                        fmt.Println("Arrêt       :", passage.NomArret)
                        fmt.Println("Destination :", passage.Destination)
                        fmt.Println("Arrivée     :", passage.Arrivee)
                        fmt.Println("-----------------------------------")
                        fmt.Println()
                }
        }

        fmt.Println("===================================")
        fmt.Println(" STATISTIQUES ")
        fmt.Println("===================================")

        fmt.Println("Nombre de passages trouvés :", nbPassages)

        if nbPassages == 0 {
                fmt.Println("Aucun passage de la ligne 83 trouvé dans les 20 premiers résultats.")
                fmt.Println("Cela est normal si aucun bus 83 n'est présent dans l'échantillon récupéré.")
        }
}
