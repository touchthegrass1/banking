create or replace procedure public.withdraw_money(card_id int, summ decimal) as $$
begin
	update public.card set balance = balance - $2 where card_id = $1;
	insert into public.transaction(transaction_type, card_id, summ, transaction_datetime) VALUES('withdraw', $1, $2, NOW());
end
$$ language plpgsql;
