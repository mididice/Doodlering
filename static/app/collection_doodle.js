$(function () {
    $.ajax({
      type: 'GET',
      url: '/tales',
      dataType: 'json'
    }).done(function (data) {
        var html = '';
        for(var i =0; i<data.length;i++){
            html += '<div class="story_box">'
            +'<span class="story_date">'+data[i].Date+'</span>'
            +'<a href="/story/'+data[i].Key+'/1" class="story_content">'+data[i].Sentence+'</a></div>';
        }
      $('.stories').html(html);
    }).fail(function (error) {
      console.log(error);
    });
    $('.bttn_off_back_white').click(function(){
        location.href="/";
    });
  });
  