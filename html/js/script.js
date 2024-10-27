fetch_text("/query?cmd=type%20themes\\.current").then((curr) => {
  fetch_text("/query?cmd=type%20themes\\" + curr + ".json").then((text) => {
    console.log("Received text:", text); // Выводим полученный текст
    set_style_by_json(text);
  });
});

document.addEventListener("DOMContentLoaded", function () {
  const content = document.getElementById("content");

  if (localStorage.getItem("pos") === null) {
    localStorage.setItem("pos", "explorer.html");
  }

  content.src = localStorage.getItem("pos");
});
