-- +goose Up
-- +goose StatementBegin
alter database postgres set timezone to 'Europe/Moscow';
create extension if not exists pgcrypto;

create type file_status as enum (
  'recieved',
  'in_queue',
  'processing',
  'done',
  'error'
);

create table files (
  id            varchar(36) primary key default gen_random_uuid(),
  status        file_status,
  timestamp     timestamp with time zone default clock_timestamp(),
  filename_hash varchar(32) -- md5 hex
);

create index files_hash_idx on files(filename_hash);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index files_hash_idx;
drop table files;
drop type file_status;
-- +goose StatementEnd
