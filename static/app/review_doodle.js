// Copyright (c) 2019 ml5
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT


// A variable to hold the canvas image we want to classify
let canvas;

// Two variable to hold the label and confidence of the result
let label;
let confidence;

let startingStrokes = [];

let startingStrokeIndex = 0;

function setup() {
  // Create a canvas with 280 x 280 px
  canvas = createCanvas(280, 280);
  canvas.parent('drawingPaper');
  // Set canvas background to white
  background('#3644eb');
  
}

function draw() {

  // Set stroke weight to 10
  strokeWeight(15);
  // Set stroke color to black
  stroke(255);
  
  if(startingStrokeIndex < startingStrokes.length) {
    let strokeXY = startingStrokes[startingStrokeIndex];
    dx = strokeXY[0];
    dy = strokeXY[1];

    line(strokeXY[0], strokeXY[1], strokeXY[2], strokeXY[3]);
  
    startingStrokeIndex++;
  }

}
var getKey = function(){
  var pathName = window.location.pathname;
  var pathNameList = pathName.split("/");
  return pathNameList[2];
}
var getSequence = function(){
  var pathName = window.location.pathname;
  var pathNameList = pathName.split("/");
  return pathNameList[3];
}
$(function () {
  var key = getKey();
  var sequence = getSequence();
  $.ajax({
    type: 'GET',
    url: '/play/'+key+'/'+sequence,
    dataType: 'json'
  }).done(function(data) {
    if(data){
      $('.sentence').html(data.Sentence);
    }
  }).fail(function (error) {
      alert(error);
  });
  $(".badge").attr("src","/static/img/badge/"+sequence+".svg");

  $.ajax({
    type: 'GET',
    url: '/end/' + key + '/' + sequence,
    dataType: 'json'
  }).done(function (data) {
    if(data.Answer){
      var label = data.Answer[0].Label;
      $('#word').text(hanguel.get(label));
      var confidence = data.Answer[0].Confidence;
      $('#confidence').html("");
    }
    if(data.Coordinate){
      for(var i =0 ; i<data.Coordinate.length;i++){
        startingStrokes.push(data.Coordinate[i])
      }
    }
  }).fail(function (error) {
    console.log(error);
  });

  $('.bttn_off_next').click(function(){
    location.href="/ending/"+key+"/"+(parseInt(sequence)+1);
    if(sequence>=10){
      location.href="/ending/"+key+"/end"
    }else{
      location.href="/ending/"+key+"/"+(parseInt(sequence)+1);
    }
  });
});

