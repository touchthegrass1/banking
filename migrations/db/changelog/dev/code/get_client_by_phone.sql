CREATE OR REPLACE FUNCTION public.get_client_by_phone(phone VARCHAR) RETURNS SETOF public.client AS $$
    BEGIN
        RETURN QUERY SELECT * FROM public.client WHERE "phone" = phone;
    END;
$$ LANGUAGE plpgsql STABLE STRICT;;