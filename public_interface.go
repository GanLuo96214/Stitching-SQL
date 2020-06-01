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
select field1 from table1
select exists(select field1 from table1)
*/
type Query interface {
	Query() (string, []interface{}, error)
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
insert into table (field1) values (value1) returning field1
*/
type ExecWithReturning interface {
	Exec
	ExecWithReturning() Returning
}
