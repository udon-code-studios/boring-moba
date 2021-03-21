// canvas variables
var canvas = document.getElementById("gameCanvas");
var ctx = canvas.getContext("2d");

// game variables
var player;
var players;

function draw() {
  console.log(players)
  // clear canvas
  ctx.clearRect(0, 0, canvas.width, canvas.height);

  for (var i = 0; i < players.length; i++) {
    ctx.beginPath();
    ctx.arc(players[i].currentPosition.x, players[i].currentPosition.y, 10, 0, Math.PI * 2);
    ctx.fillStyle = "#0095DD";
    ctx.fill();
    ctx.closePath();

    ctx.font = "16px Arial";
    ctx.fillStyle = "#0095DD";
    ctx.fillText(players[i].name, players[i].currentPosition.x - 5, players[i].currentPosition.y - 16);
  }
}

// connection variables
var ws; // WebSocket connection to game server
var wsConnected;
var sendTime;

// connect to game server and setup websocket methods
document.getElementById("connect").onclick = function(evt) {
  getPlayerID({ 'name': 'Linus Torvalds' })
    .then(data => {
      console.log(data);
      player = data;
    });

  ws = new WebSocket("ws://subparprogramming.cf:5000/player-input-ws");

  ws.onopen = function(event) {
    document.getElementById("gameConnectionStatus").innerHTML = 'Connected';
    wsConnected = true;
  };

  ws.onclose = function(evt) {
    document.getElementById("gameConnectionStatus").innerHTML = 'Disconnected';
    ws = null;
  }

  ws.onmessage = function(evt) {
    var messageData = JSON.parse(evt.data);
    console.log(messageData);
    if (messageData.type == "response") {
      document.getElementById("gameConnectionStatus").innerHTML = `Connected (${(Date.now() - sendTime)} ms)`;
    }
    else if (messageData.type == "players") {
      players = messageData.players;
      draw();
    }
  }

  ws.onerror = function(evt) {
    //print("ERROR: " + evt.data);
  }

  return false;
};

// disconnect from game server
document.getElementById("disconnect").onclick = function(evt) {
  if (!ws) {
    return false;
  }
  ws.close();
  return false;
};

// handle mouse events
document.addEventListener("mousedown", e => {
  switch (e.button) {
    case 0:
      document.getElementById("mouseDownStatus").innerHTML = 'Left button clicked.';
      break;
    case 1:
      document.getElementById("mouseDownStatus").innerHTML = 'Middle button clicked.';
      break;
    case 2:
      document.getElementById("mouseDownStatus").innerHTML = 'Right button clicked.';
      if (wsConnected) {
        sendTime = Date.now();
        ws.send(`id: ${player.id} right clicked`);
      }
      break;
    default:
      document.getElementById("mouseDownStatus").innerHTML = `Unknown button code: ${e.button}`;
  }
});

// POST method to get a player ID
async function getPlayerID(data) {
  const url = 'http://subparprogramming.cf:5000/player-create';

  const response = await fetch(url, { /* global fetch */
    method: 'POST',
    body: JSON.stringify(data)
  });

  return response.json();
}
