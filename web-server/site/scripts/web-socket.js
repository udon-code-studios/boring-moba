var output = document.getElementById("output");
var input = document.getElementById("input");
var ws;
var sendTime;

var print = function(message) {
  var d = document.createElement("div");
  d.textContent = message;
  output.appendChild(d);
};

document.getElementById("open").onclick = function(evt) {
  ws = new WebSocket("ws://subparprogramming.cf:5000/echo");

  ws.onopen = function(event) {
    print("OPEN");
  };

  ws.onclose = function(evt) {
    print("CLOSE");
    ws = null;
  }
  
  ws.onmessage = function(evt) {
    print("RESPONSE (" + (Date.now() - sendTime) + " ms): " + evt.data);
  }
  
  ws.onerror = function(evt) {
    print("ERROR: " + evt.data);
  }
  
  return false;
};

document.getElementById("send").onclick = function(evt) {
  if (!ws) {
    return false;
  }
  print("SEND: " + input.value);
  sendTime = Date.now();
  ws.send(input.value);
  return false;
};

document.getElementById("close").onclick = function(evt) {
  if (!ws) {
    return false;
  }
  ws.close();
  return false;
};
