CREATE OR REPLACE FUNCTION banking.get_client_cards(phone VARCHAR) RETURNS banking.card AS $$
    BEGIN
        SELECT * FROM banking.card WHERE client_id = (
            SELECT client_id
            FROM banking.client
            WHERE "phone" = $1
        );
    END;
$$ LANGUAGE plpgsql STABLE STRICT;;
-- rollback DROP FUNCTION IF EXISTS banking.get_client_cards(varchar);