// Copyright (c) 2019 ml5
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

/* ===
ml5 Example
Canvas Image Classification using DoodleNet and p5.js
This example uses a callback pattern to create the classifier
=== */

// Initialize the Image Classifier method with DoodleNet.
let classifier;

// A variable to hold the canvas image we want to classify
let canvas;

// Two variable to hold the label and confidence of the result
let label;
let confidence;

let possibleDraw = true;

let firstClicked = false;

let strokes = [];

let current_raw_line = [];

let prediction = [];

function preload() {
  // Load the DoodleNet Image Classification model
  classifier = ml5.imageClassifier('DoodleNet');
}

function setup() {
  init(function(){
    console.log('ready.');
  });
  // Create a canvas with 280 x 280 px
  canvas = createCanvas(280, 280);
  canvas.parent('drawingPaper');
  // Set canvas background to white
  background('#3644eb');
  // Whenever mouseReleased event happens on canvas, call "classifyCanvas" function
  //canvas.mouseReleased(classifyCanvas);
  canvas.touchEnded(classifyCanvas);
  // Create a clear canvas button
  const button = select('.bttn_redraw');
  button.mousePressed(clearCanvas);
  button.touchStarted(clearCanvas);
  // Create 'label' and 'confidence' div to hold results
}

function clearCanvas() {
  background('#3644eb');
  window.document.getElementById('word').textContent = '___';
}

function draw() {
  
// Set stroke weight to 10
strokeWeight(15);
// Set stroke color to black
stroke(255);
// If mouse is pressed, draw line between previous and current mouse positions
if (mouseIsPressed) {
  line(pmouseX, pmouseY, mouseX, mouseY);
  current_raw_line.push([pmouseX, pmouseY, mouseX, mouseY]);
  
}
  
}

function classifyCanvas() {
  classifier.classify(canvas, gotResult);
}

// A function to run when we get any errors and the results
function gotResult(error, results) {
  // Display error in the console
  if (error) {
    console.error(error);
  }
  // The results are in an array ordered by confidence.
  console.log(results);
  prediction = results;
  // Show the first label and confidence
  window.document.getElementById('word').textContent = hanguel.get(results[0].label);
  console.log(`Confidence: ${nf(results[0].confidence, 0, 2)}`); // Round the confidence to 0.01
}

var init = function() {
  screen_width = get_window_width(); //window.innerWidth
  screen_height = get_window_height(); //window.innerHeight

  // var canvas = document.getElementsByTagName("canvas")[0];
};

function get_window_width() {
  // return p.windowWidth;
  return window.innerWidth;
}

function get_window_height() {
  // return p.windowHeight;
  return window.innerHeight;
}

function setPossibleDraw(){
  possibleDraw = true;
}

$('.bttn_off_next').click(function(){
  var key = getKey();
  var sequence = getSequence();
  location.href="/story/"+key+"/"+sequence;
});

var getKey = function(){
  var pathName = window.location.pathname;
  var pathNameList = pathName.split("/");
  return pathNameList[2];
}
var getSequence = function(){
  var pathName = window.location.pathname;
  var pathNameList = pathName.split("/");
  return pathNameList[2];
}
$(function() {
  var key = getKey();
  var sequence = getSequence();
  $.ajax({
    type: 'GET',
    url: '/play/'+key+'/'+sequence,
    dataType: 'text/html'
  }).done(function(data) {
    $('.sentence').html(data);
  }).fail(function (error) {
      alert(error);
  });

  $.ajax({
    type: 'GET',
    url: '/tale/' + key + '/' + sequence,
    dataType: 'json'
  }).done(function (data) {
    if(data.answer){
      var label = data.answer[0].label;
      for(var i=0; i<data.answer.length;i++){
        $('.answer'+i).text(hanguel.get(data.answer[i]));
      }
      var confidence = data.answer.confidence;
      $('#confidence').html("");
    }
    if(data.coordinate){
      for(var i =0 ; i<data.coordinate.length;i++){
        startingStrokes.push(data.coordinate[i])
      }
    }
  }).fail(function (error) {
    alert(error);
  });
});

$(document).on(".button_assume", "click", function(){
  var label = $(this).text();
  $('#answer').label();
});