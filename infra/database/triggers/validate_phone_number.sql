CREATE OR REPLACE FUNCTION clean_phone_number()
RETURNS TRIGGER AS $$
BEGIN
    -- Remove special characters and spaces from the phone number
    IF NEW.phone IS NOT NULL THEN
        NEW.phone := regexp_replace(NEW.phone, '[^0-9]', '', 'g');
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER clean_phone_number_trigger
BEFORE INSERT OR UPDATE OF phone ON users
FOR EACH ROW
EXECUTE FUNCTION clean_phone_number();
