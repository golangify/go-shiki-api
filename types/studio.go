package types

type Studio struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	FilteredName string `json:"filtered_name"`
	Real         bool   `json:"real"`
	Image        string `json:"image"`
}
