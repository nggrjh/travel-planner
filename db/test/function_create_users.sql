BEGIN;
SELECT plan(9);

SELECT is_empty(
    'SELECT username, email, password FROM users WHERE username = ''hariiniindah'';'
);

SELECT is_empty(
    'SELECT username, email, password FROM users WHERE email = ''hariiniindah@gmail.com'';'
);

-- create success
SELECT performs_ok(
    'SELECT create_user(''hariiniindah'', ''hariiniindah@gmail.com'', ''hariiniindah'');',
    40
);

SELECT results_eq(
    'SELECT username, email, password FROM users WHERE username = ''hariiniindah'';',
    'SELECT ''hariiniindah''::VARCHAR, ''hariiniindah@gmail.com''::VARCHAR, ''hariiniindah''::VARCHAR;'
);

SELECT results_eq(
    'SELECT username, email, password FROM users WHERE email = ''hariiniindah@gmail.com'';',
    'SELECT ''hariiniindah''::VARCHAR, ''hariiniindah@gmail.com''::VARCHAR, ''hariiniindah''::VARCHAR;'
);

-- create failed when username exists
SELECT throws_matching(
   'SELECT create_user(''hariiniindah'', ''hariiniindah2@gmail.com'', ''hariiniindah'');',
   'constraint uidx_users_username violated'
);

SELECT is_empty(
    'SELECT username, email, password FROM users WHERE email = ''hariiniindah2@gmail.com'';'
);

-- create failed when email exists
SELECT throws_matching(
   'SELECT create_user(''hariiniindah2'', ''hariiniindah@gmail.com'', ''hariiniindah'');',
   'constraint uidx_users_email violated'
);

SELECT is_empty(
    'SELECT username, email, password FROM users WHERE username = ''hariiniindah2'';'
);

ROLLBACK;
