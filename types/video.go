package types

type Video struct {
	ID        uint   `json:"id"`
	Url       string `json:"url"`
	ImageUrl  string `json:"image_url"`
	PlayerUrl string `json:"player_url"`
	Name      string `json:"name"`
	Kind      string `json:"kind"`
	Hosing    string `json:"hosting"`
}
