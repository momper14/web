function loadkarten() {
    for (let i = 1; i <= 6; i++) {
        loadTemplateKarte(function ($karte) {
            templateKarteReplace($karte, i, "Titel der Karte");
        });
    }
}

function loadTemplateKarte(todo = function () {}) {
    loadTemplate("res/html/karteikasten/karte-template.html", function (data) {
        let $karte = $('<html />', {
            html: data
        });

        todo($karte);

        $('#cards').append($karte.html());
    })
}

function templateKarteReplace($karte, nr, title) {
    $karte.find('#placeholder-nr').html(nr);
    $karte.find('#placeholder-nr').removeAttr('id');

    $karte.find('#placeholder-titel').html(title);
    $karte.find('#placeholder-titel').removeAttr('id');
}