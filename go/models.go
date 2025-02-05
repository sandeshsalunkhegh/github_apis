package main

type Repository struct { // struct for the repository object
	ID          int    `json:"id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Private     bool   `json:"private"`
	Owner       Owner  `json:"owner"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Language    string `json:"language"`
}

type Owner struct { // struct for the owner object
	Login string `json:"login"`
	ID    int    `json:"id"`
}
