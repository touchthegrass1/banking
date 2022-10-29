CREATE OR REPLACE PROCEDURE public.create_card(valid_to DATE, cvc_code VARCHAR, client_id BIGINT) AS $$
begin
	insert into public.card VALUES(floor(random() * 1000000000000)::int, 0, $1, $2, 'debit', 'RUB', $3);
end;
$$ language plpgsql;