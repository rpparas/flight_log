--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2 (Debian 14.2-1.pgdg110+1)
-- Dumped by pg_dump version 14.4 (Ubuntu 14.4-1.pgdg20.04+1)

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

--
-- Name: pg_trgm; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pg_trgm WITH SCHEMA public;


--
-- Name: EXTENSION pg_trgm; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pg_trgm IS 'text similarity measurement and index searching based on trigrams';


--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: flights; Type: TABLE; Schema: public; Owner: myusername
--

CREATE TABLE public.flights (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    robot_id uuid,
    start_time timestamp with time zone,
    end_time timestamp with time zone,
    lat numeric(10,8),
    lng numeric(11,8)
);


ALTER TABLE public.flights OWNER TO myusername;

--
-- Name: robots; Type: TABLE; Schema: public; Owner: myusername
--

CREATE TABLE public.robots (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(100),
    generation smallint
);


ALTER TABLE public.robots OWNER TO myusername;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: myusername
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO myusername;

--
-- Data for Name: flights; Type: TABLE DATA; Schema: public; Owner: myusername
--

COPY public.flights (id, created_at, updated_at, deleted_at, robot_id, start_time, end_time, lat, lng) FROM stdin;
36ccca14-985d-4bb9-8458-c12a998a4420	2022-07-17 04:44:43.652966+00	2022-07-17 04:44:43.652966+00	\N	e570b6c0-9bb0-47c9-a358-b984ed402406	2022-07-16 00:00:00+00	2022-07-16 00:15:17.684885+00	0.00000000	0.00000000
\.


--
-- Data for Name: robots; Type: TABLE DATA; Schema: public; Owner: myusername
--

COPY public.robots (id, created_at, updated_at, deleted_at, name, generation) FROM stdin;
e570b6c0-9bb0-47c9-a358-b984ed402406	2022-07-17 04:29:36.840092+00	2022-07-17 04:29:36.840092+00	\N	Alpha	1
d9364276-ec19-40af-8069-0a0cdb5ae13e	2022-07-17 04:52:25.850799+00	2022-07-17 04:52:25.850799+00	\N	Alpha3	0
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: myusername
--

COPY public.schema_migrations (version, dirty) FROM stdin;
\.


--
-- Name: flights flights_pkey; Type: CONSTRAINT; Schema: public; Owner: myusername
--

ALTER TABLE ONLY public.flights
    ADD CONSTRAINT flights_pkey PRIMARY KEY (id);


--
-- Name: robots robots_pkey; Type: CONSTRAINT; Schema: public; Owner: myusername
--

ALTER TABLE ONLY public.robots
    ADD CONSTRAINT robots_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: myusername
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: idx_flight; Type: INDEX; Schema: public; Owner: myusername
--

CREATE UNIQUE INDEX idx_flight ON public.flights USING btree (robot_id, start_time, end_time, lat, lng);


--
-- Name: idx_flights_deleted_at; Type: INDEX; Schema: public; Owner: myusername
--

CREATE INDEX idx_flights_deleted_at ON public.flights USING btree (deleted_at);


--
-- Name: idx_robots_deleted_at; Type: INDEX; Schema: public; Owner: myusername
--

CREATE INDEX idx_robots_deleted_at ON public.robots USING btree (deleted_at);


--
-- Name: idx_robots_name; Type: INDEX; Schema: public; Owner: myusername
--

CREATE UNIQUE INDEX idx_robots_name ON public.robots USING btree (name);


--
-- Name: flights flights_robot_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: myusername
--

ALTER TABLE ONLY public.flights
    ADD CONSTRAINT flights_robot_id_fkey FOREIGN KEY (robot_id) REFERENCES public.robots(id);


--
-- PostgreSQL database dump complete
--

