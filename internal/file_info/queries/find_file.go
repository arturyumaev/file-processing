package queries

import "fmt"

var selectFileInfoQuery = `
select
	id,
	filename_hash,
	status,
	to_char(timestamp, '%s') as timestamp
from files
where timestamp = (
  select max(timestamp) from %s where filename_hash = encode(digest($1, 'md5'), 'hex')
)
`

var SelectFileInfo = fmt.Sprintf(selectFileInfoQuery, DATE_FORMAT, TABLE_USERS_NAME)
