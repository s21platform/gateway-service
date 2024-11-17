package option

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	optionhub "github.com/s21platform/optionhub-proto/optionhub-proto"
)

type Usecase struct {
	oC OptionClient
}

func New(oC OptionClient) *Usecase {
	return &Usecase{oC: oC}
}

func (uc *Usecase) GetOS(r *http.Request) (*optionhub.GetByIdOut, error) {
	id, err := extractOSID(r)
	if err != nil {
		return nil, fmt.Errorf("failed to extract os id: %w", err)
	}

	resp, err := uc.oC.GetOSById(r.Context(), id)
	if err != nil {
		return nil, fmt.Errorf("failed to get os in usercase: %w", err)
	}

	return resp, nil
}

func extractOSID(r *http.Request) (int64, error) {
	var requestData struct {
		ID int64 `json:"id"`
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read request body: %w", err)
	}
	defer r.Body.Close()

	if len(body) == 0 {
		return 0, fmt.Errorf("request body is empty")
	}

	if err := json.Unmarshal(body, &requestData); err != nil {
		return 0, fmt.Errorf("failed to decode request body: %w", err)
	}

	return requestData.ID, nil
}
