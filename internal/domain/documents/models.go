package documents

type Document struct {
	ID          string
	Title       string
	Description string
	Category    string
	Tags        []string
	UpdatedAt   string
	ContentB64  string
	CoverImage  string
}
