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

    //playing-end
    $('.bttn_gamestart_howto').click(function () {
        location.href = "home.html";
    });

    //ending-real, story-end
    var copyLink = function () {
        var copyText = document.getElementById("shareLink");
        document.execCommand("copy");
    }

});