$(document).ready(function(){

  // enable clipboard copyier
  var client = new ZeroClipboard($('#copy'));
  client.on("ready", function(readyEvent) {
    client.on("aftercopy", function(event) {
      $('#copied').fadeIn(200);
      setTimeout(function(){
        $('#copied').fadeOut(300);
      }, 1000);
    });
  });
});
