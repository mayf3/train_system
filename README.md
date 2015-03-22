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

MySQL 中文输入问题：
	my.cnf中设置
		[client]
		default-character-set = utf8
		[mysql]
		default-character-set = utf8
		[mysqld]
		default-character-set = utf8
	MySQL 5.5之后[mysqld]段设置要变更：
		[mysqld]
		default-storage-engine = INNODB
		character-set-server = utf8
		collation-server = utf8_general_ci
