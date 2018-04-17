var holder = document.getElementById('uploadfile');

holder.ondragover = function () {
    return false;
};

holder.ondragend = function () {
    return false;
};

holder.ondrop = function (event) {
    event.preventDefault();

    var file = event.dataTransfer.files[0];
    var reader = new FileReader();

    reader.onload = function (event) {
        var binary = event.target.result;
        var md5 = CryptoJS.MD5(binary).toString();
        console.log(md5);
    };

    reader.readAsBinaryString(file);
};

function count(event) {
    var file = event.dataTransfer.files[0];
    var reader = new FileReader();

    reader.onload = function (event) {
        var binary = event.target.result;
        var md5 = CryptoJS.MD5(binary).toString();
        console.log(md5);
    };

    reader.readAsBinaryString(file);
}

// alert("YES, It Works...!!!");
$(function () {
    $("#upload").click(function (event) {
        alert("YES, It Works...!!!");
        event.preventDefault();
        count();
    });
   
});
