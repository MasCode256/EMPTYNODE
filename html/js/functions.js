async function fetch_text(url) {
  try {
    const response = await fetch(url); // Выполняем GET-запрос
    if (!response.ok) {
      // Проверяем, успешен ли ответ
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const text = await response.text(); // Получаем текст из ответа
    return text; // Возвращаем текст
  } catch (error) {
    console.error("Error fetching text:", error); // Обработка ошибок
  }
}

async function set_style_by_json(jsonString) {
  const styles = JSON.parse(jsonString);

  // Вспомогательная функция для установки переменных
  function setVariables(prefix, obj) {
    for (const [key, value] of Object.entries(obj)) {
      const cssVariable = prefix ? `--${prefix}_${key}` : `--${key}`;
      if (typeof value === "object" && value !== null) {
        // Если значение - объект, рекурсивно вызываем функцию
        setVariables(cssVariable, value);
      } else {
        // Устанавливаем CSS-переменную
        document.documentElement.style.setProperty(cssVariable, value);
      }
    }
  }

  // Запускаем установку переменных с пустым префиксом
  setVariables("", styles);
}

const dir = document.getElementById("dir");
var dirs = [];

async function load_dirs() {
  await fetch_text("/query?cmd=type%20fsys\\.dirs").then((text) => {
    dirs = text.split("\n");
    console.log(dirs);
  });

  console.log("pre-for");

  for (let index = 0; index < dirs.length; index++) {
    console.log("for " + index);

    if (dirs[index] != "root") {
      dir.innerHTML +=
        "<div class='card'>" +
        "<img class='icon' src='/icons/dir_ico.png' alt='' />" +
        dirs[index] +
        "</div>";
    }
  }
}

async function load_applist(dir) {}
