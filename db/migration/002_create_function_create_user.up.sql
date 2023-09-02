CREATE OR REPLACE FUNCTION create_user(username_in VARCHAR,
    email_in VARCHAR,
    password_in VARCHAR)
    RETURNS VOID
    LANGUAGE plpgsql AS
$$

DECLARE
   constraint_name text;
BEGIN
   INSERT INTO users(username, email, password)
      VALUES (username_in, email_in, password_in);
EXCEPTION
   WHEN unique_violation THEN
      -- Get the name of the violated constraint from the error message
      GET STACKED DIAGNOSTICS constraint_name = CONSTRAINT_NAME;
      
      -- Handle the unique constraint violation and display the constraint name
      RAISE EXCEPTION 'constraint % violated.', constraint_name;
   WHEN others THEN
      -- Handle other exceptions if necessary
      RAISE;
END;
$$;
