  CREATE EXTENSION pgcrypto;

  CREATE TYPE "role" AS ENUM (
    'admin',
    'pro',
    'user'
  );

  CREATE TYPE "futur" AS ENUM (
    'robotics',
    'space',
    'brain',
    'animals'
  );

  CREATE TABLE "files" (
    "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
    "created_at" timestamptz NOT NULL DEFAULT (NOW()),
    "updated_at" timestamptz NOT NULL DEFAULT (NOW()),
    "deleted_at" timestamptz,
    "name" text,
    "url" text,
    "mime" text,
    "size" bigint
  );

  CREATE TABLE "users" (
    "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
    "created_at" timestamp NOT NULL DEFAULT (now()),
    "updated_at" timestamp NOT NULL DEFAULT (now()),
    "deleted_at" timestamp CONSTRAINT deletedchk CHECK (deleted_at > created_at),
    "email" text NOT NULL CONSTRAINT emailchk CHECK (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
    "password" text NOT NULL CONSTRAINT passwordchk CHECK (char_length(password) >= 9),
    "firstname" text CONSTRAINT firstnamechk CHECK (char_length(firstname) >= 3 AND char_length(firstname) <= 20 AND  firstname ~ '^[^0-9]*$') DEFAULT NULL,
    "lastname" text CONSTRAINT lastnamechk CHECK (char_length(lastname) >= 3 AND char_length(lastname) <= 20 AND lastname ~ '^[^0-9]*$') DEFAULT NULL,
    "username" text NOT NULL UNIQUE CONSTRAINT usernamechk CHECK (char_length(username) >= 3 AND char_length(username) <= 20 AND username ~ '^[a-z0-9_\-]+$'),
    "password_confirm_code" text DEFAULT NULL CONSTRAINT code_passwordchk CHECK (char_length(password_confirm_code) = 5),
    "role" role NOT NULL DEFAULT 'user',
    "avatar" text DEFAULT NULL
  );


CREATE TABLE "contacts" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp CONSTRAINT deletedchk CHECK (deleted_at > created_at),
  "email" text NOT NULL CONSTRAINT emailchk CHECK (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
  "msg" text NOT NULL,
  "user_id" uuid NOT NULL
);

CREATE TABLE "refresh_token" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "created_at" timestamptz NOT NULL DEFAULT (NOW()),
  "updated_at" timestamptz NOT NULL DEFAULT (NOW()),
  "deleted_at" timestamptz,
  "token" text NOT NULL,
  "expir_on" timestamptz NOT NULL,
  "user_id" uuid NOT NULL
);

CREATE TABLE "data" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp CONSTRAINT deletedchk CHECK (deleted_at > created_at),
  "title" text NOT NULL CONSTRAINT titlechk CHECK (char_length(title) >= 3 AND char_length(title) <= 20),
  "description" text NOT NULL,
  "user_id" uuid NOT NULL,
  "img" text DEFAULT NULL,
  "category" futur NOT NULL DEFAULT 'robotics'
);

ALTER TABLE "data" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;
ALTER TABLE "refresh_token" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;
ALTER TABLE "contacts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;
