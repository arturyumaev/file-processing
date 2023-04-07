package queries

import "fmt"

var selectFileInfoQuery = `
select
	id,
	filename,
	status,
	to_char(timestamp, '%s') as timestamp
from files
where timestamp = (
  select max(timestamp) from %s where filename = ?
)
`

var SelectFileInfo = fmt.Sprintf(selectFileInfoQuery, DATE_FORMAT, TABLE_USERS_NAME)
