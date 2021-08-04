CREATE TABLE account_details (
  id bigserial PRIMARY KEY,
  user_name TEXT NOT NULL,
  balance bigint NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE transactions (
  id bigserial PRIMARY KEY,
  account_id bigint NOT NULL,
  amount bigint NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE transfers (
  id bigserial PRIMARY KEY,
  source_account_id bigint NOT NULL,
  target_account_id bigint NOT NULL,
  amount bigint NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE transactions ADD FOREIGN KEY (account_id) REFERENCES account_details (id);

ALTER TABLE transfers ADD FOREIGN KEY (source_account_id) REFERENCES account_details (id);

ALTER TABLE transfers ADD FOREIGN KEY (target_account_id) REFERENCES account_details (id);

ALTER TABLE account_details ADD FOREIGN KEY (user_name) REFERENCES users (user_name);