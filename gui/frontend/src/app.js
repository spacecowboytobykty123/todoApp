// app.js
import { GetTasks, CreateTask, DeleteTask, UpdateTask } from "./wailsjs/go/main/App.js";

const taskList = document.getElementById("taskList");
const nameInput = document.getElementById("taskName");
const descInput = document.getElementById("taskDesc");
const statusSelect = document.getElementById("taskStatus");
const addBtn = document.querySelector("button"); // первая кнопка "Добавить"

// загрузка задач
async function loadTasks() {
    const tasks = await GetTasks();
    taskList.innerHTML = "";

    tasks.forEach(task => {
        const li = document.createElement("li");

        // кнопка удалить
        const delBtn = document.createElement("button");
        delBtn.textContent = "❌";
        delBtn.addEventListener("click", () => deleteTask(task.id));

        // кнопка обновить статус
        const toggleBtn = document.createElement("button");
        toggleBtn.textContent = "🔄";
        toggleBtn.addEventListener("click", () => toggleStatus(task.id, task.status));

        li.innerHTML = `<b>${task.name}</b> - ${task.status} <small>(${task.desc || ""})</small>`;
        li.appendChild(delBtn);
        li.appendChild(toggleBtn);

        taskList.appendChild(li);
    });
}

// добавить задачу
async function addTask() {
    const name = nameInput.value;
    const desc = descInput.value;
    const status = statusSelect.value;

    if (!name) return alert("Введите название");

    await CreateTask(name, desc, status, null);
    nameInput.value = "";
    descInput.value = "";
    await loadTasks();
}

// удалить задачу
async function deleteTask(id) {
    await DeleteTask(id);
    await loadTasks();
}

// обновить статус
async function toggleStatus(id, status) {
    let nextStatus =
        status === "todo" ? "in_progress" :
            status === "in_progress" ? "done" : "todo";

    await UpdateTask(id, null, null, nextStatus, null);
    await loadTasks();
}

// навешиваем события
addBtn.addEventListener("click", addTask);

// при старте
loadTasks();
