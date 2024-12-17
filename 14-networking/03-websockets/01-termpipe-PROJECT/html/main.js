let input = document.getElementById('input');
let code = document.getElementById('code');
let data = '';

let ws = new WebSocket('/echo');
ws.onmessage = function(event) {
  data += event.data;
  code.innerText = data;
  input.scrollIntoView();
}

ws.onclose = function() {
  let ele = document.createElement('b');
  ele.innerText = 'Connection closed';
  code.appendChild(ele);
}

document.onkeydown = function(event) {
  if (event.key === 'Enter') {
    ws.send(input.value + '\n');
    input.value = '';
  }
}
