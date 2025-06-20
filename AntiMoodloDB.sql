--
-- PostgreSQL database dump
--

-- Dumped from database version 14.18 (Homebrew)
-- Dumped by pg_dump version 14.18 (Homebrew)

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
-- Name: block; Type: TABLE; Schema: public; Owner: arturaipov
--

CREATE TABLE public.block (
    blockid integer NOT NULL,
    blockname text,
    courseid bigint
);


ALTER TABLE public.block OWNER TO arturaipov;

--
-- Name: block_blockid_seq; Type: SEQUENCE; Schema: public; Owner: arturaipov
--

CREATE SEQUENCE public.block_blockid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.block_blockid_seq OWNER TO arturaipov;

--
-- Name: block_blockid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: arturaipov
--

ALTER SEQUENCE public.block_blockid_seq OWNED BY public.block.blockid;


--
-- Name: correctanswers; Type: TABLE; Schema: public; Owner: arturaipov
--

CREATE TABLE public.correctanswers (
    id integer NOT NULL,
    questionid bigint,
    optionid bigint
);


ALTER TABLE public.correctanswers OWNER TO arturaipov;

--
-- Name: correctanswers_id_seq; Type: SEQUENCE; Schema: public; Owner: arturaipov
--

CREATE SEQUENCE public.correctanswers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.correctanswers_id_seq OWNER TO arturaipov;

--
-- Name: correctanswers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: arturaipov
--

ALTER SEQUENCE public.correctanswers_id_seq OWNED BY public.correctanswers.id;


--
-- Name: course; Type: TABLE; Schema: public; Owner: arturaipov
--

CREATE TABLE public.course (
    courseid integer NOT NULL,
    title text
);


ALTER TABLE public.course OWNER TO arturaipov;

--
-- Name: course_courseid_seq; Type: SEQUENCE; Schema: public; Owner: arturaipov
--

ALTER TABLE public.course ALTER COLUMN courseid ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.course_courseid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: matchpairs; Type: TABLE; Schema: public; Owner: arturaipov
--

CREATE TABLE public.matchpairs (
    id integer NOT NULL,
    questionid bigint,
    lefttext text,
    righttext text
);


ALTER TABLE public.matchpairs OWNER TO arturaipov;

--
-- Name: matchpairs_id_seq; Type: SEQUENCE; Schema: public; Owner: arturaipov
--

CREATE SEQUENCE public.matchpairs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.matchpairs_id_seq OWNER TO arturaipov;

--
-- Name: matchpairs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: arturaipov
--

ALTER SEQUENCE public.matchpairs_id_seq OWNED BY public.matchpairs.id;


--
-- Name: openanswers; Type: TABLE; Schema: public; Owner: arturaipov
--

CREATE TABLE public.openanswers (
    questionid integer NOT NULL,
    answertext text
);


ALTER TABLE public.openanswers OWNER TO arturaipov;

--
-- Name: questionoptions; Type: TABLE; Schema: public; Owner: arturaipov
--

CREATE TABLE public.questionoptions (
    id integer NOT NULL,
    questionid integer NOT NULL,
    optiontext text NOT NULL
);


ALTER TABLE public.questionoptions OWNER TO arturaipov;

--
-- Name: questionoptions_id_seq; Type: SEQUENCE; Schema: public; Owner: arturaipov
--

CREATE SEQUENCE public.questionoptions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.questionoptions_id_seq OWNER TO arturaipov;

--
-- Name: questionoptions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: arturaipov
--

ALTER SEQUENCE public.questionoptions_id_seq OWNED BY public.questionoptions.id;


--
-- Name: questions; Type: TABLE; Schema: public; Owner: arturaipov
--

CREATE TABLE public.questions (
    id integer NOT NULL,
    quizid bigint,
    questiontext text,
    questiontypeid integer NOT NULL
);


ALTER TABLE public.questions OWNER TO arturaipov;

--
-- Name: questions_id_seq; Type: SEQUENCE; Schema: public; Owner: arturaipov
--

CREATE SEQUENCE public.questions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.questions_id_seq OWNER TO arturaipov;

--
-- Name: questions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: arturaipov
--

ALTER SEQUENCE public.questions_id_seq OWNED BY public.questions.id;


--
-- Name: questiontypes; Type: TABLE; Schema: public; Owner: arturaipov
--

CREATE TABLE public.questiontypes (
    id integer NOT NULL,
    name text
);


ALTER TABLE public.questiontypes OWNER TO arturaipov;

--
-- Name: questiontypes_id_seq; Type: SEQUENCE; Schema: public; Owner: arturaipov
--

CREATE SEQUENCE public.questiontypes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.questiontypes_id_seq OWNER TO arturaipov;

--
-- Name: questiontypes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: arturaipov
--

ALTER SEQUENCE public.questiontypes_id_seq OWNED BY public.questiontypes.id;


--
-- Name: quizzes; Type: TABLE; Schema: public; Owner: arturaipov
--

CREATE TABLE public.quizzes (
    id integer NOT NULL,
    title text,
    maxgrade bigint,
    duration bigint,
    stateid bigint,
    startdate timestamp without time zone,
    enddate timestamp without time zone NOT NULL,
    submiteddate timestamp without time zone,
    courseid integer NOT NULL
);


ALTER TABLE public.quizzes OWNER TO arturaipov;

--
-- Name: quizzes_id_seq; Type: SEQUENCE; Schema: public; Owner: arturaipov
--

CREATE SEQUENCE public.quizzes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.quizzes_id_seq OWNER TO arturaipov;

--
-- Name: quizzes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: arturaipov
--

ALTER SEQUENCE public.quizzes_id_seq OWNED BY public.quizzes.id;


--
-- Name: states; Type: TABLE; Schema: public; Owner: arturaipov
--

CREATE TABLE public.states (
    id integer NOT NULL,
    name text
);


ALTER TABLE public.states OWNER TO arturaipov;

--
-- Name: states_id_seq; Type: SEQUENCE; Schema: public; Owner: arturaipov
--

CREATE SEQUENCE public.states_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.states_id_seq OWNER TO arturaipov;

--
-- Name: states_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: arturaipov
--

ALTER SEQUENCE public.states_id_seq OWNED BY public.states.id;


--
-- Name: usercourse; Type: TABLE; Schema: public; Owner: arturaipov
--

CREATE TABLE public.usercourse (
    usercourseid integer NOT NULL,
    userid bigint,
    courseid bigint,
    grade numeric,
    lettergrade text
);


ALTER TABLE public.usercourse OWNER TO arturaipov;

--
-- Name: usercourse_usercourseid_seq; Type: SEQUENCE; Schema: public; Owner: arturaipov
--

CREATE SEQUENCE public.usercourse_usercourseid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.usercourse_usercourseid_seq OWNER TO arturaipov;

--
-- Name: usercourse_usercourseid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: arturaipov
--

ALTER SEQUENCE public.usercourse_usercourseid_seq OWNED BY public.usercourse.usercourseid;


--
-- Name: userroles; Type: TABLE; Schema: public; Owner: arturaipov
--

CREATE TABLE public.userroles (
    roleid integer NOT NULL,
    rolename text
);


ALTER TABLE public.userroles OWNER TO arturaipov;

--
-- Name: userroles_roleid_seq; Type: SEQUENCE; Schema: public; Owner: arturaipov
--

CREATE SEQUENCE public.userroles_roleid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.userroles_roleid_seq OWNER TO arturaipov;

--
-- Name: userroles_roleid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: arturaipov
--

ALTER SEQUENCE public.userroles_roleid_seq OWNED BY public.userroles.roleid;


--
-- Name: users; Type: TABLE; Schema: public; Owner: arturaipov
--

CREATE TABLE public.users (
    userid integer NOT NULL,
    userlogin text,
    userpassword text,
    userrole bigint
);


ALTER TABLE public.users OWNER TO arturaipov;

--
-- Name: users_userid_seq; Type: SEQUENCE; Schema: public; Owner: arturaipov
--

CREATE SEQUENCE public.users_userid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_userid_seq OWNER TO arturaipov;

--
-- Name: users_userid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: arturaipov
--

ALTER SEQUENCE public.users_userid_seq OWNED BY public.users.userid;


--
-- Name: block blockid; Type: DEFAULT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.block ALTER COLUMN blockid SET DEFAULT nextval('public.block_blockid_seq'::regclass);


--
-- Name: correctanswers id; Type: DEFAULT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.correctanswers ALTER COLUMN id SET DEFAULT nextval('public.correctanswers_id_seq'::regclass);


--
-- Name: matchpairs id; Type: DEFAULT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.matchpairs ALTER COLUMN id SET DEFAULT nextval('public.matchpairs_id_seq'::regclass);


--
-- Name: questionoptions id; Type: DEFAULT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.questionoptions ALTER COLUMN id SET DEFAULT nextval('public.questionoptions_id_seq'::regclass);


--
-- Name: questions id; Type: DEFAULT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.questions ALTER COLUMN id SET DEFAULT nextval('public.questions_id_seq'::regclass);


--
-- Name: questiontypes id; Type: DEFAULT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.questiontypes ALTER COLUMN id SET DEFAULT nextval('public.questiontypes_id_seq'::regclass);


--
-- Name: quizzes id; Type: DEFAULT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.quizzes ALTER COLUMN id SET DEFAULT nextval('public.quizzes_id_seq'::regclass);


--
-- Name: states id; Type: DEFAULT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.states ALTER COLUMN id SET DEFAULT nextval('public.states_id_seq'::regclass);


--
-- Name: usercourse usercourseid; Type: DEFAULT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.usercourse ALTER COLUMN usercourseid SET DEFAULT nextval('public.usercourse_usercourseid_seq'::regclass);


--
-- Name: userroles roleid; Type: DEFAULT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.userroles ALTER COLUMN roleid SET DEFAULT nextval('public.userroles_roleid_seq'::regclass);


--
-- Name: users userid; Type: DEFAULT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.users ALTER COLUMN userid SET DEFAULT nextval('public.users_userid_seq'::regclass);


--
-- Data for Name: block; Type: TABLE DATA; Schema: public; Owner: arturaipov
--

COPY public.block (blockid, blockname, courseid) FROM stdin;
1	Algebra	1
2	Mechanics	2
3	Renaissance	3
\.


--
-- Data for Name: correctanswers; Type: TABLE DATA; Schema: public; Owner: arturaipov
--

COPY public.correctanswers (id, questionid, optionid) FROM stdin;
1	1	1
2	2	4
3	2	6
\.


--
-- Data for Name: course; Type: TABLE DATA; Schema: public; Owner: arturaipov
--

COPY public.course (courseid, title) FROM stdin;
1	Math 101
2	Physics Basics
3	History of Art
\.


--
-- Data for Name: matchpairs; Type: TABLE DATA; Schema: public; Owner: arturaipov
--

COPY public.matchpairs (id, questionid, lefttext, righttext) FROM stdin;
1	2	2	Even
2	2	3	Odd
3	2	4	Even
\.


--
-- Data for Name: openanswers; Type: TABLE DATA; Schema: public; Owner: arturaipov
--

COPY public.openanswers (questionid, answertext) FROM stdin;
3	It is a period of revival in arts and science.
\.


--
-- Data for Name: questionoptions; Type: TABLE DATA; Schema: public; Owner: arturaipov
--

COPY public.questionoptions (id, questionid, optiontext) FROM stdin;
1	1	4
2	1	3
3	1	5
4	2	2
5	2	3
6	2	4
\.


--
-- Data for Name: questions; Type: TABLE DATA; Schema: public; Owner: arturaipov
--

COPY public.questions (id, quizid, questiontext, questiontypeid) FROM stdin;
1	1	What is 2 + 2?	1
2	2	Choose all even numbers	2
3	3	Explain Renaissance art.	3
\.


--
-- Data for Name: questiontypes; Type: TABLE DATA; Schema: public; Owner: arturaipov
--

COPY public.questiontypes (id, name) FROM stdin;
1	Single Choice
2	Multiple Choice
3	Open Answer
\.


--
-- Data for Name: quizzes; Type: TABLE DATA; Schema: public; Owner: arturaipov
--

COPY public.quizzes (id, title, maxgrade, duration, stateid, startdate, enddate, submiteddate, courseid) FROM stdin;
1	Midterm Quiz	100	60	2	2025-06-13 02:44:29.536886	2025-06-20 02:44:29.536886	\N	1
2	Final Quiz	100	90	2	2025-06-13 02:44:29.536886	2025-06-23 02:44:29.536886	\N	2
3	Practice Quiz	50	30	1	2025-06-13 02:44:29.536886	2025-06-16 02:44:29.536886	\N	3
\.


--
-- Data for Name: states; Type: TABLE DATA; Schema: public; Owner: arturaipov
--

COPY public.states (id, name) FROM stdin;
1	Draft
2	Published
3	Closed
\.


--
-- Data for Name: usercourse; Type: TABLE DATA; Schema: public; Owner: arturaipov
--

COPY public.usercourse (usercourseid, userid, courseid, grade, lettergrade) FROM stdin;
1	1	1	90	A
2	2	2	75	C
3	3	3	85	B
\.


--
-- Data for Name: userroles; Type: TABLE DATA; Schema: public; Owner: arturaipov
--

COPY public.userroles (roleid, rolename) FROM stdin;
1	Student
2	Teacher
3	Admin
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: arturaipov
--

COPY public.users (userid, userlogin, userpassword, userrole) FROM stdin;
1	alice	password123	1
2	bob	secure456	2
3	charlie	admin789	3
\.


--
-- Name: block_blockid_seq; Type: SEQUENCE SET; Schema: public; Owner: arturaipov
--

SELECT pg_catalog.setval('public.block_blockid_seq', 3, true);


--
-- Name: correctanswers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: arturaipov
--

SELECT pg_catalog.setval('public.correctanswers_id_seq', 3, true);


--
-- Name: course_courseid_seq; Type: SEQUENCE SET; Schema: public; Owner: arturaipov
--

SELECT pg_catalog.setval('public.course_courseid_seq', 3, true);


--
-- Name: matchpairs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: arturaipov
--

SELECT pg_catalog.setval('public.matchpairs_id_seq', 3, true);


--
-- Name: questionoptions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: arturaipov
--

SELECT pg_catalog.setval('public.questionoptions_id_seq', 6, true);


--
-- Name: questions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: arturaipov
--

SELECT pg_catalog.setval('public.questions_id_seq', 3, true);


--
-- Name: questiontypes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: arturaipov
--

SELECT pg_catalog.setval('public.questiontypes_id_seq', 3, true);


--
-- Name: quizzes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: arturaipov
--

SELECT pg_catalog.setval('public.quizzes_id_seq', 3, true);


--
-- Name: states_id_seq; Type: SEQUENCE SET; Schema: public; Owner: arturaipov
--

SELECT pg_catalog.setval('public.states_id_seq', 3, true);


--
-- Name: usercourse_usercourseid_seq; Type: SEQUENCE SET; Schema: public; Owner: arturaipov
--

SELECT pg_catalog.setval('public.usercourse_usercourseid_seq', 3, true);


--
-- Name: userroles_roleid_seq; Type: SEQUENCE SET; Schema: public; Owner: arturaipov
--

SELECT pg_catalog.setval('public.userroles_roleid_seq', 3, true);


--
-- Name: users_userid_seq; Type: SEQUENCE SET; Schema: public; Owner: arturaipov
--

SELECT pg_catalog.setval('public.users_userid_seq', 3, true);


--
-- Name: block block_pkey; Type: CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.block
    ADD CONSTRAINT block_pkey PRIMARY KEY (blockid);


--
-- Name: correctanswers correctanswers_pkey; Type: CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.correctanswers
    ADD CONSTRAINT correctanswers_pkey PRIMARY KEY (id);


--
-- Name: course course_pkey; Type: CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.course
    ADD CONSTRAINT course_pkey PRIMARY KEY (courseid);


--
-- Name: matchpairs matchpairs_pkey; Type: CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.matchpairs
    ADD CONSTRAINT matchpairs_pkey PRIMARY KEY (id);


--
-- Name: openanswers openanswers_pkey; Type: CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.openanswers
    ADD CONSTRAINT openanswers_pkey PRIMARY KEY (questionid);


--
-- Name: questionoptions questionoptions_pkey; Type: CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.questionoptions
    ADD CONSTRAINT questionoptions_pkey PRIMARY KEY (id);


--
-- Name: questions questions_pkey; Type: CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.questions
    ADD CONSTRAINT questions_pkey PRIMARY KEY (id);


--
-- Name: questiontypes questiontypes_pkey; Type: CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.questiontypes
    ADD CONSTRAINT questiontypes_pkey PRIMARY KEY (id);


--
-- Name: quizzes quizzes_pkey; Type: CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.quizzes
    ADD CONSTRAINT quizzes_pkey PRIMARY KEY (id);


--
-- Name: states states_pkey; Type: CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.states
    ADD CONSTRAINT states_pkey PRIMARY KEY (id);


--
-- Name: usercourse usercourse_pkey; Type: CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.usercourse
    ADD CONSTRAINT usercourse_pkey PRIMARY KEY (usercourseid);


--
-- Name: userroles userroles_pkey; Type: CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.userroles
    ADD CONSTRAINT userroles_pkey PRIMARY KEY (roleid);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (userid);


--
-- Name: correctanswers correctanswers_optionid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.correctanswers
    ADD CONSTRAINT correctanswers_optionid_fkey FOREIGN KEY (optionid) REFERENCES public.questionoptions(id) ON DELETE CASCADE;


--
-- Name: correctanswers correctanswers_questionid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.correctanswers
    ADD CONSTRAINT correctanswers_questionid_fkey FOREIGN KEY (questionid) REFERENCES public.questions(id) ON DELETE CASCADE;


--
-- Name: block fk_blockcourse; Type: FK CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.block
    ADD CONSTRAINT fk_blockcourse FOREIGN KEY (courseid) REFERENCES public.course(courseid);


--
-- Name: quizzes fk_course; Type: FK CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.quizzes
    ADD CONSTRAINT fk_course FOREIGN KEY (courseid) REFERENCES public.course(courseid);


--
-- Name: usercourse fk_course; Type: FK CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.usercourse
    ADD CONSTRAINT fk_course FOREIGN KEY (courseid) REFERENCES public.course(courseid);


--
-- Name: usercourse fk_user; Type: FK CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.usercourse
    ADD CONSTRAINT fk_user FOREIGN KEY (userid) REFERENCES public.users(userid);


--
-- Name: users fk_userrole; Type: FK CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_userrole FOREIGN KEY (userrole) REFERENCES public.userroles(roleid);


--
-- Name: matchpairs matchpairs_questionid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.matchpairs
    ADD CONSTRAINT matchpairs_questionid_fkey FOREIGN KEY (questionid) REFERENCES public.questions(id) ON DELETE CASCADE;


--
-- Name: openanswers openanswers_questionid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.openanswers
    ADD CONSTRAINT openanswers_questionid_fkey FOREIGN KEY (questionid) REFERENCES public.questions(id) ON DELETE CASCADE;


--
-- Name: questionoptions questionoptions_questionid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.questionoptions
    ADD CONSTRAINT questionoptions_questionid_fkey FOREIGN KEY (questionid) REFERENCES public.questions(id) ON DELETE CASCADE;


--
-- Name: questions questions_questiontypeid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.questions
    ADD CONSTRAINT questions_questiontypeid_fkey FOREIGN KEY (questiontypeid) REFERENCES public.questiontypes(id);


--
-- Name: questions questions_quizid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.questions
    ADD CONSTRAINT questions_quizid_fkey FOREIGN KEY (quizid) REFERENCES public.quizzes(id) ON DELETE CASCADE;


--
-- Name: quizzes quizzes_stateid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: arturaipov
--

ALTER TABLE ONLY public.quizzes
    ADD CONSTRAINT quizzes_stateid_fkey FOREIGN KEY (stateid) REFERENCES public.states(id);


--
-- PostgreSQL database dump complete
--

