create or replace procedure banking.create_client(client_name varchar, phone varchar, registration_address varchar, residential_address varchar, client_type varchar, ogrn int default null, inn int default null, kpp int default null) as $$
declare
	client_id integer;
begin
	insert into banking.client(name, phone, registration_address, residential_address, client_type, ogrn, inn, kpp)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8) returning client_id into client_id;
	insert into banking.contract("contract_type", "conclusion_date", "contract_content", "client_id") VALUES('settlement_and_cash_service_agreement', CURRENT_DATE, "bla bla", client_id);
end;
$$ language plpgsql;
