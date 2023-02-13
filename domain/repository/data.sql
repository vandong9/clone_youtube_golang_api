# clone_youtube_golang_api


databse:

-- DROP SCHEMA "CloneYoutube";

CREATE SCHEMA "CloneYoutube" AUTHORIZATION postgres;



-- "CloneYoutube"."User" definition

-- Drop table

-- DROP TABLE "CloneYoutube"."User";

CREATE TABLE "CloneYoutube"."User" (
	id varchar(10) NOT NULL,
	full_name varchar(100) NULL,
	first_name varchar(50) NULL,
	last_name varchar(50) NULL,
	email varchar(100) NULL,
	"password" varchar(70) NULL,
	profile_pic varchar NULL,
	isverified bool NULL DEFAULT false,
	has_channel bool NULL DEFAULT false,
	CONSTRAINT user_pk PRIMARY KEY (id)
);
CREATE UNIQUE INDEX user_user_id_idx ON "CloneYoutube"."User" USING btree (id);


-- "CloneYoutube"."Channel" definition

-- Drop table

-- DROP TABLE "CloneYoutube"."Channel";

CREATE TABLE "CloneYoutube"."Channel" (
	id varchar(10) NOT NULL,
	"name" varchar(150) NULL,
	description text NULL,
	user_id varchar(10) NOT NULL,
	CONSTRAINT channel_pk PRIMARY KEY (id),
	CONSTRAINT channel_un UNIQUE (user_id)
);


-- "CloneYoutube"."Video" definition

-- Drop table

-- DROP TABLE "CloneYoutube"."Video";

CREATE TABLE "CloneYoutube"."Video" (
	id varchar(10) NULL,
	title varchar(150) NULL,
	description text NULL,
	streaming_url varchar NULL,
	thumbnail_url varchar NULL,
	"key" varchar(100) NULL,
	processed bool NULL DEFAULT false,
	user_id varchar(10) NULL,
	channel_id varchar(10) NULL
);


-- "CloneYoutube"."Comment" definition

-- Drop table

-- DROP TABLE "CloneYoutube"."Comment";

CREATE TABLE "CloneYoutube"."Comment" (
	id varchar(10) NULL,
	user_id varchar(10) NULL,
	video_id varchar(10) NULL,
	reply text NULL
);


-- "CloneYoutube".verifytoken definition

-- Drop table

-- DROP TABLE "CloneYoutube".verifytoken;

CREATE TABLE "CloneYoutube".verifytoken (
	"token" text NULL,
	user_id varchar(10) NULL
);