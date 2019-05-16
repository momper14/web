$(document).ready(function () {
    $('#register-input-privacy').click(onChangeAkzeptiert);
    $('#register-input-username').blur(validateName);
    $('#register-input-email').blur(validateEmail);
    $('#register-input-password').blur(function () { validatePassword(); });
    $('#register-input-repeat').blur(function () { validatePassword(); });
    $('#form-register').submit(register);
})

function onChangeAkzeptiert() {
    $("#register-input-privacy-help").toggleClass("is-invisible", this.checked);
}

function validateName() {
    let val = $("#register-input-username").val();

    if (val == "") {
        $("#register-input-username-icon-right").toggleClass("is-hidden", true);
        return
    }

    $.ajax({
        type: "POST",
        url: "/register/name/" + val,
        success: function () {
            $("#register-input-username-help").toggleClass("is-invisible", true);
            $("#register-input-username-icon-right").toggleClass("is-hidden", false);
        },
        error: function (event, _, msg) {
            if (event.status == 409) {
                $("#register-input-username-help").toggleClass("is-invisible", false);
                $("#register-input-username-icon-right").toggleClass("is-hidden", true);
            } else {
                alert(msg);
            }
        }
    });
}

function validateEmail() {
    let val = $("#register-input-email").val();

    console.log(validateEmailString(val))
    if (val == "" || !validateEmailString(val)) {
        $("#register-input-email-icon-right").toggleClass("is-hidden", true);
        return
    }

    $.ajax({
        type: "POST",
        url: "/register/email/" + val,
        success: function () {
            $("#register-input-email-help").toggleClass("is-invisible", true);
            $("#register-input-email-icon-right").toggleClass("is-hidden", false);
        },
        error: function (event, _, msg) {
            if (event.status == 409) {
                $("#register-input-email-help").toggleClass("is-invisible", false);
                $("#register-input-email-icon-right").toggleClass("is-hidden", true);
            } else {
                alert(msg);
            }
        }
    });
}

function validateEmailString(email) {
    var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(String(email).toLowerCase());
}

function validatePassword() {
    pw1 = $('#register-input-password').val()
    pw2 = $('#register-input-repeat').val()

    if (pw2 == "") {
        $("#register-input-repeat-icon-right").toggleClass("is-hidden", true);
        if (pw1 == "") {
            $("#register-input-repeat-help").toggleClass("is-invisible", true);
        }
        return false;
    }

    if (pw1 === pw2) {
        $("#register-input-repeat-icon-right").toggleClass("is-hidden", false);
        $("#register-input-repeat-help").toggleClass("is-invisible", true);
        return true;
    } else {
        $("#register-input-repeat-icon-right").toggleClass("is-hidden", true);
        $("#register-input-repeat-help").toggleClass("is-invisible", false);
        return false;
    }
}

function register(){
    validateName();
    validateEmail();

    if (!validatePassword()){
        return false;
    }

    elements = document.forms["register"].elements;
    var data = {
        "Name": elements["Name"].value,
        "EMail": elements["EMail"].value,
        "Passwort": sha512(elements["Passwort"].value),
        "Akzeptiert": elements["Akzeptiert"].checked
    }

    $.ajax({
        type: "POST",
        url: "/register",
        data: JSON.stringify(data),
        success: function () {
            window.location.href = "/meinekarteien";
        },
        error: function (event, jqXHR, msg) {
            if (event.status == 409) {
                alert("Fehlerhafte Eingaben!");
            } else {
                alert(msg);
            }
        },
        contentType: "application/json"
    });

    return false;
}