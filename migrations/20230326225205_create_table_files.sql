-- +goose Up
-- +goose StatementBegin
create type file_status as enum (
  'recieved',
  'in_queue',
  'processing',
  'done',
  'error'
);

create table files (
  id        varchar(36) primary key default gen_random_uuid(),
  status    file_status,
  timestamp timestamp,
  hash      text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table files;
drop type file_status;
-- +goose StatementEnd
