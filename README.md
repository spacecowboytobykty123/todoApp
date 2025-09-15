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

## Ссылка на видео демонстрацию проекта
https://youtu.be/Jwj8Z6fBSAw

## 📸 Скриншоты приложения

<p align="center">
  <img src="https://github.com/user-attachments/assets/2f050f4a-6b43-4502-925b-ba6d25a64cff" width="787"/>
  <br>
  <i>Вывод всех задач</i>
</p>

<p align="center">
  <img src="https://github.com/user-attachments/assets/75ef91cd-1726-47fb-92d3-5730d060adf9" width="822"/>
  <br>
  <i>Поля для создания задачи (имя, описание, дата дедлайна и статус)</i>
</p>

<p align="center">
  <img src="https://github.com/user-attachments/assets/e1d13eea-79bf-4f52-b16b-7fc1f6a8a0c6" width="800"/>
  <br>
  <i>Задача добавилась</i>
</p>

<p align="center">
  <img src="https://github.com/user-attachments/assets/aa164dd2-d259-49c1-b0a9-324de22e69e0" width="414"/>
  <br>
  <i>Поле для изменения задачи</i>
</p>

<p align="center">
  <img src="https://github.com/user-attachments/assets/651c456.png" width="651"/>
  <br>
  <i>Задача изменилась</i>
</p>

<p align="center">
  <img src="https://github.com/user-attachments/assets/cd835737-571a-406d-a5e2-86bf60fe03fa" width="801"/>
  <br>
  <i>После нажатия кнопки задача удаляется</i>
</p>

<p align="center">
  <img src="https://github.com/user-attachments/assets/ee8524cb-3c0f-4707-87c7-7c3dbcfd8562" width="811"/>
  <br>
  <i>Фильтр: только выполненные задачи</i>
</p>

<p align="center">
  <img src="https://github.com/user-attachments/assets/d63d1eba-1665-4d93-8903-1ddf3caa3a5e" width="874"/>
  <br>
  <i>Фильтр: задачи с дедлайном до 26 числа</i>
</p>














