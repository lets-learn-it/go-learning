window.onload = function() {
  // Box width
  var bw = 400;
  // Box height
  var bh = 400;
  // Padding
  var p = 10;
  var canvas = document.getElementById("canvas-lines");
  var context = canvas.getContext("2d");

  function drawBoard(){
    for (var x = 0; x <= bw; x += 40) {
        context.moveTo(0.5 + x + p, p);
        context.lineTo(0.5 + x + p, bh + p);
    }

    for (var x = 0; x <= bh; x += 40) {
        context.moveTo(p, 0.5 + x + p);
        context.lineTo(bw + p, 0.5 + x + p);
    }
    context.strokeStyle = "black";
    context.stroke();
  }
  drawBoard();

  if(typeof(EventSource) !== "undefined") {
    var source = new EventSource("/api/sse");
    source.onmessage = function(event) {
      var s = JSON.parse(event.data)
      context.beginPath();
      context.strokeStyle = "red";
      context.lineWidth = 3;
      context.moveTo(s.p1.x, s.p1.y);
      context.lineTo(s.p2.x, s.p2.y);
      context.stroke();
    };
  } else {
    document.getElementById("result").innerHTML = "Sorry, your browser does not support server-sent events...";
  }
}