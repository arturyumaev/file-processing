package queries

var SelectFileInfo = `
select
	id,
	filename,
	status,
	to_char(timestamp, 'DD.MM.YYYY HH24:MI:SS GMTOF') as timestamp
from files
where timestamp = (
  select max(timestamp) from files where filename = ?
)
`

var InsertFile = `
insert into files (status, filename)
  values
    ('recieved', ?)
`
