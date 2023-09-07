BEGIN;

SELECT plan(4);

SELECT is_empty(
    'SELECT email, password FROM users WHERE email = ''hariiniindah@gmail.com'';'
);

-- create success
SELECT performs_ok(
    'SELECT create_user(''hariiniindah@gmail.com'', ''hariiniindah'');',
    40
);

SELECT results_eq(
    'SELECT email, password FROM users WHERE email = ''hariiniindah@gmail.com'';',
    'SELECT ''hariiniindah@gmail.com''::VARCHAR, ''hariiniindah''::VARCHAR;'
);

-- create failed when email exists
SELECT throws_matching(
   'SELECT create_user(''hariiniindah@gmail.com'', ''hariiniindah'');',
   'constraint uidx_users_email violated'
);

ROLLBACK;
