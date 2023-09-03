BEGIN;
SELECT plan(2);

-- data does not exist
SELECT is_empty(
    'SELECT * FROM get_user_by_email(''hariiniindah@gmail.com'');'
);

SELECT create_user('hariiniindah@gmail.com', 'hariiniindah');

-- data exists
SELECT results_eq(
    'SELECT * FROM get_user_by_email(''hariiniindah@gmail.com'');',
    'SELECT ''hariiniindah@gmail.com''::VARCHAR, ''hariiniindah''::VARCHAR;'
);

ROLLBACK;
