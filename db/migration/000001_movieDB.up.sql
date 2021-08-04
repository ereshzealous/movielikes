CREATE table movie (
    id bigserial primary key,
    name text not null,
    production_company text not null,
    year_released int not null,
    created_at timestamptz default now()
);

CREATE TABLE users (
  user_name text PRIMARY KEY,
  hashed_password text NOT NULL,
  full_name text NOT NULL,
  email text UNIQUE NOT NULL,
  password_changed_at timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),  
  created_at timestamptz NOT NULL DEFAULT (now())
);

create table movie_likes (
    id bigserial,
    movie_id bigserial not null,
    user_name text not null,
    liked boolean not null default true,
    created_at timestamptz default now()
);

ALTER TABLE movie_likes ADD FOREIGN KEY (user_name) REFERENCES users (user_name);
ALTER TABLE movie_likes ADD FOREIGN KEY (movie_id) REFERENCES movie (id);