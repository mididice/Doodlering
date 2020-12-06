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
$(function () {

    //ending-real, story-end
    var copyLink = function () {
        var copyText = document.getElementById("shareLink");
        document.execCommand("copy");
    }

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
});