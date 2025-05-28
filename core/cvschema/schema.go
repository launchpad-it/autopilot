package cvschema

// Resume represents a structured CV in JSON format.
// See: https://jsonresume.org/schema.
type Resume struct {
	Basics       Basics        `json:"basics"`
	Work         []Work        `json:"work" validate:"omitempty,dive,required"`
	Volunteer    []Volunteer   `json:"volunteer" validate:"omitempty,dive,required"`
	Education    []Education   `json:"education" validate:"omitempty,dive,required"`
	Awards       []Award       `json:"awards" validate:"omitempty,dive,required"`
	Publications []Publication `json:"publications" validate:"omitempty,dive,required"`
	Skills       []Skill       `json:"skills" validate:"omitempty,dive,required"`
	Languages    []Language    `json:"languages" validate:"omitempty,dive,required"`
	Interests    []Interest    `json:"interests" validate:"omitempty,dive,required"`
	References   []Reference   `json:"references" validate:"omitempty,dive,required"`
	Projects     []Project     `json:"projects" validate:"omitempty,dive,required"`
	Meta         *Meta         `json:"meta"`
}

type Basics struct {
	Name     string    `json:"name"`
	Label    string    `json:"label"`
	Image    string    `json:"image" validate:"omitempty,url"`
	Email    string    `json:"email" validate:"omitempty,email"`
	Phone    string    `json:"phone"`
	URL      string    `json:"url" validate:"omitempty,url"`
	Summary  string    `json:"summary"`
	Location *Location `json:"location"`
	Profiles []Profile `json:"profiles" validate:"omitempty,dive,required"`
}

type Location struct {
	Address     string `json:"address"`
	PostalCode  string `json:"postalCode"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
}

type Profile struct {
	Network  string `json:"network"`
	UserName string `json:"username"`
	URL      string `json:"url" validate:"omitempty,url"`
}

type Work struct {
	Name        string   `json:"name"`
	Location    string   `json:"location"`
	Description string   `json:"description"`
	Position    string   `json:"position"`
	URL         string   `json:"url" validate:"omitempty,url"`
	StartDate   string   `json:"startDate" validate:"omitempty,date"`
	EndDate     string   `json:"endDate" validate:"omitempty,date"`
	Summary     string   `json:"summary"`
	Highlights  []string `json:"highlights" validate:"omitempty,dive,required"`
}

type Volunteer struct {
	Organization string   `json:"organization"`
	Position     string   `json:"position"`
	URL          string   `json:"url" validate:"omitempty,url"`
	StartDate    string   `json:"startDate" validate:"omitempty,date"`
	EndDate      string   `json:"endDate" validate:"omitempty,date"`
	Summary      string   `json:"summary"`
	Highlights   []string `json:"highlights" validate:"omitempty,dive,required"`
}

type Education struct {
	Institution string   `json:"institution"`
	Area        string   `json:"area"`
	StudyType   string   `json:"studyType"`
	StartDate   string   `json:"startDate" validate:"omitempty,date"`
	EndDate     string   `json:"endDate" validate:"omitempty,date"`
	GPA         string   `json:"gpa"`
	Courses     []string `json:"courses" validate:"omitempty,dive,required"`
}

type Award struct {
	Title   string `json:"title"`
	Date    string `json:"date" validate:"omitempty,date"`
	Awarder string `json:"awarder"`
	Summary string `json:"summary"`
}

type Publication struct {
	Name        string `json:"name"`
	Publisher   string `json:"publisher"`
	ReleaseDate string `json:"releaseDate" validate:"omitempty,date"`
	URL         string `json:"url" validate:"omitempty,url"`
	Summary     string `json:"summary"`
}

type Skill struct {
	Name     string   `json:"name"`
	Level    string   `json:"level"`
	Keywords []string `json:"keywords" validate:"omitempty,dive,required"`
}

type Language struct {
	Language string `json:"language"`
	Fluency  string `json:"fluency"`
}

type Interest struct {
	Name     string   `json:"name"`
	Keywords []string `json:"keywords" validate:"omitempty,dive,required"`
}

type Reference struct {
	Name      string `json:"name"`
	Reference string `json:"reference"`
}

type Project struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Highlights  []string `json:"highlights" validate:"omitempty,dive,required"`
	Keywords    []string `json:"keywords" validate:"omitempty,dive,required"`
	StartDate   string   `json:"startDate" validate:"omitempty,date"`
	EndDate     string   `json:"endDate" validate:"omitempty,date"`
	URL         string   `json:"url" validate:"omitempty,url"`
	Roles       []string `json:"roles" validate:"omitempty,dive,required"`
	Entity      string   `json:"entity"`
	Type        string   `json:"type"`
}

type Meta struct {
	Canonical    string `json:"canonical"`
	Version      string `json:"version" validate:"omitempty,semver"`
	LastModified string `json:"lastModified"`
}
