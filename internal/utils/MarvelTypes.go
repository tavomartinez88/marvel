package utils

type Hero struct {
	Code int `json:"code"`
	Data struct{
		Count int `json:"count"`
		Results [] struct{
			Id int64 `json:"id"`
			Name string `json:"name"`
		} `json:"results"`
	} `json:"data"`
}

type Comics struct {
	Data struct{
		Count int `json:"count"`
		Results []struct{
			Id int `json:"id"`
			Title string `json:"title"`
			Collaborators struct{
				Items []Collaborator
			} `json:"creators"`
			Characters struct{
				Items [] struct{
					Name string `json:"name"`
				} `json:"items"`
			} `json:"characters"`
		} `json:"results"`
	} `json:"data"`
}

type Collaborator struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type InfoHero struct {
	Name string `json:"name"`
}


