// Copyright (c) 2020 https://github.com/yining1023/machine-learning-for-the-web/tree/master/cnn/DoodleClassifier100
const IMAGE_SIZE = 784;
const CLASSES = ['laptop', 'rainbow', 'baseball_bat', 'ice_cream', 'flower', 'suitcase', 'tree', 'microphone', 'sword', 'helmet', 'apple', 'umbrella', 'frying_pan', 'envelope', 'triangle', 'alarm_clock', 'paper_clip', 'light_bulb', 'scissors', 'cat', 't-shirt', 'ceiling_fan', 'key', 'mountain', 'table', 'moon', 'smiley_face', 'car', 'spoon', 'bird', 'saw', 'traffic_light', 'knife', 'wristwatch', 'shovel', 'circle', 'face', 'bridge', 'camera', 'bread', 'screwdriver', 'tennis_racquet', 'cell_phone', 'airplane', 'bed', 'baseball', 'moustache', 'candle', 'tooth', 'star', 'sock', 'dumbbell', 'lollipop', 'bicycle', 'hat', 'spider', 'clock', 'shorts', 'anvil', 'pants', 'syringe', 'ladder', 'axe', 'headphones', 'grapes', 'square', 'chair', 'coffee_cup', 'lightning', 'cookie', 'wheel', 'pencil', 'cloud', 'mushroom', 'door', 'drums', 'fan', 'bench', 'sun', 'stop_sign', 'eye', 'beard', 'radio', 'snake', 'line', 'power_outlet', 'diving_board', 'rifle', 'eyeglasses', 'broom', 'donut', 'pillow', 'hot_dog', 'butterfly', 'hammer', 'basketball', 'book', 'tent', 'pizza', 'cup'];
let model;
let cnv;
let startDraw = false;
let doGuess = false;

let current_raw_line = [];
let prediction = [];

let moveClicked = true;

async function loadMyModel() {
  model = await tf.loadLayersModel('/model/model.json');
  model.summary();
}

function setup() {
  loadMyModel();

  cnv = createCanvas(280, 280);
  background('#3644eb');

  // cnv.mouseReleased(guess);
  cnv.touchEnded(guess);
  cnv.parent('drawingPaper');

  let clearButton = select('.bttn_redraw');
  clearButton.mousePressed(() => {
    clear();
    background('#3644eb');
    current_raw_line = [];
    prediction = [];
    select('#word').html("____");
    $('.bttn_off_next').attr("src", "/static/img/bttn_off_next/0.svg");
    $('.bttn_guess').attr("src", "/static/img/bttn_guess_yet.png")
    startDraw = false;
    doGuess = false;
  });
}

function guess() {

  // Get input image from the canvas
  const inputs = getInputImage();

  // Predict
  let guess = model.predict(tf.tensor([inputs]));

  // Format res to an array
  const rawProb = Array.from(guess.dataSync());

  // Get top 5 res with index and probability
  const rawProbWIndex = rawProb.map((probability, index) => {
    return {
      index,
      probability
    }
  });

  const sortProb = rawProbWIndex.sort((a, b) => b.probability - a.probability);
  const top5ClassWIndex = sortProb.slice(0, 5);
  const results = [];
  top5ClassWIndex.map(i=> results.push({['label']: hanguel.get(CLASSES[i.index]) , ['confidence']: i.probability}));
  prediction = results;
  select('#word').html(results[0].label);
  $('.bttn_guess').attr("src", "/static/img/bttn_guess.png");
}

function getInputImage() {
  let inputs = [];
  // p5 function, get image from the canvas
  let img = get();
  img.resize(28, 28);
  img.loadPixels();
  
  // Group data into [[[i00] [i01], [i02], [i03], ..., [i027]], .... [[i270], [i271], ... , [i2727]]]]
  let oneRow = [];

  for (let i = 0; i < IMAGE_SIZE; i++) {
    //54, 68, 235
    let bright = img.pixels[i * 4];
    // let green = img.pixels[i*4+1];
    // let blue = img.pixels[i*4+2];
    let onePix = [parseFloat((255 - bright) / 255)];
    
    oneRow.push(onePix);
    if (oneRow.length === 28) {
      inputs.push(oneRow);
      oneRow = [];
    }
  }
  return inputs;
}
var threshold = function(pixels, level) {

  if (level === undefined) {
    level = 0.5;
  }
  const thresh = Math.floor(level * 255);

  for (let i = 0; i < pixels.length; i += 4) {
    const r = pixels[i];
    const g = pixels[i + 1];
    const b = pixels[i + 2];
    const gray = 0.2126 * r + 0.7152 * g + 0.0722 * b;
    let val;
    if (gray >= thresh) {
      val = 255;
    } else {
      val = 0;
    }
    pixels[i] = pixels[i + 1] = pixels[i + 2] = val;
  }
};

function draw() {
  strokeWeight(10);
  stroke(255);
  if (mouseIsPressed && !doGuess) {
    line(pmouseX, pmouseY, mouseX, mouseY);
    current_raw_line.push([pmouseX, pmouseY, mouseX, mouseY]);
    startDraw = true;
  }
}

function get_window_width() {
  // return p.windowWidth;
  return window.innerWidth;
}

function get_window_height() {
  // return p.windowHeight;
  return window.innerHeight;
}
$('.bttn_off_back').click(function(){
  if(moveClicked){
    moveClicked = !moveClicked;
    var key = getKey();
    var sequence = getSequence();
    if(sequence>1){
      location.href="/playing/"+key+"/"+(parseInt(sequence)-1);
    }else{
      location.href="/home";
    }
    setTimeout(function () {
      moveClicked = true;
    }, 2000)
  }
});
$('.bttn_off_next').click(function(){
  if(!startDraw){
    return;
  }
  if(!doGuess){
    return;
  }
  var key = getKey();
  var sequence = getSequence();
  var data = {"prediction":prediction, "coordinate": current_raw_line};
  if(prediction.length<1){
    return;
  }
  if(moveClicked){
    moveClicked = !moveClicked;
   
    $.ajax({
      type: 'POST',
      url: '/play/'+key+'/'+sequence,
      dataType: 'text',
      contentType:'application/json; charset=utf-8',
      data: JSON.stringify(data)
    }).done(function(data) {
      if(sequence>=10){
        location.href="/ending/"+key
      }else{
        location.href="/playing/"+key+"/"+(parseInt(sequence)+1);
      }
    }).fail(function (error) {
      console.log(error);
    });

    setTimeout(function () {
      moveClicked = true;
    }, 3000)
  }
});
$('.bttn_guess').click(function(){
  if(startDraw){
    filter(INVERT);
    filter(THRESHOLD);
    guess();
    background('#3644eb');
    $('.bttn_off_next').attr("src", "/static/img/bttn_off_next/"+getSequence()+".svg");
    $('.bttn_guess').attr("src", "/static/img/bttn_guess_yet.png");
    doGuess = true;
  }
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
$(function() {
  var key = getKey();
  var sequence = getSequence();
  $.ajax({
    type: 'GET',
    url: '/play/'+key+'/'+sequence,
    dataType: 'json'
  }).done(function(data) {
    if(data){
      if(data.Sentence){
        $('.sentence').html(data.Sentence);
      }else{
        $('.sentence').html("<span id='word'>그림을 그려주세요.</span>");
      }
    }
  }).fail(function (error) {
      alert(error);
  });

  $(".badge").attr("src", "/static/img/badge/"+sequence+".svg");
  $('.bttn_off_back').attr("src", "/static/img/bttn_off_back/"+sequence+".svg");
});