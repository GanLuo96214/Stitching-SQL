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
