$(function () {
    $.ajax({
      type: 'GET',
      url: '/tales',
      dataType: 'json'
    }).done(function (data) {
        var html = '';
        for(var i =0; i<data.length;i++){
            html += '<div class="story_box">'
            +'<span class="story_date">'+data[i].date+'</span>'
            +'<a href="/story/'+data[i].key+'/1" class="story_content">'+data[i].sentence+'</a></div>';
        }
      $('.stories').html(html);
    }).fail(function (error) {
      console.log(error);
    });
  });
  