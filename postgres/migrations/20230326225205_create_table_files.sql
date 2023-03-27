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
  id        varchar(36) not null primary key,
  status    file_status not null,
  timestamp timestamp   not null,
  hash      text        not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table files;
drop type file_status;
-- +goose StatementEnd
