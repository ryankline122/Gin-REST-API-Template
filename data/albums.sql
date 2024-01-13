--
-- PostgreSQL database dump
--

-- Dumped from database version 14.10 (Ubuntu 14.10-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.10 (Ubuntu 14.10-0ubuntu0.22.04.1)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: album; Type: TABLE; Schema: public; Owner: dev
--

CREATE TABLE public.album (
    id bigint NOT NULL,
    title character varying(255) NOT NULL,
    artist character varying(255) NOT NULL,
    price numeric(10,2) NOT NULL
);


ALTER TABLE public.album OWNER TO dev;

--
-- Name: album_id_seq; Type: SEQUENCE; Schema: public; Owner: dev
--

CREATE SEQUENCE public.album_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.album_id_seq OWNER TO dev;

--
-- Name: album_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dev
--

ALTER SEQUENCE public.album_id_seq OWNED BY public.album.id;


--
-- Name: album id; Type: DEFAULT; Schema: public; Owner: dev
--

ALTER TABLE ONLY public.album ALTER COLUMN id SET DEFAULT nextval('public.album_id_seq'::regclass);


--
-- Data for Name: album; Type: TABLE DATA; Schema: public; Owner: dev
--

COPY public.album (id, title, artist, price) FROM stdin;
1	Appetite for Destruction	Guns n' Roses	12.99
2	Tickets to My Downfall	Machine Gun Kelly	15.99
3	Back in Black	AC/DC	9.99
\.


--
-- Name: album_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dev
--

SELECT pg_catalog.setval('public.album_id_seq', 3, true);


--
-- Name: album album_pkey; Type: CONSTRAINT; Schema: public; Owner: dev
--

ALTER TABLE ONLY public.album
    ADD CONSTRAINT album_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

