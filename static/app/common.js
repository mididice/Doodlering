//playing -> playing-end -> ending -> ending-real -> story -> story-end
var getKey = function () {
    var pathName = window.location.pathname;
    var pathNameList = pathName.split("/");
    return pathNameList[2];
}
var getSequence = function () {
    var pathName = window.location.pathname;
    var pathNameList = pathName.split("/");
    return pathNameList[2];
}
var getShareLink = function(){
    var pathName = window.location.pathname;
    pathName = pathName.replace("ending", "story");
    var pathNameList = pathName.split("/");
    pathNameList.pop();
    return( window.location.origin + pathNameList.join('/') +"/1");
}
//ending-real, story-end
var copyLink = function () {
    var copyText = document.getElementById("shareLink").value;

    var tempElem = document.createElement('textarea');
    tempElem.value = copyText;
    document.body.appendChild(tempElem);
  
    tempElem.select();
    document.execCommand("copy");
    document.body.removeChild(tempElem);
}
$(function () {
    $('#shareLink').val(getShareLink());
    $('.real_ending_center_link').text(getShareLink)
    $('.bttn_gamestart').click(function () {
        $.ajax({
            type: 'GET',
            url: '/start',
            dataType: 'json'
        }).done(function (data) {
            if (data.key) {
                location.href = "/playing/" + data.key + "/1";
            }
        }).fail(function (error) {
            alert(error);
        });
    });
    $('.bttn_gamestart_howto').click(function () {
        $.ajax({
            type: 'GET',
            url: '/start',
            dataType: 'json'
        }).done(function (data) {
            if (data.key) {
                location.href = "/playing/" + data.key + "/1";
            }
        }).fail(function (error) {
            alert(error);
        });
    });
    $('.ending_first').click(function(){
        location.href="/ending/"+getKey()+"/1";
    });

    $('.bttn_retry').click(function(){
        location.href="/home";
    });
});