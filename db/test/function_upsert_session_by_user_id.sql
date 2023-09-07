BEGIN;

SELECT plan(5);

SELECT is_empty(
    'SELECT token, refreshed_at, expires_at FROM sessions WHERE user_id = ''4779b66e-8820-4e96-b011-fe77d28402b0'';'
);

-- insert success
SELECT performs_ok(
    'SELECT upsert_session_by_user_id(''4779b66e-8820-4e96-b011-fe77d28402b0'', ''c07a6504-7a58-477d-8f29-c1c5b7179858'', ''2022-12-22 00:00:00'', ''2022-12-25 00:00:00'');',
    40
);

SELECT results_eq(
    'SELECT token, refreshed_at, expires_at FROM sessions WHERE user_id = ''4779b66e-8820-4e96-b011-fe77d28402b0'';',
    'SELECT ''c07a6504-7a58-477d-8f29-c1c5b7179858''::UUID, ''2022-12-22 00:00:00''::TIMESTAMP, ''2022-12-25 00:00:00''::TIMESTAMP;'
);

-- update success
SELECT performs_ok(
    'SELECT upsert_session_by_user_id(''4779b66e-8820-4e96-b011-fe77d28402b0'', ''171f1055-525a-4195-a929-dae820e141d1'', ''2023-02-13 00:00:00'', ''2023-02-15 00:00:00'');',
    40
);

SELECT results_eq(
    'SELECT token, refreshed_at, expires_at FROM sessions WHERE user_id = ''4779b66e-8820-4e96-b011-fe77d28402b0'';',
    'SELECT ''171f1055-525a-4195-a929-dae820e141d1''::UUID, ''2023-02-13 00:00:00''::TIMESTAMP, ''2023-02-15 00:00:00''::TIMESTAMP;'
);

ROLLBACK;
