-- удаляем таблицы, если они существуют
drop table if exists product_warehouse;
drop table if exists product;
drop table if exists warehouse;


create table Warehouse (
                           id bigint primary key ,
                           name varchar not null,
                           availability boolean not null default true
);

CREATE TABLE product (
                         name varchar not null,
    -- уникальный код продукта
                         code bigint primary key,
                         size int not null,
    -- уникальная комбинация полей name и size
                         constraint uq_product_name_size unique (name, size)
);

-- таблица где хранится информация о том, какой продукт на каком складе

create table product_warehouse (
                                   product_code BIGINT not null,
                                   warehouse_id bigint not null,
                                   count int not null,
                                   reserved int not null default 0,

                                   constraint fk_product_warehouse_product_code foreign key  (product_code) references product(code),
                                   constraint fk_product_warehouse_warehouse_id foreign key  (warehouse_id) references warehouse(id),
    -- reserved не может быть больше чем count и не может быть меньше 0
                                   constraint chk_product_warehouse_reserved check (reserved <= count and reserved >= 0 and count >= 0)

);



INSERT INTO product (name, code, size) VALUES ('product1', 23, 10);
INSert into product (name, code, size) VALUES ('product2', 24, 10);
INSERT INTO product (name, code, size) VALUES ('Шлепки', 25, 10);
INSERT INTO product (name, code, size) VALUES ('Шорты', 26, 10);
INSERT INTO product (name, code, size) VALUES ('Штаны', 27, 10);
INSERT INTO product (name, code, size) VALUES ('Шапка', 28, 10);
INSERT INTO product (name, code, size) VALUES ('Шарф', 29, 10);
INSERT INTO product (name, code, size) VALUES ('Costume', 30, 10);

INSERT INTO warehouse (id, name, availability) VALUES (1, 'Mine', true);

-- Добавим несколько резерваций в таблицу product_warehouse

INSERT INTO product_warehouse (product_code, warehouse_id, count, reserved) VALUES (23, 1, 10, 0);

-- добавим к product_warehouse еще 10 товаров с кодом 23, было 10 стало 20

UPDATE product_warehouse SET count = count + 10 WHERE product_code = 23;

-- Запрос на резервацию 25 товаров с кодом 23

UPDATE product_warehouse SET reserved = reserved + 25 WHERE product_code = 23;