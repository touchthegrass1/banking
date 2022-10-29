CREATE OR REPLACE FUNCTION public.truncate_tables(username VARCHAR, table_names varchar[]) RETURNS void LANGUAGE plpgsql AS
$$
DECLARE
    table_ text;
    schema_ text;
BEGIN
    FOR schema_, table_ IN
        SELECT *
        FROM pg_tables
        WHERE tableowner = 'banking_db_user' AND tablename = ANY(table_names)
    LOOP

    EXECUTE FORMAT('TRUNCATE TABLE %I.%I RESTART IDENTITY CASCADE', schema_, table_);

    END LOOP;
END
$$;
