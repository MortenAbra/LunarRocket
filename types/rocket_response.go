package types

type RocketResponse struct {
	Channel string `json:"channel"`
	Type    string `json:"type"`
	Mission string `json:"mission"`
	Status  string `json:"status"`
	Speed   int    `json:"speed`
}

func MapToRocketResponse(models []RocketModel) []RocketResponse {
	rocketResponse := make([]RocketResponse, 0)

	for _, v := range models {
		if v.Status == "" {
			v.Status = "Ongoing"
		}
		rocketResponse = append(rocketResponse, RocketResponse{
			Channel: v.Channel,
			Type:    v.Type,
			Mission: v.Mission,
			Status:  v.Status,
			Speed:   v.Speed,
		})
	}
	return rocketResponse
}
