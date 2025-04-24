package campaign

type CreateCampaign struct{
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	Emails    []string 		  `json:"emails"`
}