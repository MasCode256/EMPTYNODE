fetch_text("/query?cmd=type%20themes\\.current").then((curr) => {
  fetch_text("/query?cmd=type%20themes\\" + curr + ".json").then((text) => {
    console.log("Received text:", text); // Выводим полученный текст
    set_style_by_json(text);
  });
});
