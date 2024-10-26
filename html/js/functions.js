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
