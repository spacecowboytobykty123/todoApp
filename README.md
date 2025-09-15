# 📝 todoApp

Простое приложение для управления задачами с возможностью фильтрации и приоритизации.

---

## ✅ Основные возможности
- Создание, удаление и получение задач
- Установка статуса задачи:
    - `Назначена`
    - `В работе`
    - `Выполнена`
    - `Отклонена`
- Фильтрация задач по статусу и дедлайну
- Возможность задать срок выполнения (deadline)

---

## 💾 Структура проекта

### Backend
- `cmd/api/tasks.go` – сервисный слой (обработка запросов, бизнес-логика)
- `pkg/data/tasks.go` – слой базы данных (CRUD-операции)
- `cmd/api/errors.go` – кастомная обработка ошибок
- `pkg/jsonlog/jsonlog.go` – кастомный логгер
- `pkg/validator/validator.go` – валидация и обработка сообщений

### Миграции базы данных
- `migrations/` – SQL-скрипты для управления схемой БД

### Frontend
- `./src/components/` – основные компоненты приложения
- `./src/App.tsx` – точка входа фронтенда

---

## 🔌 Запуск проекта

1. **Запустить PostgreSQL**  
   Убедитесь, что база доступна по адресу:
   postgres://tasks:pass@localhost/crm?sslmode=disable
2. **Применить миграции**
```migrate -path ./migrations -database "postgres://tasks:pass@localhost/crm?sslmode=disable" up```
3. **Установить зависимости**
    ```go mod tidy
    cd frontend
    npm install
    cd .. 
   ```
###Вывод всех задач
![Вывод всех задач](https://github.com/user-attachments/assets/2f050f4a-6b43-4502-925b-ba6d25a64cff)

Поля для создания задачи(имя описание дата дэдлайна и статус)
<img width="822" height="345" alt="image" src="https://github.com/user-attachments/assets/75ef91cd-1726-47fb-92d3-5730d060adf9" />
Задача добавилась
<img width="800" height="475" alt="image" src="https://github.com/user-attachments/assets/e1d13eea-79bf-4f52-b16b-7fc1f6a8a0c6" />

Поле для изменения задачи
<img width="414" height="368" alt="image" src="https://github.com/user-attachments/assets/aa164dd2-d259-49c1-b0a9-324de22e69e0" />
Задача изменилась
<img width="651" height="456" alt="image" src="https://github.com/user-attachments/assets/2b5e94d6-29ed-4578-8969-8dc32a0aad9f" />

После нажатия кнопки задача удаляется
<img width="801" height="500" alt="image" src="https://github.com/user-attachments/assets/cd835737-571a-406d-a5e2-86bf60fe03fa" />

Показывает только выполненные задачи
<img width="811" height="395" alt="image" src="https://github.com/user-attachments/assets/ee8524cb-3c0f-4707-87c7-7c3dbcfd8562" />

Показывает только задачи до 26 числа
<img width="874" height="448" alt="image" src="https://github.com/user-attachments/assets/d63d1eba-1665-4d93-8903-1ddf3caa3a5e" />













