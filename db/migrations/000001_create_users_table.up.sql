
create table users(
    user_id serial primary key,
    name varchar(255) not null,
    email varchar(255) not null,
    password varchar(255) not null,
    role varchar(25) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp default current_timestamp
);

create table category(
    category_id serial primary key,
    name varchar(255) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp default current_timestamp
  
);

create table product(
    product_id serial primary key,
    name varchar(255) not null,
    quantity int not null,
    seller_id_fk integer not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    CONSTRAINT fk_seller FOREIGN KEY (seller_id_fk) REFERENCES users(user_id)
);

create table orders (
    order_id serial primary key,
    user_id_fk integer not null,
    product_id_fk integer not null,
    quantity integer not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    CONSTRAINT fk_user FOREIGN KEY (user_id_fk) REFERENCES users(user_id),
    CONSTRAINT fk_product FOREIGN KEY (product_id_fk) REFERENCES product(product_id)
);

create table category_product (
    product_id_fk integer not null,
    category_id_fk integer not null,
    CONSTRAINT fk_product FOREIGN KEY (product_id_fk) REFERENCES product(product_id),
    CONSTRAINT fk_category FOREIGN KEY (category_id_fk) REFERENCES category(category_id)
);
