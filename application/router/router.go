package route

import (
	"bufio"
	"errors"
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

//criando uma estrutura de uma rota
type Route struct {
	ID string `json:"routeId"`
	ClientID string `json:"clientId"`
	Positions []Position `json:"position"`
}

type Position struct {
	Lat float64 `json:"lat"`
	Long float64 `json:"long"`
}

type PartialRoutePosition struct {
	ID string 			`json:"routeId"`
	ClientID string 	`json:"clientId"`
	Position []float64 	`json:"position"`
	Finished bool 		`json:"finished"`
}

//criação de metodo passando um ponteiro
func(r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("route id not informed")
	}

	f, err := os.Open("destinations/"+ r.ID + ".txt")

	if err != nil {
		return err
	}
	//fecha o arquivo
	defer f.Close()

	//fazendo a leitura do arquivo aberto
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		//pegar a latitude e transformar em float
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}

		long, err := strconv.ParseFloat(data[1] ,64)
		if err != nil {
			return err
		}

		//add as positions
		r.Positions = append(r.Positions, Position {
			Lat: lat,
			Long: long,
		})
	}

	return nil

}

func(r *Route) ExportJsonPositions() ([]string, error) {

	var route PartialRoutePosition
	var result []string

	total := len(r.Positions)

	for k, v := range r.Positions {
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position =  []float64{v.Lat, v.Long}
		route.Finished = false

		if total-1  == k {
			route.Finished = true
		}

		jsonRoute, err := json.Marshal(route)

		if err != nil {
			return nil, err
		}

		result = append(result, string(jsonRoute))
	}

	return result, nil
}