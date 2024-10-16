package types

type Settings struct {
	ProjectName  string `json:"Project_Name"`
	ProjectPath  string `json:"Project_Path"`
	ProjectDesc  string `json:"Project_Description"`
	ImgLink      string `json:"Image_Link"`
	OutputPath   string `json:"Output_Path"`
	IncludeTests bool   `json:"Include_Tests"`
	CapitalizeItems bool `json:CapitalizeItems`
}

type Error struct {
	Message string
	Filepath string
	Comment string
}

type CommentBlock struct {
	Filepath string
	Package  string
	Text     []string
}

type Package struct {
	Name  string
	Desc  string
	Usage string
	Deps  []Dependancy
	Files []File
	Types []Type
	Vars  []Variable
	Funcs []Function
}

type Dependancy struct {
	Name       string
	Desc       string
	Link       string
	ImportPath string
}

type File struct {
	Path    string
	Name    string
	Desc    string
	Auth    string
	Version string
	Date    string
	Funcs   []Function
	Vars    []Variable
	Types   []Type
}

type Type struct {
	Name     string
	Desc     string
	Fields   []Variable
	Exported bool
}

type Function struct {
	Name      string
	Desc      string
	Params    []Variable
	Returns   []ReturnValue
	Responses []Response
	Receiver  *Type
	Examples  []Example
	Exported  bool
}

type ReturnValue struct {
	Variable
	IsError bool // Flag indicating whether this return value is an error
}

type Example struct {
	Code string
	Desc string // Description or explanation of the code
}

type Variable struct {
	Name     string
	Type     string
	Desc     string
	Exported bool
}

type Response struct {
	Code int // eg. 404, 200, 204 ...
	Desc string
}
