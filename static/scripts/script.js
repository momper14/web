$(document).ready(function () {
    $('.dropdown').click(function () {
        $(this).toggleClass("is-active");
    });
})

function loadTemplate(path, todo, async = true) {
    $.ajax({
        url: path,
        success: function (data) {
            todo(data);
        },
        async: false,
        dataType: 'html'
    });
}