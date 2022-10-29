CREATE OR REPLACE FUNCTION banking.get_client_by_phone(phone VARCHAR) RETURNS SETOF banking.client AS $$
    BEGIN
        RETURN QUERY SELECT * FROM banking.client WHERE "phone" = phone;
    END;
$$ LANGUAGE plpgsql STABLE STRICT;;