-- +goose Up
-- +goose StatementBegin
alter database postgres set timezone to 'Europe/Moscow';

create table files (
  id        bigserial primary key, -- 64bit autoincrement
  status    varchar(100),
  timestamp timestamp with time zone default clock_timestamp(),
  filename  varchar(500),
  unique(filename, status)
);

create index files_filename_idx on files(filename);
create index files_timestamp_idx on files(timestamp);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index files_filename_idx;
drop index files_timestamp_idx;
drop table files;
-- drop type file_status;
-- +goose StatementEnd
