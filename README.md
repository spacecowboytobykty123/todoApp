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

Вывод всех задач
<img width="782" height="598" alt="image" src="https://github.com/user-attachments/assets/aa89b600-bc76-4c8c-a940-6297cad504b8" />

Добавление задачи
<img width="868" height="615" alt="image" src="https://github.com/user-attachments/assets/00c087e4-3591-4d08-88bf-a5d6c6f6321c" />
<img width="822" height="121" alt="image" src="https://github.com/user-attachments/assets/af55213b-393e-41dd-b557-deec1b3947a5" />

<img width="681" height="299" alt="image" src="https://github.com/user-attachments/assets/47f080b6-0ba9-4da5-9500-51c1bf90889a" />



<img width="404" height="345" alt="image" src="https://github.com/user-attachments/assets/05b5bf86-776c-4697-9928-fba284c18870" />

<img width="774" height="602" alt="image" src="https://github.com/user-attachments/assets/2bf375f1-e542-414a-b48a-3085a4e1ff34" />



<img width="764" height="611" alt="image" src="https://github.com/user-attachments/assets/ac9417db-dfcb-446d-8401-5c69b0b4c81b" />



<img width="822" height="342" alt="image" src="https://github.com/user-attachments/assets/2b6108a5-e444-4c8e-b61e-bc5f446414f2" />


<img width="833" height="437" alt="image" src="https://github.com/user-attachments/assets/4a0d8f18-efd4-445e-8d7a-c56527e7d558" />

<img width="883" height="350" alt="image" src="https://github.com/user-attachments/assets/5460fa03-2a3a-4055-9db6-414b581372a5" />



