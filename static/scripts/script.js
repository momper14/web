$(document).ready(function () {
    $('.dropdown').click(function () {
        $(this).toggleClass("is-active");
    });
    $('#button-logout').click(logout);
    $('#form-login').submit(login);
})

function login() {
    elements = document.forms["login"].elements;
    var data = {
        "User": elements["User"].value,
        "Passwort": sha512(elements["Passwort"].value)
    }

    $.ajax({
        type: "POST",
        url: "/login",
        data: JSON.stringify(data),
        success: function () {
            window.location.href = "/meinekarteien";
        },
        error: function () {
            alert("Username oder Passwort ungültig");
        },
        contentType: "application/json"
    });

    return false;
}

function logout() {
    $.ajax({
        type: "POST",
        url: "/logout",
    }).done(function(){
        window.location.href = "/";
    });
}