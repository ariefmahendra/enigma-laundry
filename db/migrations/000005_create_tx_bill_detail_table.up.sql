create table tx_bill_detail(
                               id serial primary key,
                               bill_id integer,
                               product_id integer,
                               quantity integer,
                               product_price integer,
                               foreign key (bill_id) references tx_bill(id),
                               foreign key (product_id) references mst_product(id)
);