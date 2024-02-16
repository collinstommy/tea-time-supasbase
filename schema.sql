
SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

CREATE SCHEMA IF NOT EXISTS "public";

ALTER SCHEMA "public" OWNER TO "pg_database_owner";

SET default_tablespace = '';

SET default_table_access_method = "heap";

CREATE TABLE IF NOT EXISTS "public"."games" (
    "game_id" uuid DEFAULT extensions.uuid_generate_v4() NOT NULL,
    "url" text NOT NULL,
    "name" text NOT NULL,
    "created_by" uuid
);

ALTER TABLE "public"."games" OWNER TO "postgres";

CREATE TABLE IF NOT EXISTS "public"."likes" (
    "like_id" uuid DEFAULT extensions.uuid_generate_v4() NOT NULL,
    "user_id" uuid,
    "content_type" text NOT NULL,
    "content_id" uuid NOT NULL,
    CONSTRAINT "likes_content_type_check" CHECK ((content_type = ANY (ARRAY['playlist'::text, 'post'::text])))
);

ALTER TABLE "public"."likes" OWNER TO "postgres";

CREATE TABLE IF NOT EXISTS "public"."playlists" (
    "playlist_id" uuid DEFAULT extensions.uuid_generate_v4() NOT NULL,
    "spotify_id" text NOT NULL,
    "user_id" uuid,
    "artists" TEXT NOT NULL DEFAULT '',
    "name" TEXT NOT NULL DEFAULT '',
    "track_count" INTEGER NOT NULL DEFAULT 0
);

ALTER TABLE "public"."playlists" OWNER TO "postgres";

CREATE TABLE IF NOT EXISTS "public"."posts" (
    "post_id" uuid DEFAULT extensions.uuid_generate_v4() NOT NULL,
    "url" text NOT NULL,
    "user_id" uuid,
    "type" text NOT NULL,
    CONSTRAINT "posts_type_check" CHECK ((type = ANY (ARRAY['blog_post'::text, 'article'::text, 'podcast'::text, 'podcast_episode'::text, 'blog'::text])))
);

ALTER TABLE "public"."posts" OWNER TO "postgres";

CREATE TABLE IF NOT EXISTS "public"."session_participants" (
    "session_id" uuid NOT NULL,
    "user_id" uuid NOT NULL,
    "result" text NOT NULL,
    CONSTRAINT "session_participants_result_check" CHECK ((result = ANY (ARRAY['winner'::text, 'loser'::text])))
);

ALTER TABLE "public"."session_participants" OWNER TO "postgres";

CREATE TABLE IF NOT EXISTS "public"."sessions" (
    "session_id" uuid DEFAULT extensions.uuid_generate_v4() NOT NULL,
    "game_id" uuid,
    "start_time" timestamp without time zone NOT NULL
);

ALTER TABLE "public"."sessions" OWNER TO "postgres";

ALTER TABLE ONLY "public"."likes"
    ADD CONSTRAINT "content_type_content_id_unique" UNIQUE ("content_type", "content_id");

ALTER TABLE ONLY "public"."games"
    ADD CONSTRAINT "games_pkey" PRIMARY KEY ("game_id");

ALTER TABLE ONLY "public"."likes"
    ADD CONSTRAINT "likes_pkey" PRIMARY KEY ("like_id");

ALTER TABLE ONLY "public"."playlists"
    ADD CONSTRAINT "playlists_pkey" PRIMARY KEY ("playlist_id");

ALTER TABLE ONLY "public"."posts"
    ADD CONSTRAINT "posts_pkey" PRIMARY KEY ("post_id");

ALTER TABLE ONLY "public"."session_participants"
    ADD CONSTRAINT "session_participants_pkey" PRIMARY KEY ("session_id", "user_id");

ALTER TABLE ONLY "public"."sessions"
    ADD CONSTRAINT "sessions_pkey" PRIMARY KEY ("session_id");

ALTER TABLE ONLY "public"."games"
    ADD CONSTRAINT "games_created_by_fkey" FOREIGN KEY (created_by) REFERENCES auth.users(id);

ALTER TABLE ONLY "public"."likes"
    ADD CONSTRAINT "likes_user_id_fkey" FOREIGN KEY (user_id) REFERENCES auth.users(id);

ALTER TABLE ONLY "public"."playlists"
    ADD CONSTRAINT "playlists_user_id_fkey" FOREIGN KEY (user_id) REFERENCES auth.users(id);

ALTER TABLE ONLY "public"."posts"
    ADD CONSTRAINT "posts_user_id_fkey" FOREIGN KEY (user_id) REFERENCES auth.users(id);

ALTER TABLE ONLY "public"."session_participants"
    ADD CONSTRAINT "session_participants_session_id_fkey" FOREIGN KEY (session_id) REFERENCES public.sessions(session_id);

ALTER TABLE ONLY "public"."session_participants"
    ADD CONSTRAINT "session_participants_user_id_fkey" FOREIGN KEY (user_id) REFERENCES auth.users(id);

ALTER TABLE ONLY "public"."sessions"
    ADD CONSTRAINT "sessions_game_id_fkey" FOREIGN KEY (game_id) REFERENCES public.games(game_id);

GRANT USAGE ON SCHEMA "public" TO "postgres";
GRANT USAGE ON SCHEMA "public" TO "anon";
GRANT USAGE ON SCHEMA "public" TO "authenticated";
GRANT USAGE ON SCHEMA "public" TO "service_role";

GRANT ALL ON TABLE "public"."games" TO "anon";
GRANT ALL ON TABLE "public"."games" TO "authenticated";
GRANT ALL ON TABLE "public"."games" TO "service_role";

GRANT ALL ON TABLE "public"."likes" TO "anon";
GRANT ALL ON TABLE "public"."likes" TO "authenticated";
GRANT ALL ON TABLE "public"."likes" TO "service_role";

GRANT ALL ON TABLE "public"."playlists" TO "anon";
GRANT ALL ON TABLE "public"."playlists" TO "authenticated";
GRANT ALL ON TABLE "public"."playlists" TO "service_role";

GRANT ALL ON TABLE "public"."posts" TO "anon";
GRANT ALL ON TABLE "public"."posts" TO "authenticated";
GRANT ALL ON TABLE "public"."posts" TO "service_role";

GRANT ALL ON TABLE "public"."session_participants" TO "anon";
GRANT ALL ON TABLE "public"."session_participants" TO "authenticated";
GRANT ALL ON TABLE "public"."session_participants" TO "service_role";

GRANT ALL ON TABLE "public"."sessions" TO "anon";
GRANT ALL ON TABLE "public"."sessions" TO "authenticated";
GRANT ALL ON TABLE "public"."sessions" TO "service_role";

ALTER DEFAULT PRIVILEGES FOR ROLE "postgres" IN SCHEMA "public" GRANT ALL ON SEQUENCES  TO "postgres";
ALTER DEFAULT PRIVILEGES FOR ROLE "postgres" IN SCHEMA "public" GRANT ALL ON SEQUENCES  TO "anon";
ALTER DEFAULT PRIVILEGES FOR ROLE "postgres" IN SCHEMA "public" GRANT ALL ON SEQUENCES  TO "authenticated";
ALTER DEFAULT PRIVILEGES FOR ROLE "postgres" IN SCHEMA "public" GRANT ALL ON SEQUENCES  TO "service_role";

ALTER DEFAULT PRIVILEGES FOR ROLE "postgres" IN SCHEMA "public" GRANT ALL ON FUNCTIONS  TO "postgres";
ALTER DEFAULT PRIVILEGES FOR ROLE "postgres" IN SCHEMA "public" GRANT ALL ON FUNCTIONS  TO "anon";
ALTER DEFAULT PRIVILEGES FOR ROLE "postgres" IN SCHEMA "public" GRANT ALL ON FUNCTIONS  TO "authenticated";
ALTER DEFAULT PRIVILEGES FOR ROLE "postgres" IN SCHEMA "public" GRANT ALL ON FUNCTIONS  TO "service_role";

ALTER DEFAULT PRIVILEGES FOR ROLE "postgres" IN SCHEMA "public" GRANT ALL ON TABLES  TO "postgres";
ALTER DEFAULT PRIVILEGES FOR ROLE "postgres" IN SCHEMA "public" GRANT ALL ON TABLES  TO "anon";
ALTER DEFAULT PRIVILEGES FOR ROLE "postgres" IN SCHEMA "public" GRANT ALL ON TABLES  TO "authenticated";
ALTER DEFAULT PRIVILEGES FOR ROLE "postgres" IN SCHEMA "public" GRANT ALL ON TABLES  TO "service_role";

RESET ALL;
