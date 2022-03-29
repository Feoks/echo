CREATE SCHEMA IF NOT EXISTS "echo";

CREATE TABLE "echo"."tasks" (
    "id" serial8,
    "name" varchar(255) NOT NULL,
    "is_active" bool NOT NULL,
    PRIMARY KEY ("id")
);
