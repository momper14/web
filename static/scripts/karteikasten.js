var url;
$(document).ready(function () {
    $('#karteikasten-filter-kategorie').click(onChange);
})

function onChange() {

    let selected = $(':selected', $(this))
    let unterkategorie = selected.attr("value");
    if (unterkategorie === "") {
        let tmp = window.location.href;

        if (tmp != tmp.split('?')[0]) {
            window.location.href = tmp.split('?')[0];
        }
    } else {
        let tmp = window.location.href;

        let oberkategorie = selected.parent().attr("label")

        let url = window.location.href.split('?')[0];
        if (url.charAt(url.length - 1) == '#') {
            url = url.substring(0, url.length - 1);
        }

        url = url
            + "?" + "oberkategorie" + "=" + oberkategorie
            + "&" + "unterkategorie" + "=" + unterkategorie;

        if (url != tmp) {
            window.location.href = url;
        }
    }
}