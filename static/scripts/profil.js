
$(document).ready(function () {
    $('#profil-button-delete').click(function () {
        $('#modal').toggleClass("is-active", true);
    });
    $('#profil-modal-button-keep').click(function () {
        $('#modal').toggleClass("is-active", false);
    });
    $('#profil-input-email').blur(validateEmail);
    $('#profil-input-passwort').blur(validatePassword);
    $('#profil-input-neu').blur(function () { validatePasswordNeu(); });
    $('#profil-input-wiederholen').blur(function () { validatePasswordNeu(); });
    $('#edit-button-save').click(updateProfil);
    $('#profil-modal-button-delete').click(deleteProfile);
})

function validateEmail() {
    let val = $("#profil-input-email").val();

    if (val == "" || !validateEmailString(val)) {
        $("#profil-input-email-icon-right").toggleClass("is-hidden", true);
        return
    }

    $.ajax({
        type: "POST",
        url: "/profil/email/" + val,
        success: function () {
            $("#edit-mail-help").toggleClass("is-invisible", true);
            $("#profil-input-email-icon-right").toggleClass("is-hidden", false);
        },
        error: function (event, _, msg) {
            if (event.status == 409) {
                $("#edit-mail-help").toggleClass("is-invisible", false);
                $("#profil-input-email-icon-right").toggleClass("is-hidden", true);
            } else {
                alert(msg);
            }
        }
    });
}

function validatePassword() {
    pw1 = $('#profil-input-passwort').val()

    if (pw1 == "") {
        $("#edit-oldpw-help").toggleClass("is-invisible", true);
        $("#profil-input-passwort-icon-right").toggleClass("is-hidden", true);
    } else {
        $.ajax({
            type: "POST",
            url: "/profil/passwort/" + sha512(pw1),
            success: function () {
                $("#edit-oldpw-help").toggleClass("is-invisible", false);
                $("#profil-input-passwort-icon-right").toggleClass("is-hidden", true);
            },
            error: function (event, _, msg) {
                if (event.status == 409) {
                    $("#edit-oldpw-help").toggleClass("is-invisible", true);
                    $("#profil-input-passwort-icon-right").toggleClass("is-hidden", false);
                } else {
                    alert(msg);
                }
            }
        });
    }
}

function validatePasswordNeu() {
    pw1 = $('#profil-input-neu').val()
    pw2 = $('#profil-input-wiederholen').val()

    if (pw1 == "") {
        $("#edit-newpw-help").toggleClass("is-invisible", true);
        $("#profil-input-neu-icon-right").toggleClass("is-hidden", true);
        return true
    } else {
        $.ajax({
            type: "POST",
            url: "/profil/passwort/" + sha512(pw1),
            success: function () {
                $("#edit-newpw-help").toggleClass("is-invisible", true);
                $("#profil-input-neu-icon-right").toggleClass("is-hidden", false);
            },
            error: function (event, _, msg) {
                if (event.status == 409) {
                    $("#edit-newpw-help").toggleClass("is-invisible", false);
                    $("#profil-input-neu-icon-right").toggleClass("is-hidden", true);
                } else {
                    alert(msg);
                }
            }
        });
    }

    if (pw2 == "") {
        $("#profil-input-wiederholen-icon-right").toggleClass("is-hidden", true);
        if (pw1 == "") {
            $("#edit-reppw-help").toggleClass("is-invisible", true);
        }
        return false;
    }

    if (pw1 === pw2) {
        $("#profil-input-wiederholen-icon-right").toggleClass("is-hidden", false);
        $("#edit-reppw-help").toggleClass("is-invisible", true);
        return true;
    } else {
        $("#profil-input-wiederholen-icon-right").toggleClass("is-hidden", true);
        $("#edit-reppw-help").toggleClass("is-invisible", false);
        return false;
    }
}

function updateProfil() {
    validateEmail();
    validatePassword();
    if (!validatePasswordNeu()) {
        return false;
    }

    var data = {}

    email = $('#profil-input-email').val()
    pwn = $('#profil-input-neu').val()
    if (email == "" && pwn == "") {
        alert("Es gibt keine Änderungen!");
        return false;
    }

    if (email != "" && !validateEmailString(email)) {
        alert("EMail ist ungültig!");
        return false;
    } else {
        data["EMail"] = email;
    }

    if (pwn != "") {
        pwo = $('#profil-input-passwort').val()

        if (pwo == "") {
            alert("Altes Passwort ungültig!");
            return false;
        }

        data["Passwort"] = sha512(pwo);
        data["Neu"] = sha512(pwn);
    }

    $.ajax({
        type: "PUT",
        url: "/profil",
        data: JSON.stringify(data),
        success: function () {
            window.location.href = "/profil";
        },
        error: function (event, _, msg) {
            if (event.status == 401) {
                alert("Altes Passwort falsch!");
            } else if (event.status == 409) {
                alert("Fehlerhafte Eingaben!");
            } else {
                alert(msg);
            }
        },
        contentType: "application/json"
    });
}

function deleteProfile() {
    $.ajax({
        type: "DELETE",
        url: "/profil",
        success: function () {
            window.location.href = "/";
        },
        error: function (event, _, msg) {
            alert(msg);
            window.location.href = "/profil";
        },
        contentType: "application/json"
    });
}