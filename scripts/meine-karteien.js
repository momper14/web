const BESCHREIBUNG = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. <strong> Pellentesque risus mi </strong>, tempus quis placerat ut, porta nec nulla.Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum <a> felis venenatis </a> efficitur. Aenean ac <em>eleifend lacus </em>, in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque.Donec dui urna, vehicula et sem eget, facilisis sodales sem.";
const KATEGORIE_PREFIX = "placeholder-kategorie-";

function loadKategorien() {
    loadTemplateKategorie(function ($kategorie) {
        templateKategorieReplace($kategorie, "Selbst erstellte Karteikästen");
    });
    loadTemplateKategorie(function ($kategorie) {
        templateKategorieReplace($kategorie, "Gelernte Karteikästen anderer Nutzer");
    });
    loadKaesten();
}

function loadKaesten() {
    loadTemplateKasten("selbst-erstellte-karteikästen", function ($kasten) {
        templateKastenReplace($kasten, "Naturwissenschaften", "Mathematik", "Geometrische Formen und Körper", 23, BESCHREIBUNG, "Öffentlich", 76);
    });
    loadTemplateKasten("selbst-erstellte-karteikästen", function ($kasten) {
        templateKastenReplace($kasten, "Naturwissenschaften", "Chemie", "Atome A-Z", 23, BESCHREIBUNG, "Privat", 20);
    });
    loadTemplateKasten("selbst-erstellte-karteikästen", function ($kasten) {
        templateKastenReplace($kasten, "Sprachen", "Latein", "Vokabeln Lekrtion 1", 23, BESCHREIBUNG, "Öffentlich", 0);
    });
    loadTemplateKasten("selbst-erstellte-karteikästen", function ($kasten) {
        templateKastenReplace($kasten, "Gesellschaft", "Verkehrskunde", "Theoriefragen Fahrprüfung", 23, BESCHREIBUNG, "Öffentlich", 6);
    });

    loadTemplateKasten("gelernte-karteikästen-anderer-nutzer", function ($kasten) {
        templateKastenReplace($kasten, "Naturwissenschaften", "Physik", "Lorem Ipsum", 23, BESCHREIBUNG, "Öffentlich", 100);
    });
}

function loadTemplateKasten(kategorie, todo = function () {}) {
    loadTemplate("res/html/karteikasten/mein-kasten-template.html", function (data) {
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
    }, false)
}

function templateKastenReplace($kasten, kategorie, unterkategorie, title, anzahl, beschreibung, sichtbarkeit, fortschritt) {
    $kasten.find('#placeholder-kategorie').html(kategorie);
    $kasten.find('#placeholder-kategorie').removeAttr('id');

    $kasten.find('#placeholder-unterkategorie').html(unterkategorie);
    $kasten.find('#placeholder-unterkategorie').removeAttr('id');

    $kasten.find('#placeholder-title').html(title);
    $kasten.find('#placeholder-title').removeAttr('id');

    $kasten.find('#placeholder-anzahl').html(anzahl);
    $kasten.find('#placeholder-anzahl').removeAttr('id');

    $kasten.find('#placeholder-beschreibung').html(beschreibung);
    $kasten.find('#placeholder-beschreibung').removeAttr('id');

    $kasten.find('#placeholder-sichtbarkeit').html(sichtbarkeit);
    $kasten.find('#placeholder-sichtbarkeit').removeAttr('id');

    $kasten.find('#placeholder-fortschritt').html(fortschritt);
    $kasten.find('#placeholder-fortschritt').removeAttr('id');
}

function templateKategorieReplace($kategorie, title) {
    $kategorie.find('#placeholder-title').html(title);
    $kategorie.find('#placeholder-title').removeAttr('id');
    title = title.toLowerCase().replace(/ /g, '-');
    $kategorie.find('#placeholder-kaesten').attr('id', `${KATEGORIE_PREFIX}${title}`);
}