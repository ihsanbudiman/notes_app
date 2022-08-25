-- Adminer 4.8.1 PostgreSQL 14.5 (Debian 14.5-1.pgdg110+1) dump

DROP TABLE IF EXISTS "files";
DROP SEQUENCE IF EXISTS files_id_seq;
CREATE SEQUENCE files_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."files" (
    "id" integer DEFAULT nextval('files_id_seq') NOT NULL,
    "folder_sha_id" character varying(10) NOT NULL,
    "name" character varying(255) NOT NULL,
    "type" character varying(25) NOT NULL,
    "created_at" timestamp DEFAULT now() NOT NULL,
    "updated_at" timestamp DEFAULT now() NOT NULL,
    "sha_id" character varying(10) NOT NULL,
    "path" text NOT NULL,
    "user_id" integer NOT NULL,
    CONSTRAINT "files_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "files_folder_sha_id" ON "public"."files" USING btree ("folder_sha_id");

CREATE INDEX "files_path" ON "public"."files" USING btree ("path");

CREATE INDEX "files_sha_id" ON "public"."files" USING btree ("sha_id");

CREATE INDEX "files_type" ON "public"."files" USING btree ("type");

CREATE INDEX "files_user_id" ON "public"."files" USING btree ("user_id");

COMMENT ON COLUMN "public"."files"."type" IS 'folder, note';


DROP TABLE IF EXISTS "folders";
DROP SEQUENCE IF EXISTS folders_id_seq;
CREATE SEQUENCE folders_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."folders" (
    "id" integer DEFAULT nextval('folders_id_seq') NOT NULL,
    "sha_id" character varying(10) NOT NULL,
    "parent_id" integer,
    "created_at" timestamp DEFAULT now() NOT NULL,
    "updated_at" timestamp DEFAULT now() NOT NULL,
    CONSTRAINT "folders_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "folders_parent_id" ON "public"."folders" USING btree ("parent_id");

CREATE INDEX "folders_sha_id" ON "public"."folders" USING btree ("sha_id");


DROP TABLE IF EXISTS "notes";
DROP SEQUENCE IF EXISTS notes_id_seq;
CREATE SEQUENCE notes_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."notes" (
    "id" integer DEFAULT nextval('notes_id_seq') NOT NULL,
    "file_sha_id" character varying(10) NOT NULL,
    "note" text,
    "created_at" timestamp DEFAULT now() NOT NULL,
    "updated_at" timestamp DEFAULT now() NOT NULL,
    CONSTRAINT "notes_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "notes_file_sha_id" ON "public"."notes" USING btree ("file_sha_id");


DROP TABLE IF EXISTS "users";
DROP SEQUENCE IF EXISTS users_id_seq;
CREATE SEQUENCE users_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."users" (
    "id" integer DEFAULT nextval('users_id_seq') NOT NULL,
    "username" character varying(50) NOT NULL,
    "email" character varying(70),
    "phone_number" character varying(15),
    "password" character varying(255) NOT NULL,
    "created_at" timestamp DEFAULT now() NOT NULL,
    "updated_at" timestamp DEFAULT now() NOT NULL,
    "name" character varying(50) NOT NULL,
    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "users_email" ON "public"."users" USING btree ("email");

CREATE INDEX "users_phone_number" ON "public"."users" USING btree ("phone_number");

CREATE INDEX "users_username" ON "public"."users" USING btree ("username");


-- 2022-08-23 09:05:42.61381+00