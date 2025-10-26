const todos = document.getElementById("todos");
const tasks = document.getElementById("tasks");
const taskList = document.getElementById("task-list");
const todolistInput = document.getElementById("todolist");
const taskInput = document.getElementById("taskInput");
const heading = document.getElementById("heading");
const button = document.getElementById("clearTask")

let task = [];

todos.addEventListener("submit", createTodoList);
button.addEventListener("click", clearTask);
tasks.addEventListener("submit", addTask);

window.onload = renderTasks;

function createTodoList(event) {
    event.preventDefault();
    const todoListName = todolistInput.value.trim();
    if (todoListName === '') return
    fetch("/load", {
        method: "Post",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({name: todoListName})
    })
        .then(res => res.json())
        .then(data => {
            if (Array.isArray(data)) {
                task = data;
            } else {
                task = [];
            }
            renderTasks();
        })
        .catch(err => console.error(err))
    heading.textContent = `Tasks: ${todoListName}`;
    todolistInput.value = '';
    taskInput.disabled = false;
    taskInput.autofocus = true;
}

function clearTask(event) {
    event.preventDefault();
    fetch("/clear", {
        method: "Delete"
    })
        .then(_ => {
            task = [];
            renderTasks();
        })
        .catch(err => console.error(err));
    heading.textContent = "Tasks";
    taskInput.disabled = true;
}

function addTask(event) {
    event.preventDefault();
    const taskName = taskInput.value.trim();
    if (taskName === '') return
    fetch("/add", {
        method: "Post", headers: {"Content-Type": "application/json"}, body: JSON.stringify({name: taskName})
    })
        .then(res => res.json())
        .then(data => {
            task.push(data.name);
            renderTasks();
            taskInput.value = '';
        })
        .catch(err => console.error(err));
}

function renderTasks() {
    taskList.innerHTML = '';
    if (!Array.isArray(task)) return;
    task.forEach((name, index) => {
        const div = document.createElement('div');
        div.className = 'task';

        const number = document.createElement('span');
        number.textContent = `${index + 1}. `;

        const input = document.createElement('input');
        input.type = 'text';
        input.className = 'task-data';
        input.value = name;
        input.disabled = true;

        const removeButton = document.createElement('button');
        removeButton.className = 'remove-btn';
        removeButton.textContent = 'Remove';
        removeButton.onclick = () => {
            fetch("/remove", {
                method: "Delete", headers: {"Content-Type": "application/json"}, body: JSON.stringify({index: index})
            })
                .then(() => {
                    task.splice(index, 1);
                    renderTasks();
                })
                .catch(err => console.error(err));
        };

        const updateButton = document.createElement('button');
        updateButton.className = 'update-btn';
        updateButton.textContent = 'Update';
        updateButton.onclick = () => {
            input.disabled = false;
            input.focus();
            updateButton.style.display = 'none';
            saveButton.style.display = 'inline';
        };

        const saveButton = document.createElement('button');
        saveButton.className = 'save-btn';
        saveButton.textContent = 'Save';
        saveButton.style.display = 'none';
        saveButton.onclick = () => {
            input.disabled = true;
            saveButton.style.display = 'none';
            updateButton.style.display = 'inline';

            fetch('/update', {
                method: 'Put',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify({name: input.value, index: index})
            })
                .then(() => {
                    task[index] = input.value;
                })
                .catch(err => console.error(err))
        };

        div.appendChild(number);
        div.appendChild(input);
        div.appendChild(updateButton);
        div.appendChild(saveButton);
        div.appendChild(removeButton);

        taskList.appendChild(div);
    });
}