package StitchingSQLGo

/*
Where field interface
write field into SqlBuilder string builder
example:
field_1 or table_1.field
*/
type Field interface {
	Field(*SqlBuilder, bool) error
}

/*
Where table interface
write table into SqlBuilder string builder
example:
table
*/
type Table interface {
	Table(*SqlBuilder) error
}

/*
example:
insert into table (field) values (value)
*/
type Exec interface {
	Exec() (string, []interface{}, error)
}

/*
example:
insert into table (field) values (value)
*/
type ExecWithReturning interface {
	Exec
	Returning() Returning
}
