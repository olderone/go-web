var lastReceived = 0;
var isWait = false;

var fetch = function () {
    if (isWait) return;
    isWait = true;
    $.getJSON("/lp/fetch?lastReceived=" + lastReceived, function (data) {
        if (data == null) return;
        $.each(data, function (i, event) {
            var li = document.createElement('li');

            switch (event.Type) {
            case 0: // JOIN
                if (event.User == $('#uname').text()) {
                    li.innerText = 'You joined the chat room.';
                } else {
                    li.innerText = event.User + ' joined the chat room.';
                }
                break;
            case 1: // LEAVE
                li.innerText = event.User + ' left the chat room.';
                break;
            case 2: // MESSAGE
                var username = document.createElement('strong');
                var content = document.createElement('span');
                var time = document.createElement('p');

                username.innerText = event.User;
                content.innerText = event.Content;
                time.innerText = event.Time;

                li.appendChild(time);
                li.appendChild(username);
                li.appendChild(document.createTextNode(': '));
                li.appendChild(content);

                break;
            }

            $('#chatbox li').first().before(li);

            lastReceived = event.Timestamp;
        });
        isWait = false;
    });
}

// 每3秒取一次
setInterval(fetch, 2000);

fetch();

$(document).ready(function () {

    var postConecnt = function () {
        var uname = $('#uname').text();
        var content = $('#sendbox').val();
        $.post("/lp/post", {
            uname: uname,
            content: content
        });
        $('#sendbox').val("");
    }

    $('#sendbtn').click(function () {
        postConecnt();
    });
});
