CREATE TABLE events(
    id UUID Primary key DEFAULT gen_random_uuid(),
    user_id INT,
    activity VARCHAR(255) NOT NULL,
    product_id INT,
    happened_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
