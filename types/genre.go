package types

type Genre struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Russian   string `json:"russian"`
	Kind      string `json:"kind"`
	EntryType string `json:"entry_type"`
}
