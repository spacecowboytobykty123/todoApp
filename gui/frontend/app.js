import { GetTasks, CreateTask, DeleteTask, UpdateTask } from "./wailsjs/go/main/App.js";

const taskList = document.getElementById("taskList");
const nameInput = document.getElementById("taskName");
const descInput = document.getElementById("taskDesc");
const statusSelect = document.getElementById("taskStatus");
const addBtn = document.getElementById("addTask"); // кнопка "Добавить"

const filterStatus = document.getElementById("filterStatus");
const filterDeadline = document.getElementById("filterDeadline");
const applyFiltersBtn = document.getElementById("applyFilters");

const modal = document.createElement("div");
modal.className = "modal";
modal.innerHTML = `
  <div class="modal-content">
    <h3>✏️ Редактировать задачу</h3>
    <input id="editName" placeholder="Название" />
    <input id="editDesc" placeholder="Описание" />
    <select id="editStatus">
      <option value="Назначена">Назначена</option>
      <option value="В работе">В работе</option>
      <option value="Выполненный">Выполненный</option>
      <option value="Отклонена">Отклонена</option>
    </select>
    <div style="text-align:right; margin-top:10px;">
      <button id="cancelEdit">Отмена</button>
      <button id="saveEdit">Сохранить</button>
    </div>
  </div>
`;
document.body.appendChild(modal);

let editingTaskId = null;



// загрузка задач
async function loadTasks() {
    const status = filterStatus.value || "";
    const deadline = filterDeadline.value ? new Date(filterDeadline.value).toISOString() : null;

    const tasks = await GetTasks(status, deadline);
    renderTasks(tasks);
}

// Отрисовка задач
function renderTasks(tasks) {
    taskList.innerHTML = "";

    tasks.forEach(task => {
        const li = document.createElement("li");

        const delBtn = document.createElement("button");
        delBtn.textContent = "❌";
        delBtn.addEventListener("click", () => deleteTask(task.id));

        const editBtn = document.createElement("button");
        editBtn.textContent = "🔄";
        editBtn.addEventListener("click", () => openModal(task));

        li.innerHTML = `<b>${task.name}</b> - <span class="status-${task.status.replace(" ", "\\ ")}">${task.status}</span> <small>(дедлайн: ${task.deadline || "—"})</small>`;
        li.appendChild(delBtn);
        li.appendChild(editBtn);

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

// открыть модалку
function openModal(task) {
    editingTaskId = task.id;
    document.getElementById("editName").value = task.name;
    document.getElementById("editDesc").value = task.desc || "";
    document.getElementById("editStatus").value = task.status;
    modal.style.display = "flex";
}

// закрыть модалку
function closeModal() {
    modal.style.display = "none";
    editingTaskId = null;
}

// сохранить изменения
async function saveTask() {
    const name = document.getElementById("editName").value;
    const desc = document.getElementById("editDesc").value;
    const status = document.getElementById("editStatus").value;

    if (!editingTaskId) return;

    await UpdateTask(editingTaskId, name, desc, status, null);
    closeModal();
    await loadTasks();
}

async function loadTasksWithFilters() {
    const status = filterStatus.value || "";
    const deadline = filterDeadline.value || "";

    const tasks = await GetTasks(status, deadline);
    renderTasks(tasks);
}

applyFiltersBtn.addEventListener("click", loadTasksWithFilters);


// навешиваем события
addBtn.addEventListener("click", addTask);
document.getElementById("cancelEdit").addEventListener("click", closeModal);
document.getElementById("saveEdit").addEventListener("click", saveTask);

// при старте
loadTasks();
