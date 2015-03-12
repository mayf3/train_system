create table if not exists member (
    name char(20) not null,
    primary key (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table if not exists tables (
    contest_name varchar(200) not null,
    problem_number int(30),
    source varchar(200),
    create_time datetime not null,
    table_id int(30) unsigned not null AUTO_INCREMENT,
    primary key (table_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table if not exists infomation (
    info_id int(30) unsigned not null AUTO_INCREMENT,
    table_id int(30) unsigned not null,
    rank int(2) unsigned not null,
    info char(250) not null,
    primary key (info_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
