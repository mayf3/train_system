init sql:
	go run main.go orm syncdb -db=default -force=true -v=true

model construct:

Tables:
	Id				int			`orm:"pk;auto"`
	ContestName		string
	ProblemNumber	int			`orm:"null"`
	Source			string		`orm:"null"`
	Create_time		time.Time	`orm:"auto_now_add;type(datetime)"`
	Information		[]*Information	`orm:"reverse(many)"`

Information:
	Id				int			`orm:"pk;auto"`
	Table			*Tables		`orm:"rel(fk)"`
	Rank			int			`orm:"null"`
	info			string		`orm:"null"`

Member:
	Id		int		`orm:"pk;auto"`
	Name	string
	Grade	int
