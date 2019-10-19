CREATE keyspace hiromaily WITH replication = {'class':'SimpleStrategy', 'replication_factor':1};

USE hiromaily;

create table t_users (
    id uuid,
    first_name varchar,
    last_name varchar,
    email varchar,
    password varchar,
    age int,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (id)
);

CREATE INDEX ix01_t_users ON t_users ( first_name );
CREATE INDEX ix02_t_users ON t_users ( updated_at );

DESCRIBE table t_users;

INSERT INTO t_users
 (
  id, first_name, last_name, email, password, age, created_at, updated_at
 )
VALUES
 (
  now(), 'harry', 'asakura', 'aa@test.jp', 'xxxx', 29, dateof(now()), dateof(now())
 );
