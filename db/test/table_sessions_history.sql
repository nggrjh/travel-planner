BEGIN;

SELECT plan(25);

SELECT has_table('sessions_history');

SELECT has_column('sessions_history', 'id');
SELECT col_type_is('sessions_history', 'id', 'uuid');
SELECT col_not_null('sessions_history', 'id');

SELECT has_column('sessions_history', 'session_id');
SELECT col_type_is('sessions_history', 'session_id', 'uuid');
SELECT col_not_null('sessions_history', 'session_id');

SELECT has_column('sessions_history', 'token');
SELECT col_type_is('sessions_history', 'token', 'uuid');
SELECT col_not_null('sessions_history', 'token');

SELECT has_column('sessions_history', 'refreshed_at');
SELECT col_type_is('sessions_history', 'refreshed_at', 'timestamp without time zone');
SELECT col_not_null('sessions_history', 'refreshed_at');

SELECT has_column('sessions_history', 'expires_at');
SELECT col_type_is('sessions_history', 'expires_at', 'timestamp without time zone');
SELECT col_not_null('sessions_history', 'expires_at');

SELECT has_column('sessions_history', 'created_at');
SELECT col_type_is('sessions_history', 'created_at', 'timestamp without time zone');
SELECT col_default_is('sessions_history', 'created_at', 'now()');

SELECT has_column('sessions_history', 'operation');
SELECT col_type_is('sessions_history', 'operation', 'character varying');
SELECT col_not_null('sessions_history', 'operation');

SELECT has_column('sessions_history', 'modified_by');
SELECT col_type_is('sessions_history', 'modified_by', 'character varying');
SELECT col_not_null('sessions_history', 'modified_by');

ROLLBACK;
