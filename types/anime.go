package types

type Anime struct {
	ID                 uint                 `json:"id"`
	Name               string               `json:"name"`
	Russian            string               `json:"russian"`
	Image              Image                `json:"image"`
	Url                string               `json:"url"`
	Kind               string               `json:"kind"`
	Score              string               `json:"score"`
	Status             string               `json:"status"`
	Episodes           uint                 `json:"episodes"`
	EpisodesAired      uint                 `json:"episodes_aired"`
	AiredOn            string               `json:"aired_on"`
	ReleasedOn         string               `json:"released_on"`
	Rating             string               `json:"rating"`
	English            []string             `json:"english"`
	Japanese           []string             `json:"japanese"`
	Synonyms           []string             `json:"synonyms"`
	LicenseNameRu      string               `json:"license_name_ru"`
	Duration           uint                 `json:"duration"`
	Description        string               `json:"description"`
	DescriptionHtml    string               `json:"description_html"`
	DescriptionSource  *string              `json:"description_source"`
	Franchise          string               `json:"franchise"`
	Favoured           bool                 `json:"favoured"`
	Anons              bool                 `json:"anons"`
	Ongoing            bool                 `json:"ongoing"`
	ThreadID           uint                 `json:"thread_id"`
	TopicID            uint                 `json:"topic_id"`
	MyanimelistID      uint                 `json:"myanimelist_id"`
	RatesStatutesStats []RatesStatutesStats `json:"rates_statuses_stats"`
	UpdatedAt          string               `json:"updated_at"`
	NextEpisodeAt      *string              `json:"next_episode_at"`
	Fansubbers         []string             `json:"fansubbers"`
	Fandubbers         []string             `json:"fandubbers"`
	Licensors          []string             `json:"licensors"`
	Genres             []Genre              `json:"genres"`
	Studios            []Studio             `json:"studios"`
	Videos             []Video              `json:"videos"`
	Screenshots        []Screenshot         `json:"screenshots"`
}
