package lottery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const api = "http://api.elpais.com/ws/LoteriaNavidadPremiados"

func CheckNumbers(numbers ...int) (res []SearchResponse, err error) {
	for _, n := range numbers {
		ns := strconv.Itoa(n)
		u := fmt.Sprintf("%s?n=%s", api, ns)

		resp, err := http.Get(u)
		if err != nil {
			return nil, err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		elems := strings.Split(string(body), "=")
		if len(elems) != 2 || elems[0] != "busqueda" {
			return nil, fmt.Errorf("unknown response format: %s", body)
		}

		sr, err := decodeSearchResponse(elems[1])
		if err != nil {
			return nil, err
		}

		res = append(res, sr)
	}

	return res, nil
}

type SearchResponse struct {
	Num       int `json:"numero"`
	Prize     int `json:"premio"`
	Timestamp int `json:"timestamp"`
	Status    int `json:"status"`
	Error     int `json:"error"`
}

func decodeSearchResponse(j string) (SearchResponse, error) {
	var res SearchResponse
	err := json.Unmarshal([]byte(j), &res)
	return res, err
}
