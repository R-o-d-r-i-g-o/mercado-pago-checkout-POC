CREATE OR REPLACE FUNCTION validate_email()
RETURNS TRIGGER AS $$
BEGIN
    -- Check if the email is a valid email address
    IF NEW.email IS NOT NULL AND NEW.email !~* '^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$' THEN
        RAISE EXCEPTION 'Invalid email address';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER validate_email_trigger
BEFORE INSERT OR UPDATE OF email ON users
FOR EACH ROW
EXECUTE FUNCTION validate_email();
