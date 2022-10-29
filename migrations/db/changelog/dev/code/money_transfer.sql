create or replace procedure public.money_transfer(card_id_from int, card_id_to int, summ decimal) as $$
begin
	set transaction isolation level serializable;
	update public.card set balance = balance - $3 where card_id = $1;
	update public.card set balance = balance + $3 where card_id = $2;
	insert into public.transaction(transaction_type, card_id_from, card_id_to, summ, transaction_datetime) VALUES('transfer', $1, $2, $3, NOW());
end
$$ language plpgsql;
