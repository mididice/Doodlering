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

function preload() {
  // Load the DoodleNet Image Classification model
  classifier = ml5.imageClassifier('DoodleNet');
}

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

$(function () {
  var key = getKey();
  var sequence = getSequence();
  $.ajax({
    type: 'GET',
    url: '/play/' + key + '/' + sequence,
    dataType: 'text/html'
  }).done(function (data) {
    $('.sentence').html(data);
  }).fail(function (error) {
    alert(error);
  });

  $.ajax({
    type: 'GET',
    url: '/end/' + key + '/' + sequence,
    dataType: 'json'
  }).done(function (data) {
    if(data.answer){
      var label = data.answer.label;
      $('#word').text(label);
      var confidence = data.answer.confidence;
      $('#confidence').html("");
    }
    if(data.coordinate){
      for(var i =0 ; i<data.coordinate;i++){
        startingStrokes.push(data.coordinate[i])
      }
    }
  }).fail(function (error) {
    alert(error);
  });
});

$('.bttn_off_next').click(function(){
  location.href="/ending/"+getKey()+"/"+getSequence();
});