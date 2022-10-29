ALTER TABLE "transaction" DROP CONSTRAINT transaction_card_id_fkey;
ALTER TABLE "transaction" DROP CONSTRAINT transaction_card_id_from_fkey;
ALTER TABLE "transaction" DROP CONSTRAINT transaction_card_id_to_fkey;

ALTER TABLE "transaction" ALTER COLUMN card_id TYPE VARCHAR(12) USING card_id::VARCHAR;
ALTER TABLE "transaction" ALTER COLUMN card_from_id TYPE VARCHAR(12) USING card_from_id::VARCHAR;
ALTER TABLE "transaction" ALTER COLUMN card_to_id TYPE VARCHAR(12) USING card_to_id::VARCHAR;

ALTER TABLE card ALTER COLUMN card_id TYPE VARCHAR(12) USING card_id::VARCHAR;

ALTER TABLE "transaction" ADD CONSTRAINT transaction_card_id_fkey FOREIGN KEY (card_id) REFERENCES card(card_id);
ALTER TABLE "transaction" ADD CONSTRAINT transaction_card_from_id_fkey FOREIGN KEY (card_from_id) REFERENCES card(card_id);
ALTER TABLE "transaction" ADD CONSTRAINT transaction_card_to_id_fkey FOREIGN KEY (card_to_id) REFERENCES card(card_id);
