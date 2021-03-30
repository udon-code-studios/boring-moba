// canvas variables
var canvas = document.getElementById("gameCanvas");
var ctx = canvas.getContext("2d");

// game variables
var player;
var players;

function draw() {
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
    ctx.fillText(players[i].displayName, players[i].currentPosition.x - 5, players[i].currentPosition.y - 16);
  }
}

// connection variables
var ws; // WebSocket connection to game server
var wsConnected;

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
    players = messageData.players;
    draw();
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
      var location = getCursorPosition(canvas, e)
      console.log(location)
      document.getElementById("mouseDownStatus").innerHTML = 'Right button clicked.';
      if (wsConnected) {
        ws.send(JSON.stringify({'id': player.id, 'newTargetPosition': location}));
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

function getCursorPosition(canvas, event) {
    const rect = canvas.getBoundingClientRect()
    const x = Math.round(event.clientX - rect.left)
    const y = Math.round(event.clientY - rect.top)
    return {'x': x, 'y': y}
}