-- +goose Up
-- +goose StatementBegin
alter database postgres set timezone to 'Europe/Moscow';

create table files (
  id        bigserial primary key, -- 64bit autoincrement
  status    varchar(100),
  timestamp timestamp with time zone default clock_timestamp(),
  filename  varchar(500)
);

create index files_filename_idx on files(filename);
create index files_timestamp_idx on files(timestamp);

-- trigger function
create or replace function trigger_function_valid_status_type()
  returns trigger
  language plpgsql
as $$
begin
	if new.status not in ('recieved', 'in_queue', 'processing', 'done', 'error') then
    raise exception 'wrong status type';
  end if;

	return new;
end $$;

-- trigger
create trigger trigger_valid_status_type
  before insert or update
  on files
  for each row execute function trigger_function_valid_status_type();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop trigger trigger_valid_status_type on files;
drop function trigger_function_valid_status_type;
drop index files_filename_idx;
drop index files_timestamp_idx;
drop table files;
-- drop type file_status;
-- +goose StatementEnd
