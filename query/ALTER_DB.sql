alter table carts add column is_active int DEFAULT 1;

alter table orders add column payment varchar(25);
alter table orders add column payment_status int default 1;