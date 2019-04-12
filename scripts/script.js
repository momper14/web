const USERNAME = "username";

function load_no_login() {
    $('#placeholder-navbar').load("res/html/navbar_no_login.html", function () {
        $('#button-login').click(function () {
            login();
        })
    });

    $('#placeholder-sidemenu').load("res/html/sidemenu_no_login.html", function () {
        $('#Karteikasten-zahl').text(22);
    });
}

function load_login() {
    $('#placeholder-navbar').load("res/html/navbar_login.html", function () {
        $('#navbar-username').text(getCookie(USERNAME));
        $('#button-logout').click(function () {
            logout();
        });
    });

    $('#placeholder-sidemenu').load("res/html/sidemenu_login.html", function () {
        $('#Karteikasten-zahl').text(22);
        $('#meine-Karteikasten-zahl').text(7);
    });
}

function setCookie(cname, cvalue, exdays, path = "") {
    var d = new Date();
    d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
    var expires = "expires=" + d.toUTCString();
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/" + path;
}

function getCookie(cname) {
    var name = cname + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');
    for (var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

function deleteCookie(cname) {
    setCookie(cname, "", -1);
}

function loadMenu() {
    var user = getCookie(USERNAME);
    if (user != "") {
        load_login();
    } else {
        load_no_login();
    }
}

function login() {
    var username = $('#input-username').val(),
        password = $('input-password').val();

    setCookie(USERNAME, username, 1);

    load_login();
}

function logout() {
    deleteCookie(USERNAME);
    load_no_login();
}