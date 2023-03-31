package model

type Job struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	Company     string `json:"company"`
	CompanyURL  string `json:"company_url"`
	Location    string `json:"location"`
	Title       string `json:"title"`
	Description string `json:"Description"`
	HowToApply  string `json:"how_to_apply"`
	CompanyLogo string `json:"company_logo"`
	CreatedAt   string `json:"created_at"`
}

type ParamsJob struct {
	Description string `json:"description" form:"description"`
	Location    string `json:"location" form:"location"`
	Type        string `json:"type" form:"type"`
	Page        int    `json:"page" form:"page"`
}
