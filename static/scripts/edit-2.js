var fragemde;
var antwortmde;
var uuidlength = 32;

$(document).ready(function () {
    $('#edit2-button-cancel').click(reset);
    $('#edit2-form-edit').submit(send);
    $('#edit2-modal-button-keep').click(function () {
        $('#edit2-modal').toggleClass("is-active", false);
    });
    $('.button-loeschen').submit(loeschen);
})

function send() {
    let data = {
        "Titel": $('#edit2-input-title').val(),
        "Frage": fragemde.markdown(fragemde.value()),
        "Antwort": antwortmde.markdown(antwortmde.value())
    }

    if (data.Frage == "") {
        alert("Frage darf nicht leer sein!");
        return false;
    }

    if (data.Antwort == "") {
        alert("Antwort darf nicht leer sein!");
        return false;
    }

    $.ajax({
        type: method(),
        data: JSON.stringify(data),
        success: reset,
        error: defaultErrorHandling,
        contentType: "application/json"
    });

    return false;
}

function method() {
    let split = window.location.href.split("/");
    let last = split[split.length - 1];

    if ((lastIndex = last.lastIndexOf("?")) != -1) {
        last = last.substr(0, lastIndex);
    }

    if (last.length == uuidlength) {
        return "POST";
    } else {
        return "PUT";
    }
}

function reset() {
    window.location.href = $('#edit2-button-cancel').attr("value");
}

function loeschen() {

    let url = $(this).attr("action");

    function remove(url) {
        $.ajax({
            type: "DELETE",
            url: url,
            success: function () { location.reload(); },
            error: defaultErrorHandling
        });
    }

    $('#edit2-modal-button-delete').click(function () { remove(url); });

    $('#edit2-modal').toggleClass("is-active", true);

    return false;
}