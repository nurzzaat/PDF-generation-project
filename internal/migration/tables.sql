create table users(
	id serial primary key,
	email text default '',
	password text default '',
	firstname text default '',
	lastname text default '',
	createdat text default ''
);

create table syllabus(
	id serial primary key,
	userid int,
	subject text default '',
	faculty text default '',
	kafedra text default '' ,
	specialist text default '',
	coursenumber int default 1,
	creditnumber int default 0,
	allhours int default 0,
	lecturehour int default 0,
	practicehour int default 0,
	sro int default 0,
	madeBy text default '',
	madeByMajor text default '',
	discuss1  text default '',
	discussBy1 text default '',
	discussBy1Major text default '',
discuss2  text default '',
	discussBy2 text default '',
	discussBy2Major text default '',
confirmedBy text default '',
	confirmedByMajor text default '',
foreign key (userid) references users(id)

);

create table modules(
	syllabusid int ,
	id serial primary key,
	orderid int default 1,
	title text default '',
	foreign key(syllabusid) references syllabus(id)
);

create table topic(
	id serial primary key,
	moduleid int ,
	title  text default '',
	lk int ,
	spz int,
	sro int,
	literature text default '',
	foreign key (moduleid)references modules(id)
);

create table literature(
	syllabusid int,
	id serial primary key,
	type text default 'main',
	title text default '',
	foreign key(syllabusid) references syllabus(id)
);

alter table syllabus 
add column srop text default '';

create table questions(
	id serial primary key,
	syllabusid int,
	title text default '',
	foreign key (syllabusid) references syllabus(id) on delete cascade
);

	alter table syllabus
	add column facultyOfProf text default '';
	alter table syllabus
	add column email text default '';
	alter table syllabus
	add column address text default '';
	alter table syllabus
	add column timeofcons text default '';
	alter table questions
add column sequen text default '';