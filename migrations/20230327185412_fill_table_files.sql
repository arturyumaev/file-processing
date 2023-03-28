-- +goose Up
-- +goose StatementBegin
create procedure gen_data(s file_status, h text)
language plpgsql as $$
begin
  insert into files (status, filename_hash)
    values
      (s, encode(digest(h, 'md5'), 'hex'));

  perform pg_sleep(2);
end $$;

call gen_data('recieved',   'file1');
call gen_data('in_queue',   'file1');
call gen_data('processing', 'file1');
call gen_data('processing', 'file1');
call gen_data('done',       'file1');

call gen_data('recieved', 'file2');
call gen_data('error',    'file2');

call gen_data('recieved', 'file3');
call gen_data('in_queue', 'file3');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from files;
drop procedure gen_data;
-- +goose StatementEnd
