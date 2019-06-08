
$(document).ready(function () {
    $('#lern-button-aufdecken').click(function () { aufdecken(); });
    $('#lern-button-skip').click(function () { location.reload(); });
})

function aufdecken() {
    $('#answer').toggleClass('is-hidden', 'false');

    $('#button-1').toggleClass('is-hidden', 'true');
    $('#button-2').toggleClass('is-hidden', 'false');

    $('#lern-button-richtig').click(function () { sendeErgebnis(true) });
    $('#lern-button-falsch').click(function () { sendeErgebnis(false) });
}

function sendeErgebnis(ergebnis) {

    var data = {
        Index: parseInt($('#card-info').attr("index")),
        Ergebnis: ergebnis
    }


    $.ajax({
        type: "POST",
        data: JSON.stringify(data),
        success: function () { location.reload(); },
        error: defaultErrorHandling,
        contentType: "application/json"
    });
}