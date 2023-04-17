-- +goose Up
-- +goose StatementBegin
insert into files (status, timestamp, filename)
  values
    ('recieved',   '2023-04-08 21:00:00.000000+03', 'file1'),
    ('in_queue',   '2023-04-08 21:05:00.000000+03', 'file1'),
    ('processing', '2023-04-08 21:10:00.000000+03', 'file1'),
    ('processing', '2023-04-08 21:15:00.000000+03', 'file1'),
    ('done',       '2023-04-08 21:20:00.000000+03', 'file1'),
    ('recieved',   '2023-04-08 21:25:00.000000+03', 'file2'),
    ('error',      '2023-04-08 21:30:00.000000+03', 'file2'),
    ('recieved',   '2023-04-08 21:35:00.000000+03', 'file3'),
    ('in_queue',   '2023-04-08 21:40:00.000000+03', 'file3');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from files;
-- +goose StatementEnd
