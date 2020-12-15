// Copyright (c) 2019 ml5
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

/* ===
ml5 Example
Canvas Image Classification using DoodleNet and p5.js
This example uses a callback pattern to create the classifier
=== */

let canvas;

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
  return pathNameList[3];
}
var shuffleArray = function(array){
  for (var i = array.length - 1; i > 0; i--) {
      var j = Math.floor(Math.random() * (i + 1));
      var temp = array[i];
      array[i] = array[j];
      array[j] = temp;
  }
  return array;
}
$(function() {
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
      console.log(error);
  });
  $(".badge").attr("src","/static/img/badge/"+sequence+".svg");
  $('.bttn_off_back').attr("src", "/static/img/bttn_off_back/"+sequence+".svg");

  $.ajax({
    type: 'GET',
    url: '/tale/' + key + '/' + sequence,
    dataType: 'json'
  }).done(function (data) {
    if(data.candidate){
      $('#answer').val(data.candidate[0]);
      var candidates = shuffleArray(data.candidate);
      for(var i=0; i<candidates.length;i++){
        $('.answer'+i).text(candidates[i]);
      }
    }
    if(data.coordinate){
      for(var i =0 ; i<data.coordinate.length;i++){
        startingStrokes.push(data.coordinate[i])
      }
    }
  }).fail(function (error) {
    console.log(error);
  });
  $('.bttn_off_back').click(function(){
    var sequence = getSequence();
    if(sequence>1){
      location.href="/story/"+key+"/"+(parseInt(sequence)-1);
    }else{
      location.href="/home";
    }
  });
  $('.bttn_off_next').click(function(){
    if(sequence>=10){
      location.href="/story/"+key+"/end"
    }else{
      location.href="/story/"+key+"/"+(parseInt(sequence)+1);
    }
  });

  $('.button_assume').click(function(){
    var label = $(this).text();
    $('#word').text(label);

    $(this).css("background-color", "#111111");
    $(this).css("color", "#c8fc50");
    var answer = $('#answer').val();
    $('#aiChoose').text(answer);
    $('#youChoose').text(label);
    $('.button_assume').hide();
    $('.button_result').show();
  });
});
