package campaign

type CreateCampaign struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Emails  []string `json:"emails"`
}

type Campaign struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Emails    []string `json:"emails"`
	CreatedAt string   `json:"created_at"`
}
