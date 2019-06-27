$(document).ready(function () {
    $('#karteikasten-filter-kategorie').click(onChange);
})

function onChange() {
    let selected = $(':selected', $(this))
    unterkategorie = selected.attr("value")
    if (unterkategorie === "") {
        window.location.href = window.location.href.split('?')[0]
    } else {
        oberkategorie = selected.parent().attr("label")

        path = window.location.href.split('?')[0]
        if (path.charAt(path.length - 1) == '#') {
            path = path.substring(0, path.length - 1);
        }

        window.location.href = path
            + "?" + "oberkategorie" + "=" + oberkategorie
            + "&" + "unterkategorie" + "=" + unterkategorie;
    }
}