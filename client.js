const ws = new WebSocket("ws://localhost:8080/chat");
var interval;
console.log(ws.url);

ws.addEventListener(
    "open", () => {
        console.log("connected.");
        ws.send("ping");
        // interval = setInterval(() => { ws.send("heyman!")}, 1000);
    }
);

ws.addEventListener(
    "error", (e) => {console.log(`error: ${e}`);}
);

function updateResponse(data) {
    let content = document.getElementById("content");
    content.textContent = content.textContent + data;
}

ws.addEventListener("message", (e) => {
//   console.log(`RECEIVED: ${e.data}`);
  updateResponse(e.data);
});

ws.addEventListener("close", () => {
  console.log("DISCONNECTED");
//   clearInterval(interval);
});
