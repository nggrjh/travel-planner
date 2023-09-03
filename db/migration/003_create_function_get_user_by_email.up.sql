CREATE OR REPLACE FUNCTION get_user_by_email(
   email_in VARCHAR
)
   RETURNS table (
      email_out VARCHAR,
      password_out VARCHAR
   )
   LANGUAGE plpgsql AS
$$

BEGIN
   RETURN QUERY
      SELECT email, password 
      FROM users
      WHERE email = email_in;
END;
$$;
