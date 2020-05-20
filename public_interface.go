package StitchingSQLGo

/*
Where field interface
return Where field string
example:
field_1
*/
type SQLField interface {
	SQLTable() string
	SQLField() string
}

/*
Where table interface
return Where table string
example:
table
*/
type SQLTable interface {
	SQLTable() string
}

/*
example:
insert into table (field) values (value)
*/
type Exec interface {
	Exec() (string, []interface{}, error)
}
