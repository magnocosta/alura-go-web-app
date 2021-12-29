CREATE TABLE products (
  id serial primary key,
  name varchar,
  description varchar,
  price decimal,
  quantity integer
);

INSERT INTO products (name, description, price, quantity) VALUES
('Camiseta Preta', 'Com estampa dragon ball Z', 50, 10),
('Camiseta Branca', 'Com estampa dragon ball Z', 80, 10);
