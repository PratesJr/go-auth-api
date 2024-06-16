
CREATE TABLE IF NOT EXISTS "auth_users"."users"
(
 "id"         uuid NOT NULL,
 "name"       varchar(50) NOT NULL,
 "email"      varchar(50) NOT NULL,
 "birth"      date NOT NULL,
 "password"   varchar NOT NULL,
 "created_at" timestamp with time zone NOT NULL,
 "updated_at" timestamp with time zone NOT NULL,
 "deleted_at" timestamp with time zone NULL,
 CONSTRAINT "PK_1" PRIMARY KEY ( "id" )
);

