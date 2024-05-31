package types

import (
	"encoding/json"

	"gofr.dev/pkg/gofr/http"
)

type ExoPlanetType string

const (
	GasGiant    ExoPlanetType = "GasGiant"
	Terrestrial ExoPlanetType = "Terrestrial"
)

func (e *ExoPlanetType) UnmarshalJSON(b []byte) error {
	var s string

	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	switch s {
	case string(GasGiant):
		*e = GasGiant
		return nil
	case string(Terrestrial):
		*e = Terrestrial
		return nil
	default:
		return http.ErrorInvalidParam{Params: []string{"ExoPlanet"}}

	}
}
