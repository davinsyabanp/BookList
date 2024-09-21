# books
 braincore


''' CREATE TABLE categories (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY, -- Gunakan BIGINT UNSIGNED
    name VARCHAR(255) UNIQUE NOT NULL
);
CREATE TABLE books (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    publication_date DATE NOT NULL,
    publisher VARCHAR(255) NOT NULL,
    num_pages BIGINT NOT NULL,
    category_id BIGINT UNSIGNED ,  -- Sesuaikan tipe data dengan categories.id
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);
'''
	
	
