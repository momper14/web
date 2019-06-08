$(document).ready(function () {
    $('#form-edit').submit(create);
})

function create() {
    let elements = document.forms["form-edit"].elements;
    let selected = $(':selected', $('#edit-optgroup-kat'))

    let data = {
        "Kategorie": selected.parent().attr("label"),
        "Unterkategorie": selected.attr("value"),
        "Titel": elements["Titel"].value,
        "Beschreibung": elements["Beschreibung"].value,
        "Public": elements["Sichtbarkeit"].value === 'true'
    }

    let arr = window.location.href.split("/");
    let last = arr[arr.length - 1]

    if (last == "edit" || last == "edit?") {
        $.ajax({
            type: "POST",
            data: JSON.stringify(data),
            success: function (xhr) {
                window.location.href = "/karteikasten/edit-2/" + xhr;
            },
            error: defaultErrorHandling,
            contentType: "application/json"
        });
    } else {
        $.ajax({
            type: "PUT",
            data: JSON.stringify(data),
            success: function () {
                window.location.href = "/karteikasten/edit-2/" + last;
            },
            error: defaultErrorHandling,
            contentType: "application/json"
        });
    }

    return false;
}