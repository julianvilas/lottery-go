package lottery

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const api = "http://api.elpais.com/ws/LoteriaNavidadPremiados"

func CheckNumber(n int) (int, error) {
	ns := strconv.Itoa(n)
	u := fmt.Sprintf("%s?n=%s", api, ns)

	resp, err := http.Get(u)
	if err != nil {
		return 0, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	elems := strings.Split(string(body), "=")
	if len(elems) != 2 || elems[0] != "busqueda" {
		return 0, fmt.Errorf("unknown response format: %s", body)
	}

	sr, err := decodeSearchResponse(elems[1])
	if err != nil {
		return 0, err
	}

	if sr.Error != 0 {
		return 0, fmt.Errorf("api returned error: %v", sr.Error)
	}

	if sr.Status < 0 || sr.Status > 4 {
		return 0, fmt.Errorf("api returned unknown status: %v", sr.Status)
	}

	if sr.Status == 0 {
		return 0, errors.New("raffle hasn't started yet")
	}

	return sr.Prize, nil
}

type searchResponse struct {
	Num       int `json:"numero"`
	Prize     int `json:"premio"`
	Timestamp int `json:"timestamp"`
	Status    int `json:"status"`
	Error     int `json:"error"`
}

func decodeSearchResponse(j string) (searchResponse, error) {
	var res searchResponse
	err := json.Unmarshal([]byte(j), &res)
	return res, err
}
