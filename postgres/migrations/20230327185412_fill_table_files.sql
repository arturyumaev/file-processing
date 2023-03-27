-- +goose Up
-- +goose StatementBegin
create function gen_uuid()
returns uuid
language plpgsql as $$
begin
  return (
    select uuid_in(md5(random()::text || random()::text)::cstring)
  );
end $$;

create procedure gen_data(s file_status, h text)
language plpgsql as $$
begin
  insert into files (id, status, timestamp, hash)
    values
      ((select gen_uuid()), s, (select clock_timestamp()), h);

  perform pg_sleep(2);
end $$;

call gen_data('recieved',   'hash1');
call gen_data('in_queue',   'hash1');
call gen_data('processing', 'hash1');
call gen_data('processing', 'hash1');
call gen_data('done',       'hash1');

call gen_data('recieved', 'hash2');
call gen_data('error',    'hash2');

call gen_data('recieved', 'hash3');
call gen_data('in_queue', 'hash3');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from files;
drop function gen_uuid;
drop procedure gen_data;
-- +goose StatementEnd
