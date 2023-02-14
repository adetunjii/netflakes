CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "body" varchar NOT NULL,
  "movie_id" bigint NOT NULL,
  "movie_url" varchar NOT NULL,
  "sender_ip" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" varchar NOT NULL
);

CREATE INDEX idx_comment_movie_id ON "comments" ("movie_id");

CREATE INDEX idx_comment_sender_ip ON "comments" ("sender_ip");

CREATE INDEX idx_comment_created_by ON "comments" ("created_by");

CREATE INDEX idx_comment_created_at ON "comments" ("created_at");
