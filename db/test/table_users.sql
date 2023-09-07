BEGIN;

SELECT plan(16);

SELECT has_table('users');

SELECT has_column('users', 'id');
SELECT col_type_is('users', 'id', 'uuid');
SELECT col_not_null('users', 'id');

SELECT has_column('users', 'email');
SELECT col_type_is('users', 'email', 'character varying');
SELECT col_not_null('users', 'email');

SELECT has_column('users', 'password');
SELECT col_type_is('users', 'password', 'character varying');
SELECT col_not_null('users', 'password');

SELECT has_column('users', 'created_at');
SELECT col_type_is('users', 'created_at', 'timestamp without time zone');
SELECT col_default_is('users', 'created_at', 'now()');

SELECT has_column('users', 'last_modified_at');
SELECT col_type_is('users', 'last_modified_at', 'timestamp without time zone');

SELECT index_is_unique('users', 'uidx_users_email');

ROLLBACK;
