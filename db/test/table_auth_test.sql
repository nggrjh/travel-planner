BEGIN;
SELECT plan(20);

SELECT has_table('auth');

SELECT has_column('auth', 'id');
SELECT col_type_is('auth', 'id', 'uuid');
SELECT col_not_null('auth', 'id');

SELECT has_column('auth', 'username');
SELECT col_type_is('auth', 'username', 'character varying');
SELECT col_not_null('auth', 'username');

SELECT has_column('auth', 'email');
SELECT col_type_is('auth', 'email', 'character varying');
SELECT col_not_null('auth', 'email');

SELECT has_column('auth', 'password');
SELECT col_type_is('auth', 'password', 'character varying');
SELECT col_is_null('auth', 'password');

SELECT has_column('auth', 'created_at');
SELECT col_type_is('auth', 'created_at', 'timestamp without time zone');
SELECT col_default_is('auth', 'created_at', 'now()');

SELECT has_column('auth', 'last_modified_at');
SELECT col_type_is('auth', 'last_modified_at', 'timestamp without time zone');

SELECT index_is_unique('auth', 'uidx_auth_username');
SELECT index_is_unique('auth', 'uidx_auth_email');

ROLLBACK;
