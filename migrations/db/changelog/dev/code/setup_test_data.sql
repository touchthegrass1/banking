CREATE OR REPLACE PROCEDURE public.p_setup_test_data() AS $$
DECLARE
    streets VARCHAR[] := array[
        'Ленина',
        'Гагарина',
        '1 Мая',
        'Маркса',
        'Энгельса',
        'Пушкина',
        'Лермонтова',
        'Лесная',
        'Бакинская',
        'Боевая',
        'Магистральная'
    ];
    first_names VARCHAR[] := array[
        'Пётр', 'Никита', 'Илья', 'Василий'
    ];
    last_names VARCHAR[] := array[
        'Петров', 'Иванов', 'Попов', 'Васильев', 'Королёв'
    ];
    start_datetime TIMESTAMP WITH TIME ZONE := '2010-01-01 00:00:00.000000'::TIMESTAMP without time zone at time zone 'Europe/Saratov';
    card_valid_to_start_datetime TIMESTAMP WITH TIME ZONE := '2020-01-01 00:00:00.000000'::TIMESTAMP without time zone at time zone 'Europe/Saratov';
    currencies VARCHAR[] := array['USD', 'EUR', 'RUB', 'GBP', 'JPY'];
BEGIN
    WITH DataCTE AS (
        SELECT
            generate_series(1, 10000) AS id,
            md5(random()::text) AS password,
            start_datetime + random() * (timestamp '2022-10-30 00:00:00' - timestamp '2010-01-01 00:00:00.000000') AS last_login,
            false AS is_superuser,
            (first_names)[floor(random() * 4 + 1)] AS first_name,
            (last_names)[floor(random() * 5 + 1)] AS last_name,
            false AS is_staff,
            true AS is_active,
            FLOOR(RANDOM() * 10)::varchar || FLOOR((RANDOM() * (9999999999 - 1000000000) + 1000000000))::varchar AS phone,
            'example@gmail.com' AS email
    )
    INSERT INTO public.user(id, password, last_login, is_superuser, first_name, last_name, is_staff, is_active, date_joined, phone, email)
    SELECT
        DataCTE.id,
        DataCTE.password,
        DataCTE.last_login,
        DataCTE.is_superuser,
        DataCTE.first_name,
        DataCTE.last_name,
        DataCTE.is_staff,
        DataCTE.is_active,
        DataCTE.last_login + random() * (timestamp '2022-10-30 00:00:00' - timestamp '2015-01-01 00:00:00.000000'),
        DataCTE.phone,
        DataCTE.email
    FROM DataCTE;

    WITH DataCTE AS (
        SELECT
            generate_series(1, 10000) AS client_id,
            ('ул. ' || (streets)[floor(random() * 8 + 1)::int] || ' д. ' || FLOOR(RANDOM()*(100 - 1) + 1)::varchar) AS registration_address,
            'individual'::client_type AS client_type,
            FLOOR(RANDOM() * 10)::varchar || FLOOR((RANDOM() * (99999999999 - 10000000000) + 10000000000))::varchar AS ogrn,
            FLOOR(RANDOM() * 10)::varchar || FLOOR((RANDOM() * (99999999999 - 10000000000) + 10000000000))::varchar AS inn,
            FLOOR(RANDOM() * 10)::varchar || FLOOR((RANDOM() * (99999999999 - 10000000000) + 10000000000))::varchar AS kpp
    )
    INSERT INTO public.client(client_id, registration_address, residential_address, client_type, ogrn, inn, kpp, user_id)
    SELECT
        DataCTE.client_id,
        DataCTE.registration_address,
        DataCTE.registration_address,
        DataCTE.client_type,
        DataCTE.ogrn,
        DataCTE.inn,
        DataCTE.kpp,
        DataCTE.client_id
    FROM DataCTE;

    WITH DataCTE AS (
        SELECT
            generate_series(1, 10000) AS id,
            FLOOR(RANDOM() * (999999999999::bigint - 100000000000::bigint) + 100000000000::bigint)::bigint AS card_id,
            RANDOM() * 100000 + 1 AS balance,
            card_valid_to_start_datetime + random() * (timestamp '2022-10-30 00:00:00' - timestamp '2018-01-01 00:00:00.000000') AS valid_to,
            (RANDOM() * (999 - 100) + 100)::int AS cvc_code,
            'debit'::card_type AS card_type,
            (currencies)[(RANDOM() * (array_length(currencies, 1)) + 1)::int] AS currency
    )
    INSERT INTO public.card(card_id, balance, valid_to, cvc_code, card_type, currency, client_id)
    SELECT
        DataCTE.card_id,
        DataCTE.balance,
        DataCTE.valid_to,
        DataCTE.cvc_code,
        DataCTE.card_type,
        DataCTE.currency,
        DataCTE.id
    FROM DataCTE;
END;
$$
LANGUAGE plpgsql;