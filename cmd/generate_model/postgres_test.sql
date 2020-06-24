-- Role: generate_model_test
-- DROP ROLE generate_model_test;

CREATE ROLE generate_model_test WITH
  LOGIN
  NOSUPERUSER
  NOINHERIT
  NOCREATEDB
  NOCREATEROLE
  NOREPLICATION
  ENCRYPTED PASSWORD 'md56b274a8d1fb8cdf32d5586e63a69d894'; -- password 73ac4322ba724d72980dbbda43317dcd

-- Database: generate_model_test

-- DROP DATABASE generate_model_test;

CREATE DATABASE generate_model_test
    WITH
    OWNER = generate_model_test
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

-- SCHEMA: generate_model_test

-- DROP SCHEMA generate_model_test ;

CREATE SCHEMA generate_model_test
    AUTHORIZATION generate_model_test;

-- Table: generate_model_test.generate_model_test_table

-- DROP TABLE generate_model_test.generate_model_test_table;

CREATE TABLE generate_model_test.generate_model_test_table
(
    generate_model_test_table_char "char",
    generate_model_test_table_array_char "char"[],
    generate_model_test_table_aclitem aclitem,
    generate_model_test_table_array_aclitem aclitem[],
    generate_model_test_table_bigint bigint,
    generate_model_test_table_array_bigint bigint[],
    generate_model_test_table_bigserial bigserial,
    generate_model_test_table_bit bit(1),
    generate_model_test_table_array_bit bit(1)[],
    generate_model_test_table_bit_varying bit varying,
    generate_model_test_table_array_bit_varying bit varying[],
    generate_model_test_table_boolean boolean,
    generate_model_test_table_array_boolean boolean[],
    generate_model_test_table_box box,
    generate_model_test_table_array_box box[],
    generate_model_test_table_bytea bytea,
    generate_model_test_table_array_bytea bytea[],
    generate_model_test_table_character character(1) COLLATE pg_catalog."default",
    generate_model_test_table_array_character character(1)[] COLLATE pg_catalog."default",
    generate_model_test_table_character_varying character varying[] COLLATE pg_catalog."default",
    generate_model_test_table_array_character_varying character varying[] COLLATE pg_catalog."default",
    generate_model_test_table_cid cid,
    generate_model_test_table_array_cid cid[],
    generate_model_test_table_cidr cidr,
    generate_model_test_table_array_cidr cidr[],
    generate_model_test_table_circle circle,
    generate_model_test_table_array_circle circle[],
    generate_model_test_table_date date,
    generate_model_test_table_array_date date[],
    generate_model_test_table_daterange daterange,
    generate_model_test_table_array_daterange daterange[],
    generate_model_test_table_double_precision double precision,
    generate_model_test_table_array_double_precision double precision[],
    generate_model_test_table_gtsvector gtsvector,
    generate_model_test_table_array_gtsvector gtsvector[],
    generate_model_test_table_inet inet,
    generate_model_test_table_array_inet inet[],
    generate_model_test_table_int2vector int2vector,
    generate_model_test_table_array_int2vector int2vector[],
    generate_model_test_table_int4range int4range,
    generate_model_test_table_array_int4range int4range[],
    generate_model_test_table_int8range int8range,
    generate_model_test_table_array_int8range int8range[],
    generate_model_test_table_integer integer,
    generate_model_test_table_array_integer integer[],
    generate_model_test_table_interval interval,
    generate_model_test_table_array_interval interval[],
    generate_model_test_table_json json,
    generate_model_test_table_array_json json[],
    generate_model_test_table_jsonb jsonb,
    generate_model_test_table_array_jsonb jsonb[],
    generate_model_test_table_jsonpath jsonpath,
    generate_model_test_table_array_jsonpath jsonpath[],
    generate_model_test_table_line line,
    generate_model_test_table_array_line line[],
    generate_model_test_table_lseg lseg,
    generate_model_test_table_array_lseg lseg[],
    generate_model_test_table_macaddr macaddr,
    generate_model_test_table_macaddr8 macaddr8,
    generate_model_test_table_array_macaddr8 macaddr8[],
    generate_model_test_table_array_macaddr macaddr[],
    generate_model_test_table_money money,
    generate_model_test_table_array_money money[],
    generate_model_test_table_name name COLLATE pg_catalog."C",
    generate_model_test_table_array_name name[] COLLATE pg_catalog."C",
    generate_model_test_table_numeric numeric,
    generate_model_test_table_array_numeric numeric[],
    generate_model_test_table_numrange numrange,
    generate_model_test_table_array_numrange numrange[],
    generate_model_test_table_oid oid,
    generate_model_test_table_array_oid oid[],
    generate_model_test_table_oidvector oidvector,
    generate_model_test_table_array_oidvector oidvector[],
    generate_model_test_table_path path,
    generate_model_test_table_array_path path[],
    generate_model_test_table_pg_dependencies pg_dependencies COLLATE pg_catalog."default",
    generate_model_test_table_pg_lsn pg_lsn,
    generate_model_test_table_array_pg_lsn pg_lsn[],
    generate_model_test_table_pg_mcv_list pg_mcv_list COLLATE pg_catalog."default",
    generate_model_test_table_pg_ndistinct pg_ndistinct COLLATE pg_catalog."default",
    generate_model_test_table_pg_node_tree pg_node_tree COLLATE pg_catalog."default",
    generate_model_test_table_point point,
    generate_model_test_table_array_point point[],
    generate_model_test_table_polygon polygon,
    generate_model_test_table_array_polygon polygon[],
    generate_model_test_table_real real,
    generate_model_test_table_array_real real[],
    generate_model_test_table_refcursor refcursor,
    generate_model_test_table_array_refcursor refcursor[],
    generate_model_test_table_regclass regclass,
    generate_model_test_table_array_regclass regclass[],
    generate_model_test_table_regconfig regconfig,
    generate_model_test_table_array_regconfig regconfig[],
    generate_model_test_table_regdictionary regdictionary,
    generate_model_test_table_array_regdictionary regdictionary[],
    generate_model_test_table_regnamespace regnamespace,
    generate_model_test_table_array_regnamespace regnamespace[],
    generate_model_test_table_regoper regoper,
    generate_model_test_table_array_regoper regoper[],
    generate_model_test_table_regoperator regoperator,
    generate_model_test_table_array_regoperator regoperator[],
    generate_model_test_table_regproc regproc,
    generate_model_test_table_array_regproc regproc[],
    generate_model_test_table_regprocedure regprocedure,
    generate_model_test_table_array_regprocedure regprocedure[],
    generate_model_test_table_regrole regrole,
    generate_model_test_table_array_regrole regrole[],
    generate_model_test_table_regtype regtype,
    generate_model_test_table_array_regtype regtype[],
    generate_model_test_table_serial serial,
    generate_model_test_table_smallint smallint,
    generate_model_test_table_array_smallint smallint[],
    generate_model_test_table_smallserial smallserial,
    generate_model_test_table_text text COLLATE pg_catalog."default",
    generate_model_test_table_array_text text[] COLLATE pg_catalog."default",
    generate_model_test_table_tid tid,
    generate_model_test_table_array_tid tid[],
    generate_model_test_table_time_with_time_zone time with time zone,
    generate_model_test_table_array_time_with_time_zone time with time zone[],
    generate_model_test_table_time_without_time_zone time without time zone,
    generate_model_test_table_array_time_without_time_zone time without time zone[],
    generate_model_test_table_timestamp_with_time_zone timestamp with time zone,
    generate_model_test_table_array_timestamp_with_time_zone timestamp with time zone[],
    generate_model_test_table_timestamp_without_time_zone timestamp without time zone,
    generate_model_test_table_array_timestamp_without_time_zone timestamp without time zone[],
    generate_model_test_table_tsquery tsquery,
    generate_model_test_table_array_tsquery tsquery[],
    generate_model_test_table_tsrange tsrange,
    generate_model_test_table_array_tsrange tsrange[],
    generate_model_test_table_tstzrange tstzrange,
    generate_model_test_table_array_tstzrange tstzrange[],
    generate_model_test_table_tsvector tsvector,
    generate_model_test_table_array_tsvector tsvector[],
    generate_model_test_table_txid_snapshot txid_snapshot,
    generate_model_test_table_array_txid_snapshot txid_snapshot[],
    generate_model_test_table_uuid uuid,
    generate_model_test_table_array_uuid uuid[],
    generate_model_test_table_xid xid,
    generate_model_test_table_array_xid xid[],
    generate_model_test_table_xml xml,
    generate_model_test_table_array_xml xml[],
    pk bigserial,
    CONSTRAINT generate_model_test_table_pkey PRIMARY KEY (pk)
)

TABLESPACE pg_default;

ALTER TABLE generate_model_test.generate_model_test_table
    OWNER to generate_model_test;

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_char
    IS 'this is generate_model_test_table_char comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_char
    IS 'this is generate_model_test_table_array_char comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_aclitem
    IS 'this is generate_model_test_table_aclitem comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_aclitem
    IS 'this is generate_model_test_table_array_aclitem comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_bigint
    IS 'this is generate_model_test_table_bigint comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_bigint
    IS 'this is generate_model_test_table_array_bigint comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_bigserial
    IS 'this is generate_model_test_table_bigserial comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_bit
    IS 'this is generate_model_test_table_bit comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_bit
    IS 'this is generate_model_test_table_array_bit comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_bit_varying
    IS 'this is generate_model_test_table_bit_varying comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_bit_varying
    IS 'this is generate_model_test_table_array_bit_varying comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_boolean
    IS 'this is generate_model_test_table_boolean comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_boolean
    IS 'this is generate_model_test_table_array_boolean comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_box
    IS 'this is generate_model_test_table_box comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_box
    IS 'this is generate_model_test_table_array_box comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_bytea
    IS 'this is generate_model_test_table_bytea comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_bytea
    IS 'this is generate_model_test_table_array_bytea comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_character
    IS 'this is generate_model_test_table_character comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_character
    IS 'this is generate_model_test_table_array_character comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_character_varying
    IS 'this is generate_model_test_table_character_varying comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_character_varying
    IS 'this is generate_model_test_table_array_character_varying comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_cid
    IS 'this is generate_model_test_table_cid comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_cid
    IS 'this is generate_model_test_table_array_cid comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_cidr
    IS 'this is generate_model_test_table_cidr comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_cidr
    IS 'this is generate_model_test_table_array_cidr comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_circle
    IS 'this is generate_model_test_table_circle comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_circle
    IS 'this is generate_model_test_table_array_circle comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_date
    IS 'this is generate_model_test_table_date comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_date
    IS 'this is generate_model_test_table_array_date comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_daterange
    IS 'this is generate_model_test_table_daterange comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_daterange
    IS 'this is generate_model_test_table_array_daterange comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_double_precision
    IS 'this is generate_model_test_table_double_precision comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_double_precision
    IS 'this is generate_model_test_table_array_double_precision comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_gtsvector
    IS 'this is generate_model_test_table_gtsvector comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_gtsvector
    IS 'this is generate_model_test_table_array_gtsvector comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_inet
    IS 'this is generate_model_test_table_inet comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_inet
    IS 'this is generate_model_test_table_array_inet comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_int2vector
    IS 'this is generate_model_test_table_int2vector comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_int2vector
    IS 'this is generate_model_test_table_array_int2vector comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_int4range
    IS 'this is generate_model_test_table_int4range comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_int4range
    IS 'this is generate_model_test_table_array_int4range comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_int8range
    IS 'this is generate_model_test_table_int8range comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_int8range
    IS 'this is generate_model_test_table_array_int8range comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_integer
    IS 'this is generate_model_test_table_integer comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_integer
    IS 'this is generate_model_test_table_array_integer comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_interval
    IS 'this is generate_model_test_table_interval comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_interval
    IS 'this is generate_model_test_table_array_interval comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_json
    IS 'this is generate_model_test_table_json comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_json
    IS 'this is generate_model_test_table_array_json comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_jsonb
    IS 'this is generate_model_test_table_jsonb comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_jsonb
    IS 'this is generate_model_test_table_array_jsonb comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_jsonpath
    IS 'this is generate_model_test_table_jsonpath comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_jsonpath
    IS 'this is generate_model_test_table_array_jsonpath comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_line
    IS 'this is generate_model_test_table_line comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_line
    IS 'this is generate_model_test_table_array_line comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_lseg
    IS 'this is generate_model_test_table_lseg comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_lseg
    IS 'this is generate_model_test_table_array_lseg comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_macaddr
    IS 'this is generate_model_test_table_macaddr comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_macaddr8
    IS 'this is generate_model_test_table_macaddr8 comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_macaddr8
    IS 'this is generate_model_test_table_array_macaddr8 comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_macaddr
    IS 'this is generate_model_test_table_array_macaddr comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_money
    IS 'this is generate_model_test_table_money comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_money
    IS 'this is generate_model_test_table_array_money comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_name
    IS 'this is generate_model_test_table_name comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_name
    IS 'this is generate_model_test_table_array_name comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_numeric
    IS 'this is generate_model_test_table_numeric comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_numeric
    IS 'this is generate_model_test_table_array_numeric comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_numrange
    IS 'this is generate_model_test_table_numrange comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_numrange
    IS 'this is generate_model_test_table_array_numrange comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_oid
    IS 'this is generate_model_test_table_oid comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_oid
    IS 'this is generate_model_test_table_array_oid comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_oidvector
    IS 'this is generate_model_test_table_oidvector comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_oidvector
    IS 'this is generate_model_test_table_array_oidvector comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_path
    IS 'this is generate_model_test_table_path comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_path
    IS 'this is generate_model_test_table_array_path comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_pg_dependencies
    IS 'this is generate_model_test_table_pg_dependencies comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_pg_lsn
    IS 'this is generate_model_test_table_pg_lsn comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_pg_lsn
    IS 'this is generate_model_test_table_array_pg_lsn comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_pg_mcv_list
    IS 'this is generate_model_test_table_pg_mcv_list comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_pg_ndistinct
    IS 'this is generate_model_test_table_pg_ndistinct comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_pg_node_tree
    IS 'this is generate_model_test_table_pg_node_tree comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_point
    IS 'this is generate_model_test_table_point comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_point
    IS 'this is generate_model_test_table_array_point comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_polygon
    IS 'this is generate_model_test_table_polygon comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_polygon
    IS 'this is generate_model_test_table_array_polygon comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_real
    IS 'this is generate_model_test_table_real comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_real
    IS 'this is generate_model_test_table_array_real comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_refcursor
    IS 'this is generate_model_test_table_refcursor comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_refcursor
    IS 'this is generate_model_test_table_array_refcursor comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_regclass
    IS 'this is generate_model_test_table_regclass comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_regclass
    IS 'this is generate_model_test_table_array_regclass comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_regconfig
    IS 'this is generate_model_test_table_regconfig comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_regconfig
    IS 'this is generate_model_test_table_array_regconfig comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_regdictionary
    IS 'this is generate_model_test_table_regdictionary comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_regdictionary
    IS 'this is generate_model_test_table_array_regdictionary comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_regnamespace
    IS 'this is generate_model_test_table_regnamespace comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_regnamespace
    IS 'this is generate_model_test_table_array_regnamespace comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_regoper
    IS 'this is generate_model_test_table_regoper comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_regoper
    IS 'this is generate_model_test_table_array_regoper comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_regoperator
    IS 'this is generate_model_test_table_regoperator comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_regoperator
    IS 'this is generate_model_test_table_array_regoperator comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_regproc
    IS 'this is generate_model_test_table_regproc comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_regproc
    IS 'this is generate_model_test_table_array_regproc comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_regprocedure
    IS 'this is generate_model_test_table_regprocedure comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_regprocedure
    IS 'this is generate_model_test_table_array_regprocedure comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_regrole
    IS 'this is generate_model_test_table_regrole comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_regrole
    IS 'this is generate_model_test_table_array_regrole comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_regtype
    IS 'this is generate_model_test_table_regtype comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_regtype
    IS 'this is generate_model_test_table_array_regtype comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_serial
    IS 'this is generate_model_test_table_serial comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_smallint
    IS 'this is generate_model_test_table_smallint comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_smallint
    IS 'this is generate_model_test_table_array_smallint comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_smallserial
    IS 'this is generate_model_test_table_smallserial comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_text
    IS 'this is generate_model_test_table_text comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_text
    IS 'this is generate_model_test_table_array_text comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_tid
    IS 'this is generate_model_test_table_tid comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_tid
    IS 'this is generate_model_test_table_array_tid comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_time_with_time_zone
    IS 'this is generate_model_test_table_time_with_time_zone comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_time_with_time_zone
    IS 'this is generate_model_test_table_array_time_with_time_zone comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_time_without_time_zone
    IS 'this is generate_model_test_table_time_without_time_zone comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_time_without_time_zone
    IS 'this is generate_model_test_table_array_time_without_time_zone comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_timestamp_with_time_zone
    IS 'this is generate_model_test_table_timestamp_with_time_zone comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_timestamp_with_time_zone
    IS 'this is generate_model_test_table_array_timestamp_with_time_zone comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_timestamp_without_time_zone
    IS 'this is generate_model_test_table_timestamp_without_time_zone comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_timestamp_without_time_zone
    IS 'this is generate_model_test_table_array_timestamp_without_time_zone comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_tsquery
    IS 'this is generate_model_test_table_tsquery comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_tsquery
    IS 'this is generate_model_test_table_array_tsquery comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_tsrange
    IS 'this is generate_model_test_table_tsrange comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_tsrange
    IS 'this is generate_model_test_table_array_tsrange comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_tstzrange
    IS 'this is generate_model_test_table_tstzrange comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_tstzrange
    IS 'this is generate_model_test_table_array_tstzrange comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_tsvector
    IS 'this is generate_model_test_table_tsvector comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_tsvector
    IS 'this is generate_model_test_table_array_tsvector comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_txid_snapshot
    IS 'this is generate_model_test_table_txid_snapshot comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_txid_snapshot
    IS 'this is generate_model_test_table_array_txid_snapshot comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_uuid
    IS 'this is generate_model_test_table_uuid comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_uuid
    IS 'this is generate_model_test_table_array_uuid comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_xid
    IS 'this is generate_model_test_table_xid comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_xid
    IS 'this is generate_model_test_table_array_xid comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_xml
    IS 'this is generate_model_test_table_xml comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.generate_model_test_table_array_xml
    IS 'this is generate_model_test_table_array_xml comment';

COMMENT ON COLUMN generate_model_test.generate_model_test_table.pk
    IS 'this is pk comment';
