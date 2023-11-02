package domain

type AddressValidatorReq struct {
	Address struct {
		AddressLines string `json:"addressLines"`
	} `json:"address"`
}

type AddressValidatorResponse struct {
	Result struct {
		Address struct {
			AddressComponents []struct {
				ComponentName struct {
					Text         string `json:"text"`
					LanguageCode string `json:"languageCode"`
				} `json:"componentName"`
				ComponentType     string `json:"componentType"`
				ConfirmationLevel string `json:"confirmationLevel"`
				Inferred          bool   `json:"inferred,omitempty"`
			} `json:"addressComponents"`
			MissingComponentTypes []string `json:"missingComponentTypes"`
		} `json:"address"`
		Geocode struct {
			Location struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"location"`
		} `json:"geocode"`
	} `json:"result"`
}
