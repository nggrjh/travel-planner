BEGIN;
SELECT plan(9);

SELECT is_empty(
    'SELECT username, email, password FROM auth WHERE username = ''hariiniindah'';'
);

SELECT is_empty(
    'SELECT username, email, password FROM auth WHERE email = ''hariiniindah@gmail.com'';'
);

-- insert success
SELECT performs_ok(
    'SELECT insert_auth(''hariiniindah'', ''hariiniindah@gmail.com'', ''hariiniindah'');',
    40
);

SELECT results_eq(
    'SELECT username, email, password FROM auth WHERE username = ''hariiniindah'';',
    'SELECT ''hariiniindah''::VARCHAR, ''hariiniindah@gmail.com''::VARCHAR, ''hariiniindah''::VARCHAR;'
);

SELECT results_eq(
    'SELECT username, email, password FROM auth WHERE email = ''hariiniindah@gmail.com'';',
    'SELECT ''hariiniindah''::VARCHAR, ''hariiniindah@gmail.com''::VARCHAR, ''hariiniindah''::VARCHAR;'
);

-- insert failed when username exists
SELECT throws_matching(
   'SELECT insert_auth(''hariiniindah'', ''hariiniindah2@gmail.com'', ''hariiniindah'');',
   'constraint uidx_auth_username violated'
);

SELECT is_empty(
    'SELECT username, email, password FROM auth WHERE email = ''hariiniindah2@gmail.com'';'
);

-- insert failed when email exists
SELECT throws_matching(
   'SELECT insert_auth(''hariiniindah2'', ''hariiniindah@gmail.com'', ''hariiniindah'');',
   'constraint uidx_auth_email violated'
);

SELECT is_empty(
    'SELECT username, email, password FROM auth WHERE username = ''hariiniindah2'';'
);

ROLLBACK;
