Cài đặt MySQL bằng Docker
![image](https://github.com/user-attachments/assets/d5803651-71d5-42b8-9329-d4e90094558d)

Cấu hình MySQL
![image](https://github.com/user-attachments/assets/1a069bfc-6497-49b1-9b3f-65bbab0b10b4)
![image](https://github.com/user-attachments/assets/c728cfd5-d070-46e4-81f6-cbb8be6d4dc1)

Tạo table và insert dữ liệu
![image](https://github.com/user-attachments/assets/926d0bfb-de9f-4724-930a-eb15022ca65d)
CREATE TABLE todo_items (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    image JSON NULL,
    description TEXT NULL,
    status ENUM('Doing', 'Done', 'Deleted') DEFAULT ‘Doing’,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
CREATE INDEX idx_todo_items_status ON todo_items(status);

Insert dữ liệu vào table
![image](https://github.com/user-attachments/assets/a2e552aa-2b51-49fe-b3c9-3f48cb017507)
INSERT INTO todo_items (id, title, image, description, status) VALUES
(1, 'This is taks 1', NULL, 'This is task 1 description', 'Doing'),
(2, 'This is task 2', NULL, 'This is task 2 description', 'Done'),
(3, 'This is task 3', NULL, 'This is task 3 description', 'Doing'),
(4, 'This is task 4', NULL, 'This is task 4 description', 'Deleted'),
(5, 'This is task 5', NULL, 'This is task 5 description', 'Done'),
(6, 'This is task 6', NULL, 'This is task 6 description', 'Doing');




