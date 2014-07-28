package vcs

type PostReceiveHook struct {
	Before  string      `json:"before"`
	After   string      `json:"after"`
	Ref     string      `json:"ref"`
	Repo    *CommitRepo `json:"repository"`
	Commits []*Commit   `json:"commits"`
	Head    *Commit     `json:"head_commit"`
	Deleted bool        `json:"deleted"`
}

type CommitRepo struct {
	Url   string `json:"url"`
	Name  string `json:"name"`
	Desc  string `json:"description"`
	Owner *Owner `json:"owner"`
}

type Commit struct {
	Id        string   `json:"id"`
	Url       string   `json:"url"`
	Message   string   `json:"message"`
	Timestamp string   `json:timestamp`
	Author    *Author  `json:"author"`
	Added     []string `json:"added"`
}

// Owner represents the owner of a Github Repository.
type Owner struct {
	Type  string `json:"type"`
	Login string `json:"login"`
	Name  string `json:"name"`
}
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
