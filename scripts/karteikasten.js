const BESCHREIBUNG = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. <strong> Pellentesque risus mi </strong>, tempus quis placerat ut, porta nec nulla.Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum <a> felis venenatis </a> efficitur. Aenean ac <em>eleifend lacus </em>, in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque.Donec dui urna, vehicula et sem eget, facilisis sodales sem.";
const KATEGORIE_PREFIX = "placeholder-kategorie-";

function loadKategorien() {
    loadTemplateKategorie(function ($kategorie) {
        templateKategorieReplace($kategorie, "Naturwissenschaften");
    });
    loadTemplateKategorie(function ($kategorie) {
        templateKategorieReplace($kategorie, "Sprachen");
    });
    loadKaesten();
}

function loadKaesten() {
    loadTemplateKasten("Naturwissenschaften", function ($kasten) {
        templateKastenReplace($kasten, "Mathematik", "Geometrische Formen und Körper", 23, BESCHREIBUNG);
    });
    loadTemplateKasten("Naturwissenschaften", function ($kasten) {
        templateKastenReplace($kasten, "Chemie", "Atome A-Z", 23, BESCHREIBUNG);
    });
    loadTemplateKasten("Naturwissenschaften", function ($kasten) {
        templateKastenReplace($kasten, "Physik", "Licht in Wellen und Teilchen - Modelle und Versuche", 23, BESCHREIBUNG);
    });
    loadTemplateKasten("Naturwissenschaften", function ($kasten) {
        templateKastenReplace($kasten, "Mathematik", "Geometrische Formen und Körper", 23, BESCHREIBUNG);
    });

    loadTemplateKasten("Sprachen", function ($kasten) {
        templateKastenReplace($kasten, "Latein", "Vokabel Lektion 1", 23, BESCHREIBUNG);
    });
    loadTemplateKasten("Sprachen", function ($kasten) {
        templateKastenReplace($kasten, "Englisch", "Unit 2", 23, BESCHREIBUNG);
    });
}

function loadTemplate(path, todo) {
    $.ajax({
        url: path,
        success: function (data) {
            todo(data);
        },
        dataType: 'html'
    });
}

function loadTemplateKasten(kategorie, todo = function () {}) {
    loadTemplate("res/html/karteikasten/kasten-template.html", function (data) {
        let $kasten = $('<html />', {
            html: data
        });

        todo($kasten);

        $(`#${KATEGORIE_PREFIX}${kategorie}`).append($kasten.html());
    })
}

function loadTemplateKategorie(todo = function () {}) {
    loadTemplate("res/html/karteikasten/kategorie-template.html", function (data) {
        let $kategorie = $('<html />', {
            html: data
        });

        todo($kategorie);

        $('#placeholder-kategorien').append($kategorie.html());
    })
}

function templateKastenReplace($kasten, unterkategorie, title, anzahl, beschreibung) {
    $kasten.find('#placeholder-unterkategorie').html(unterkategorie);
    $kasten.find('#placeholder-unterkategorie').removeAttr('id');

    $kasten.find('#placeholder-title').html(title);
    $kasten.find('#placeholder-title').removeAttr('id');

    $kasten.find('#placeholder-anzahl').html(anzahl);
    $kasten.find('#placeholder-anzahl').removeAttr('id');

    $kasten.find('#placeholder-beschreibung').html(beschreibung);
    $kasten.find('#placeholder-beschreibung').removeAttr('id');

    //$kasten.find('#placeholder-unterkategorie').html("Mathematik");
    $kasten.find('#button').removeAttr('id');
}

function templateKategorieReplace($kategorie, title) {
    $kategorie.find('#placeholder-title').html(title);
    $kategorie.find('#placeholder-title').removeAttr('id');
    $kategorie.find('#placeholder-kaesten').attr('id', `${KATEGORIE_PREFIX}${title}`);
}