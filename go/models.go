package main

type Repository struct { // struct for the repository object
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	FullName       string   `json:"full_name"`
	Private        bool     `json:"private"`
	Owner          Owner    `json:"owner"`
	Description    string   `json:"description"`
	Homepage       string   `json:"homepage"`
	Language       string   `json:"language"`
	HTML_URL       string   `json:"html_url"`
	Languages_URL  string   `json:"languages_url"`
	Git_URL        string   `json:"git_url"`
	SSH_URL        string   `json:"ssh_url"`
	Clone_URL      string   `json:"clone_url"`
	Forks_Count    int      `json:"forks_count"`
	Visibility     string   `json:"visibility"`
	Forks          int      `json:"forks"`
	Default_Branch string   `json:"default_branch"`
	Topics         []string `json:"topics"`
}

type Owner struct { // struct for the owner object
	Login     string `json:"login"`
	ID        int    `json:"id"`
	URL       string `json:"url"`
	HTML_URL  string `json:"html_url"`
	Repos_URL string `json:"repos_url"`
	Type      string `json:"type"`
}
