package handlers

import "encoding/json"

func extractResponseCT(data any) (*configTablePayload, error) {
	var ct configTablePayload

	myData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(myData, &ct)
	if err != nil {
		return nil, err
	}

	return &ct, nil
}

func extractResponseDeleteCT(data any) (*ctEntriesPayload, error) {
	var ct ctEntriesPayload

	myData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(myData, &ct)
	if err != nil {
		return nil, err
	}

	return &ct, nil
}
