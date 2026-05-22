package types

var InboxProject = Project{
	ID:          "main",
	Name:        "main",
	Description: "The default project in tuido",
	Color:       DefaultColor,
	SortOrder:   0,
}

var NullProject = Project{}

type Project struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Color       Color  `json:"color" db:"color"`
	SortOrder   int64  `json:"sortOrder" db:"sort_order"`
}

type ProjectData struct {
	Project Project `json:"project"`
	Tasks   []Task  `json:"tasks"`
}
