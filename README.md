init sql:
	go run main.go orm syncdb -db=default -force=true -v=true

model construct:

type Tables struct {
	Id				int `orm:"pk;auto"`
	ContestName		string
	ProblemNumber	int
	Source			string
	CreateTime		time.Time      `orm:"auto_now_add;type(datetime)"`
	Information		[]*Information `orm:"reverse(many)"`
}

type Information struct {
	Id		int			`orm:"pk;auto"`
	Rank	int
	Table	*Tables		`orm:"rel(fk)"`
	Member	[]*Member	`orm:"reverse(many)"`
	Problem []*Problem	`orm:"reverse(many)"`
}

type Member struct{
	Id			int				`orm:"pk;auto"`
	Order		int
	Person		*Person			`orm:"rel(fk)"`
	Information	*Information	`orm:"rel(fk)"`
}

type Problem struct {
	Id			int				`orm:"pk;auto"`
	Number		int
	Participant	int
	Status		int
	Information	*Information	`orm:"rel(fk)"`
}

type Person struct{
	Id		int		`orm:"pk;auto"`
	Name	string
	Grade	int
	Member	[]*Member	`orm:"reverse(many)"`
}
