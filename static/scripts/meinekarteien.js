$(document).ready(function () {
    $('#meine-modal-button-keep').click(function () {
        $('#meine-modal').toggleClass("is-active", false);
    });
    $('.button-loeschen').submit(loeschen);
})

function loeschen() {

    let url = $(this).attr("action");

    function remove(url) {
        console.log("test");
        console.log(url);
        $.ajax({
            type: "REMOVE",
            url: url,
            success: function () { location.reload(); },
            error: function (xhr, _, msg) {
                alert(xhr.status + "\n\n" + msg);
            }
        });
    }

    $('#meine-modal-button-delete').click(function () { remove(url); });

    $('#meine-modal').toggleClass("is-active", true);

    return false;
}