CREATE OR REPLACE FUNCTION public.get_client_cards(phone VARCHAR) RETURNS public.card AS $$
    BEGIN
        SELECT * FROM public.card WHERE client_id = (
            SELECT client_id
            FROM public.client
            WHERE "phone" = $1
        );
    END;
$$ LANGUAGE plpgsql STABLE STRICT;;