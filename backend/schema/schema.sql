CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "full_name" varchar,
  "email" varchar,
  "password" varchar,
  "created_at" timestamp
);

CREATE TABLE "blogs" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar,
  "content" varchar,
  "cover" varchar,
  "user_id" int,
  "created_at" timestamp
);

CREATE TABLE "followers" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int,
  "follow" int,
  "created_at" timestamp
);

ALTER TABLE "blogs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "followers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");