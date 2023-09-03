BEGIN;
SELECT plan(21);

SELECT has_table('sessions');

SELECT has_column('sessions', 'id');
SELECT col_type_is('sessions', 'id', 'uuid');
SELECT col_not_null('sessions', 'id');

SELECT has_column('sessions', 'user_id');
SELECT col_type_is('sessions', 'user_id', 'uuid');
SELECT col_not_null('sessions', 'user_id');

SELECT has_column('sessions', 'token');
SELECT col_type_is('sessions', 'token', 'uuid');
SELECT col_not_null('sessions', 'token');

SELECT has_column('sessions', 'expires_at');
SELECT col_type_is('sessions', 'expires_at', 'timestamp without time zone');
SELECT col_not_null('sessions', 'expires_at');

SELECT has_column('sessions', 'created_at');
SELECT col_type_is('sessions', 'created_at', 'timestamp without time zone');
SELECT col_default_is('sessions', 'created_at', 'now()');

SELECT has_column('sessions', 'last_updated_at');
SELECT col_type_is('sessions', 'last_updated_at', 'timestamp without time zone');
SELECT col_not_null('sessions', 'last_updated_at');

SELECT index_is_unique('sessions', 'uidx_sessions_user_id');
SELECT has_index('sessions', 'idx_sessions_token');

ROLLBACK;
