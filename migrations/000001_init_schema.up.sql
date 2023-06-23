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
                                   product_name varchar not null,
                                   count int not null,
                                   reserved int not null default 0,

    -- уникальная комбинация полей product_code и warehouse_id
                                   constraint uq_product_warehouse_product_code_warehouse_id unique (product_code, warehouse_id),

                                   constraint fk_product_warehouse_product_code foreign key  (product_code) references product(code),
                                   constraint fk_product_warehouse_warehouse_id foreign key  (warehouse_id) references warehouse(id),
    -- reserved не может быть больше чем count и не может быть меньше 0
                                   constraint chk_product_warehouse_reserved check (reserved <= count and reserved >= 0 and count >= 0)

);



INSERT INTO warehouse (id, name, availability) VALUES (1, 'Mine', true);
INSERT INTO warehouse (id, name, availability) VALUES (2, 'Pops', true);


INSERT INTO product (name, code, size) VALUES ('product1', 23, 10);
INSert into product (name, code, size) VALUES ('product2', 24, 10);
INSert into product (name, code, size) VALUES ('Анастасия', 22, 10);
INSERT INTO product (name, code, size) VALUES ('Шлепки', 25, 10);
INSERT INTO product (name, code, size) VALUES ('Шорты', 26, 10);
INSERT INTO product (name, code, size) VALUES ('Штаны', 27, 10);
INSERT INTO product (name, code, size) VALUES ('Шапка', 28, 10);
INSERT INTO product (name, code, size) VALUES ('Шарф', 29, 10);
INSERT INTO product (name, code, size) VALUES ('Costume', 30, 10);
INSERT INTO product (name, code, size) VALUES ('Платье', 1, 5);
INSERT INTO product (name, code, size) VALUES ('Кроссовки Adidas', 53, 34);

INSERT INTO product (name, code, size) VALUES ('Ноутбук HP Pavilion x360', 319, 1);
INSERT INTO product (name, code, size) VALUES ('Куртка Columbia', 78, 58);
INSERT INTO product (name, code, size) VALUES ('Подушка для беременных', 143, 2);
INSERT INTO product (name, code, size) VALUES ('Книга "Мастер и Маргарита"', 92, 1);
INSERT INTO product (name, code, size) VALUES ('Сковорода Tefal', 51, 28);
INSERT INTO product (name, code, size) VALUES ('Мяч для йоги Reebok', 116, 2);
INSERT INTO product (name, code, size) VALUES ('Чайник Bosch', 205,1);
INSERT INTO product (name, code, size) VALUES ('Беспроводные наушники Sony', 191, 1);
INSERT INTO product (name, code, size) VALUES ('Джинсы Levis', 35, 3);
INSERT INTO product (name, code, size) VALUES ('Рюкзак The North Face', 64, 2);

INSERT INTO product_warehouse (product_code, warehouse_id, product_name, count, reserved)
SELECT
    p.code AS product_code,
    CASE
        WHEN p.code IN (78, 143, 92, 51, 205) THEN 1
        ELSE 2
        END AS warehouse_id,
    p.name AS product_name,
    10 AS count,
    3 AS reserved
FROM product p
WHERE p.code IN (319, 78, 143, 92, 51, 116, 205, 191, 35, 64);