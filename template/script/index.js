
$(document).ready(function () {
    $("#test-request").on('click', function () {   
        $.ajax({
            url: '/req',
            dataType: 'json',
            type: 'POST',
            success: function (result) {
                $("#response").val(JSON.stringify(result));
            },
            error: function (result) {
                window.alert("Failed, check response");
                $("#response").val(JSON.stringify(result));
            }
        });    
    });
});