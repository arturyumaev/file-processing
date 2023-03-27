-- +goose Up
-- +goose StatementBegin
create procedure gen_data(s file_status, h text)
language plpgsql as $$
begin
  insert into files (status, timestamp, hash)
    values
      (s, (select clock_timestamp()), h);

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
drop procedure gen_data;
-- +goose StatementEnd
