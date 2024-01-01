create table tx_bill (
                         id serial primary key,
                         bill_date date,
                         entry_date date,
                         finish_date date,
                         employee_id integer,
                         customer_id integer,
                         foreign key (employee_id) references mst_employee(id),
                         foreign key (customer_id) references mst_customer(id)
);