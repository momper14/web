$(document).ready(function () {
    $('#meine-modal-button-keep').click(function () {
        $('#meine-modal').toggleClass("is-active", false);
    });
    $('.button-loeschen').submit(loeschen);
})

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

    $('#meine-modal-button-delete').click(function () { remove(url); });

    $('#meine-modal').toggleClass("is-active", true);

    return false;
}